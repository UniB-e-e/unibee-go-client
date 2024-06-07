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

// checks if the UnibeeApiMerchantUserListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserListRes{}

// UnibeeApiMerchantUserListRes struct for UnibeeApiMerchantUserListRes
type UnibeeApiMerchantUserListRes struct {
	// Total
	Total *int32 `json:"total,omitempty"`
	// User Account Object List
	UserAccounts []UnibeeApiBeanDetailUserAccountDetail `json:"userAccounts,omitempty"`
}

// NewUnibeeApiMerchantUserListRes instantiates a new UnibeeApiMerchantUserListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserListRes() *UnibeeApiMerchantUserListRes {
	this := UnibeeApiMerchantUserListRes{}
	return &this
}

// NewUnibeeApiMerchantUserListResWithDefaults instantiates a new UnibeeApiMerchantUserListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserListResWithDefaults() *UnibeeApiMerchantUserListRes {
	this := UnibeeApiMerchantUserListRes{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantUserListRes) SetTotal(v int32) {
	o.Total = &v
}

// GetUserAccounts returns the UserAccounts field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListRes) GetUserAccounts() []UnibeeApiBeanDetailUserAccountDetail {
	if o == nil || IsNil(o.UserAccounts) {
		var ret []UnibeeApiBeanDetailUserAccountDetail
		return ret
	}
	return o.UserAccounts
}

// GetUserAccountsOk returns a tuple with the UserAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListRes) GetUserAccountsOk() ([]UnibeeApiBeanDetailUserAccountDetail, bool) {
	if o == nil || IsNil(o.UserAccounts) {
		return nil, false
	}
	return o.UserAccounts, true
}

// HasUserAccounts returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListRes) HasUserAccounts() bool {
	if o != nil && !IsNil(o.UserAccounts) {
		return true
	}

	return false
}

// SetUserAccounts gets a reference to the given []UnibeeApiBeanDetailUserAccountDetail and assigns it to the UserAccounts field.
func (o *UnibeeApiMerchantUserListRes) SetUserAccounts(v []UnibeeApiBeanDetailUserAccountDetail) {
	o.UserAccounts = v
}

func (o UnibeeApiMerchantUserListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.UserAccounts) {
		toSerialize["userAccounts"] = o.UserAccounts
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserListRes struct {
	value *UnibeeApiMerchantUserListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantUserListRes) Get() *UnibeeApiMerchantUserListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserListRes) Set(val *UnibeeApiMerchantUserListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserListRes(val *UnibeeApiMerchantUserListRes) *NullableUnibeeApiMerchantUserListRes {
	return &NullableUnibeeApiMerchantUserListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


