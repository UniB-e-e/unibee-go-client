/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}

// UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq struct for UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq
type UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq struct {
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq

// NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq instantiates a new UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(subscriptionId string) *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	this := UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	this := UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq := _UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(varUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq struct {
	value *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) Get() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) Set(val *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(val *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq {
	return &NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


