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

// checks if the UnibeeApiMerchantSubscriptionDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionDetailRes{}

// UnibeeApiMerchantSubscriptionDetailRes struct for UnibeeApiMerchantSubscriptionDetailRes
type UnibeeApiMerchantSubscriptionDetailRes struct {
	// Plan Addon
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	Gateway *UnibeeApiBeanGatewaySimplify `json:"gateway,omitempty"`
	Plan *UnibeeApiBeanPlanSimplify `json:"plan,omitempty"`
	Subscription *UnibeeApiBeanSubscriptionSimplify `json:"subscription,omitempty"`
	UnfinishedSubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"unfinishedSubscriptionPendingUpdate,omitempty"`
	User *UnibeeApiBeanUserAccountSimplify `json:"user,omitempty"`
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
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
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

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGateway() UnibeeApiBeanGatewaySimplify {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGatewaySimplify
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool) {
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

// SetGateway gets a reference to the given UnibeeApiBeanGatewaySimplify and assigns it to the Gateway field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetGateway(v UnibeeApiBeanGatewaySimplify) {
	o.Gateway = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlan() UnibeeApiBeanPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool) {
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

// SetPlan gets a reference to the given UnibeeApiBeanPlanSimplify and assigns it to the Plan field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetPlan(v UnibeeApiBeanPlanSimplify) {
	o.Plan = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscription() UnibeeApiBeanSubscriptionSimplify {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscriptionSimplify
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool) {
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

// SetSubscription gets a reference to the given UnibeeApiBeanSubscriptionSimplify and assigns it to the Subscription field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetSubscription(v UnibeeApiBeanSubscriptionSimplify) {
	o.Subscription = &v
}

// GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.UnfinishedSubscriptionPendingUpdate
}

// GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
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

// SetUnfinishedSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the UnfinishedSubscriptionPendingUpdate field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.UnfinishedSubscriptionPendingUpdate = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUser() UnibeeApiBeanUserAccountSimplify {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccountSimplify
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUserOk() (*UnibeeApiBeanUserAccountSimplify, bool) {
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

// SetUser gets a reference to the given UnibeeApiBeanUserAccountSimplify and assigns it to the User field.
func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUser(v UnibeeApiBeanUserAccountSimplify) {
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


