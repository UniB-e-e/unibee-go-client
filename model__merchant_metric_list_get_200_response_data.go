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

// checks if the MerchantMetricListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMetricListGet200ResponseData{}

// MerchantMetricListGet200ResponseData struct for MerchantMetricListGet200ResponseData
type MerchantMetricListGet200ResponseData struct {
	// MerchantMetrics
	MerchantMetrics []UnibeeApiBeanMerchantMetricSimplify `json:"merchantMetrics,omitempty"`
}

// NewMerchantMetricListGet200ResponseData instantiates a new MerchantMetricListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMetricListGet200ResponseData() *MerchantMetricListGet200ResponseData {
	this := MerchantMetricListGet200ResponseData{}
	return &this
}

// NewMerchantMetricListGet200ResponseDataWithDefaults instantiates a new MerchantMetricListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMetricListGet200ResponseDataWithDefaults() *MerchantMetricListGet200ResponseData {
	this := MerchantMetricListGet200ResponseData{}
	return &this
}

// GetMerchantMetrics returns the MerchantMetrics field value if set, zero value otherwise.
func (o *MerchantMetricListGet200ResponseData) GetMerchantMetrics() []UnibeeApiBeanMerchantMetricSimplify {
	if o == nil || IsNil(o.MerchantMetrics) {
		var ret []UnibeeApiBeanMerchantMetricSimplify
		return ret
	}
	return o.MerchantMetrics
}

// GetMerchantMetricsOk returns a tuple with the MerchantMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricListGet200ResponseData) GetMerchantMetricsOk() ([]UnibeeApiBeanMerchantMetricSimplify, bool) {
	if o == nil || IsNil(o.MerchantMetrics) {
		return nil, false
	}
	return o.MerchantMetrics, true
}

// HasMerchantMetrics returns a boolean if a field has been set.
func (o *MerchantMetricListGet200ResponseData) HasMerchantMetrics() bool {
	if o != nil && !IsNil(o.MerchantMetrics) {
		return true
	}

	return false
}

// SetMerchantMetrics gets a reference to the given []UnibeeApiBeanMerchantMetricSimplify and assigns it to the MerchantMetrics field.
func (o *MerchantMetricListGet200ResponseData) SetMerchantMetrics(v []UnibeeApiBeanMerchantMetricSimplify) {
	o.MerchantMetrics = v
}

func (o MerchantMetricListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMetricListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMetrics) {
		toSerialize["merchantMetrics"] = o.MerchantMetrics
	}
	return toSerialize, nil
}

type NullableMerchantMetricListGet200ResponseData struct {
	value *MerchantMetricListGet200ResponseData
	isSet bool
}

func (v NullableMerchantMetricListGet200ResponseData) Get() *MerchantMetricListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantMetricListGet200ResponseData) Set(val *MerchantMetricListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMetricListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMetricListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMetricListGet200ResponseData(val *MerchantMetricListGet200ResponseData) *NullableMerchantMetricListGet200ResponseData {
	return &NullableMerchantMetricListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMetricListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMetricListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


