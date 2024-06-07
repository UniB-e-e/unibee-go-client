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

// checks if the MerchantRoleListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantRoleListGet200ResponseData{}

// MerchantRoleListGet200ResponseData struct for MerchantRoleListGet200ResponseData
type MerchantRoleListGet200ResponseData struct {
	// Merchant Roles
	MerchantRoles []UnibeeApiBeanMerchantRoleSimplify `json:"merchantRoles,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantRoleListGet200ResponseData instantiates a new MerchantRoleListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantRoleListGet200ResponseData() *MerchantRoleListGet200ResponseData {
	this := MerchantRoleListGet200ResponseData{}
	return &this
}

// NewMerchantRoleListGet200ResponseDataWithDefaults instantiates a new MerchantRoleListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantRoleListGet200ResponseDataWithDefaults() *MerchantRoleListGet200ResponseData {
	this := MerchantRoleListGet200ResponseData{}
	return &this
}

// GetMerchantRoles returns the MerchantRoles field value if set, zero value otherwise.
func (o *MerchantRoleListGet200ResponseData) GetMerchantRoles() []UnibeeApiBeanMerchantRoleSimplify {
	if o == nil || IsNil(o.MerchantRoles) {
		var ret []UnibeeApiBeanMerchantRoleSimplify
		return ret
	}
	return o.MerchantRoles
}

// GetMerchantRolesOk returns a tuple with the MerchantRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRoleListGet200ResponseData) GetMerchantRolesOk() ([]UnibeeApiBeanMerchantRoleSimplify, bool) {
	if o == nil || IsNil(o.MerchantRoles) {
		return nil, false
	}
	return o.MerchantRoles, true
}

// HasMerchantRoles returns a boolean if a field has been set.
func (o *MerchantRoleListGet200ResponseData) HasMerchantRoles() bool {
	if o != nil && !IsNil(o.MerchantRoles) {
		return true
	}

	return false
}

// SetMerchantRoles gets a reference to the given []UnibeeApiBeanMerchantRoleSimplify and assigns it to the MerchantRoles field.
func (o *MerchantRoleListGet200ResponseData) SetMerchantRoles(v []UnibeeApiBeanMerchantRoleSimplify) {
	o.MerchantRoles = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantRoleListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRoleListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantRoleListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantRoleListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantRoleListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantRoleListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantRoles) {
		toSerialize["merchantRoles"] = o.MerchantRoles
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantRoleListGet200ResponseData struct {
	value *MerchantRoleListGet200ResponseData
	isSet bool
}

func (v NullableMerchantRoleListGet200ResponseData) Get() *MerchantRoleListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantRoleListGet200ResponseData) Set(val *MerchantRoleListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantRoleListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantRoleListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantRoleListGet200ResponseData(val *MerchantRoleListGet200ResponseData) *NullableMerchantRoleListGet200ResponseData {
	return &NullableMerchantRoleListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantRoleListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantRoleListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


