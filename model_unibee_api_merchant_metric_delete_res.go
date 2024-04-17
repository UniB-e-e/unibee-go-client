/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404171839 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricDeleteRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricDeleteRes{}

// UnibeeApiMerchantMetricDeleteRes struct for UnibeeApiMerchantMetricDeleteRes
type UnibeeApiMerchantMetricDeleteRes struct {
	MerchantMetric *UnibeeApiBeanMerchantMetricSimplify `json:"merchantMetric,omitempty"`
}

// NewUnibeeApiMerchantMetricDeleteRes instantiates a new UnibeeApiMerchantMetricDeleteRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricDeleteRes() *UnibeeApiMerchantMetricDeleteRes {
	this := UnibeeApiMerchantMetricDeleteRes{}
	return &this
}

// NewUnibeeApiMerchantMetricDeleteResWithDefaults instantiates a new UnibeeApiMerchantMetricDeleteRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricDeleteResWithDefaults() *UnibeeApiMerchantMetricDeleteRes {
	this := UnibeeApiMerchantMetricDeleteRes{}
	return &this
}

// GetMerchantMetric returns the MerchantMetric field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricDeleteRes) GetMerchantMetric() UnibeeApiBeanMerchantMetricSimplify {
	if o == nil || IsNil(o.MerchantMetric) {
		var ret UnibeeApiBeanMerchantMetricSimplify
		return ret
	}
	return *o.MerchantMetric
}

// GetMerchantMetricOk returns a tuple with the MerchantMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeleteRes) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetricSimplify, bool) {
	if o == nil || IsNil(o.MerchantMetric) {
		return nil, false
	}
	return o.MerchantMetric, true
}

// HasMerchantMetric returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricDeleteRes) HasMerchantMetric() bool {
	if o != nil && !IsNil(o.MerchantMetric) {
		return true
	}

	return false
}

// SetMerchantMetric gets a reference to the given UnibeeApiBeanMerchantMetricSimplify and assigns it to the MerchantMetric field.
func (o *UnibeeApiMerchantMetricDeleteRes) SetMerchantMetric(v UnibeeApiBeanMerchantMetricSimplify) {
	o.MerchantMetric = &v
}

func (o UnibeeApiMerchantMetricDeleteRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricDeleteRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetric) {
		toSerialize["merchantMetric"] = o.MerchantMetric
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricDeleteRes struct {
	value *UnibeeApiMerchantMetricDeleteRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricDeleteRes) Get() *UnibeeApiMerchantMetricDeleteRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricDeleteRes) Set(val *UnibeeApiMerchantMetricDeleteRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricDeleteRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricDeleteRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricDeleteRes(val *UnibeeApiMerchantMetricDeleteRes) *NullableUnibeeApiMerchantMetricDeleteRes {
	return &NullableUnibeeApiMerchantMetricDeleteRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricDeleteRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricDeleteRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


