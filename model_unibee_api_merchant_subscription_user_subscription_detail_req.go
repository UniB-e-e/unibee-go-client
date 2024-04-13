/*
OpenAPI UniBee

This is UniBee api server

API version: buildtime:202404131246 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}

// UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq struct for UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq
type UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq struct {
	// UserId
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq

// NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq instantiates a new UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(userId int64) *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	this := UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
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

	varUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq := _UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(varUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq struct {
	value *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) Get() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) Set(val *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(val *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq {
	return &NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


