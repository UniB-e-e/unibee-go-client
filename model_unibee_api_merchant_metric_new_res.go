/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404191455 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricNewRes{}

// UnibeeApiMerchantMetricNewRes struct for UnibeeApiMerchantMetricNewRes
type UnibeeApiMerchantMetricNewRes struct {
	MerchantMetric *UnibeeApiBeanMerchantMetricSimplify `json:"merchantMetric,omitempty"`
}

// NewUnibeeApiMerchantMetricNewRes instantiates a new UnibeeApiMerchantMetricNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricNewRes() *UnibeeApiMerchantMetricNewRes {
	this := UnibeeApiMerchantMetricNewRes{}
	return &this
}

// NewUnibeeApiMerchantMetricNewResWithDefaults instantiates a new UnibeeApiMerchantMetricNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricNewResWithDefaults() *UnibeeApiMerchantMetricNewRes {
	this := UnibeeApiMerchantMetricNewRes{}
	return &this
}

// GetMerchantMetric returns the MerchantMetric field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewRes) GetMerchantMetric() UnibeeApiBeanMerchantMetricSimplify {
	if o == nil || IsNil(o.MerchantMetric) {
		var ret UnibeeApiBeanMerchantMetricSimplify
		return ret
	}
	return *o.MerchantMetric
}

// GetMerchantMetricOk returns a tuple with the MerchantMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewRes) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetricSimplify, bool) {
	if o == nil || IsNil(o.MerchantMetric) {
		return nil, false
	}
	return o.MerchantMetric, true
}

// HasMerchantMetric returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewRes) HasMerchantMetric() bool {
	if o != nil && !IsNil(o.MerchantMetric) {
		return true
	}

	return false
}

// SetMerchantMetric gets a reference to the given UnibeeApiBeanMerchantMetricSimplify and assigns it to the MerchantMetric field.
func (o *UnibeeApiMerchantMetricNewRes) SetMerchantMetric(v UnibeeApiBeanMerchantMetricSimplify) {
	o.MerchantMetric = &v
}

func (o UnibeeApiMerchantMetricNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetric) {
		toSerialize["merchantMetric"] = o.MerchantMetric
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricNewRes struct {
	value *UnibeeApiMerchantMetricNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricNewRes) Get() *UnibeeApiMerchantMetricNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricNewRes) Set(val *UnibeeApiMerchantMetricNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricNewRes(val *UnibeeApiMerchantMetricNewRes) *NullableUnibeeApiMerchantMetricNewRes {
	return &NullableUnibeeApiMerchantMetricNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


