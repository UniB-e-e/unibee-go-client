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

// checks if the UnibeeApiMerchantSubscriptionRenewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionRenewReq{}

// UnibeeApiMerchantSubscriptionRenewReq struct for UnibeeApiMerchantSubscriptionRenewReq
type UnibeeApiMerchantSubscriptionRenewReq struct {
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
	// UserId
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantSubscriptionRenewReq UnibeeApiMerchantSubscriptionRenewReq

// NewUnibeeApiMerchantSubscriptionRenewReq instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionRenewReq(subscriptionId string, userId int64) *UnibeeApiMerchantSubscriptionRenewReq {
	this := UnibeeApiMerchantSubscriptionRenewReq{}
	this.SubscriptionId = subscriptionId
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults() *UnibeeApiMerchantSubscriptionRenewReq {
	this := UnibeeApiMerchantSubscriptionRenewReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantSubscriptionRenewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionRenewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionRenewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionRenewReq := _UnibeeApiMerchantSubscriptionRenewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionRenewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionRenewReq(varUnibeeApiMerchantSubscriptionRenewReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionRenewReq struct {
	value *UnibeeApiMerchantSubscriptionRenewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionRenewReq) Get() *UnibeeApiMerchantSubscriptionRenewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewReq) Set(val *UnibeeApiMerchantSubscriptionRenewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionRenewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionRenewReq(val *UnibeeApiMerchantSubscriptionRenewReq) *NullableUnibeeApiMerchantSubscriptionRenewReq {
	return &NullableUnibeeApiMerchantSubscriptionRenewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionRenewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


