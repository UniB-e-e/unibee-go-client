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

// checks if the MerchantAuthSsoLoginPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantAuthSsoLoginPost200ResponseData{}

// MerchantAuthSsoLoginPost200ResponseData struct for MerchantAuthSsoLoginPost200ResponseData
type MerchantAuthSsoLoginPost200ResponseData struct {
	MerchantMember *UnibeeInternalModelEntityOverseaPayMerchantMember `json:"merchantMember,omitempty"`
	// Token
	Token *string `json:"token,omitempty"`
}

// NewMerchantAuthSsoLoginPost200ResponseData instantiates a new MerchantAuthSsoLoginPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAuthSsoLoginPost200ResponseData() *MerchantAuthSsoLoginPost200ResponseData {
	this := MerchantAuthSsoLoginPost200ResponseData{}
	return &this
}

// NewMerchantAuthSsoLoginPost200ResponseDataWithDefaults instantiates a new MerchantAuthSsoLoginPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAuthSsoLoginPost200ResponseDataWithDefaults() *MerchantAuthSsoLoginPost200ResponseData {
	this := MerchantAuthSsoLoginPost200ResponseData{}
	return &this
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *MerchantAuthSsoLoginPost200ResponseData) GetMerchantMember() UnibeeInternalModelEntityOverseaPayMerchantMember {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeInternalModelEntityOverseaPayMerchantMember
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAuthSsoLoginPost200ResponseData) GetMerchantMemberOk() (*UnibeeInternalModelEntityOverseaPayMerchantMember, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *MerchantAuthSsoLoginPost200ResponseData) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeInternalModelEntityOverseaPayMerchantMember and assigns it to the MerchantMember field.
func (o *MerchantAuthSsoLoginPost200ResponseData) SetMerchantMember(v UnibeeInternalModelEntityOverseaPayMerchantMember) {
	o.MerchantMember = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MerchantAuthSsoLoginPost200ResponseData) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAuthSsoLoginPost200ResponseData) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MerchantAuthSsoLoginPost200ResponseData) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MerchantAuthSsoLoginPost200ResponseData) SetToken(v string) {
	o.Token = &v
}

func (o MerchantAuthSsoLoginPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantAuthSsoLoginPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableMerchantAuthSsoLoginPost200ResponseData struct {
	value *MerchantAuthSsoLoginPost200ResponseData
	isSet bool
}

func (v NullableMerchantAuthSsoLoginPost200ResponseData) Get() *MerchantAuthSsoLoginPost200ResponseData {
	return v.value
}

func (v *NullableMerchantAuthSsoLoginPost200ResponseData) Set(val *MerchantAuthSsoLoginPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAuthSsoLoginPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAuthSsoLoginPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAuthSsoLoginPost200ResponseData(val *MerchantAuthSsoLoginPost200ResponseData) *NullableMerchantAuthSsoLoginPost200ResponseData {
	return &NullableMerchantAuthSsoLoginPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantAuthSsoLoginPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAuthSsoLoginPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

