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

// checks if the MerchantSubscriptionPendingUpdateListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionPendingUpdateListGet200ResponseData{}

// MerchantSubscriptionPendingUpdateListGet200ResponseData struct for MerchantSubscriptionPendingUpdateListGet200ResponseData
type MerchantSubscriptionPendingUpdateListGet200ResponseData struct {
	// SubscriptionPendingUpdateDetails
	SubscriptionPendingUpdateDetails []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"subscriptionPendingUpdateDetails,omitempty"`
}

// NewMerchantSubscriptionPendingUpdateListGet200ResponseData instantiates a new MerchantSubscriptionPendingUpdateListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionPendingUpdateListGet200ResponseData() *MerchantSubscriptionPendingUpdateListGet200ResponseData {
	this := MerchantSubscriptionPendingUpdateListGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionPendingUpdateListGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionPendingUpdateListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionPendingUpdateListGet200ResponseDataWithDefaults() *MerchantSubscriptionPendingUpdateListGet200ResponseData {
	this := MerchantSubscriptionPendingUpdateListGet200ResponseData{}
	return &this
}

// GetSubscriptionPendingUpdateDetails returns the SubscriptionPendingUpdateDetails field value if set, zero value otherwise.
func (o *MerchantSubscriptionPendingUpdateListGet200ResponseData) GetSubscriptionPendingUpdateDetails() []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.SubscriptionPendingUpdateDetails) {
		var ret []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return o.SubscriptionPendingUpdateDetails
}

// GetSubscriptionPendingUpdateDetailsOk returns a tuple with the SubscriptionPendingUpdateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionPendingUpdateListGet200ResponseData) GetSubscriptionPendingUpdateDetailsOk() ([]UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.SubscriptionPendingUpdateDetails) {
		return nil, false
	}
	return o.SubscriptionPendingUpdateDetails, true
}

// HasSubscriptionPendingUpdateDetails returns a boolean if a field has been set.
func (o *MerchantSubscriptionPendingUpdateListGet200ResponseData) HasSubscriptionPendingUpdateDetails() bool {
	if o != nil && !IsNil(o.SubscriptionPendingUpdateDetails) {
		return true
	}

	return false
}

// SetSubscriptionPendingUpdateDetails gets a reference to the given []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the SubscriptionPendingUpdateDetails field.
func (o *MerchantSubscriptionPendingUpdateListGet200ResponseData) SetSubscriptionPendingUpdateDetails(v []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.SubscriptionPendingUpdateDetails = v
}

func (o MerchantSubscriptionPendingUpdateListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionPendingUpdateListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionPendingUpdateDetails) {
		toSerialize["subscriptionPendingUpdateDetails"] = o.SubscriptionPendingUpdateDetails
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionPendingUpdateListGet200ResponseData struct {
	value *MerchantSubscriptionPendingUpdateListGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionPendingUpdateListGet200ResponseData) Get() *MerchantSubscriptionPendingUpdateListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionPendingUpdateListGet200ResponseData) Set(val *MerchantSubscriptionPendingUpdateListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionPendingUpdateListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionPendingUpdateListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionPendingUpdateListGet200ResponseData(val *MerchantSubscriptionPendingUpdateListGet200ResponseData) *NullableMerchantSubscriptionPendingUpdateListGet200ResponseData {
	return &NullableMerchantSubscriptionPendingUpdateListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionPendingUpdateListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionPendingUpdateListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


