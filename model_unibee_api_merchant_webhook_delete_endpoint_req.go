/*
OpenAPI UniBee

This is UniBee api server

API version: buildtime:202404131246 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantWebhookDeleteEndpointReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookDeleteEndpointReq{}

// UnibeeApiMerchantWebhookDeleteEndpointReq struct for UnibeeApiMerchantWebhookDeleteEndpointReq
type UnibeeApiMerchantWebhookDeleteEndpointReq struct {
	// EndpointId
	EndpointId int64 `json:"endpointId"`
}

type _UnibeeApiMerchantWebhookDeleteEndpointReq UnibeeApiMerchantWebhookDeleteEndpointReq

// NewUnibeeApiMerchantWebhookDeleteEndpointReq instantiates a new UnibeeApiMerchantWebhookDeleteEndpointReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookDeleteEndpointReq(endpointId int64) *UnibeeApiMerchantWebhookDeleteEndpointReq {
	this := UnibeeApiMerchantWebhookDeleteEndpointReq{}
	this.EndpointId = endpointId
	return &this
}

// NewUnibeeApiMerchantWebhookDeleteEndpointReqWithDefaults instantiates a new UnibeeApiMerchantWebhookDeleteEndpointReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookDeleteEndpointReqWithDefaults() *UnibeeApiMerchantWebhookDeleteEndpointReq {
	this := UnibeeApiMerchantWebhookDeleteEndpointReq{}
	return &this
}

// GetEndpointId returns the EndpointId field value
func (o *UnibeeApiMerchantWebhookDeleteEndpointReq) GetEndpointId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookDeleteEndpointReq) GetEndpointIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *UnibeeApiMerchantWebhookDeleteEndpointReq) SetEndpointId(v int64) {
	o.EndpointId = v
}

func (o UnibeeApiMerchantWebhookDeleteEndpointReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookDeleteEndpointReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpointId"] = o.EndpointId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantWebhookDeleteEndpointReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpointId",
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

	varUnibeeApiMerchantWebhookDeleteEndpointReq := _UnibeeApiMerchantWebhookDeleteEndpointReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantWebhookDeleteEndpointReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantWebhookDeleteEndpointReq(varUnibeeApiMerchantWebhookDeleteEndpointReq)

	return err
}

type NullableUnibeeApiMerchantWebhookDeleteEndpointReq struct {
	value *UnibeeApiMerchantWebhookDeleteEndpointReq
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookDeleteEndpointReq) Get() *UnibeeApiMerchantWebhookDeleteEndpointReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookDeleteEndpointReq) Set(val *UnibeeApiMerchantWebhookDeleteEndpointReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookDeleteEndpointReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookDeleteEndpointReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookDeleteEndpointReq(val *UnibeeApiMerchantWebhookDeleteEndpointReq) *NullableUnibeeApiMerchantWebhookDeleteEndpointReq {
	return &NullableUnibeeApiMerchantWebhookDeleteEndpointReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookDeleteEndpointReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookDeleteEndpointReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


