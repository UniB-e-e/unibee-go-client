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

// checks if the UnibeeApiMerchantPaymentMethodListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentMethodListRes{}

// UnibeeApiMerchantPaymentMethodListRes struct for UnibeeApiMerchantPaymentMethodListRes
type UnibeeApiMerchantPaymentMethodListRes struct {
	// Payment Method Object List
	MethodList []UnibeeApiBeanPaymentMethod `json:"methodList,omitempty"`
}

// NewUnibeeApiMerchantPaymentMethodListRes instantiates a new UnibeeApiMerchantPaymentMethodListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentMethodListRes() *UnibeeApiMerchantPaymentMethodListRes {
	this := UnibeeApiMerchantPaymentMethodListRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentMethodListResWithDefaults instantiates a new UnibeeApiMerchantPaymentMethodListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentMethodListResWithDefaults() *UnibeeApiMerchantPaymentMethodListRes {
	this := UnibeeApiMerchantPaymentMethodListRes{}
	return &this
}

// GetMethodList returns the MethodList field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentMethodListRes) GetMethodList() []UnibeeApiBeanPaymentMethod {
	if o == nil || IsNil(o.MethodList) {
		var ret []UnibeeApiBeanPaymentMethod
		return ret
	}
	return o.MethodList
}

// GetMethodListOk returns a tuple with the MethodList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentMethodListRes) GetMethodListOk() ([]UnibeeApiBeanPaymentMethod, bool) {
	if o == nil || IsNil(o.MethodList) {
		return nil, false
	}
	return o.MethodList, true
}

// HasMethodList returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentMethodListRes) HasMethodList() bool {
	if o != nil && !IsNil(o.MethodList) {
		return true
	}

	return false
}

// SetMethodList gets a reference to the given []UnibeeApiBeanPaymentMethod and assigns it to the MethodList field.
func (o *UnibeeApiMerchantPaymentMethodListRes) SetMethodList(v []UnibeeApiBeanPaymentMethod) {
	o.MethodList = v
}

func (o UnibeeApiMerchantPaymentMethodListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentMethodListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MethodList) {
		toSerialize["methodList"] = o.MethodList
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentMethodListRes struct {
	value *UnibeeApiMerchantPaymentMethodListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentMethodListRes) Get() *UnibeeApiMerchantPaymentMethodListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentMethodListRes) Set(val *UnibeeApiMerchantPaymentMethodListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentMethodListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentMethodListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentMethodListRes(val *UnibeeApiMerchantPaymentMethodListRes) *NullableUnibeeApiMerchantPaymentMethodListRes {
	return &NullableUnibeeApiMerchantPaymentMethodListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentMethodListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentMethodListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


