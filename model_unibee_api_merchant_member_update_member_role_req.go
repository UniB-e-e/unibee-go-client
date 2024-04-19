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

// checks if the UnibeeApiMerchantMemberUpdateMemberRoleReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberUpdateMemberRoleReq{}

// UnibeeApiMerchantMemberUpdateMemberRoleReq struct for UnibeeApiMerchantMemberUpdateMemberRoleReq
type UnibeeApiMerchantMemberUpdateMemberRoleReq struct {
	// The unique id of member
	MemberId *int64 `json:"memberId,omitempty"`
	// The permission role of member
	Role *string `json:"role,omitempty"`
}

// NewUnibeeApiMerchantMemberUpdateMemberRoleReq instantiates a new UnibeeApiMerchantMemberUpdateMemberRoleReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberUpdateMemberRoleReq() *UnibeeApiMerchantMemberUpdateMemberRoleReq {
	this := UnibeeApiMerchantMemberUpdateMemberRoleReq{}
	return &this
}

// NewUnibeeApiMerchantMemberUpdateMemberRoleReqWithDefaults instantiates a new UnibeeApiMerchantMemberUpdateMemberRoleReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberUpdateMemberRoleReqWithDefaults() *UnibeeApiMerchantMemberUpdateMemberRoleReq {
	this := UnibeeApiMerchantMemberUpdateMemberRoleReq{}
	return &this
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UnibeeApiMerchantMemberUpdateMemberRoleReq) SetRole(v string) {
	o.Role = &v
}

func (o UnibeeApiMerchantMemberUpdateMemberRoleReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberUpdateMemberRoleReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberUpdateMemberRoleReq struct {
	value *UnibeeApiMerchantMemberUpdateMemberRoleReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberUpdateMemberRoleReq) Get() *UnibeeApiMerchantMemberUpdateMemberRoleReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberUpdateMemberRoleReq) Set(val *UnibeeApiMerchantMemberUpdateMemberRoleReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberUpdateMemberRoleReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberUpdateMemberRoleReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberUpdateMemberRoleReq(val *UnibeeApiMerchantMemberUpdateMemberRoleReq) *NullableUnibeeApiMerchantMemberUpdateMemberRoleReq {
	return &NullableUnibeeApiMerchantMemberUpdateMemberRoleReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberUpdateMemberRoleReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberUpdateMemberRoleReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


