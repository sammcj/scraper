// Package gdb interacts with thegamedb.net's API.
package gdb

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	gamesdb "github.com/J-Swift/thegamesdb-swagger-client-go"
)

var apiClient = gamesdb.NewAPIClient(gamesdb.NewConfiguration())

// configureAPIKey sets up the API client with the provided API key
func configureAPIKey(apikey string) {
	apiClient.GetConfig().AddDefaultHeader("apikey", apikey)
}

// Publishers

type publishers struct {
	mux        sync.Mutex
	publishers *map[string]gamesdb.Publisher
}

var publishersCache = publishers{}

func getPublishers(ctx context.Context, apikey string) map[string]gamesdb.Publisher {
	configureAPIKey(apikey)
	request := apiClient.PublishersApi.Publishers(ctx)
	result, resp, err := request.Execute()

	if err != nil || resp.StatusCode != 200 {
		return make(map[string]gamesdb.Publisher)
	}

	return result.Data.Publishers
}

func getCachedPublishers(ctx context.Context, apikey string) map[string]gamesdb.Publisher {
	publishers := publishersCache.publishers
	if publishers != nil {
		return *publishers
	}

	publishersCache.mux.Lock()
	defer publishersCache.mux.Unlock()

	publishers = publishersCache.publishers
	if publishers == nil {
		apiPublishers := getPublishers(ctx, apikey)
		publishers = &apiPublishers
		publishersCache.publishers = publishers
	}

	return *publishers
}

// Developers

type developers struct {
	mux        sync.Mutex
	developers *map[string]gamesdb.Developer
}

var developersCache = developers{}

func getDevelopers(ctx context.Context, apikey string) map[string]gamesdb.Developer {
	configureAPIKey(apikey)
	request := apiClient.DevelopersApi.Developers(ctx)
	result, resp, err := request.Execute()

	if err != nil || resp.StatusCode != 200 {
		return make(map[string]gamesdb.Developer)
	}

	return result.Data.Developers
}

func getCachedDevelopers(ctx context.Context, apikey string) map[string]gamesdb.Developer {
	developers := developersCache.developers
	if developers != nil {
		return *developers
	}

	developersCache.mux.Lock()
	defer developersCache.mux.Unlock()

	developers = developersCache.developers
	if developers == nil {
		apiDevelopers := getDevelopers(ctx, apikey)
		developers = &apiDevelopers
		developersCache.developers = developers
	}

	return *developers
}

// Genres

type genres struct {
	mux    sync.Mutex
	genres *map[string]gamesdb.Genre
}

var genresCache = genres{}

func getGenres(ctx context.Context, apikey string) map[string]gamesdb.Genre {
	configureAPIKey(apikey)
	request := apiClient.GenresApi.Genres(ctx)
	result, resp, err := request.Execute()

	if err != nil || resp.StatusCode != 200 {
		return make(map[string]gamesdb.Genre)
	}

	return result.Data.Genres
}

func getCachedGenres(ctx context.Context, apikey string) map[string]gamesdb.Genre {
	genres := genresCache.genres
	if genres != nil {
		return *genres
	}

	genresCache.mux.Lock()
	defer genresCache.mux.Unlock()

	genres = genresCache.genres
	if genres == nil {
		apiGenres := getGenres(ctx, apikey)
		genres = &apiGenres
		genresCache.genres = genres
	}

	return *genres
}

// ParsedDeveloper is a normalized GamesDB Developer
type ParsedDeveloper struct {
	ID   int
	Name string
}

func toParsedDeveloper(apiDeveloper gamesdb.Developer) ParsedDeveloper {
	return ParsedDeveloper{
		ID:   int(apiDeveloper.Id),
		Name: apiDeveloper.Name,
	}
}

// ParsedGenre is a normalized GamesDB Genre
type ParsedGenre struct {
	ID   int
	Name string
}

func toParsedGenre(apiGenre gamesdb.Genre) ParsedGenre {
	return ParsedGenre{
		ID:   int(apiGenre.Id),
		Name: apiGenre.Name,
	}
}

// ParsedPublisher is a normalized GamesDB Publisher
type ParsedPublisher struct {
	ID   int
	Name string
}

func toParsedPublisher(apiPublisher gamesdb.Publisher) ParsedPublisher {
	return ParsedPublisher{
		ID:   int(apiPublisher.Id),
		Name: apiPublisher.Name,
	}
}

// ParsedGameImage is a normalized GamesDB GameImage
type ParsedGameImage struct {
	ID       int
	Type     string
	Side     string
	Filename string
}

