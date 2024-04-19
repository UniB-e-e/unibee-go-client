/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404191455 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionCancelReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCancelReq{}

// UnibeeApiMerchantSubscriptionCancelReq Cancel subscription immediately, no proration invoice will generate
type UnibeeApiMerchantSubscriptionCancelReq struct {
	// Default false
	// Deprecated
	InvoiceNow *bool `json:"invoiceNow,omitempty"`
	// Prorate Generate Invoice，Default false
	// Deprecated
	Prorate *bool `json:"prorate,omitempty"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionCancelReq UnibeeApiMerchantSubscriptionCancelReq

// NewUnibeeApiMerchantSubscriptionCancelReq instantiates a new UnibeeApiMerchantSubscriptionCancelReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCancelReq(subscriptionId string) *UnibeeApiMerchantSubscriptionCancelReq {
	this := UnibeeApiMerchantSubscriptionCancelReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelReq {
	this := UnibeeApiMerchantSubscriptionCancelReq{}
	return &this
}

// GetInvoiceNow returns the InvoiceNow field value if set, zero value otherwise.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetInvoiceNow() bool {
	if o == nil || IsNil(o.InvoiceNow) {
		var ret bool
		return ret
	}
	return *o.InvoiceNow
}

// GetInvoiceNowOk returns a tuple with the InvoiceNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetInvoiceNowOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceNow) {
		return nil, false
	}
	return o.InvoiceNow, true
}

// HasInvoiceNow returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasInvoiceNow() bool {
	if o != nil && !IsNil(o.InvoiceNow) {
		return true
	}

	return false
}

// SetInvoiceNow gets a reference to the given bool and assigns it to the InvoiceNow field.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetInvoiceNow(v bool) {
	o.InvoiceNow = &v
}

// GetProrate returns the Prorate field value if set, zero value otherwise.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProrate() bool {
	if o == nil || IsNil(o.Prorate) {
		var ret bool
		return ret
	}
	return *o.Prorate
}

// GetProrateOk returns a tuple with the Prorate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProrateOk() (*bool, bool) {
	if o == nil || IsNil(o.Prorate) {
		return nil, false
	}
	return o.Prorate, true
}

// HasProrate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) HasProrate() bool {
	if o != nil && !IsNil(o.Prorate) {
		return true
	}

	return false
}

// SetProrate gets a reference to the given bool and assigns it to the Prorate field.
// Deprecated
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetProrate(v bool) {
	o.Prorate = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionCancelReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionCancelReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCancelReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceNow) {
		toSerialize["invoiceNow"] = o.InvoiceNow
	}
	if !IsNil(o.Prorate) {
		toSerialize["prorate"] = o.Prorate
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionCancelReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionCancelReq := _UnibeeApiMerchantSubscriptionCancelReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionCancelReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionCancelReq(varUnibeeApiMerchantSubscriptionCancelReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionCancelReq struct {
	value *UnibeeApiMerchantSubscriptionCancelReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCancelReq) Get() *UnibeeApiMerchantSubscriptionCancelReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelReq) Set(val *UnibeeApiMerchantSubscriptionCancelReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCancelReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCancelReq(val *UnibeeApiMerchantSubscriptionCancelReq) *NullableUnibeeApiMerchantSubscriptionCancelReq {
	return &NullableUnibeeApiMerchantSubscriptionCancelReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCancelReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


