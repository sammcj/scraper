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

// GenresAllOfData struct for GenresAllOfData
type GenresAllOfData struct {
	Count  int32            `json:"count"`
	Genres map[string]Genre `json:"genres"`
}

// NewGenresAllOfData instantiates a new GenresAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenresAllOfData(count int32, genres map[string]Genre) *GenresAllOfData {
	this := GenresAllOfData{}
	this.Count = count
	this.Genres = genres
	return &this
}

// NewGenresAllOfDataWithDefaults instantiates a new GenresAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenresAllOfDataWithDefaults() *GenresAllOfData {
	this := GenresAllOfData{}
	return &this
}

// GetCount returns the Count field value
func (o *GenresAllOfData) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GenresAllOfData) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GenresAllOfData) SetCount(v int32) {
	o.Count = v
}

// GetGenres returns the Genres field value
func (o *GenresAllOfData) GetGenres() map[string]Genre {
	if o == nil {
		var ret map[string]Genre
		return ret
	}

	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value
// and a boolean to check if the value has been set.
func (o *GenresAllOfData) GetGenresOk() (*map[string]Genre, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Genres, true
}

// SetGenres sets field value
func (o *GenresAllOfData) SetGenres(v map[string]Genre) {
	o.Genres = v
}

func (o GenresAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["genres"] = o.Genres
	}
	return json.Marshal(toSerialize)
}

type NullableGenresAllOfData struct {
	value *GenresAllOfData
	isSet bool
}

func (v NullableGenresAllOfData) Get() *GenresAllOfData {
	return v.value
}

func (v *NullableGenresAllOfData) Set(val *GenresAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableGenresAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableGenresAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenresAllOfData(val *GenresAllOfData) *NullableGenresAllOfData {
	return &NullableGenresAllOfData{value: val, isSet: true}
}

func (v NullableGenresAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenresAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}