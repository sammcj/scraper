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

// ImageBaseUrlMeta struct for ImageBaseUrlMeta
type ImageBaseUrlMeta struct {
	Original           string `json:"original"`
	Small              string `json:"small"`
	Thumb              string `json:"thumb"`
	CroppedCenterThumb string `json:"cropped_center_thumb"`
	Medium             string `json:"medium"`
	Large              string `json:"large"`
}

// NewImageBaseUrlMeta instantiates a new ImageBaseUrlMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBaseUrlMeta(original string, small string, thumb string, croppedCenterThumb string, medium string, large string) *ImageBaseUrlMeta {
	this := ImageBaseUrlMeta{}
	this.Original = original
	this.Small = small
	this.Thumb = thumb
	this.CroppedCenterThumb = croppedCenterThumb
	this.Medium = medium
	this.Large = large
	return &this
}

// NewImageBaseUrlMetaWithDefaults instantiates a new ImageBaseUrlMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBaseUrlMetaWithDefaults() *ImageBaseUrlMeta {
	this := ImageBaseUrlMeta{}
	return &this
}

// GetOriginal returns the Original field value
func (o *ImageBaseUrlMeta) GetOriginal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Original
}

// GetOriginalOk returns a tuple with the Original field value
// and a boolean to check if the value has been set.
func (o *ImageBaseUrlMeta) GetOriginalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Original, true
}

// SetOriginal sets field value
func (o *ImageBaseUrlMeta) SetOriginal(v string) {
	o.Original = v
}

// GetSmall returns the Small field value
func (o *ImageBaseUrlMeta) GetSmall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Small
}

// GetSmallOk returns a tuple with the Small field value
// and a boolean to check if the value has been set.
func (o *ImageBaseUrlMeta) GetSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Small, true
}

// SetSmall sets field value
func (o *ImageBaseUrlMeta) SetSmall(v string) {
	o.Small = v
}

// GetThumb returns the Thumb field value
func (o *ImageBaseUrlMeta) GetThumb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Thumb
}

// GetThumbOk returns a tuple with the Thumb field value
// and a boolean to check if the value has been set.
func (o *ImageBaseUrlMeta) GetThumbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thumb, true
}

// SetThumb sets field value
func (o *ImageBaseUrlMeta) SetThumb(v string) {
	o.Thumb = v
}

// GetCroppedCenterThumb returns the CroppedCenterThumb field value
func (o *ImageBaseUrlMeta) GetCroppedCenterThumb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CroppedCenterThumb
}

// GetCroppedCenterThumbOk returns a tuple with the CroppedCenterThumb field value
// and a boolean to check if the value has been set.
func (o *ImageBaseUrlMeta) GetCroppedCenterThumbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CroppedCenterThumb, true
}

// SetCroppedCenterThumb sets field value
func (o *ImageBaseUrlMeta) SetCroppedCenterThumb(v string) {
	o.CroppedCenterThumb = v
}

// GetMedium returns the Medium field value
func (o *ImageBaseUrlMeta) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *ImageBaseUrlMeta) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Medium, true
}

// SetMedium sets field value
func (o *ImageBaseUrlMeta) SetMedium(v string) {
	o.Medium = v
}

// GetLarge returns the Large field value
func (o *ImageBaseUrlMeta) GetLarge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Large
}

// GetLargeOk returns a tuple with the Large field value
// and a boolean to check if the value has been set.
func (o *ImageBaseUrlMeta) GetLargeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Large, true
}

// SetLarge sets field value
func (o *ImageBaseUrlMeta) SetLarge(v string) {
	o.Large = v
}

func (o ImageBaseUrlMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["original"] = o.Original
	}
	if true {
		toSerialize["small"] = o.Small
	}
	if true {
		toSerialize["thumb"] = o.Thumb
	}
	if true {
		toSerialize["cropped_center_thumb"] = o.CroppedCenterThumb
	}
	if true {
		toSerialize["medium"] = o.Medium
	}
	if true {
		toSerialize["large"] = o.Large
	}
	return json.Marshal(toSerialize)
}

type NullableImageBaseUrlMeta struct {
	value *ImageBaseUrlMeta
	isSet bool
}

func (v NullableImageBaseUrlMeta) Get() *ImageBaseUrlMeta {
	return v.value
}

func (v *NullableImageBaseUrlMeta) Set(val *ImageBaseUrlMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBaseUrlMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBaseUrlMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBaseUrlMeta(val *ImageBaseUrlMeta) *NullableImageBaseUrlMeta {
	return &NullableImageBaseUrlMeta{value: val, isSet: true}
}

func (v NullableImageBaseUrlMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBaseUrlMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}