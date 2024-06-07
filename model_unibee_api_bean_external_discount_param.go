/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406070109 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanExternalDiscountParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanExternalDiscountParam{}

// UnibeeApiBeanExternalDiscountParam struct for UnibeeApiBeanExternalDiscountParam
type UnibeeApiBeanExternalDiscountParam struct {
	// the count limitation of subscription recurring cycle, recurring need enable if cycleLimit set
	CycleLimit *int32 `json:"cycleLimit,omitempty"`
	// Amount of discount
	DiscountAmount *int32 `json:"discountAmount,omitempty"`
	// Percentage of discount, 100=1%, ignore if discountAmount set
	DiscountPercentage *int32 `json:"discountPercentage,omitempty"`
	// end of discount available utc time
	EndTime *int32 `json:"endTime,omitempty"`
	// Metadata，Map
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Discount recurring enable, default false
	Recurring *bool `json:"recurring,omitempty"`
}

// NewUnibeeApiBeanExternalDiscountParam instantiates a new UnibeeApiBeanExternalDiscountParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanExternalDiscountParam() *UnibeeApiBeanExternalDiscountParam {
	this := UnibeeApiBeanExternalDiscountParam{}
	return &this
}

// NewUnibeeApiBeanExternalDiscountParamWithDefaults instantiates a new UnibeeApiBeanExternalDiscountParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanExternalDiscountParamWithDefaults() *UnibeeApiBeanExternalDiscountParam {
	this := UnibeeApiBeanExternalDiscountParam{}
	return &this
}

// GetCycleLimit returns the CycleLimit field value if set, zero value otherwise.
func (o *UnibeeApiBeanExternalDiscountParam) GetCycleLimit() int32 {
	if o == nil || IsNil(o.CycleLimit) {
		var ret int32
		return ret
	}
	return *o.CycleLimit
}

// GetCycleLimitOk returns a tuple with the CycleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanExternalDiscountParam) GetCycleLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.CycleLimit) {
		return nil, false
	}
	return o.CycleLimit, true
}

// HasCycleLimit returns a boolean if a field has been set.
func (o *UnibeeApiBeanExternalDiscountParam) HasCycleLimit() bool {
	if o != nil && !IsNil(o.CycleLimit) {
		return true
	}

	return false
}

// SetCycleLimit gets a reference to the given int32 and assigns it to the CycleLimit field.
func (o *UnibeeApiBeanExternalDiscountParam) SetCycleLimit(v int32) {
	o.CycleLimit = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountAmount() int32 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int32
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanExternalDiscountParam) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int32 and assigns it to the DiscountAmount field.
func (o *UnibeeApiBeanExternalDiscountParam) SetDiscountAmount(v int32) {
	o.DiscountAmount = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountPercentage() int32 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int32
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanExternalDiscountParam) GetDiscountPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *UnibeeApiBeanExternalDiscountParam) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int32 and assigns it to the DiscountPercentage field.
func (o *UnibeeApiBeanExternalDiscountParam) SetDiscountPercentage(v int32) {
	o.DiscountPercentage = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanExternalDiscountParam) GetEndTime() int32 {
	if o == nil || IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanExternalDiscountParam) GetEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanExternalDiscountParam) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *UnibeeApiBeanExternalDiscountParam) SetEndTime(v int32) {
	o.EndTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanExternalDiscountParam) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanExternalDiscountParam) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanExternalDiscountParam) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiBeanExternalDiscountParam) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *UnibeeApiBeanExternalDiscountParam) GetRecurring() bool {
	if o == nil || IsNil(o.Recurring) {
		var ret bool
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanExternalDiscountParam) GetRecurringOk() (*bool, bool) {
	if o == nil || IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *UnibeeApiBeanExternalDiscountParam) HasRecurring() bool {
	if o != nil && !IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given bool and assigns it to the Recurring field.
func (o *UnibeeApiBeanExternalDiscountParam) SetRecurring(v bool) {
	o.Recurring = &v
}

func (o UnibeeApiBeanExternalDiscountParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanExternalDiscountParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CycleLimit) {
		toSerialize["cycleLimit"] = o.CycleLimit
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanExternalDiscountParam struct {
	value *UnibeeApiBeanExternalDiscountParam
	isSet bool
}

func (v NullableUnibeeApiBeanExternalDiscountParam) Get() *UnibeeApiBeanExternalDiscountParam {
	return v.value
}

func (v *NullableUnibeeApiBeanExternalDiscountParam) Set(val *UnibeeApiBeanExternalDiscountParam) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanExternalDiscountParam) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanExternalDiscountParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanExternalDiscountParam(val *UnibeeApiBeanExternalDiscountParam) *NullableUnibeeApiBeanExternalDiscountParam {
	return &NullableUnibeeApiBeanExternalDiscountParam{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanExternalDiscountParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanExternalDiscountParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


