/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404171839 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantWebhookEndpointLogListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookEndpointLogListRes{}

// UnibeeApiMerchantWebhookEndpointLogListRes struct for UnibeeApiMerchantWebhookEndpointLogListRes
type UnibeeApiMerchantWebhookEndpointLogListRes struct {
	// EndpointLogList
	EndpointLogList []UnibeeApiBeanMerchantWebhookLogSimplify `json:"endpointLogList,omitempty"`
}

// NewUnibeeApiMerchantWebhookEndpointLogListRes instantiates a new UnibeeApiMerchantWebhookEndpointLogListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookEndpointLogListRes() *UnibeeApiMerchantWebhookEndpointLogListRes {
	this := UnibeeApiMerchantWebhookEndpointLogListRes{}
	return &this
}

// NewUnibeeApiMerchantWebhookEndpointLogListResWithDefaults instantiates a new UnibeeApiMerchantWebhookEndpointLogListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookEndpointLogListResWithDefaults() *UnibeeApiMerchantWebhookEndpointLogListRes {
	this := UnibeeApiMerchantWebhookEndpointLogListRes{}
	return &this
}

// GetEndpointLogList returns the EndpointLogList field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookEndpointLogListRes) GetEndpointLogList() []UnibeeApiBeanMerchantWebhookLogSimplify {
	if o == nil || IsNil(o.EndpointLogList) {
		var ret []UnibeeApiBeanMerchantWebhookLogSimplify
		return ret
	}
	return o.EndpointLogList
}

// GetEndpointLogListOk returns a tuple with the EndpointLogList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListRes) GetEndpointLogListOk() ([]UnibeeApiBeanMerchantWebhookLogSimplify, bool) {
	if o == nil || IsNil(o.EndpointLogList) {
		return nil, false
	}
	return o.EndpointLogList, true
}

// HasEndpointLogList returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookEndpointLogListRes) HasEndpointLogList() bool {
	if o != nil && !IsNil(o.EndpointLogList) {
		return true
	}

	return false
}

// SetEndpointLogList gets a reference to the given []UnibeeApiBeanMerchantWebhookLogSimplify and assigns it to the EndpointLogList field.
func (o *UnibeeApiMerchantWebhookEndpointLogListRes) SetEndpointLogList(v []UnibeeApiBeanMerchantWebhookLogSimplify) {
	o.EndpointLogList = v
}

func (o UnibeeApiMerchantWebhookEndpointLogListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookEndpointLogListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointLogList) {
		toSerialize["endpointLogList"] = o.EndpointLogList
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantWebhookEndpointLogListRes struct {
	value *UnibeeApiMerchantWebhookEndpointLogListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookEndpointLogListRes) Get() *UnibeeApiMerchantWebhookEndpointLogListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookEndpointLogListRes) Set(val *UnibeeApiMerchantWebhookEndpointLogListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookEndpointLogListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookEndpointLogListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookEndpointLogListRes(val *UnibeeApiMerchantWebhookEndpointLogListRes) *NullableUnibeeApiMerchantWebhookEndpointLogListRes {
	return &NullableUnibeeApiMerchantWebhookEndpointLogListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookEndpointLogListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookEndpointLogListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


