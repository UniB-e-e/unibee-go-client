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

// checks if the UnibeeApiMerchantRoleListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantRoleListRes{}

// UnibeeApiMerchantRoleListRes struct for UnibeeApiMerchantRoleListRes
type UnibeeApiMerchantRoleListRes struct {
	// Merchant Roles
	MerchantRoles []UnibeeApiBeanMerchantRoleSimplify `json:"merchantRoles,omitempty"`
}

// NewUnibeeApiMerchantRoleListRes instantiates a new UnibeeApiMerchantRoleListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantRoleListRes() *UnibeeApiMerchantRoleListRes {
	this := UnibeeApiMerchantRoleListRes{}
	return &this
}

// NewUnibeeApiMerchantRoleListResWithDefaults instantiates a new UnibeeApiMerchantRoleListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantRoleListResWithDefaults() *UnibeeApiMerchantRoleListRes {
	this := UnibeeApiMerchantRoleListRes{}
	return &this
}

// GetMerchantRoles returns the MerchantRoles field value if set, zero value otherwise.
func (o *UnibeeApiMerchantRoleListRes) GetMerchantRoles() []UnibeeApiBeanMerchantRoleSimplify {
	if o == nil || IsNil(o.MerchantRoles) {
		var ret []UnibeeApiBeanMerchantRoleSimplify
		return ret
	}
	return o.MerchantRoles
}

// GetMerchantRolesOk returns a tuple with the MerchantRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleListRes) GetMerchantRolesOk() ([]UnibeeApiBeanMerchantRoleSimplify, bool) {
	if o == nil || IsNil(o.MerchantRoles) {
		return nil, false
	}
	return o.MerchantRoles, true
}

// HasMerchantRoles returns a boolean if a field has been set.
func (o *UnibeeApiMerchantRoleListRes) HasMerchantRoles() bool {
	if o != nil && !IsNil(o.MerchantRoles) {
		return true
	}

	return false
}

// SetMerchantRoles gets a reference to the given []UnibeeApiBeanMerchantRoleSimplify and assigns it to the MerchantRoles field.
func (o *UnibeeApiMerchantRoleListRes) SetMerchantRoles(v []UnibeeApiBeanMerchantRoleSimplify) {
	o.MerchantRoles = v
}

func (o UnibeeApiMerchantRoleListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantRoleListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantRoles) {
		toSerialize["merchantRoles"] = o.MerchantRoles
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantRoleListRes struct {
	value *UnibeeApiMerchantRoleListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantRoleListRes) Get() *UnibeeApiMerchantRoleListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantRoleListRes) Set(val *UnibeeApiMerchantRoleListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantRoleListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantRoleListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantRoleListRes(val *UnibeeApiMerchantRoleListRes) *NullableUnibeeApiMerchantRoleListRes {
	return &NullableUnibeeApiMerchantRoleListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantRoleListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantRoleListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

