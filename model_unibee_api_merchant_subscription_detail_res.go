/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionDetailRes{}

// UnibeeApiMerchantSubscriptionDetailRes struct for UnibeeApiMerchantSubscriptionDetailRes
type UnibeeApiMerchantSubscriptionDetailRes struct {
	// Plan Addon
	Addons []UnibeeInternalLogicGatewayRoPlanAddonVo `json:"addons,omitempty"`
	Gateway *UnibeeInternalLogicGatewayRoGatewaySimplify `json:"gateway,omitempty"`
	Plan *UnibeeInternalLogicGatewayRoPlanSimplify `json:"plan,omitempty"`
	Subscription *UnibeeInternalLogicGatewayRoSubscriptionSimplify `json:"subscription,omitempty"`
	UnfinishedSubscriptionPendingUpdate *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo `json:"unfinishedSubscriptionPendingUpdate,omitempty"`
	User *UnibeeInternalLogicGatewayRoUserAccountSimplify `json:"user,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionDetailRes instantiates a new UnibeeApiMerchantSubscriptionDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionDetailRes() *UnibeeApiMerchantSubscriptionDetailRes {
	this := UnibeeApiMerchantSubscriptionDetailRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionDetailResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionDetailResWithDefaults() *UnibeeApiMerchantSubscriptionDetailRes {
	this := UnibeeApiMerchantSubscriptionDetailRes{}
	return &this
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeInternalLogicGatewayRoPlanAddonVo
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddonsOk() ([]UnibeeInternalLogicGatewayRoPlanAddonVo, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeInternalLogicGatewayRoPlanAddonVo and assigns it to the Addons field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo) {
	o.Addons = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGateway() UnibeeInternalLogicGatewayRoGatewaySimplify {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeInternalLogicGatewayRoGatewaySimplify
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGatewayOk() (*UnibeeInternalLogicGatewayRoGatewaySimplify, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeInternalLogicGatewayRoGatewaySimplify and assigns it to the Gateway field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetGateway(v UnibeeInternalLogicGatewayRoGatewaySimplify) {
	o.Gateway = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the Plan field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.Plan = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscription() UnibeeInternalLogicGatewayRoSubscriptionSimplify {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeInternalLogicGatewayRoSubscriptionSimplify
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscriptionOk() (*UnibeeInternalLogicGatewayRoSubscriptionSimplify, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeInternalLogicGatewayRoSubscriptionSimplify and assigns it to the Subscription field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetSubscription(v UnibeeInternalLogicGatewayRoSubscriptionSimplify) {
	o.Subscription = &v
}

// GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdate() UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		var ret UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo
		return ret
	}
	return *o.UnfinishedSubscriptionPendingUpdate
}

// GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo, bool) {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return nil, false
	}
	return o.UnfinishedSubscriptionPendingUpdate, true
}

// HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) HasUnfinishedSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetUnfinishedSubscriptionPendingUpdate gets a reference to the given UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo and assigns it to the UnfinishedSubscriptionPendingUpdate field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUnfinishedSubscriptionPendingUpdate(v UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) {
	o.UnfinishedSubscriptionPendingUpdate = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify {
	if o == nil || IsNil(o.User) {
		var ret UnibeeInternalLogicGatewayRoUserAccountSimplify
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeInternalLogicGatewayRoUserAccountSimplify and assigns it to the User field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify) {
	o.User = &v
}

func (o UnibeeApiMerchantSubscriptionDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		toSerialize["unfinishedSubscriptionPendingUpdate"] = o.UnfinishedSubscriptionPendingUpdate
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionDetailRes struct {
	value *UnibeeApiMerchantSubscriptionDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionDetailRes) Get() *UnibeeApiMerchantSubscriptionDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionDetailRes) Set(val *UnibeeApiMerchantSubscriptionDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionDetailRes(val *UnibeeApiMerchantSubscriptionDetailRes) *NullableUnibeeApiMerchantSubscriptionDetailRes {
	return &NullableUnibeeApiMerchantSubscriptionDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

