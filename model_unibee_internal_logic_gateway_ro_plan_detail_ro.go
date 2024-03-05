/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoPlanDetailRo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoPlanDetailRo{}

// UnibeeInternalLogicGatewayRoPlanDetailRo struct for UnibeeInternalLogicGatewayRoPlanDetailRo
type UnibeeInternalLogicGatewayRoPlanDetailRo struct {
	// AddonIds
	AddonIds []int64 `json:"addonIds,omitempty"`
	// Addons
	Addons []UnibeeInternalLogicGatewayRoPlanSimplify `json:"addons,omitempty"`
	// MetricPlanLimits
	MetricPlanLimits []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo `json:"metricPlanLimits,omitempty"`
	Plan *UnibeeInternalLogicGatewayRoPlanSimplify `json:"plan,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoPlanDetailRo instantiates a new UnibeeInternalLogicGatewayRoPlanDetailRo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoPlanDetailRo() *UnibeeInternalLogicGatewayRoPlanDetailRo {
	this := UnibeeInternalLogicGatewayRoPlanDetailRo{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoPlanDetailRoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoPlanDetailRo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoPlanDetailRoWithDefaults() *UnibeeInternalLogicGatewayRoPlanDetailRo {
	this := UnibeeInternalLogicGatewayRoPlanDetailRo{}
	return &this
}

// GetAddonIds returns the AddonIds field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetAddonIds() []int64 {
	if o == nil || IsNil(o.AddonIds) {
		var ret []int64
		return ret
	}
	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AddonIds) {
		return nil, false
	}
	return o.AddonIds, true
}

// HasAddonIds returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) HasAddonIds() bool {
	if o != nil && !IsNil(o.AddonIds) {
		return true
	}

	return false
}

// SetAddonIds gets a reference to the given []int64 and assigns it to the AddonIds field.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetAddons() []UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetAddonsOk() ([]UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the Addons field.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) SetAddons(v []UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.Addons = v
}

// GetMetricPlanLimits returns the MetricPlanLimits field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetMetricPlanLimits() []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo {
	if o == nil || IsNil(o.MetricPlanLimits) {
		var ret []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo
		return ret
	}
	return o.MetricPlanLimits
}

// GetMetricPlanLimitsOk returns a tuple with the MetricPlanLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetMetricPlanLimitsOk() ([]UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo, bool) {
	if o == nil || IsNil(o.MetricPlanLimits) {
		return nil, false
	}
	return o.MetricPlanLimits, true
}

// HasMetricPlanLimits returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) HasMetricPlanLimits() bool {
	if o != nil && !IsNil(o.MetricPlanLimits) {
		return true
	}

	return false
}

// SetMetricPlanLimits gets a reference to the given []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo and assigns it to the MetricPlanLimits field.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) SetMetricPlanLimits(v []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo) {
	o.MetricPlanLimits = v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the Plan field.
func (o *UnibeeInternalLogicGatewayRoPlanDetailRo) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.Plan = &v
}

func (o UnibeeInternalLogicGatewayRoPlanDetailRo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoPlanDetailRo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonIds) {
		toSerialize["addonIds"] = o.AddonIds
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.MetricPlanLimits) {
		toSerialize["metricPlanLimits"] = o.MetricPlanLimits
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoPlanDetailRo struct {
	value *UnibeeInternalLogicGatewayRoPlanDetailRo
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoPlanDetailRo) Get() *UnibeeInternalLogicGatewayRoPlanDetailRo {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoPlanDetailRo) Set(val *UnibeeInternalLogicGatewayRoPlanDetailRo) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoPlanDetailRo) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoPlanDetailRo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoPlanDetailRo(val *UnibeeInternalLogicGatewayRoPlanDetailRo) *NullableUnibeeInternalLogicGatewayRoPlanDetailRo {
	return &NullableUnibeeInternalLogicGatewayRoPlanDetailRo{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoPlanDetailRo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoPlanDetailRo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

