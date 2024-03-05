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

// checks if the UnibeeApiMerchantAuthLoginRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthLoginRes{}

// UnibeeApiMerchantAuthLoginRes struct for UnibeeApiMerchantAuthLoginRes
type UnibeeApiMerchantAuthLoginRes struct {
	MerchantMember *UnibeeInternalModelEntityOverseaPayMerchantMember `json:"merchantMember,omitempty"`
	// Token
	Token *string `json:"token,omitempty"`
}

// NewUnibeeApiMerchantAuthLoginRes instantiates a new UnibeeApiMerchantAuthLoginRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthLoginRes() *UnibeeApiMerchantAuthLoginRes {
	this := UnibeeApiMerchantAuthLoginRes{}
	return &this
}

// NewUnibeeApiMerchantAuthLoginResWithDefaults instantiates a new UnibeeApiMerchantAuthLoginRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthLoginResWithDefaults() *UnibeeApiMerchantAuthLoginRes {
	this := UnibeeApiMerchantAuthLoginRes{}
	return &this
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthLoginRes) GetMerchantMember() UnibeeInternalModelEntityOverseaPayMerchantMember {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeInternalModelEntityOverseaPayMerchantMember
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginRes) GetMerchantMemberOk() (*UnibeeInternalModelEntityOverseaPayMerchantMember, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthLoginRes) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeInternalModelEntityOverseaPayMerchantMember and assigns it to the MerchantMember field.
func (o *UnibeeApiMerchantAuthLoginRes) SetMerchantMember(v UnibeeInternalModelEntityOverseaPayMerchantMember) {
	o.MerchantMember = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthLoginRes) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginRes) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthLoginRes) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UnibeeApiMerchantAuthLoginRes) SetToken(v string) {
	o.Token = &v
}

func (o UnibeeApiMerchantAuthLoginRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthLoginRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantAuthLoginRes struct {
	value *UnibeeApiMerchantAuthLoginRes
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthLoginRes) Get() *UnibeeApiMerchantAuthLoginRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthLoginRes) Set(val *UnibeeApiMerchantAuthLoginRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthLoginRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthLoginRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthLoginRes(val *UnibeeApiMerchantAuthLoginRes) *NullableUnibeeApiMerchantAuthLoginRes {
	return &NullableUnibeeApiMerchantAuthLoginRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthLoginRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthLoginRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

