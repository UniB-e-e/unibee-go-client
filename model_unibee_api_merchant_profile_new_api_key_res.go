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

// checks if the UnibeeApiMerchantProfileNewApiKeyRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProfileNewApiKeyRes{}

// UnibeeApiMerchantProfileNewApiKeyRes struct for UnibeeApiMerchantProfileNewApiKeyRes
type UnibeeApiMerchantProfileNewApiKeyRes struct {
	// ApiKey
	ApiKey *string `json:"apiKey,omitempty"`
}

// NewUnibeeApiMerchantProfileNewApiKeyRes instantiates a new UnibeeApiMerchantProfileNewApiKeyRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProfileNewApiKeyRes() *UnibeeApiMerchantProfileNewApiKeyRes {
	this := UnibeeApiMerchantProfileNewApiKeyRes{}
	return &this
}

// NewUnibeeApiMerchantProfileNewApiKeyResWithDefaults instantiates a new UnibeeApiMerchantProfileNewApiKeyRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProfileNewApiKeyResWithDefaults() *UnibeeApiMerchantProfileNewApiKeyRes {
	this := UnibeeApiMerchantProfileNewApiKeyRes{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProfileNewApiKeyRes) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProfileNewApiKeyRes) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProfileNewApiKeyRes) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *UnibeeApiMerchantProfileNewApiKeyRes) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o UnibeeApiMerchantProfileNewApiKeyRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProfileNewApiKeyRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProfileNewApiKeyRes struct {
	value *UnibeeApiMerchantProfileNewApiKeyRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProfileNewApiKeyRes) Get() *UnibeeApiMerchantProfileNewApiKeyRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProfileNewApiKeyRes) Set(val *UnibeeApiMerchantProfileNewApiKeyRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProfileNewApiKeyRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProfileNewApiKeyRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProfileNewApiKeyRes(val *UnibeeApiMerchantProfileNewApiKeyRes) *NullableUnibeeApiMerchantProfileNewApiKeyRes {
	return &NullableUnibeeApiMerchantProfileNewApiKeyRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProfileNewApiKeyRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProfileNewApiKeyRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

