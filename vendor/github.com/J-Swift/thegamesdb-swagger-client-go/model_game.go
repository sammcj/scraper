/*
TheGamesDB API

API Documentation

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gamesdb

import (
	"encoding/json"
)

// Game struct for Game
type Game struct {
	Id          *int32   `json:"id,omitempty"`
	GameTitle   *string  `json:"game_title,omitempty"`
	ReleaseDate *string  `json:"release_date,omitempty"`
	Platform    *int32   `json:"platform,omitempty"`
	Players     *int32   `json:"players,omitempty"`
	Overview    *string  `json:"overview,omitempty"`
	LastUpdated *string  `json:"last_updated,omitempty"`
	Rating      *string  `json:"rating,omitempty"`
	Coop        *string  `json:"coop,omitempty"`
	Youtube     *string  `json:"youtube,omitempty"`
	Os          *string  `json:"os,omitempty"`
	Processor   *string  `json:"processor,omitempty"`
	Ram         *string  `json:"ram,omitempty"`
	Hdd         *string  `json:"hdd,omitempty"`
	Video       *string  `json:"video,omitempty"`
	Sound       *string  `json:"sound,omitempty"`
	Developers  []int32  `json:"developers,omitempty"`
	Genres      []int32  `json:"genres,omitempty"`
	Publishers  []int32  `json:"publishers,omitempty"`
	Alternates  []string `json:"alternates,omitempty"`
}

// NewGame instantiates a new Game object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGame() *Game {
	this := Game{}
	return &this
}

// NewGameWithDefaults instantiates a new Game object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameWithDefaults() *Game {
	this := Game{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Game) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Game) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Game) SetId(v int32) {
	o.Id = &v
}

// GetGameTitle returns the GameTitle field value if set, zero value otherwise.
func (o *Game) GetGameTitle() string {
	if o == nil || o.GameTitle == nil {
		var ret string
		return ret
	}
	return *o.GameTitle
}

// GetGameTitleOk returns a tuple with the GameTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetGameTitleOk() (*string, bool) {
	if o == nil || o.GameTitle == nil {
		return nil, false
	}
	return o.GameTitle, true
}

// HasGameTitle returns a boolean if a field has been set.
func (o *Game) HasGameTitle() bool {
	if o != nil && o.GameTitle != nil {
		return true
	}

	return false
}

// SetGameTitle gets a reference to the given string and assigns it to the GameTitle field.
func (o *Game) SetGameTitle(v string) {
	o.GameTitle = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *Game) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *Game) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *Game) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *Game) GetPlatform() int32 {
	if o == nil || o.Platform == nil {
		var ret int32
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetPlatformOk() (*int32, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *Game) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given int32 and assigns it to the Platform field.
func (o *Game) SetPlatform(v int32) {
	o.Platform = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *Game) GetPlayers() int32 {
	if o == nil || o.Players == nil {
		var ret int32
		return ret
	}
	return *o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetPlayersOk() (*int32, bool) {
	if o == nil || o.Players == nil {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *Game) HasPlayers() bool {
	if o != nil && o.Players != nil {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given int32 and assigns it to the Players field.
func (o *Game) SetPlayers(v int32) {
	o.Players = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *Game) GetOverview() string {
	if o == nil || o.Overview == nil {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetOverviewOk() (*string, bool) {
	if o == nil || o.Overview == nil {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *Game) HasOverview() bool {
	if o != nil && o.Overview != nil {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *Game) SetOverview(v string) {
	o.Overview = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Game) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Game) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *Game) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *Game) GetRating() string {
	if o == nil || o.Rating == nil {
		var ret string
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetRatingOk() (*string, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *Game) HasRating() bool {
	if o != nil && o.Rating != nil {
		return true
	}

	return false
}

// SetRating gets a reference to the given string and assigns it to the Rating field.
func (o *Game) SetRating(v string) {
	o.Rating = &v
}

// GetCoop returns the Coop field value if set, zero value otherwise.
func (o *Game) GetCoop() string {
	if o == nil || o.Coop == nil {
		var ret string
		return ret
	}
	return *o.Coop
}

// GetCoopOk returns a tuple with the Coop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetCoopOk() (*string, bool) {
	if o == nil || o.Coop == nil {
		return nil, false
	}
	return o.Coop, true
}

// HasCoop returns a boolean if a field has been set.
func (o *Game) HasCoop() bool {
	if o != nil && o.Coop != nil {
		return true
	}

	return false
}

// SetCoop gets a reference to the given string and assigns it to the Coop field.
func (o *Game) SetCoop(v string) {
	o.Coop = &v
}

// GetYoutube returns the Youtube field value if set, zero value otherwise.
func (o *Game) GetYoutube() string {
	if o == nil || o.Youtube == nil {
		var ret string
		return ret
	}
	return *o.Youtube
}

// GetYoutubeOk returns a tuple with the Youtube field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetYoutubeOk() (*string, bool) {
	if o == nil || o.Youtube == nil {
		return nil, false
	}
	return o.Youtube, true
}

// HasYoutube returns a boolean if a field has been set.
func (o *Game) HasYoutube() bool {
	if o != nil && o.Youtube != nil {
		return true
	}

	return false
}

// SetYoutube gets a reference to the given string and assigns it to the Youtube field.
func (o *Game) SetYoutube(v string) {
	o.Youtube = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *Game) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *Game) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *Game) SetOs(v string) {
	o.Os = &v
}

// GetProcessor returns the Processor field value if set, zero value otherwise.
func (o *Game) GetProcessor() string {
	if o == nil || o.Processor == nil {
		var ret string
		return ret
	}
	return *o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetProcessorOk() (*string, bool) {
	if o == nil || o.Processor == nil {
		return nil, false
	}
	return o.Processor, true
}

// HasProcessor returns a boolean if a field has been set.
func (o *Game) HasProcessor() bool {
	if o != nil && o.Processor != nil {
		return true
	}

	return false
}

// SetProcessor gets a reference to the given string and assigns it to the Processor field.
func (o *Game) SetProcessor(v string) {
	o.Processor = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *Game) GetRam() string {
	if o == nil || o.Ram == nil {
		var ret string
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetRamOk() (*string, bool) {
	if o == nil || o.Ram == nil {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *Game) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// SetRam gets a reference to the given string and assigns it to the Ram field.
func (o *Game) SetRam(v string) {
	o.Ram = &v
}

// GetHdd returns the Hdd field value if set, zero value otherwise.
func (o *Game) GetHdd() string {
	if o == nil || o.Hdd == nil {
		var ret string
		return ret
	}
	return *o.Hdd
}

// GetHddOk returns a tuple with the Hdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetHddOk() (*string, bool) {
	if o == nil || o.Hdd == nil {
		return nil, false
	}
	return o.Hdd, true
}

// HasHdd returns a boolean if a field has been set.
func (o *Game) HasHdd() bool {
	if o != nil && o.Hdd != nil {
		return true
	}

	return false
}

// SetHdd gets a reference to the given string and assigns it to the Hdd field.
func (o *Game) SetHdd(v string) {
	o.Hdd = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *Game) GetVideo() string {
	if o == nil || o.Video == nil {
		var ret string
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetVideoOk() (*string, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *Game) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given string and assigns it to the Video field.
func (o *Game) SetVideo(v string) {
	o.Video = &v
}

// GetSound returns the Sound field value if set, zero value otherwise.
func (o *Game) GetSound() string {
	if o == nil || o.Sound == nil {
		var ret string
		return ret
	}
	return *o.Sound
}

// GetSoundOk returns a tuple with the Sound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetSoundOk() (*string, bool) {
	if o == nil || o.Sound == nil {
		return nil, false
	}
	return o.Sound, true
}

// HasSound returns a boolean if a field has been set.
func (o *Game) HasSound() bool {
	if o != nil && o.Sound != nil {
		return true
	}

	return false
}

// SetSound gets a reference to the given string and assigns it to the Sound field.
func (o *Game) SetSound(v string) {
	o.Sound = &v
}

// GetDevelopers returns the Developers field value if set, zero value otherwise.
func (o *Game) GetDevelopers() []int32 {
	if o == nil || o.Developers == nil {
		var ret []int32
		return ret
	}
	return o.Developers
}

// GetDevelopersOk returns a tuple with the Developers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetDevelopersOk() ([]int32, bool) {
	if o == nil || o.Developers == nil {
		return nil, false
	}
	return o.Developers, true
}

// HasDevelopers returns a boolean if a field has been set.
func (o *Game) HasDevelopers() bool {
	if o != nil && o.Developers != nil {
		return true
	}

	return false
}

// SetDevelopers gets a reference to the given []int32 and assigns it to the Developers field.
func (o *Game) SetDevelopers(v []int32) {
	o.Developers = v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *Game) GetGenres() []int32 {
	if o == nil || o.Genres == nil {
		var ret []int32
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetGenresOk() ([]int32, bool) {
	if o == nil || o.Genres == nil {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *Game) HasGenres() bool {
	if o != nil && o.Genres != nil {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []int32 and assigns it to the Genres field.
func (o *Game) SetGenres(v []int32) {
	o.Genres = v
}

// GetPublishers returns the Publishers field value if set, zero value otherwise.
func (o *Game) GetPublishers() []int32 {
	if o == nil || o.Publishers == nil {
		var ret []int32
		return ret
	}
	return o.Publishers
}

// GetPublishersOk returns a tuple with the Publishers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetPublishersOk() ([]int32, bool) {
	if o == nil || o.Publishers == nil {
		return nil, false
	}
	return o.Publishers, true
}

// HasPublishers returns a boolean if a field has been set.
func (o *Game) HasPublishers() bool {
	if o != nil && o.Publishers != nil {
		return true
	}

	return false
}

// SetPublishers gets a reference to the given []int32 and assigns it to the Publishers field.
func (o *Game) SetPublishers(v []int32) {
	o.Publishers = v
}

// GetAlternates returns the Alternates field value if set, zero value otherwise.
func (o *Game) GetAlternates() []string {
	if o == nil || o.Alternates == nil {
		var ret []string
		return ret
	}
	return o.Alternates
}

// GetAlternatesOk returns a tuple with the Alternates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Game) GetAlternatesOk() ([]string, bool) {
	if o == nil || o.Alternates == nil {
		return nil, false
	}
	return o.Alternates, true
}

// HasAlternates returns a boolean if a field has been set.
func (o *Game) HasAlternates() bool {
	if o != nil && o.Alternates != nil {
		return true
	}

	return false
}

// SetAlternates gets a reference to the given []string and assigns it to the Alternates field.
func (o *Game) SetAlternates(v []string) {
	o.Alternates = v
}

func (o Game) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GameTitle != nil {
		toSerialize["game_title"] = o.GameTitle
	}
	if o.ReleaseDate != nil {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Players != nil {
		toSerialize["players"] = o.Players
	}
	if o.Overview != nil {
		toSerialize["overview"] = o.Overview
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	if o.Coop != nil {
		toSerialize["coop"] = o.Coop
	}
	if o.Youtube != nil {
		toSerialize["youtube"] = o.Youtube
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.Processor != nil {
		toSerialize["processor"] = o.Processor
	}
	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}
	if o.Hdd != nil {
		toSerialize["hdd"] = o.Hdd
	}
	if o.Video != nil {
		toSerialize["video"] = o.Video
	}
	if o.Sound != nil {
		toSerialize["sound"] = o.Sound
	}
	if o.Developers != nil {
		toSerialize["developers"] = o.Developers
	}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if o.Publishers != nil {
		toSerialize["publishers"] = o.Publishers
	}
	if o.Alternates != nil {
		toSerialize["alternates"] = o.Alternates
	}
	return json.Marshal(toSerialize)
}

type NullableGame struct {
	value *Game
	isSet bool
}

func (v NullableGame) Get() *Game {
	return v.value
}

func (v *NullableGame) Set(val *Game) {
	v.value = val
	v.isSet = true
}

func (v NullableGame) IsSet() bool {
	return v.isSet
}

func (v *NullableGame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGame(val *Game) *NullableGame {
	return &NullableGame{value: val, isSet: true}
}

func (v NullableGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}