package ds

import (
	"compress/gzip"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/nfnt/resize"
)

const (
	hashURL  = "https://storage.googleapis.com/stevenselph.appspot.com/hash.csv.gz"
	hashName = "hash.csv"
	hashMeta = "hash.meta"
)

// ErrNotFound is the error returned when a rom can't be found in the source.
var ErrNotFound = errors.New("hash not found")

// ErrImgNotFound is the error returned when a rom image can't be found.
var ErrImgNotFound = errors.New("image not found")

// ImgType represents the different image types that sources may provide.
type ImgType string

// Image types for Datasource options. Not all types are valid for all sources.
const (
	ImgBoxart    ImgType = "b"
	ImgBoxart3D  ImgType = "3b"
	ImgScreen    ImgType = "s"
	ImgFanart    ImgType = "f"
	ImgBanner    ImgType = "a"
	ImgLogo      ImgType = "l"
	ImgTitle     ImgType = "t"
	ImgMarquee   ImgType = "m"
	ImgCabinet   ImgType = "c"
	ImgFlyer     ImgType = "fly"
	ImgCart      ImgType = "cart"
	ImgCartLabel ImgType = "clabel"
	ImgMix3      ImgType = "mix3"
	ImgMix4      ImgType = "mix4"
)

// VidType represents the different video types that sources may provide.
type VidType string

// Video types for Datasource options. Not all types are valid for all sources.
const (
	VidStandard VidType = "v"
)

// Game is the standard Game that all sources will return.
// They don't have to populate all values.
type Game struct {
	ID          string
	Source      string
	GameTitle   string
	Overview    string
	Images      map[ImgType]Image
	Thumbs      map[ImgType]Image
	Videos      map[VidType]Video
	Rating      float64
	ReleaseDate string
	Developer   string
	Publisher   string
	Genre       string
	Players     int64
	CloneOf     string
}

// NewGame returns a new Game.
func NewGame() *Game {
	g := &Game{}
	g.Images = make(map[ImgType]Image)
	g.Thumbs = make(map[ImgType]Image)
	g.Videos = make(map[VidType]Video)
	return g
}

// DS is the interface all DataSoures should implement.
type DS interface {
	// GetName takes the path of a ROM and returns the Pretty name if it differs from the Sources normal name.
	GetName(string) string
	// GetGame takes an id and returns the Game.
	GetGame(context.Context, string) (*Game, error)
}

type Video interface {
	Save(ctx context.Context, p string) error
	Ext() string
}

type HTTPVideo struct {
	URL   string
	E     string
	Limit chan struct{}
}

