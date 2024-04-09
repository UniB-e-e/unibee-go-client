/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: buildtime:202404090336 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantInvoiceMarkRefundRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceMarkRefundRes{}

// UnibeeApiMerchantInvoiceMarkRefundRes struct for UnibeeApiMerchantInvoiceMarkRefundRes
type UnibeeApiMerchantInvoiceMarkRefundRes struct {
	Refund *UnibeeApiBeanRefundSimplify `json:"refund,omitempty"`
}

// NewUnibeeApiMerchantInvoiceMarkRefundRes instantiates a new UnibeeApiMerchantInvoiceMarkRefundRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceMarkRefundRes() *UnibeeApiMerchantInvoiceMarkRefundRes {
	this := UnibeeApiMerchantInvoiceMarkRefundRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceMarkRefundResWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkRefundRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceMarkRefundResWithDefaults() *UnibeeApiMerchantInvoiceMarkRefundRes {
	this := UnibeeApiMerchantInvoiceMarkRefundRes{}
	return &this
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceMarkRefundRes) GetRefund() UnibeeApiBeanRefundSimplify {
	if o == nil || IsNil(o.Refund) {
		var ret UnibeeApiBeanRefundSimplify
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundRes) GetRefundOk() (*UnibeeApiBeanRefundSimplify, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundRes) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given UnibeeApiBeanRefundSimplify and assigns it to the Refund field.
func (o *UnibeeApiMerchantInvoiceMarkRefundRes) SetRefund(v UnibeeApiBeanRefundSimplify) {
	o.Refund = &v
}

func (o UnibeeApiMerchantInvoiceMarkRefundRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceMarkRefundRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceMarkRefundRes struct {
	value *UnibeeApiMerchantInvoiceMarkRefundRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundRes) Get() *UnibeeApiMerchantInvoiceMarkRefundRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundRes) Set(val *UnibeeApiMerchantInvoiceMarkRefundRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceMarkRefundRes(val *UnibeeApiMerchantInvoiceMarkRefundRes) *NullableUnibeeApiMerchantInvoiceMarkRefundRes {
	return &NullableUnibeeApiMerchantInvoiceMarkRefundRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


