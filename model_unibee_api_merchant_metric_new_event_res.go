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

// checks if the UnibeeApiMerchantMetricNewEventRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricNewEventRes{}

// UnibeeApiMerchantMetricNewEventRes struct for UnibeeApiMerchantMetricNewEventRes
type UnibeeApiMerchantMetricNewEventRes struct {
	MerchantMetricEvent *UnibeeApiBeanMerchantMetricEventSimplify `json:"merchantMetricEvent,omitempty"`
}

// NewUnibeeApiMerchantMetricNewEventRes instantiates a new UnibeeApiMerchantMetricNewEventRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricNewEventRes() *UnibeeApiMerchantMetricNewEventRes {
	this := UnibeeApiMerchantMetricNewEventRes{}
	return &this
}

// NewUnibeeApiMerchantMetricNewEventResWithDefaults instantiates a new UnibeeApiMerchantMetricNewEventRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricNewEventResWithDefaults() *UnibeeApiMerchantMetricNewEventRes {
	this := UnibeeApiMerchantMetricNewEventRes{}
	return &this
}

// GetMerchantMetricEvent returns the MerchantMetricEvent field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewEventRes) GetMerchantMetricEvent() UnibeeApiBeanMerchantMetricEventSimplify {
	if o == nil || IsNil(o.MerchantMetricEvent) {
		var ret UnibeeApiBeanMerchantMetricEventSimplify
		return ret
	}
	return *o.MerchantMetricEvent
}

// GetMerchantMetricEventOk returns a tuple with the MerchantMetricEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewEventRes) GetMerchantMetricEventOk() (*UnibeeApiBeanMerchantMetricEventSimplify, bool) {
	if o == nil || IsNil(o.MerchantMetricEvent) {
		return nil, false
	}
	return o.MerchantMetricEvent, true
}

// HasMerchantMetricEvent returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewEventRes) HasMerchantMetricEvent() bool {
	if o != nil && !IsNil(o.MerchantMetricEvent) {
		return true
	}

	return false
}

// SetMerchantMetricEvent gets a reference to the given UnibeeApiBeanMerchantMetricEventSimplify and assigns it to the MerchantMetricEvent field.
func (o *UnibeeApiMerchantMetricNewEventRes) SetMerchantMetricEvent(v UnibeeApiBeanMerchantMetricEventSimplify) {
	o.MerchantMetricEvent = &v
}

func (o UnibeeApiMerchantMetricNewEventRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricNewEventRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetricEvent) {
		toSerialize["merchantMetricEvent"] = o.MerchantMetricEvent
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricNewEventRes struct {
	value *UnibeeApiMerchantMetricNewEventRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricNewEventRes) Get() *UnibeeApiMerchantMetricNewEventRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricNewEventRes) Set(val *UnibeeApiMerchantMetricNewEventRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricNewEventRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricNewEventRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricNewEventRes(val *UnibeeApiMerchantMetricNewEventRes) *NullableUnibeeApiMerchantMetricNewEventRes {
	return &NullableUnibeeApiMerchantMetricNewEventRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricNewEventRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricNewEventRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


