/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404171839 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiSystemPaymentPaymentCallbackAgainReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemPaymentPaymentCallbackAgainReq{}

// UnibeeApiSystemPaymentPaymentCallbackAgainReq struct for UnibeeApiSystemPaymentPaymentCallbackAgainReq
type UnibeeApiSystemPaymentPaymentCallbackAgainReq struct {
	// PaymentId
	PaymentId string `json:"paymentId"`
}

type _UnibeeApiSystemPaymentPaymentCallbackAgainReq UnibeeApiSystemPaymentPaymentCallbackAgainReq

// NewUnibeeApiSystemPaymentPaymentCallbackAgainReq instantiates a new UnibeeApiSystemPaymentPaymentCallbackAgainReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemPaymentPaymentCallbackAgainReq(paymentId string) *UnibeeApiSystemPaymentPaymentCallbackAgainReq {
	this := UnibeeApiSystemPaymentPaymentCallbackAgainReq{}
	this.PaymentId = paymentId
	return &this
}

// NewUnibeeApiSystemPaymentPaymentCallbackAgainReqWithDefaults instantiates a new UnibeeApiSystemPaymentPaymentCallbackAgainReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemPaymentPaymentCallbackAgainReqWithDefaults() *UnibeeApiSystemPaymentPaymentCallbackAgainReq {
	this := UnibeeApiSystemPaymentPaymentCallbackAgainReq{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *UnibeeApiSystemPaymentPaymentCallbackAgainReq) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemPaymentPaymentCallbackAgainReq) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *UnibeeApiSystemPaymentPaymentCallbackAgainReq) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o UnibeeApiSystemPaymentPaymentCallbackAgainReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemPaymentPaymentCallbackAgainReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentId"] = o.PaymentId
	return toSerialize, nil
}

func (o *UnibeeApiSystemPaymentPaymentCallbackAgainReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"paymentId",
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

	varUnibeeApiSystemPaymentPaymentCallbackAgainReq := _UnibeeApiSystemPaymentPaymentCallbackAgainReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemPaymentPaymentCallbackAgainReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemPaymentPaymentCallbackAgainReq(varUnibeeApiSystemPaymentPaymentCallbackAgainReq)

	return err
}

type NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq struct {
	value *UnibeeApiSystemPaymentPaymentCallbackAgainReq
	isSet bool
}

func (v NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq) Get() *UnibeeApiSystemPaymentPaymentCallbackAgainReq {
	return v.value
}

func (v *NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq) Set(val *UnibeeApiSystemPaymentPaymentCallbackAgainReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemPaymentPaymentCallbackAgainReq(val *UnibeeApiSystemPaymentPaymentCallbackAgainReq) *NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq {
	return &NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemPaymentPaymentCallbackAgainReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


