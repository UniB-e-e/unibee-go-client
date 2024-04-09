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

// checks if the UnibeeApiMerchantSearchPrecisionMatchObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSearchPrecisionMatchObject{}

// UnibeeApiMerchantSearchPrecisionMatchObject struct for UnibeeApiMerchantSearchPrecisionMatchObject
type UnibeeApiMerchantSearchPrecisionMatchObject struct {
	Data map[string]interface{} `json:"data,omitempty"`
	// match Id user_id|invoice_id
	Id *string `json:"id,omitempty"`
	// match Type, user|invoice
	Type *string `json:"type,omitempty"`
}

// NewUnibeeApiMerchantSearchPrecisionMatchObject instantiates a new UnibeeApiMerchantSearchPrecisionMatchObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSearchPrecisionMatchObject() *UnibeeApiMerchantSearchPrecisionMatchObject {
	this := UnibeeApiMerchantSearchPrecisionMatchObject{}
	return &this
}

// NewUnibeeApiMerchantSearchPrecisionMatchObjectWithDefaults instantiates a new UnibeeApiMerchantSearchPrecisionMatchObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSearchPrecisionMatchObjectWithDefaults() *UnibeeApiMerchantSearchPrecisionMatchObject {
	this := UnibeeApiMerchantSearchPrecisionMatchObject{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnibeeApiMerchantSearchPrecisionMatchObject) SetType(v string) {
	o.Type = &v
}

func (o UnibeeApiMerchantSearchPrecisionMatchObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSearchPrecisionMatchObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSearchPrecisionMatchObject struct {
	value *UnibeeApiMerchantSearchPrecisionMatchObject
	isSet bool
}

func (v NullableUnibeeApiMerchantSearchPrecisionMatchObject) Get() *UnibeeApiMerchantSearchPrecisionMatchObject {
	return v.value
}

func (v *NullableUnibeeApiMerchantSearchPrecisionMatchObject) Set(val *UnibeeApiMerchantSearchPrecisionMatchObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSearchPrecisionMatchObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSearchPrecisionMatchObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSearchPrecisionMatchObject(val *UnibeeApiMerchantSearchPrecisionMatchObject) *NullableUnibeeApiMerchantSearchPrecisionMatchObject {
	return &NullableUnibeeApiMerchantSearchPrecisionMatchObject{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSearchPrecisionMatchObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSearchPrecisionMatchObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


