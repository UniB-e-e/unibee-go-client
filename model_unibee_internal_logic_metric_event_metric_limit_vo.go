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

// checks if the UnibeeInternalLogicMetricEventMetricLimitVo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicMetricEventMetricLimitVo{}

// UnibeeInternalLogicMetricEventMetricLimitVo struct for UnibeeInternalLogicMetricEventMetricLimitVo
type UnibeeInternalLogicMetricEventMetricLimitVo struct {
	MerchantId *int32 `json:"MerchantId,omitempty"`
	MetricId *int64 `json:"MetricId,omitempty"`
	PlanLimits []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo `json:"PlanLimits,omitempty"`
	TotalLimit *int32 `json:"TotalLimit,omitempty"`
	UserId *int64 `json:"UserId,omitempty"`
	// aggregation property
	AggregationProperty *string `json:"aggregationProperty,omitempty"`
	// 0-count，1-count unique, 2-latest, 3-max, 4-sum
	AggregationType *int32 `json:"aggregationType,omitempty"`
	// code
	Code *string `json:"code,omitempty"`
	// metric name
	MetricName *string `json:"metricName,omitempty"`
	// 1-limit_metered，2-charge_metered(come later),3-charge_recurring(come later)
	Type *int32 `json:"type,omitempty"`
}

// NewUnibeeInternalLogicMetricEventMetricLimitVo instantiates a new UnibeeInternalLogicMetricEventMetricLimitVo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicMetricEventMetricLimitVo() *UnibeeInternalLogicMetricEventMetricLimitVo {
	this := UnibeeInternalLogicMetricEventMetricLimitVo{}
	return &this
}

// NewUnibeeInternalLogicMetricEventMetricLimitVoWithDefaults instantiates a new UnibeeInternalLogicMetricEventMetricLimitVo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicMetricEventMetricLimitVoWithDefaults() *UnibeeInternalLogicMetricEventMetricLimitVo {
	this := UnibeeInternalLogicMetricEventMetricLimitVo{}
	return &this
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetMerchantId() int32 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int32
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetMerchantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int32 and assigns it to the MerchantId field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetMerchantId(v int32) {
	o.MerchantId = &v
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetMetricId() int64 {
	if o == nil || IsNil(o.MetricId) {
		var ret int64
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetMetricIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MetricId) {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasMetricId() bool {
	if o != nil && !IsNil(o.MetricId) {
		return true
	}

	return false
}

// SetMetricId gets a reference to the given int64 and assigns it to the MetricId field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetMetricId(v int64) {
	o.MetricId = &v
}

// GetPlanLimits returns the PlanLimits field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetPlanLimits() []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo {
	if o == nil || IsNil(o.PlanLimits) {
		var ret []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo
		return ret
	}
	return o.PlanLimits
}

// GetPlanLimitsOk returns a tuple with the PlanLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetPlanLimitsOk() ([]UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo, bool) {
	if o == nil || IsNil(o.PlanLimits) {
		return nil, false
	}
	return o.PlanLimits, true
}

// HasPlanLimits returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasPlanLimits() bool {
	if o != nil && !IsNil(o.PlanLimits) {
		return true
	}

	return false
}

// SetPlanLimits gets a reference to the given []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo and assigns it to the PlanLimits field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetPlanLimits(v []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo) {
	o.PlanLimits = v
}

// GetTotalLimit returns the TotalLimit field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetTotalLimit() int32 {
	if o == nil || IsNil(o.TotalLimit) {
		var ret int32
		return ret
	}
	return *o.TotalLimit
}

// GetTotalLimitOk returns a tuple with the TotalLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetTotalLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalLimit) {
		return nil, false
	}
	return o.TotalLimit, true
}

// HasTotalLimit returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasTotalLimit() bool {
	if o != nil && !IsNil(o.TotalLimit) {
		return true
	}

	return false
}

// SetTotalLimit gets a reference to the given int32 and assigns it to the TotalLimit field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetTotalLimit(v int32) {
	o.TotalLimit = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetUserId(v int64) {
	o.UserId = &v
}

// GetAggregationProperty returns the AggregationProperty field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetAggregationProperty() string {
	if o == nil || IsNil(o.AggregationProperty) {
		var ret string
		return ret
	}
	return *o.AggregationProperty
}

// GetAggregationPropertyOk returns a tuple with the AggregationProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetAggregationPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationProperty) {
		return nil, false
	}
	return o.AggregationProperty, true
}

// HasAggregationProperty returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasAggregationProperty() bool {
	if o != nil && !IsNil(o.AggregationProperty) {
		return true
	}

	return false
}

// SetAggregationProperty gets a reference to the given string and assigns it to the AggregationProperty field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetAggregationProperty(v string) {
	o.AggregationProperty = &v
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetAggregationType() int32 {
	if o == nil || IsNil(o.AggregationType) {
		var ret int32
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetAggregationTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given int32 and assigns it to the AggregationType field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetAggregationType(v int32) {
	o.AggregationType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetCode(v string) {
	o.Code = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetMetricName(v string) {
	o.MetricName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeInternalLogicMetricEventMetricLimitVo) SetType(v int32) {
	o.Type = &v
}

func (o UnibeeInternalLogicMetricEventMetricLimitVo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicMetricEventMetricLimitVo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantId) {
		toSerialize["MerchantId"] = o.MerchantId
	}
	if !IsNil(o.MetricId) {
		toSerialize["MetricId"] = o.MetricId
	}
	if !IsNil(o.PlanLimits) {
		toSerialize["PlanLimits"] = o.PlanLimits
	}
	if !IsNil(o.TotalLimit) {
		toSerialize["TotalLimit"] = o.TotalLimit
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if !IsNil(o.AggregationProperty) {
		toSerialize["aggregationProperty"] = o.AggregationProperty
	}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregationType"] = o.AggregationType
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicMetricEventMetricLimitVo struct {
	value *UnibeeInternalLogicMetricEventMetricLimitVo
	isSet bool
}

func (v NullableUnibeeInternalLogicMetricEventMetricLimitVo) Get() *UnibeeInternalLogicMetricEventMetricLimitVo {
	return v.value
}

func (v *NullableUnibeeInternalLogicMetricEventMetricLimitVo) Set(val *UnibeeInternalLogicMetricEventMetricLimitVo) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicMetricEventMetricLimitVo) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicMetricEventMetricLimitVo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicMetricEventMetricLimitVo(val *UnibeeInternalLogicMetricEventMetricLimitVo) *NullableUnibeeInternalLogicMetricEventMetricLimitVo {
	return &NullableUnibeeInternalLogicMetricEventMetricLimitVo{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicMetricEventMetricLimitVo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicMetricEventMetricLimitVo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

