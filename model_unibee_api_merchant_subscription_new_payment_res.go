/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionNewPaymentRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionNewPaymentRes{}

// UnibeeApiMerchantSubscriptionNewPaymentRes struct for UnibeeApiMerchantSubscriptionNewPaymentRes
type UnibeeApiMerchantSubscriptionNewPaymentRes struct {
	Action map[string]interface{} `json:"action,omitempty"`
	// The external unique id of payment
	ExternalPaymentId *string `json:"externalPaymentId,omitempty"`
	Link *string `json:"link,omitempty"`
	// The unique id of payment
	PaymentId *string `json:"paymentId,omitempty"`
	// Status, 10-Created|20-Success|30-Failed|40-Cancelled
	Status *int32 `json:"status,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionNewPaymentRes instantiates a new UnibeeApiMerchantSubscriptionNewPaymentRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionNewPaymentRes() *UnibeeApiMerchantSubscriptionNewPaymentRes {
	this := UnibeeApiMerchantSubscriptionNewPaymentRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionNewPaymentResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionNewPaymentRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionNewPaymentResWithDefaults() *UnibeeApiMerchantSubscriptionNewPaymentRes {
	this := UnibeeApiMerchantSubscriptionNewPaymentRes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetAction() map[string]interface{} {
	if o == nil || IsNil(o.Action) {
		var ret map[string]interface{}
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetActionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Action) {
		return map[string]interface{}{}, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given map[string]interface{} and assigns it to the Action field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetAction(v map[string]interface{}) {
	o.Action = v
}

// GetExternalPaymentId returns the ExternalPaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetExternalPaymentId() string {
	if o == nil || IsNil(o.ExternalPaymentId) {
		var ret string
		return ret
	}
	return *o.ExternalPaymentId
}

// GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetExternalPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPaymentId) {
		return nil, false
	}
	return o.ExternalPaymentId, true
}

// HasExternalPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasExternalPaymentId() bool {
	if o != nil && !IsNil(o.ExternalPaymentId) {
		return true
	}

	return false
}

// SetExternalPaymentId gets a reference to the given string and assigns it to the ExternalPaymentId field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetExternalPaymentId(v string) {
	o.ExternalPaymentId = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetLink(v string) {
	o.Link = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantSubscriptionNewPaymentRes) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiMerchantSubscriptionNewPaymentRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionNewPaymentRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ExternalPaymentId) {
		toSerialize["externalPaymentId"] = o.ExternalPaymentId
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionNewPaymentRes struct {
	value *UnibeeApiMerchantSubscriptionNewPaymentRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionNewPaymentRes) Get() *UnibeeApiMerchantSubscriptionNewPaymentRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionNewPaymentRes) Set(val *UnibeeApiMerchantSubscriptionNewPaymentRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionNewPaymentRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionNewPaymentRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionNewPaymentRes(val *UnibeeApiMerchantSubscriptionNewPaymentRes) *NullableUnibeeApiMerchantSubscriptionNewPaymentRes {
	return &NullableUnibeeApiMerchantSubscriptionNewPaymentRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionNewPaymentRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionNewPaymentRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

