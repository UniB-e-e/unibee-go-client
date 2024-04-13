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

// checks if the UnibeeApiMerchantMemberListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberListRes{}

// UnibeeApiMerchantMemberListRes struct for UnibeeApiMerchantMemberListRes
type UnibeeApiMerchantMemberListRes struct {
	// Merchant Members
	MerchantMembers []UnibeeApiBeanMerchantMemberSimplify `json:"merchantMembers,omitempty"`
}

// NewUnibeeApiMerchantMemberListRes instantiates a new UnibeeApiMerchantMemberListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberListRes() *UnibeeApiMerchantMemberListRes {
	this := UnibeeApiMerchantMemberListRes{}
	return &this
}

// NewUnibeeApiMerchantMemberListResWithDefaults instantiates a new UnibeeApiMerchantMemberListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberListResWithDefaults() *UnibeeApiMerchantMemberListRes {
	this := UnibeeApiMerchantMemberListRes{}
	return &this
}

// GetMerchantMembers returns the MerchantMembers field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberListRes) GetMerchantMembers() []UnibeeApiBeanMerchantMemberSimplify {
	if o == nil || IsNil(o.MerchantMembers) {
		var ret []UnibeeApiBeanMerchantMemberSimplify
		return ret
	}
	return o.MerchantMembers
}

// GetMerchantMembersOk returns a tuple with the MerchantMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberListRes) GetMerchantMembersOk() ([]UnibeeApiBeanMerchantMemberSimplify, bool) {
	if o == nil || IsNil(o.MerchantMembers) {
		return nil, false
	}
	return o.MerchantMembers, true
}

// HasMerchantMembers returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberListRes) HasMerchantMembers() bool {
	if o != nil && !IsNil(o.MerchantMembers) {
		return true
	}

	return false
}

// SetMerchantMembers gets a reference to the given []UnibeeApiBeanMerchantMemberSimplify and assigns it to the MerchantMembers field.
func (o *UnibeeApiMerchantMemberListRes) SetMerchantMembers(v []UnibeeApiBeanMerchantMemberSimplify) {
	o.MerchantMembers = v
}

func (o UnibeeApiMerchantMemberListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMembers) {
		toSerialize["merchantMembers"] = o.MerchantMembers
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberListRes struct {
	value *UnibeeApiMerchantMemberListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberListRes) Get() *UnibeeApiMerchantMemberListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberListRes) Set(val *UnibeeApiMerchantMemberListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberListRes(val *UnibeeApiMerchantMemberListRes) *NullableUnibeeApiMerchantMemberListRes {
	return &NullableUnibeeApiMerchantMemberListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


