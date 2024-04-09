/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: buildtime:202404090336 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiSystemRefundBulkChannelSyncReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemRefundBulkChannelSyncReq{}

// UnibeeApiSystemRefundBulkChannelSyncReq struct for UnibeeApiSystemRefundBulkChannelSyncReq
type UnibeeApiSystemRefundBulkChannelSyncReq struct {
	// merchantId
	MerchantId string `json:"merchantId"`
}

type _UnibeeApiSystemRefundBulkChannelSyncReq UnibeeApiSystemRefundBulkChannelSyncReq

// NewUnibeeApiSystemRefundBulkChannelSyncReq instantiates a new UnibeeApiSystemRefundBulkChannelSyncReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemRefundBulkChannelSyncReq(merchantId string) *UnibeeApiSystemRefundBulkChannelSyncReq {
	this := UnibeeApiSystemRefundBulkChannelSyncReq{}
	this.MerchantId = merchantId
	return &this
}

// NewUnibeeApiSystemRefundBulkChannelSyncReqWithDefaults instantiates a new UnibeeApiSystemRefundBulkChannelSyncReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemRefundBulkChannelSyncReqWithDefaults() *UnibeeApiSystemRefundBulkChannelSyncReq {
	this := UnibeeApiSystemRefundBulkChannelSyncReq{}
	return &this
}

// GetMerchantId returns the MerchantId field value
func (o *UnibeeApiSystemRefundBulkChannelSyncReq) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemRefundBulkChannelSyncReq) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *UnibeeApiSystemRefundBulkChannelSyncReq) SetMerchantId(v string) {
	o.MerchantId = v
}

func (o UnibeeApiSystemRefundBulkChannelSyncReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemRefundBulkChannelSyncReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchantId"] = o.MerchantId
	return toSerialize, nil
}

func (o *UnibeeApiSystemRefundBulkChannelSyncReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantId",
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

	varUnibeeApiSystemRefundBulkChannelSyncReq := _UnibeeApiSystemRefundBulkChannelSyncReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemRefundBulkChannelSyncReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemRefundBulkChannelSyncReq(varUnibeeApiSystemRefundBulkChannelSyncReq)

	return err
}

type NullableUnibeeApiSystemRefundBulkChannelSyncReq struct {
	value *UnibeeApiSystemRefundBulkChannelSyncReq
	isSet bool
}

func (v NullableUnibeeApiSystemRefundBulkChannelSyncReq) Get() *UnibeeApiSystemRefundBulkChannelSyncReq {
	return v.value
}

func (v *NullableUnibeeApiSystemRefundBulkChannelSyncReq) Set(val *UnibeeApiSystemRefundBulkChannelSyncReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemRefundBulkChannelSyncReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemRefundBulkChannelSyncReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemRefundBulkChannelSyncReq(val *UnibeeApiSystemRefundBulkChannelSyncReq) *NullableUnibeeApiSystemRefundBulkChannelSyncReq {
	return &NullableUnibeeApiSystemRefundBulkChannelSyncReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemRefundBulkChannelSyncReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemRefundBulkChannelSyncReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


