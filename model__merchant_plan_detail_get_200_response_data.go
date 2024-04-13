/*
OpenAPI UniBee

This is UniBee api server

API version: buildtime:202404131246 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPlanDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPlanDetailGet200ResponseData{}

// MerchantPlanDetailGet200ResponseData struct for MerchantPlanDetailGet200ResponseData
type MerchantPlanDetailGet200ResponseData struct {
	Plan *UnibeeApiBeanDetailPlanDetail `json:"plan,omitempty"`
}

// NewMerchantPlanDetailGet200ResponseData instantiates a new MerchantPlanDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPlanDetailGet200ResponseData() *MerchantPlanDetailGet200ResponseData {
	this := MerchantPlanDetailGet200ResponseData{}
	return &this
}

// NewMerchantPlanDetailGet200ResponseDataWithDefaults instantiates a new MerchantPlanDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPlanDetailGet200ResponseDataWithDefaults() *MerchantPlanDetailGet200ResponseData {
	this := MerchantPlanDetailGet200ResponseData{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *MerchantPlanDetailGet200ResponseData) GetPlan() UnibeeApiBeanDetailPlanDetail {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanDetailPlanDetail
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPlanDetailGet200ResponseData) GetPlanOk() (*UnibeeApiBeanDetailPlanDetail, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *MerchantPlanDetailGet200ResponseData) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanDetailPlanDetail and assigns it to the Plan field.
func (o *MerchantPlanDetailGet200ResponseData) SetPlan(v UnibeeApiBeanDetailPlanDetail) {
	o.Plan = &v
}

func (o MerchantPlanDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPlanDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableMerchantPlanDetailGet200ResponseData struct {
	value *MerchantPlanDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantPlanDetailGet200ResponseData) Get() *MerchantPlanDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPlanDetailGet200ResponseData) Set(val *MerchantPlanDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPlanDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPlanDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPlanDetailGet200ResponseData(val *MerchantPlanDetailGet200ResponseData) *NullableMerchantPlanDetailGet200ResponseData {
	return &NullableMerchantPlanDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPlanDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPlanDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


