/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: buildtime:202404090336 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantMemberListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMemberListGet200ResponseData{}

// MerchantMemberListGet200ResponseData struct for MerchantMemberListGet200ResponseData
type MerchantMemberListGet200ResponseData struct {
	// Merchant Members
	MerchantMembers []UnibeeApiBeanMerchantMemberSimplify `json:"merchantMembers,omitempty"`
}

// NewMerchantMemberListGet200ResponseData instantiates a new MerchantMemberListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMemberListGet200ResponseData() *MerchantMemberListGet200ResponseData {
	this := MerchantMemberListGet200ResponseData{}
	return &this
}

// NewMerchantMemberListGet200ResponseDataWithDefaults instantiates a new MerchantMemberListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMemberListGet200ResponseDataWithDefaults() *MerchantMemberListGet200ResponseData {
	this := MerchantMemberListGet200ResponseData{}
	return &this
}

// GetMerchantMembers returns the MerchantMembers field value if set, zero value otherwise.
func (o *MerchantMemberListGet200ResponseData) GetMerchantMembers() []UnibeeApiBeanMerchantMemberSimplify {
	if o == nil || IsNil(o.MerchantMembers) {
		var ret []UnibeeApiBeanMerchantMemberSimplify
		return ret
	}
	return o.MerchantMembers
}

// GetMerchantMembersOk returns a tuple with the MerchantMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMemberListGet200ResponseData) GetMerchantMembersOk() ([]UnibeeApiBeanMerchantMemberSimplify, bool) {
	if o == nil || IsNil(o.MerchantMembers) {
		return nil, false
	}
	return o.MerchantMembers, true
}

// HasMerchantMembers returns a boolean if a field has been set.
func (o *MerchantMemberListGet200ResponseData) HasMerchantMembers() bool {
	if o != nil && !IsNil(o.MerchantMembers) {
		return true
	}

	return false
}

// SetMerchantMembers gets a reference to the given []UnibeeApiBeanMerchantMemberSimplify and assigns it to the MerchantMembers field.
func (o *MerchantMemberListGet200ResponseData) SetMerchantMembers(v []UnibeeApiBeanMerchantMemberSimplify) {
	o.MerchantMembers = v
}

func (o MerchantMemberListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMemberListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMembers) {
		toSerialize["merchantMembers"] = o.MerchantMembers
	}
	return toSerialize, nil
}

type NullableMerchantMemberListGet200ResponseData struct {
	value *MerchantMemberListGet200ResponseData
	isSet bool
}

func (v NullableMerchantMemberListGet200ResponseData) Get() *MerchantMemberListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantMemberListGet200ResponseData) Set(val *MerchantMemberListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMemberListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMemberListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMemberListGet200ResponseData(val *MerchantMemberListGet200ResponseData) *NullableMerchantMemberListGet200ResponseData {
	return &NullableMerchantMemberListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMemberListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMemberListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