func toParsedGameImage(apiGameImage gamesdb.GameImage) ParsedGameImage {
	var id int
	if apiGameImage.Id != nil {
		id = int(*apiGameImage.Id)
	}

	var imgType, side, filename string
	if apiGameImage.Type != nil {
		imgType = *apiGameImage.Type
	}
	if apiGameImage.Side != nil {
		side = *apiGameImage.Side
	}
	if apiGameImage.Filename != nil {
		filename = *apiGameImage.Filename
	}

	return ParsedGameImage{
		ID:       id,
		Type:     imgType,
		Side:     side,
		Filename: filename,
	}
}

// ParsedImageSizeBaseUrls is a normalized GamesDB ImageBaseUrlMeta
type ParsedImageSizeBaseUrls struct {
	Original string
	Thumb    string
}

func toParsedImageSizeBaseUrls(apiBaseURLMeta gamesdb.ImageBaseUrlMeta) ParsedImageSizeBaseUrls {
	return ParsedImageSizeBaseUrls{
		Original: apiBaseURLMeta.Original,
		Thumb:    apiBaseURLMeta.Thumb,
	}
}

// ParsedGame is  a normalized GamesDB Game
type ParsedGame struct {
	ID          int
	Name        string
	ReleaseDate string
	//Platform    int
	Players    int
	Overview   string
	Developers []ParsedDeveloper
	Genres     []ParsedGenre
	Publishers []ParsedPublisher

	Images        map[string][]ParsedGameImage
	ImageBaseUrls ParsedImageSizeBaseUrls
}

// GetGame gets the game information from the DB.
func GetGame(ctx context.Context, apikey string, gameID string) (*ParsedGame, error) {
	configureAPIKey(apikey)

	if gameID == "" {
		return nil, fmt.Errorf("must provide an ID or Name")
	}

	// Create request with fields filter
	request := apiClient.GamesApi.GamesByGameID(ctx)
	request = request.Id(gameID)
	request = request.Fields("players,publishers,genres,overview,platform")

	result, resp, err := request.Execute()
	if err != nil {
		log.Printf("ERR: TheGamesDB API error: %v", err)
		if resp != nil {
			log.Printf("ERR: Response status: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("getting game error: %s", err)
	}

	if len(result.Data.Games) == 0 {
		return nil, fmt.Errorf("game not found")
	}

	apiGame := result.Data.Games[0]
	res := &ParsedGame{
		ID:          int(*apiGame.Id),
		Name:        *apiGame.GameTitle,
		ReleaseDate: *apiGame.ReleaseDate,
		Players:     int(*apiGame.Players),
		Overview:    *apiGame.Overview,
	}

	allGenres := getCachedGenres(ctx, apikey)
	genres := []ParsedGenre{}
	for _, genreID := range apiGame.Genres {
		if apiGenre, ok := allGenres[strconv.Itoa(int(genreID))]; ok {
			genres = append(genres, toParsedGenre(apiGenre))
		}
	}
	res.Genres = genres

	allDevelopers := getCachedDevelopers(ctx, apikey)
	developers := []ParsedDeveloper{}
	for _, developerID := range apiGame.Developers {
		if apiDeveloper, ok := allDevelopers[strconv.Itoa(int(developerID))]; ok {
			developers = append(developers, toParsedDeveloper(apiDeveloper))
		}
	}
	res.Developers = developers

	allPublishers := getCachedPublishers(ctx, apikey)
	publishers := []ParsedPublisher{}
	for _, publisherID := range apiGame.Publishers {
		if apiPublisher, ok := allPublishers[strconv.Itoa(int(publisherID))]; ok {
			publishers = append(publishers, toParsedPublisher(apiPublisher))
		}
	}
	res.Publishers = publishers

	// For now, skip image handling as the API structure has changed
	// TODO: Update image handling once we understand the new API structure
	res.Images = make(map[string][]ParsedGameImage)
	res.ImageBaseUrls = ParsedImageSizeBaseUrls{
		Original: "",
		Thumb:    "",
	}

	return res, nil
}

// IsUp returns if thegamedb.net is up.
func IsUp(ctx context.Context, apikey string) bool {
	configureAPIKey(apikey)

	request := apiClient.GamesApi.GamesByGameID(ctx)
	request = request.Id("1")
	request = request.Fields("id") // Minimal fields to reduce response size

	result, _, err := request.Execute()
	if err != nil {
		log.Printf("ERR: TheGamesDB API error: %v", err)
		return false
	}

	// Check if we got a valid response with games data
	if result == nil || len(result.Data.Games) == 0 {
		log.Printf("ERR: TheGamesDB API returned no games data")
		return false
	}

	return true
}
