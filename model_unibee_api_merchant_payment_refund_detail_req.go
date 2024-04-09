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

// checks if the UnibeeApiMerchantPaymentRefundDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentRefundDetailReq{}

// UnibeeApiMerchantPaymentRefundDetailReq struct for UnibeeApiMerchantPaymentRefundDetailReq
type UnibeeApiMerchantPaymentRefundDetailReq struct {
	// RefundId
	RefundId *string `json:"refundId,omitempty"`
}

// NewUnibeeApiMerchantPaymentRefundDetailReq instantiates a new UnibeeApiMerchantPaymentRefundDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentRefundDetailReq() *UnibeeApiMerchantPaymentRefundDetailReq {
	this := UnibeeApiMerchantPaymentRefundDetailReq{}
	return &this
}

// NewUnibeeApiMerchantPaymentRefundDetailReqWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentRefundDetailReqWithDefaults() *UnibeeApiMerchantPaymentRefundDetailReq {
	this := UnibeeApiMerchantPaymentRefundDetailReq{}
	return &this
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundDetailReq) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundDetailReq) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundDetailReq) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *UnibeeApiMerchantPaymentRefundDetailReq) SetRefundId(v string) {
	o.RefundId = &v
}

func (o UnibeeApiMerchantPaymentRefundDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentRefundDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundId) {
		toSerialize["refundId"] = o.RefundId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentRefundDetailReq struct {
	value *UnibeeApiMerchantPaymentRefundDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentRefundDetailReq) Get() *UnibeeApiMerchantPaymentRefundDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentRefundDetailReq) Set(val *UnibeeApiMerchantPaymentRefundDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentRefundDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentRefundDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentRefundDetailReq(val *UnibeeApiMerchantPaymentRefundDetailReq) *NullableUnibeeApiMerchantPaymentRefundDetailReq {
	return &NullableUnibeeApiMerchantPaymentRefundDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentRefundDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentRefundDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


