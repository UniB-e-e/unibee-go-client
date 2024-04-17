/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404171839 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanActivateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanActivateReq{}

// UnibeeApiMerchantPlanActivateReq struct for UnibeeApiMerchantPlanActivateReq
type UnibeeApiMerchantPlanActivateReq struct {
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanActivateReq UnibeeApiMerchantPlanActivateReq

// NewUnibeeApiMerchantPlanActivateReq instantiates a new UnibeeApiMerchantPlanActivateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanActivateReq(planId int64) *UnibeeApiMerchantPlanActivateReq {
	this := UnibeeApiMerchantPlanActivateReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanActivateReqWithDefaults instantiates a new UnibeeApiMerchantPlanActivateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanActivateReqWithDefaults() *UnibeeApiMerchantPlanActivateReq {
	this := UnibeeApiMerchantPlanActivateReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanActivateReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanActivateReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanActivateReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanActivateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanActivateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanActivateReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPlanActivateReq := _UnibeeApiMerchantPlanActivateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanActivateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanActivateReq(varUnibeeApiMerchantPlanActivateReq)

	return err
}

type NullableUnibeeApiMerchantPlanActivateReq struct {
	value *UnibeeApiMerchantPlanActivateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanActivateReq) Get() *UnibeeApiMerchantPlanActivateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanActivateReq) Set(val *UnibeeApiMerchantPlanActivateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanActivateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanActivateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanActivateReq(val *UnibeeApiMerchantPlanActivateReq) *NullableUnibeeApiMerchantPlanActivateReq {
	return &NullableUnibeeApiMerchantPlanActivateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanActivateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanActivateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


