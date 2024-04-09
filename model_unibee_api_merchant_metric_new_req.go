/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: buildtime:202404090336 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricNewReq{}

// UnibeeApiMerchantMetricNewReq struct for UnibeeApiMerchantMetricNewReq
type UnibeeApiMerchantMetricNewReq struct {
	// AggregationProperty, Will Needed When AggregationType != count
	AggregationProperty *string `json:"aggregationProperty,omitempty"`
	// AggregationType,1-count，2-count unique, 3-latest, 4-max, 5-sum
	AggregationType *int32 `json:"aggregationType,omitempty"`
	// Code
	Code string `json:"code"`
	// MetricDescription
	MetricDescription *string `json:"metricDescription,omitempty"`
	// MetricName
	MetricName string `json:"metricName"`
}

type _UnibeeApiMerchantMetricNewReq UnibeeApiMerchantMetricNewReq

// NewUnibeeApiMerchantMetricNewReq instantiates a new UnibeeApiMerchantMetricNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricNewReq(code string, metricName string) *UnibeeApiMerchantMetricNewReq {
	this := UnibeeApiMerchantMetricNewReq{}
	this.Code = code
	this.MetricName = metricName
	return &this
}

// NewUnibeeApiMerchantMetricNewReqWithDefaults instantiates a new UnibeeApiMerchantMetricNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricNewReqWithDefaults() *UnibeeApiMerchantMetricNewReq {
	this := UnibeeApiMerchantMetricNewReq{}
	return &this
}

// GetAggregationProperty returns the AggregationProperty field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewReq) GetAggregationProperty() string {
	if o == nil || IsNil(o.AggregationProperty) {
		var ret string
		return ret
	}
	return *o.AggregationProperty
}

// GetAggregationPropertyOk returns a tuple with the AggregationProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewReq) GetAggregationPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationProperty) {
		return nil, false
	}
	return o.AggregationProperty, true
}

// HasAggregationProperty returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewReq) HasAggregationProperty() bool {
	if o != nil && !IsNil(o.AggregationProperty) {
		return true
	}

	return false
}

// SetAggregationProperty gets a reference to the given string and assigns it to the AggregationProperty field.
func (o *UnibeeApiMerchantMetricNewReq) SetAggregationProperty(v string) {
	o.AggregationProperty = &v
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewReq) GetAggregationType() int32 {
	if o == nil || IsNil(o.AggregationType) {
		var ret int32
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewReq) GetAggregationTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewReq) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given int32 and assigns it to the AggregationType field.
func (o *UnibeeApiMerchantMetricNewReq) SetAggregationType(v int32) {
	o.AggregationType = &v
}

// GetCode returns the Code field value
func (o *UnibeeApiMerchantMetricNewReq) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewReq) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UnibeeApiMerchantMetricNewReq) SetCode(v string) {
	o.Code = v
}

// GetMetricDescription returns the MetricDescription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewReq) GetMetricDescription() string {
	if o == nil || IsNil(o.MetricDescription) {
		var ret string
		return ret
	}
	return *o.MetricDescription
}

// GetMetricDescriptionOk returns a tuple with the MetricDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewReq) GetMetricDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetricDescription) {
		return nil, false
	}
	return o.MetricDescription, true
}

// HasMetricDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewReq) HasMetricDescription() bool {
	if o != nil && !IsNil(o.MetricDescription) {
		return true
	}

	return false
}

// SetMetricDescription gets a reference to the given string and assigns it to the MetricDescription field.
func (o *UnibeeApiMerchantMetricNewReq) SetMetricDescription(v string) {
	o.MetricDescription = &v
}

// GetMetricName returns the MetricName field value
func (o *UnibeeApiMerchantMetricNewReq) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewReq) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *UnibeeApiMerchantMetricNewReq) SetMetricName(v string) {
	o.MetricName = v
}

func (o UnibeeApiMerchantMetricNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationProperty) {
		toSerialize["aggregationProperty"] = o.AggregationProperty
	}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregationType"] = o.AggregationType
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.MetricDescription) {
		toSerialize["metricDescription"] = o.MetricDescription
	}
	toSerialize["metricName"] = o.MetricName
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"metricName",
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

	varUnibeeApiMerchantMetricNewReq := _UnibeeApiMerchantMetricNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricNewReq(varUnibeeApiMerchantMetricNewReq)

	return err
}

type NullableUnibeeApiMerchantMetricNewReq struct {
	value *UnibeeApiMerchantMetricNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricNewReq) Get() *UnibeeApiMerchantMetricNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricNewReq) Set(val *UnibeeApiMerchantMetricNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricNewReq(val *UnibeeApiMerchantMetricNewReq) *NullableUnibeeApiMerchantMetricNewReq {
	return &NullableUnibeeApiMerchantMetricNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


