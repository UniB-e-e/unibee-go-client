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

// checks if the MerchantAuthSsoRegisterVerifyPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantAuthSsoRegisterVerifyPost200ResponseData{}

// MerchantAuthSsoRegisterVerifyPost200ResponseData struct for MerchantAuthSsoRegisterVerifyPost200ResponseData
type MerchantAuthSsoRegisterVerifyPost200ResponseData struct {
	MerchantMember *UnibeeApiBeanMerchantMemberSimplify `json:"merchantMember,omitempty"`
}

// NewMerchantAuthSsoRegisterVerifyPost200ResponseData instantiates a new MerchantAuthSsoRegisterVerifyPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAuthSsoRegisterVerifyPost200ResponseData() *MerchantAuthSsoRegisterVerifyPost200ResponseData {
	this := MerchantAuthSsoRegisterVerifyPost200ResponseData{}
	return &this
}

// NewMerchantAuthSsoRegisterVerifyPost200ResponseDataWithDefaults instantiates a new MerchantAuthSsoRegisterVerifyPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAuthSsoRegisterVerifyPost200ResponseDataWithDefaults() *MerchantAuthSsoRegisterVerifyPost200ResponseData {
	this := MerchantAuthSsoRegisterVerifyPost200ResponseData{}
	return &this
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *MerchantAuthSsoRegisterVerifyPost200ResponseData) GetMerchantMember() UnibeeApiBeanMerchantMemberSimplify {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanMerchantMemberSimplify
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAuthSsoRegisterVerifyPost200ResponseData) GetMerchantMemberOk() (*UnibeeApiBeanMerchantMemberSimplify, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *MerchantAuthSsoRegisterVerifyPost200ResponseData) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanMerchantMemberSimplify and assigns it to the MerchantMember field.
func (o *MerchantAuthSsoRegisterVerifyPost200ResponseData) SetMerchantMember(v UnibeeApiBeanMerchantMemberSimplify) {
	o.MerchantMember = &v
}

func (o MerchantAuthSsoRegisterVerifyPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantAuthSsoRegisterVerifyPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	return toSerialize, nil
}

type NullableMerchantAuthSsoRegisterVerifyPost200ResponseData struct {
	value *MerchantAuthSsoRegisterVerifyPost200ResponseData
	isSet bool
}

func (v NullableMerchantAuthSsoRegisterVerifyPost200ResponseData) Get() *MerchantAuthSsoRegisterVerifyPost200ResponseData {
	return v.value
}

func (v *NullableMerchantAuthSsoRegisterVerifyPost200ResponseData) Set(val *MerchantAuthSsoRegisterVerifyPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAuthSsoRegisterVerifyPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAuthSsoRegisterVerifyPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAuthSsoRegisterVerifyPost200ResponseData(val *MerchantAuthSsoRegisterVerifyPost200ResponseData) *NullableMerchantAuthSsoRegisterVerifyPost200ResponseData {
	return &NullableMerchantAuthSsoRegisterVerifyPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantAuthSsoRegisterVerifyPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAuthSsoRegisterVerifyPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


