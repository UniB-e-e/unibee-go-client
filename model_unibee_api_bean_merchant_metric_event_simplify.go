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

// checks if the UnibeeApiBeanMerchantMetricEventSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantMetricEventSimplify{}

// UnibeeApiBeanMerchantMetricEventSimplify struct for UnibeeApiBeanMerchantMetricEventSimplify
type UnibeeApiBeanMerchantMetricEventSimplify struct {
	// aggregation property data (Json)
	AggregationPropertyData *string `json:"aggregationPropertyData,omitempty"`
	// aggregation property int, use for metric of max|sum type
	AggregationPropertyInt *int64 `json:"aggregationPropertyInt,omitempty"`
	// aggregation property string, use for metric of count|count_unique type
	AggregationPropertyString *string `json:"aggregationPropertyString,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// external_event_id, should be unique
	ExternalEventId *string `json:"externalEventId,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	// merchantId
	MerchantId *int64 `json:"merchantId,omitempty"`
	// metric_id
	MetricId *int64 `json:"metricId,omitempty"`
	MetricLimit *int64 `json:"metricLimit,omitempty"`
	SubscriptionIds *string `json:"subscriptionIds,omitempty"`
	// matched subscription's current_period_end
	SubscriptionPeriodEnd *int64 `json:"subscriptionPeriodEnd,omitempty"`
	// matched subscription's current_period_start
	SubscriptionPeriodStart *int64 `json:"subscriptionPeriodStart,omitempty"`
	Used *int64 `json:"used,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanMerchantMetricEventSimplify instantiates a new UnibeeApiBeanMerchantMetricEventSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantMetricEventSimplify() *UnibeeApiBeanMerchantMetricEventSimplify {
	this := UnibeeApiBeanMerchantMetricEventSimplify{}
	return &this
}

// NewUnibeeApiBeanMerchantMetricEventSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantMetricEventSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantMetricEventSimplifyWithDefaults() *UnibeeApiBeanMerchantMetricEventSimplify {
	this := UnibeeApiBeanMerchantMetricEventSimplify{}
	return &this
}

// GetAggregationPropertyData returns the AggregationPropertyData field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyData() string {
	if o == nil || IsNil(o.AggregationPropertyData) {
		var ret string
		return ret
	}
	return *o.AggregationPropertyData
}

// GetAggregationPropertyDataOk returns a tuple with the AggregationPropertyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyDataOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationPropertyData) {
		return nil, false
	}
	return o.AggregationPropertyData, true
}

// HasAggregationPropertyData returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasAggregationPropertyData() bool {
	if o != nil && !IsNil(o.AggregationPropertyData) {
		return true
	}

	return false
}

// SetAggregationPropertyData gets a reference to the given string and assigns it to the AggregationPropertyData field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetAggregationPropertyData(v string) {
	o.AggregationPropertyData = &v
}

// GetAggregationPropertyInt returns the AggregationPropertyInt field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyInt() int64 {
	if o == nil || IsNil(o.AggregationPropertyInt) {
		var ret int64
		return ret
	}
	return *o.AggregationPropertyInt
}

// GetAggregationPropertyIntOk returns a tuple with the AggregationPropertyInt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyIntOk() (*int64, bool) {
	if o == nil || IsNil(o.AggregationPropertyInt) {
		return nil, false
	}
	return o.AggregationPropertyInt, true
}

// HasAggregationPropertyInt returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasAggregationPropertyInt() bool {
	if o != nil && !IsNil(o.AggregationPropertyInt) {
		return true
	}

	return false
}

// SetAggregationPropertyInt gets a reference to the given int64 and assigns it to the AggregationPropertyInt field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetAggregationPropertyInt(v int64) {
	o.AggregationPropertyInt = &v
}

// GetAggregationPropertyString returns the AggregationPropertyString field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyString() string {
	if o == nil || IsNil(o.AggregationPropertyString) {
		var ret string
		return ret
	}
	return *o.AggregationPropertyString
}

// GetAggregationPropertyStringOk returns a tuple with the AggregationPropertyString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyStringOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationPropertyString) {
		return nil, false
	}
	return o.AggregationPropertyString, true
}

// HasAggregationPropertyString returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasAggregationPropertyString() bool {
	if o != nil && !IsNil(o.AggregationPropertyString) {
		return true
	}

	return false
}

// SetAggregationPropertyString gets a reference to the given string and assigns it to the AggregationPropertyString field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetAggregationPropertyString(v string) {
	o.AggregationPropertyString = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetExternalEventId returns the ExternalEventId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetExternalEventId() string {
	if o == nil || IsNil(o.ExternalEventId) {
		var ret string
		return ret
	}
	return *o.ExternalEventId
}

// GetExternalEventIdOk returns a tuple with the ExternalEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetExternalEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalEventId) {
		return nil, false
	}
	return o.ExternalEventId, true
}

// HasExternalEventId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasExternalEventId() bool {
	if o != nil && !IsNil(o.ExternalEventId) {
		return true
	}

	return false
}

// SetExternalEventId gets a reference to the given string and assigns it to the ExternalEventId field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetExternalEventId(v string) {
	o.ExternalEventId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricId() int64 {
	if o == nil || IsNil(o.MetricId) {
		var ret int64
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MetricId) {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasMetricId() bool {
	if o != nil && !IsNil(o.MetricId) {
		return true
	}

	return false
}

// SetMetricId gets a reference to the given int64 and assigns it to the MetricId field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetMetricId(v int64) {
	o.MetricId = &v
}

// GetMetricLimit returns the MetricLimit field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricLimit() int64 {
	if o == nil || IsNil(o.MetricLimit) {
		var ret int64
		return ret
	}
	return *o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MetricLimit) {
		return nil, false
	}
	return o.MetricLimit, true
}

// HasMetricLimit returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasMetricLimit() bool {
	if o != nil && !IsNil(o.MetricLimit) {
		return true
	}

	return false
}

// SetMetricLimit gets a reference to the given int64 and assigns it to the MetricLimit field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetMetricLimit(v int64) {
	o.MetricLimit = &v
}

// GetSubscriptionIds returns the SubscriptionIds field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionIds() string {
	if o == nil || IsNil(o.SubscriptionIds) {
		var ret string
		return ret
	}
	return *o.SubscriptionIds
}

// GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionIdsOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionIds) {
		return nil, false
	}
	return o.SubscriptionIds, true
}

// HasSubscriptionIds returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasSubscriptionIds() bool {
	if o != nil && !IsNil(o.SubscriptionIds) {
		return true
	}

	return false
}

// SetSubscriptionIds gets a reference to the given string and assigns it to the SubscriptionIds field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetSubscriptionIds(v string) {
	o.SubscriptionIds = &v
}

// GetSubscriptionPeriodEnd returns the SubscriptionPeriodEnd field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodEnd() int64 {
	if o == nil || IsNil(o.SubscriptionPeriodEnd) {
		var ret int64
		return ret
	}
	return *o.SubscriptionPeriodEnd
}

// GetSubscriptionPeriodEndOk returns a tuple with the SubscriptionPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodEndOk() (*int64, bool) {
	if o == nil || IsNil(o.SubscriptionPeriodEnd) {
		return nil, false
	}
	return o.SubscriptionPeriodEnd, true
}

// HasSubscriptionPeriodEnd returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasSubscriptionPeriodEnd() bool {
	if o != nil && !IsNil(o.SubscriptionPeriodEnd) {
		return true
	}

	return false
}

// SetSubscriptionPeriodEnd gets a reference to the given int64 and assigns it to the SubscriptionPeriodEnd field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetSubscriptionPeriodEnd(v int64) {
	o.SubscriptionPeriodEnd = &v
}

// GetSubscriptionPeriodStart returns the SubscriptionPeriodStart field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodStart() int64 {
	if o == nil || IsNil(o.SubscriptionPeriodStart) {
		var ret int64
		return ret
	}
	return *o.SubscriptionPeriodStart
}

// GetSubscriptionPeriodStartOk returns a tuple with the SubscriptionPeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodStartOk() (*int64, bool) {
	if o == nil || IsNil(o.SubscriptionPeriodStart) {
		return nil, false
	}
	return o.SubscriptionPeriodStart, true
}

// HasSubscriptionPeriodStart returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasSubscriptionPeriodStart() bool {
	if o != nil && !IsNil(o.SubscriptionPeriodStart) {
		return true
	}

	return false
}

// SetSubscriptionPeriodStart gets a reference to the given int64 and assigns it to the SubscriptionPeriodStart field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetSubscriptionPeriodStart(v int64) {
	o.SubscriptionPeriodStart = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUsed() int64 {
	if o == nil || IsNil(o.Used) {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUsedOk() (*int64, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetUsed(v int64) {
	o.Used = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanMerchantMetricEventSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantMetricEventSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationPropertyData) {
		toSerialize["aggregationPropertyData"] = o.AggregationPropertyData
	}
	if !IsNil(o.AggregationPropertyInt) {
		toSerialize["aggregationPropertyInt"] = o.AggregationPropertyInt
	}
	if !IsNil(o.AggregationPropertyString) {
		toSerialize["aggregationPropertyString"] = o.AggregationPropertyString
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.ExternalEventId) {
		toSerialize["externalEventId"] = o.ExternalEventId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.MetricId) {
		toSerialize["metricId"] = o.MetricId
	}
	if !IsNil(o.MetricLimit) {
		toSerialize["metricLimit"] = o.MetricLimit
	}
	if !IsNil(o.SubscriptionIds) {
		toSerialize["subscriptionIds"] = o.SubscriptionIds
	}
	if !IsNil(o.SubscriptionPeriodEnd) {
		toSerialize["subscriptionPeriodEnd"] = o.SubscriptionPeriodEnd
	}
	if !IsNil(o.SubscriptionPeriodStart) {
		toSerialize["subscriptionPeriodStart"] = o.SubscriptionPeriodStart
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantMetricEventSimplify struct {
	value *UnibeeApiBeanMerchantMetricEventSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantMetricEventSimplify) Get() *UnibeeApiBeanMerchantMetricEventSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantMetricEventSimplify) Set(val *UnibeeApiBeanMerchantMetricEventSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantMetricEventSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantMetricEventSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantMetricEventSimplify(val *UnibeeApiBeanMerchantMetricEventSimplify) *NullableUnibeeApiBeanMerchantMetricEventSimplify {
	return &NullableUnibeeApiBeanMerchantMetricEventSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantMetricEventSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantMetricEventSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

