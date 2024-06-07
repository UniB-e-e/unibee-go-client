/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406070109 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantGatewayListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGatewayListGet200ResponseData{}

// MerchantGatewayListGet200ResponseData struct for MerchantGatewayListGet200ResponseData
type MerchantGatewayListGet200ResponseData struct {
	// Payment Gateway Object List
	Gateways []UnibeeApiBeanGatewaySimplify `json:"gateways,omitempty"`
}

// NewMerchantGatewayListGet200ResponseData instantiates a new MerchantGatewayListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGatewayListGet200ResponseData() *MerchantGatewayListGet200ResponseData {
	this := MerchantGatewayListGet200ResponseData{}
	return &this
}

// NewMerchantGatewayListGet200ResponseDataWithDefaults instantiates a new MerchantGatewayListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGatewayListGet200ResponseDataWithDefaults() *MerchantGatewayListGet200ResponseData {
	this := MerchantGatewayListGet200ResponseData{}
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *MerchantGatewayListGet200ResponseData) GetGateways() []UnibeeApiBeanGatewaySimplify {
	if o == nil || IsNil(o.Gateways) {
		var ret []UnibeeApiBeanGatewaySimplify
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGatewayListGet200ResponseData) GetGatewaysOk() ([]UnibeeApiBeanGatewaySimplify, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *MerchantGatewayListGet200ResponseData) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []UnibeeApiBeanGatewaySimplify and assigns it to the Gateways field.
func (o *MerchantGatewayListGet200ResponseData) SetGateways(v []UnibeeApiBeanGatewaySimplify) {
	o.Gateways = v
}

func (o MerchantGatewayListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGatewayListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	return toSerialize, nil
}

type NullableMerchantGatewayListGet200ResponseData struct {
	value *MerchantGatewayListGet200ResponseData
	isSet bool
}

func (v NullableMerchantGatewayListGet200ResponseData) Get() *MerchantGatewayListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantGatewayListGet200ResponseData) Set(val *MerchantGatewayListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGatewayListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGatewayListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGatewayListGet200ResponseData(val *MerchantGatewayListGet200ResponseData) *NullableMerchantGatewayListGet200ResponseData {
	return &NullableMerchantGatewayListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGatewayListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGatewayListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