func (v HTTPVideo) Save(ctx context.Context, p string) error {
	if v.Limit != nil {
		v.Limit <- struct{}{}
		defer func() {
			<-v.Limit
		}()
	}
	req, err := http.NewRequest("GET", v.URL, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return ErrImgNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%v from %s", resp.StatusCode, v.URL)
	}
	defer resp.Body.Close()
	out, err := os.Create(p)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func (v HTTPVideo) Ext() string {
	return v.E
}

type Image interface {
	Get(ctx context.Context, w, h uint) (image.Image, error)
	Save(ctx context.Context, p string, w, h uint) error
}

type HTTPImage struct {
	URL   string
	Limit chan struct{}
}

func (i HTTPImage) Get(ctx context.Context, w, h uint) (image.Image, error) {
	if i.Limit != nil {
		i.Limit <- struct{}{}
		defer func() {
			<-i.Limit
		}()
	}
	return getImage(ctx, i.URL, w, h)
}

func (i HTTPImage) Save(ctx context.Context, p string, w, h uint) error {
	img, err := i.Get(ctx, w, h)
	if err != nil {
		return err
	}
	out, err := os.Create(p)
	if err != nil {
		return err
	}
	defer out.Close()
	e := filepath.Ext(p)
	switch e {
	case ".jpg":
		return jpeg.Encode(out, img, nil)
	case ".png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("invalid image type")
	}
}

func getImage(ctx context.Context, url string, w, h uint) (image.Image, error) {
	if w == 0 {
		w = math.MaxUint32
	}
	if h == 0 {
		h = math.MaxUint32
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrImgNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v from %s", resp.StatusCode, url)
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	img = resize.Thumbnail(w, h, img, resize.Bilinear)
	return img, nil
}

// mkDir checks if directory exists and if it doesn't create it.
func mkDir(d string) error {
	fi, err := os.Stat(d)
	switch {
	case os.IsNotExist(err):
		return os.MkdirAll(d, 0775)
	case err != nil:
		return err
	case fi.IsDir():
		return nil
	}
	return fmt.Errorf("%s is a file not a directory", d)
}

// HashMap a mapping of hash to names and GDB IDs.
type HashMap struct {
	data map[string][]string
}

// GetID returns the id for the given hash.
func (hm *HashMap) ID(s string) (string, bool) {
	if hm == nil {
		return "", false
	}
	d, ok := hm.data[s]
	if !ok || d[0] == "" {
		return "", false
	}
	return d[0], true
}

// GetName returns the no-intro name for the given hash.
func (hm *HashMap) Name(s string) (string, bool) {
	if hm == nil {
		return "", false
	}
	d, ok := hm.data[s]
	if !ok || d[1] == "" {
		return "", false
	}
	return d[1], true
}

func (hm *HashMap) System(s string) (int, bool) {
	if hm == nil {
		return 0, false
	}
	d, ok := hm.data[s]
	if !ok || d[2] == "" {
		return 0, false
	}
	i, err := strconv.Atoi(d[2])
	if err != nil {
		return 0, false
	}
	return i, true
}

// DefaultCachePath returns the path used for all cached data.
// Current <HOME>/.sselph-scraper or <HOMEDIR>\Application Data\sselph-scraper
func DefaultCachePath() (string, error) {
	hd, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	var p string
	if runtime.GOOS == "windows" {
		p = filepath.Join(hd, "Application Data", "sselph-scraper")
	} else {
		p = filepath.Join(hd, ".sselph-scraper")
	}
	err = mkDir(p)
	if err != nil {
		return "", err
	}
	return p, nil
}

// updateHash downloads the latest hash file.
func updateHash(ctx context.Context, version, p string) error {
	log.Print("INFO: Checking for new hash.csv.")
	req, err := http.NewRequest("GET", hashURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("if-none-match", version)
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotModified {
		log.Printf("INFO: hash.csv %s up to date.", version)
		return nil
	}
	if version != "" && resp.StatusCode == 429 {
		log.Printf("WARN: Using cached hash.csv. Server over quota.")
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got %v response", resp.Status)
	}
	newVersion := resp.Header.Get("etag")
	log.Printf("INFO: Upgrading hash.csv: %s -> %s.", version, newVersion)
	bz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer bz.Close()
	b, err := io.ReadAll(bz)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(p, hashName), b, 0664)
	if err != nil {
		return err
	}
	os.WriteFile(filepath.Join(p, hashMeta), []byte(newVersion), 0664)
	return nil
}

// exists checks if a file exists and contains data.
func exists(s string) bool {
	fi, err := os.Stat(s)
	return !os.IsNotExist(err) && fi.Size() > 0
}

// CachedHashMap gets the mapping of hashes to IDs.
func CachedHashMap(ctx context.Context, p string, u bool) (*HashMap, error) {
	var err error
	if p == "" {
		p, err = DefaultCachePath()
		if err != nil {
			return nil, err
		}
	}
	fp := filepath.Join(p, hashName)
	if exists(fp) && !u {
		return FileHashMap(fp)
	}
	mp := filepath.Join(p, hashMeta)
	var version string
	if exists(fp) && exists(mp) {
		b, err := os.ReadFile(mp)
		if err != nil {
			return nil, err
		}
		version = strings.Trim(string(b[:]), "\n\r")
	}
	err = updateHash(ctx, version, p)
	if err != nil {
		return nil, err
	}
	return FileHashMap(fp)
}

// FileHashMap creates a hash map from a csv file.
func FileHashMap(p string) (*HashMap, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	c := csv.NewReader(f)
	r, err := c.ReadAll()
	if err != nil {
		return nil, err
	}
	ret := &HashMap{data: make(map[string][]string)}
	for _, v := range r {
		ret.data[strings.ToLower(v[0])] = []string{v[1], v[3], v[2]}
	}
	return ret, nil
}
