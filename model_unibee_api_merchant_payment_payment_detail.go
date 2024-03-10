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

// checks if the UnibeeApiMerchantPaymentPaymentDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentPaymentDetail{}

// UnibeeApiMerchantPaymentPaymentDetail struct for UnibeeApiMerchantPaymentPaymentDetail
type UnibeeApiMerchantPaymentPaymentDetail struct {
	Payment *UnibeeInternalLogicGatewayRoPaymentSimplify `json:"payment,omitempty"`
	User *UnibeeInternalLogicGatewayRoUserAccountSimplify `json:"user,omitempty"`
}

// NewUnibeeApiMerchantPaymentPaymentDetail instantiates a new UnibeeApiMerchantPaymentPaymentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentPaymentDetail() *UnibeeApiMerchantPaymentPaymentDetail {
	this := UnibeeApiMerchantPaymentPaymentDetail{}
	return &this
}

// NewUnibeeApiMerchantPaymentPaymentDetailWithDefaults instantiates a new UnibeeApiMerchantPaymentPaymentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentPaymentDetailWithDefaults() *UnibeeApiMerchantPaymentPaymentDetail {
	this := UnibeeApiMerchantPaymentPaymentDetail{}
	return &this
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentPaymentDetail) GetPayment() UnibeeInternalLogicGatewayRoPaymentSimplify {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeInternalLogicGatewayRoPaymentSimplify
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentPaymentDetail) GetPaymentOk() (*UnibeeInternalLogicGatewayRoPaymentSimplify, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentPaymentDetail) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeInternalLogicGatewayRoPaymentSimplify and assigns it to the Payment field.
func (o *UnibeeApiMerchantPaymentPaymentDetail) SetPayment(v UnibeeInternalLogicGatewayRoPaymentSimplify) {
	o.Payment = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentPaymentDetail) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify {
	if o == nil || IsNil(o.User) {
		var ret UnibeeInternalLogicGatewayRoUserAccountSimplify
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentPaymentDetail) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentPaymentDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeInternalLogicGatewayRoUserAccountSimplify and assigns it to the User field.
func (o *UnibeeApiMerchantPaymentPaymentDetail) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify) {
	o.User = &v
}

func (o UnibeeApiMerchantPaymentPaymentDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentPaymentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentPaymentDetail struct {
	value *UnibeeApiMerchantPaymentPaymentDetail
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentPaymentDetail) Get() *UnibeeApiMerchantPaymentPaymentDetail {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentPaymentDetail) Set(val *UnibeeApiMerchantPaymentPaymentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentPaymentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentPaymentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentPaymentDetail(val *UnibeeApiMerchantPaymentPaymentDetail) *NullableUnibeeApiMerchantPaymentPaymentDetail {
	return &NullableUnibeeApiMerchantPaymentPaymentDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentPaymentDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentPaymentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

