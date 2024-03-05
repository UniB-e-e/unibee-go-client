/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthLoginReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthLoginReq{}

// UnibeeApiMerchantAuthLoginReq struct for UnibeeApiMerchantAuthLoginReq
type UnibeeApiMerchantAuthLoginReq struct {
	// Email
	Email string `json:"email"`
	// Password
	Password string `json:"password"`
}

type _UnibeeApiMerchantAuthLoginReq UnibeeApiMerchantAuthLoginReq

// NewUnibeeApiMerchantAuthLoginReq instantiates a new UnibeeApiMerchantAuthLoginReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthLoginReq(email string, password string) *UnibeeApiMerchantAuthLoginReq {
	this := UnibeeApiMerchantAuthLoginReq{}
	this.Email = email
	this.Password = password
	return &this
}

// NewUnibeeApiMerchantAuthLoginReqWithDefaults instantiates a new UnibeeApiMerchantAuthLoginReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthLoginReqWithDefaults() *UnibeeApiMerchantAuthLoginReq {
	this := UnibeeApiMerchantAuthLoginReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthLoginReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthLoginReq) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *UnibeeApiMerchantAuthLoginReq) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginReq) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UnibeeApiMerchantAuthLoginReq) SetPassword(v string) {
	o.Password = v
}

func (o UnibeeApiMerchantAuthLoginReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthLoginReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthLoginReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantAuthLoginReq := _UnibeeApiMerchantAuthLoginReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthLoginReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthLoginReq(varUnibeeApiMerchantAuthLoginReq)

	return err
}

type NullableUnibeeApiMerchantAuthLoginReq struct {
	value *UnibeeApiMerchantAuthLoginReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthLoginReq) Get() *UnibeeApiMerchantAuthLoginReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthLoginReq) Set(val *UnibeeApiMerchantAuthLoginReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthLoginReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthLoginReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthLoginReq(val *UnibeeApiMerchantAuthLoginReq) *NullableUnibeeApiMerchantAuthLoginReq {
	return &NullableUnibeeApiMerchantAuthLoginReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthLoginReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthLoginReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

