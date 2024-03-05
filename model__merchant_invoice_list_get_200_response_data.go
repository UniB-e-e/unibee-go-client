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

// checks if the MerchantInvoiceListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantInvoiceListGet200ResponseData{}

// MerchantInvoiceListGet200ResponseData struct for MerchantInvoiceListGet200ResponseData
type MerchantInvoiceListGet200ResponseData struct {
	// invoice Detail List
	Invoices []UnibeeInternalLogicGatewayRoInvoiceDetailRo `json:"invoices,omitempty"`
}

// NewMerchantInvoiceListGet200ResponseData instantiates a new MerchantInvoiceListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantInvoiceListGet200ResponseData() *MerchantInvoiceListGet200ResponseData {
	this := MerchantInvoiceListGet200ResponseData{}
	return &this
}

// NewMerchantInvoiceListGet200ResponseDataWithDefaults instantiates a new MerchantInvoiceListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantInvoiceListGet200ResponseDataWithDefaults() *MerchantInvoiceListGet200ResponseData {
	this := MerchantInvoiceListGet200ResponseData{}
	return &this
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *MerchantInvoiceListGet200ResponseData) GetInvoices() []UnibeeInternalLogicGatewayRoInvoiceDetailRo {
	if o == nil || IsNil(o.Invoices) {
		var ret []UnibeeInternalLogicGatewayRoInvoiceDetailRo
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantInvoiceListGet200ResponseData) GetInvoicesOk() ([]UnibeeInternalLogicGatewayRoInvoiceDetailRo, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *MerchantInvoiceListGet200ResponseData) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []UnibeeInternalLogicGatewayRoInvoiceDetailRo and assigns it to the Invoices field.
func (o *MerchantInvoiceListGet200ResponseData) SetInvoices(v []UnibeeInternalLogicGatewayRoInvoiceDetailRo) {
	o.Invoices = v
}

func (o MerchantInvoiceListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantInvoiceListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	return toSerialize, nil
}

type NullableMerchantInvoiceListGet200ResponseData struct {
	value *MerchantInvoiceListGet200ResponseData
	isSet bool
}

func (v NullableMerchantInvoiceListGet200ResponseData) Get() *MerchantInvoiceListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantInvoiceListGet200ResponseData) Set(val *MerchantInvoiceListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantInvoiceListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantInvoiceListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantInvoiceListGet200ResponseData(val *MerchantInvoiceListGet200ResponseData) *NullableMerchantInvoiceListGet200ResponseData {
	return &NullableMerchantInvoiceListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantInvoiceListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantInvoiceListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

