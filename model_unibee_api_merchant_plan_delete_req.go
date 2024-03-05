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

// checks if the UnibeeApiMerchantPlanDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanDeleteReq{}

// UnibeeApiMerchantPlanDeleteReq struct for UnibeeApiMerchantPlanDeleteReq
type UnibeeApiMerchantPlanDeleteReq struct {
	// PlanId
	PlanId int32 `json:"planId"`
}

type _UnibeeApiMerchantPlanDeleteReq UnibeeApiMerchantPlanDeleteReq

// NewUnibeeApiMerchantPlanDeleteReq instantiates a new UnibeeApiMerchantPlanDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanDeleteReq(planId int32) *UnibeeApiMerchantPlanDeleteReq {
	this := UnibeeApiMerchantPlanDeleteReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanDeleteReqWithDefaults instantiates a new UnibeeApiMerchantPlanDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanDeleteReqWithDefaults() *UnibeeApiMerchantPlanDeleteReq {
	this := UnibeeApiMerchantPlanDeleteReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanDeleteReq) GetPlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanDeleteReq) GetPlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanDeleteReq) SetPlanId(v int32) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanDeleteReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planId",
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

	varUnibeeApiMerchantPlanDeleteReq := _UnibeeApiMerchantPlanDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanDeleteReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanDeleteReq(varUnibeeApiMerchantPlanDeleteReq)

	return err
}

type NullableUnibeeApiMerchantPlanDeleteReq struct {
	value *UnibeeApiMerchantPlanDeleteReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanDeleteReq) Get() *UnibeeApiMerchantPlanDeleteReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanDeleteReq) Set(val *UnibeeApiMerchantPlanDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanDeleteReq(val *UnibeeApiMerchantPlanDeleteReq) *NullableUnibeeApiMerchantPlanDeleteReq {
	return &NullableUnibeeApiMerchantPlanDeleteReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

