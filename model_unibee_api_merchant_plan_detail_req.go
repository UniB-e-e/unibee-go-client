/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406070109 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanDetailReq{}

// UnibeeApiMerchantPlanDetailReq struct for UnibeeApiMerchantPlanDetailReq
type UnibeeApiMerchantPlanDetailReq struct {
	// PlanId
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanDetailReq UnibeeApiMerchantPlanDetailReq

// NewUnibeeApiMerchantPlanDetailReq instantiates a new UnibeeApiMerchantPlanDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanDetailReq(planId int64) *UnibeeApiMerchantPlanDetailReq {
	this := UnibeeApiMerchantPlanDetailReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanDetailReqWithDefaults instantiates a new UnibeeApiMerchantPlanDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanDetailReqWithDefaults() *UnibeeApiMerchantPlanDetailReq {
	this := UnibeeApiMerchantPlanDetailReq{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanDetailReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanDetailReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanDetailReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanDetailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantPlanDetailReq := _UnibeeApiMerchantPlanDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanDetailReq(varUnibeeApiMerchantPlanDetailReq)

	return err
}

type NullableUnibeeApiMerchantPlanDetailReq struct {
	value *UnibeeApiMerchantPlanDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanDetailReq) Get() *UnibeeApiMerchantPlanDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanDetailReq) Set(val *UnibeeApiMerchantPlanDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanDetailReq(val *UnibeeApiMerchantPlanDetailReq) *NullableUnibeeApiMerchantPlanDetailReq {
	return &NullableUnibeeApiMerchantPlanDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


