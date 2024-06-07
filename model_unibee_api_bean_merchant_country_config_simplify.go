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

// checks if the UnibeeApiBeanMerchantCountryConfigSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantCountryConfigSimplify{}

// UnibeeApiBeanMerchantCountryConfigSimplify struct for UnibeeApiBeanMerchantCountryConfigSimplify
type UnibeeApiBeanMerchantCountryConfigSimplify struct {
	CountryCode *string `json:"countryCode,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUnibeeApiBeanMerchantCountryConfigSimplify instantiates a new UnibeeApiBeanMerchantCountryConfigSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantCountryConfigSimplify() *UnibeeApiBeanMerchantCountryConfigSimplify {
	this := UnibeeApiBeanMerchantCountryConfigSimplify{}
	return &this
}

// NewUnibeeApiBeanMerchantCountryConfigSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantCountryConfigSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantCountryConfigSimplifyWithDefaults() *UnibeeApiBeanMerchantCountryConfigSimplify {
	this := UnibeeApiBeanMerchantCountryConfigSimplify{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanMerchantCountryConfigSimplify) SetName(v string) {
	o.Name = &v
}

func (o UnibeeApiBeanMerchantCountryConfigSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantCountryConfigSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantCountryConfigSimplify struct {
	value *UnibeeApiBeanMerchantCountryConfigSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantCountryConfigSimplify) Get() *UnibeeApiBeanMerchantCountryConfigSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantCountryConfigSimplify) Set(val *UnibeeApiBeanMerchantCountryConfigSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantCountryConfigSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantCountryConfigSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantCountryConfigSimplify(val *UnibeeApiBeanMerchantCountryConfigSimplify) *NullableUnibeeApiBeanMerchantCountryConfigSimplify {
	return &NullableUnibeeApiBeanMerchantCountryConfigSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantCountryConfigSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantCountryConfigSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


