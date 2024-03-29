/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantInvoiceCancelReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceCancelReq{}

// UnibeeApiMerchantInvoiceCancelReq struct for UnibeeApiMerchantInvoiceCancelReq
type UnibeeApiMerchantInvoiceCancelReq struct {
	// InvoiceId
	InvoiceId *string `json:"invoiceId,omitempty"`
}

// NewUnibeeApiMerchantInvoiceCancelReq instantiates a new UnibeeApiMerchantInvoiceCancelReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceCancelReq() *UnibeeApiMerchantInvoiceCancelReq {
	this := UnibeeApiMerchantInvoiceCancelReq{}
	return &this
}

// NewUnibeeApiMerchantInvoiceCancelReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceCancelReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceCancelReqWithDefaults() *UnibeeApiMerchantInvoiceCancelReq {
	this := UnibeeApiMerchantInvoiceCancelReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceCancelReq) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceCancelReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceCancelReq) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiMerchantInvoiceCancelReq) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

func (o UnibeeApiMerchantInvoiceCancelReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceCancelReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceCancelReq struct {
	value *UnibeeApiMerchantInvoiceCancelReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceCancelReq) Get() *UnibeeApiMerchantInvoiceCancelReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceCancelReq) Set(val *UnibeeApiMerchantInvoiceCancelReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceCancelReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceCancelReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceCancelReq(val *UnibeeApiMerchantInvoiceCancelReq) *NullableUnibeeApiMerchantInvoiceCancelReq {
	return &NullableUnibeeApiMerchantInvoiceCancelReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceCancelReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceCancelReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


