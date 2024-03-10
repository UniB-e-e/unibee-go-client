/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanPlanDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanPlanDetail{}

// UnibeeApiMerchantPlanPlanDetail struct for UnibeeApiMerchantPlanPlanDetail
type UnibeeApiMerchantPlanPlanDetail struct {
	// AddonIds
	AddonIds []int64 `json:"addonIds,omitempty"`
	// Addons
	Addons []UnibeeInternalLogicGatewayRoPlanSimplify `json:"addons,omitempty"`
	// MetricPlanLimits
	MetricPlanLimits []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo `json:"metricPlanLimits,omitempty"`
	Plan *UnibeeInternalLogicGatewayRoPlanSimplify `json:"plan,omitempty"`
}

// NewUnibeeApiMerchantPlanPlanDetail instantiates a new UnibeeApiMerchantPlanPlanDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanPlanDetail() *UnibeeApiMerchantPlanPlanDetail {
	this := UnibeeApiMerchantPlanPlanDetail{}
	return &this
}

// NewUnibeeApiMerchantPlanPlanDetailWithDefaults instantiates a new UnibeeApiMerchantPlanPlanDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanPlanDetailWithDefaults() *UnibeeApiMerchantPlanPlanDetail {
	this := UnibeeApiMerchantPlanPlanDetail{}
	return &this
}

// GetAddonIds returns the AddonIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanPlanDetail) GetAddonIds() []int64 {
	if o == nil || IsNil(o.AddonIds) {
		var ret []int64
		return ret
	}
	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) GetAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AddonIds) {
		return nil, false
	}
	return o.AddonIds, true
}

// HasAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) HasAddonIds() bool {
	if o != nil && !IsNil(o.AddonIds) {
		return true
	}

	return false
}

// SetAddonIds gets a reference to the given []int64 and assigns it to the AddonIds field.
func (o *UnibeeApiMerchantPlanPlanDetail) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanPlanDetail) GetAddons() []UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) GetAddonsOk() ([]UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the Addons field.
func (o *UnibeeApiMerchantPlanPlanDetail) SetAddons(v []UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.Addons = v
}

// GetMetricPlanLimits returns the MetricPlanLimits field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanPlanDetail) GetMetricPlanLimits() []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo {
	if o == nil || IsNil(o.MetricPlanLimits) {
		var ret []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo
		return ret
	}
	return o.MetricPlanLimits
}

// GetMetricPlanLimitsOk returns a tuple with the MetricPlanLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) GetMetricPlanLimitsOk() ([]UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo, bool) {
	if o == nil || IsNil(o.MetricPlanLimits) {
		return nil, false
	}
	return o.MetricPlanLimits, true
}

// HasMetricPlanLimits returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) HasMetricPlanLimits() bool {
	if o != nil && !IsNil(o.MetricPlanLimits) {
		return true
	}

	return false
}

// SetMetricPlanLimits gets a reference to the given []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo and assigns it to the MetricPlanLimits field.
func (o *UnibeeApiMerchantPlanPlanDetail) SetMetricPlanLimits(v []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo) {
	o.MetricPlanLimits = v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanPlanDetail) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanPlanDetail) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the Plan field.
func (o *UnibeeApiMerchantPlanPlanDetail) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.Plan = &v
}

func (o UnibeeApiMerchantPlanPlanDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanPlanDetail) ToMap() (map[string]interface{}, error) {
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

type NullableUnibeeApiMerchantPlanPlanDetail struct {
	value *UnibeeApiMerchantPlanPlanDetail
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanPlanDetail) Get() *UnibeeApiMerchantPlanPlanDetail {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanPlanDetail) Set(val *UnibeeApiMerchantPlanPlanDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanPlanDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanPlanDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanPlanDetail(val *UnibeeApiMerchantPlanPlanDetail) *NullableUnibeeApiMerchantPlanPlanDetail {
	return &NullableUnibeeApiMerchantPlanPlanDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanPlanDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanPlanDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

