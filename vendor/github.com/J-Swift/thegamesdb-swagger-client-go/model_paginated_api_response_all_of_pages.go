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

// PaginatedApiResponseAllOfPages struct for PaginatedApiResponseAllOfPages
type PaginatedApiResponseAllOfPages struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}

// NewPaginatedApiResponseAllOfPages instantiates a new PaginatedApiResponseAllOfPages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiResponseAllOfPages(previous string, current string, next string) *PaginatedApiResponseAllOfPages {
	this := PaginatedApiResponseAllOfPages{}
	this.Previous = previous
	this.Current = current
	this.Next = next
	return &this
}

// NewPaginatedApiResponseAllOfPagesWithDefaults instantiates a new PaginatedApiResponseAllOfPages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiResponseAllOfPagesWithDefaults() *PaginatedApiResponseAllOfPages {
	this := PaginatedApiResponseAllOfPages{}
	return &this
}

// GetPrevious returns the Previous field value
func (o *PaginatedApiResponseAllOfPages) GetPrevious() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
func (o *PaginatedApiResponseAllOfPages) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Previous, true
}

// SetPrevious sets field value
func (o *PaginatedApiResponseAllOfPages) SetPrevious(v string) {
	o.Previous = v
}

// GetCurrent returns the Current field value
func (o *PaginatedApiResponseAllOfPages) GetCurrent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *PaginatedApiResponseAllOfPages) GetCurrentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *PaginatedApiResponseAllOfPages) SetCurrent(v string) {
	o.Current = v
}

// GetNext returns the Next field value
func (o *PaginatedApiResponseAllOfPages) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *PaginatedApiResponseAllOfPages) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *PaginatedApiResponseAllOfPages) SetNext(v string) {
	o.Next = v
}

func (o PaginatedApiResponseAllOfPages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["previous"] = o.Previous
	}
	if true {
		toSerialize["current"] = o.Current
	}
	if true {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedApiResponseAllOfPages struct {
	value *PaginatedApiResponseAllOfPages
	isSet bool
}

func (v NullablePaginatedApiResponseAllOfPages) Get() *PaginatedApiResponseAllOfPages {
	return v.value
}

func (v *NullablePaginatedApiResponseAllOfPages) Set(val *PaginatedApiResponseAllOfPages) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedApiResponseAllOfPages) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedApiResponseAllOfPages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedApiResponseAllOfPages(val *PaginatedApiResponseAllOfPages) *NullablePaginatedApiResponseAllOfPages {
	return &NullablePaginatedApiResponseAllOfPages{value: val, isSet: true}
}

func (v NullablePaginatedApiResponseAllOfPages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedApiResponseAllOfPages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}