/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionUpdatePreviewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUpdatePreviewReq{}

// UnibeeApiMerchantSubscriptionUpdatePreviewReq struct for UnibeeApiMerchantSubscriptionUpdatePreviewReq
type UnibeeApiMerchantSubscriptionUpdatePreviewReq struct {
	// addonParams
	AddonParams []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo `json:"addonParams,omitempty"`
	// New PlanId
	NewPlanId int32 `json:"newPlanId"`
	// Quantity，Default 1
	Quantity *int64 `json:"quantity,omitempty"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
	// Effect Immediate，1-Immediate，2-Next Period
	WithImmediateEffect *int32 `json:"withImmediateEffect,omitempty"`
}

type _UnibeeApiMerchantSubscriptionUpdatePreviewReq UnibeeApiMerchantSubscriptionUpdatePreviewReq

// NewUnibeeApiMerchantSubscriptionUpdatePreviewReq instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(newPlanId int32, subscriptionId string) *UnibeeApiMerchantSubscriptionUpdatePreviewReq {
	this := UnibeeApiMerchantSubscriptionUpdatePreviewReq{}
	this.NewPlanId = newPlanId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults() *UnibeeApiMerchantSubscriptionUpdatePreviewReq {
	this := UnibeeApiMerchantSubscriptionUpdatePreviewReq{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetAddonParams() []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetAddonParamsOk() ([]UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo and assigns it to the AddonParams field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetAddonParams(v []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo) {
	o.AddonParams = v
}

// GetNewPlanId returns the NewPlanId field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetNewPlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NewPlanId
}

// GetNewPlanIdOk returns a tuple with the NewPlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetNewPlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPlanId, true
}

// SetNewPlanId sets field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetNewPlanId(v int32) {
	o.NewPlanId = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetWithImmediateEffect returns the WithImmediateEffect field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetWithImmediateEffect() int32 {
	if o == nil || IsNil(o.WithImmediateEffect) {
		var ret int32
		return ret
	}
	return *o.WithImmediateEffect
}

// GetWithImmediateEffectOk returns a tuple with the WithImmediateEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetWithImmediateEffectOk() (*int32, bool) {
	if o == nil || IsNil(o.WithImmediateEffect) {
		return nil, false
	}
	return o.WithImmediateEffect, true
}

// HasWithImmediateEffect returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasWithImmediateEffect() bool {
	if o != nil && !IsNil(o.WithImmediateEffect) {
		return true
	}

	return false
}

// SetWithImmediateEffect gets a reference to the given int32 and assigns it to the WithImmediateEffect field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetWithImmediateEffect(v int32) {
	o.WithImmediateEffect = &v
}

func (o UnibeeApiMerchantSubscriptionUpdatePreviewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUpdatePreviewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	toSerialize["newPlanId"] = o.NewPlanId
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	if !IsNil(o.WithImmediateEffect) {
		toSerialize["withImmediateEffect"] = o.WithImmediateEffect
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPlanId",
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

	varUnibeeApiMerchantSubscriptionUpdatePreviewReq := _UnibeeApiMerchantSubscriptionUpdatePreviewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionUpdatePreviewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionUpdatePreviewReq(varUnibeeApiMerchantSubscriptionUpdatePreviewReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq struct {
	value *UnibeeApiMerchantSubscriptionUpdatePreviewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) Get() *UnibeeApiMerchantSubscriptionUpdatePreviewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) Set(val *UnibeeApiMerchantSubscriptionUpdatePreviewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUpdatePreviewReq(val *UnibeeApiMerchantSubscriptionUpdatePreviewReq) *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq {
	return &NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

