/*
OpenAPI UniBee

This is UniBee api server

API version: buildtime:202404131246 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMemberProfileRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberProfileRes{}

// UnibeeApiMerchantMemberProfileRes struct for UnibeeApiMerchantMemberProfileRes
type UnibeeApiMerchantMemberProfileRes struct {
	MerchantMember *UnibeeApiBeanMerchantMemberSimplify `json:"merchantMember,omitempty"`
}

// NewUnibeeApiMerchantMemberProfileRes instantiates a new UnibeeApiMerchantMemberProfileRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberProfileRes() *UnibeeApiMerchantMemberProfileRes {
	this := UnibeeApiMerchantMemberProfileRes{}
	return &this
}

// NewUnibeeApiMerchantMemberProfileResWithDefaults instantiates a new UnibeeApiMerchantMemberProfileRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberProfileResWithDefaults() *UnibeeApiMerchantMemberProfileRes {
	this := UnibeeApiMerchantMemberProfileRes{}
	return &this
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberProfileRes) GetMerchantMember() UnibeeApiBeanMerchantMemberSimplify {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanMerchantMemberSimplify
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberProfileRes) GetMerchantMemberOk() (*UnibeeApiBeanMerchantMemberSimplify, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberProfileRes) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanMerchantMemberSimplify and assigns it to the MerchantMember field.
func (o *UnibeeApiMerchantMemberProfileRes) SetMerchantMember(v UnibeeApiBeanMerchantMemberSimplify) {
	o.MerchantMember = &v
}

func (o UnibeeApiMerchantMemberProfileRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberProfileRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberProfileRes struct {
	value *UnibeeApiMerchantMemberProfileRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberProfileRes) Get() *UnibeeApiMerchantMemberProfileRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberProfileRes) Set(val *UnibeeApiMerchantMemberProfileRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberProfileRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberProfileRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberProfileRes(val *UnibeeApiMerchantMemberProfileRes) *NullableUnibeeApiMerchantMemberProfileRes {
	return &NullableUnibeeApiMerchantMemberProfileRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberProfileRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberProfileRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


