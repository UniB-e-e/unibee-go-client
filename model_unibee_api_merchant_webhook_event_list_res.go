/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: buildtime:202404090336 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantWebhookEventListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookEventListRes{}

// UnibeeApiMerchantWebhookEventListRes struct for UnibeeApiMerchantWebhookEventListRes
type UnibeeApiMerchantWebhookEventListRes struct {
	// EventList
	EventList []string `json:"eventList,omitempty"`
}

// NewUnibeeApiMerchantWebhookEventListRes instantiates a new UnibeeApiMerchantWebhookEventListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookEventListRes() *UnibeeApiMerchantWebhookEventListRes {
	this := UnibeeApiMerchantWebhookEventListRes{}
	return &this
}

// NewUnibeeApiMerchantWebhookEventListResWithDefaults instantiates a new UnibeeApiMerchantWebhookEventListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookEventListResWithDefaults() *UnibeeApiMerchantWebhookEventListRes {
	this := UnibeeApiMerchantWebhookEventListRes{}
	return &this
}

// GetEventList returns the EventList field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookEventListRes) GetEventList() []string {
	if o == nil || IsNil(o.EventList) {
		var ret []string
		return ret
	}
	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookEventListRes) GetEventListOk() ([]string, bool) {
	if o == nil || IsNil(o.EventList) {
		return nil, false
	}
	return o.EventList, true
}

// HasEventList returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookEventListRes) HasEventList() bool {
	if o != nil && !IsNil(o.EventList) {
		return true
	}

	return false
}

// SetEventList gets a reference to the given []string and assigns it to the EventList field.
func (o *UnibeeApiMerchantWebhookEventListRes) SetEventList(v []string) {
	o.EventList = v
}

func (o UnibeeApiMerchantWebhookEventListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookEventListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventList) {
		toSerialize["eventList"] = o.EventList
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantWebhookEventListRes struct {
	value *UnibeeApiMerchantWebhookEventListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookEventListRes) Get() *UnibeeApiMerchantWebhookEventListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookEventListRes) Set(val *UnibeeApiMerchantWebhookEventListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookEventListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookEventListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookEventListRes(val *UnibeeApiMerchantWebhookEventListRes) *NullableUnibeeApiMerchantWebhookEventListRes {
	return &NullableUnibeeApiMerchantWebhookEventListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookEventListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookEventListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


