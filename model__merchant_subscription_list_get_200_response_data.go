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

// checks if the MerchantSubscriptionListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionListGet200ResponseData{}

// MerchantSubscriptionListGet200ResponseData struct for MerchantSubscriptionListGet200ResponseData
type MerchantSubscriptionListGet200ResponseData struct {
	// Subscriptions
	Subscriptions []UnibeeApiBeanDetailSubscriptionDetail `json:"subscriptions,omitempty"`
}

// NewMerchantSubscriptionListGet200ResponseData instantiates a new MerchantSubscriptionListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionListGet200ResponseData() *MerchantSubscriptionListGet200ResponseData {
	this := MerchantSubscriptionListGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionListGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionListGet200ResponseDataWithDefaults() *MerchantSubscriptionListGet200ResponseData {
	this := MerchantSubscriptionListGet200ResponseData{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *MerchantSubscriptionListGet200ResponseData) GetSubscriptions() []UnibeeApiBeanDetailSubscriptionDetail {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []UnibeeApiBeanDetailSubscriptionDetail
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionListGet200ResponseData) GetSubscriptionsOk() ([]UnibeeApiBeanDetailSubscriptionDetail, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *MerchantSubscriptionListGet200ResponseData) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []UnibeeApiBeanDetailSubscriptionDetail and assigns it to the Subscriptions field.
func (o *MerchantSubscriptionListGet200ResponseData) SetSubscriptions(v []UnibeeApiBeanDetailSubscriptionDetail) {
	o.Subscriptions = v
}

func (o MerchantSubscriptionListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionListGet200ResponseData struct {
	value *MerchantSubscriptionListGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionListGet200ResponseData) Get() *MerchantSubscriptionListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionListGet200ResponseData) Set(val *MerchantSubscriptionListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionListGet200ResponseData(val *MerchantSubscriptionListGet200ResponseData) *NullableMerchantSubscriptionListGet200ResponseData {
	return &NullableMerchantSubscriptionListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


