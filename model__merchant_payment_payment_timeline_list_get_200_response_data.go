/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MerchantPaymentPaymentTimelineListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentPaymentTimelineListGet200ResponseData{}

// MerchantPaymentPaymentTimelineListGet200ResponseData struct for MerchantPaymentPaymentTimelineListGet200ResponseData
type MerchantPaymentPaymentTimelineListGet200ResponseData struct {
	// PaymentTimeLines
	PaymentTimeLines []UnibeeInternalModelEntityOverseaPayPaymentTimeline `json:"paymentTimeLines,omitempty"`
}

// NewMerchantPaymentPaymentTimelineListGet200ResponseData instantiates a new MerchantPaymentPaymentTimelineListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentPaymentTimelineListGet200ResponseData() *MerchantPaymentPaymentTimelineListGet200ResponseData {
	this := MerchantPaymentPaymentTimelineListGet200ResponseData{}
	return &this
}

// NewMerchantPaymentPaymentTimelineListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentPaymentTimelineListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentPaymentTimelineListGet200ResponseDataWithDefaults() *MerchantPaymentPaymentTimelineListGet200ResponseData {
	this := MerchantPaymentPaymentTimelineListGet200ResponseData{}
	return &this
}

// GetPaymentTimeLines returns the PaymentTimeLines field value if set, zero value otherwise.
func (o *MerchantPaymentPaymentTimelineListGet200ResponseData) GetPaymentTimeLines() []UnibeeInternalModelEntityOverseaPayPaymentTimeline {
	if o == nil || IsNil(o.PaymentTimeLines) {
		var ret []UnibeeInternalModelEntityOverseaPayPaymentTimeline
		return ret
	}
	return o.PaymentTimeLines
}

// GetPaymentTimeLinesOk returns a tuple with the PaymentTimeLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentPaymentTimelineListGet200ResponseData) GetPaymentTimeLinesOk() ([]UnibeeInternalModelEntityOverseaPayPaymentTimeline, bool) {
	if o == nil || IsNil(o.PaymentTimeLines) {
		return nil, false
	}
	return o.PaymentTimeLines, true
}

// HasPaymentTimeLines returns a boolean if a field has been set.
func (o *MerchantPaymentPaymentTimelineListGet200ResponseData) HasPaymentTimeLines() bool {
	if o != nil && !IsNil(o.PaymentTimeLines) {
		return true
	}

	return false
}

// SetPaymentTimeLines gets a reference to the given []UnibeeInternalModelEntityOverseaPayPaymentTimeline and assigns it to the PaymentTimeLines field.
func (o *MerchantPaymentPaymentTimelineListGet200ResponseData) SetPaymentTimeLines(v []UnibeeInternalModelEntityOverseaPayPaymentTimeline) {
	o.PaymentTimeLines = v
}

func (o MerchantPaymentPaymentTimelineListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentPaymentTimelineListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentTimeLines) {
		toSerialize["paymentTimeLines"] = o.PaymentTimeLines
	}
	return toSerialize, nil
}

type NullableMerchantPaymentPaymentTimelineListGet200ResponseData struct {
	value *MerchantPaymentPaymentTimelineListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentPaymentTimelineListGet200ResponseData) Get() *MerchantPaymentPaymentTimelineListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentPaymentTimelineListGet200ResponseData) Set(val *MerchantPaymentPaymentTimelineListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentPaymentTimelineListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentPaymentTimelineListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentPaymentTimelineListGet200ResponseData(val *MerchantPaymentPaymentTimelineListGet200ResponseData) *NullableMerchantPaymentPaymentTimelineListGet200ResponseData {
	return &NullableMerchantPaymentPaymentTimelineListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentPaymentTimelineListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentPaymentTimelineListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

