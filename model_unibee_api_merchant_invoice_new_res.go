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

// checks if the UnibeeApiMerchantInvoiceNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceNewRes{}

// UnibeeApiMerchantInvoiceNewRes struct for UnibeeApiMerchantInvoiceNewRes
type UnibeeApiMerchantInvoiceNewRes struct {
	Invoice *UnibeeInternalLogicGatewayRoInvoiceDetailRo `json:"invoice,omitempty"`
}

// NewUnibeeApiMerchantInvoiceNewRes instantiates a new UnibeeApiMerchantInvoiceNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceNewRes() *UnibeeApiMerchantInvoiceNewRes {
	this := UnibeeApiMerchantInvoiceNewRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceNewResWithDefaults instantiates a new UnibeeApiMerchantInvoiceNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceNewResWithDefaults() *UnibeeApiMerchantInvoiceNewRes {
	this := UnibeeApiMerchantInvoiceNewRes{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceNewRes) GetInvoice() UnibeeInternalLogicGatewayRoInvoiceDetailRo {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeInternalLogicGatewayRoInvoiceDetailRo
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceNewRes) GetInvoiceOk() (*UnibeeInternalLogicGatewayRoInvoiceDetailRo, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceNewRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeInternalLogicGatewayRoInvoiceDetailRo and assigns it to the Invoice field.
func (o *UnibeeApiMerchantInvoiceNewRes) SetInvoice(v UnibeeInternalLogicGatewayRoInvoiceDetailRo) {
	o.Invoice = &v
}

func (o UnibeeApiMerchantInvoiceNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceNewRes struct {
	value *UnibeeApiMerchantInvoiceNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceNewRes) Get() *UnibeeApiMerchantInvoiceNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceNewRes) Set(val *UnibeeApiMerchantInvoiceNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceNewRes(val *UnibeeApiMerchantInvoiceNewRes) *NullableUnibeeApiMerchantInvoiceNewRes {
	return &NullableUnibeeApiMerchantInvoiceNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

