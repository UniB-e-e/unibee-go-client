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

// checks if the MerchantInvoiceMarkRefundPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantInvoiceMarkRefundPost200ResponseData{}

// MerchantInvoiceMarkRefundPost200ResponseData struct for MerchantInvoiceMarkRefundPost200ResponseData
type MerchantInvoiceMarkRefundPost200ResponseData struct {
	Refund *UnibeeApiBeanRefundSimplify `json:"refund,omitempty"`
}

// NewMerchantInvoiceMarkRefundPost200ResponseData instantiates a new MerchantInvoiceMarkRefundPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantInvoiceMarkRefundPost200ResponseData() *MerchantInvoiceMarkRefundPost200ResponseData {
	this := MerchantInvoiceMarkRefundPost200ResponseData{}
	return &this
}

// NewMerchantInvoiceMarkRefundPost200ResponseDataWithDefaults instantiates a new MerchantInvoiceMarkRefundPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantInvoiceMarkRefundPost200ResponseDataWithDefaults() *MerchantInvoiceMarkRefundPost200ResponseData {
	this := MerchantInvoiceMarkRefundPost200ResponseData{}
	return &this
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *MerchantInvoiceMarkRefundPost200ResponseData) GetRefund() UnibeeApiBeanRefundSimplify {
	if o == nil || IsNil(o.Refund) {
		var ret UnibeeApiBeanRefundSimplify
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantInvoiceMarkRefundPost200ResponseData) GetRefundOk() (*UnibeeApiBeanRefundSimplify, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *MerchantInvoiceMarkRefundPost200ResponseData) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given UnibeeApiBeanRefundSimplify and assigns it to the Refund field.
func (o *MerchantInvoiceMarkRefundPost200ResponseData) SetRefund(v UnibeeApiBeanRefundSimplify) {
	o.Refund = &v
}

func (o MerchantInvoiceMarkRefundPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantInvoiceMarkRefundPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	return toSerialize, nil
}

type NullableMerchantInvoiceMarkRefundPost200ResponseData struct {
	value *MerchantInvoiceMarkRefundPost200ResponseData
	isSet bool
}

func (v NullableMerchantInvoiceMarkRefundPost200ResponseData) Get() *MerchantInvoiceMarkRefundPost200ResponseData {
	return v.value
}

func (v *NullableMerchantInvoiceMarkRefundPost200ResponseData) Set(val *MerchantInvoiceMarkRefundPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantInvoiceMarkRefundPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantInvoiceMarkRefundPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantInvoiceMarkRefundPost200ResponseData(val *MerchantInvoiceMarkRefundPost200ResponseData) *NullableMerchantInvoiceMarkRefundPost200ResponseData {
	return &NullableMerchantInvoiceMarkRefundPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantInvoiceMarkRefundPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantInvoiceMarkRefundPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

