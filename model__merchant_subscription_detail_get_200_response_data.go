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

// checks if the MerchantSubscriptionDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionDetailGet200ResponseData{}

// MerchantSubscriptionDetailGet200ResponseData struct for MerchantSubscriptionDetailGet200ResponseData
type MerchantSubscriptionDetailGet200ResponseData struct {
	// AddonParams
	AddonParams []UnibeeApiBeanPlanAddonParam `json:"addonParams,omitempty"`
	// Plan Addon
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	Gateway *UnibeeApiBeanGatewaySimplify `json:"gateway,omitempty"`
	LatestInvoice *UnibeeApiBeanInvoiceSimplify `json:"latestInvoice,omitempty"`
	Plan *UnibeeApiBeanPlanSimplify `json:"plan,omitempty"`
	Subscription *UnibeeApiBeanSubscriptionSimplify `json:"subscription,omitempty"`
	UnfinishedSubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"unfinishedSubscriptionPendingUpdate,omitempty"`
	User *UnibeeApiBeanUserAccountSimplify `json:"user,omitempty"`
}

// NewMerchantSubscriptionDetailGet200ResponseData instantiates a new MerchantSubscriptionDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionDetailGet200ResponseData() *MerchantSubscriptionDetailGet200ResponseData {
	this := MerchantSubscriptionDetailGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionDetailGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionDetailGet200ResponseDataWithDefaults() *MerchantSubscriptionDetailGet200ResponseData {
	this := MerchantSubscriptionDetailGet200ResponseData{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetGateway() UnibeeApiBeanGatewaySimplify {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGatewaySimplify
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGatewaySimplify and assigns it to the Gateway field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetGateway(v UnibeeApiBeanGatewaySimplify) {
	o.Gateway = &v
}

// GetLatestInvoice returns the LatestInvoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetLatestInvoice() UnibeeApiBeanInvoiceSimplify {
	if o == nil || IsNil(o.LatestInvoice) {
		var ret UnibeeApiBeanInvoiceSimplify
		return ret
	}
	return *o.LatestInvoice
}

// GetLatestInvoiceOk returns a tuple with the LatestInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetLatestInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool) {
	if o == nil || IsNil(o.LatestInvoice) {
		return nil, false
	}
	return o.LatestInvoice, true
}

// HasLatestInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasLatestInvoice() bool {
	if o != nil && !IsNil(o.LatestInvoice) {
		return true
	}

	return false
}

// SetLatestInvoice gets a reference to the given UnibeeApiBeanInvoiceSimplify and assigns it to the LatestInvoice field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetLatestInvoice(v UnibeeApiBeanInvoiceSimplify) {
	o.LatestInvoice = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetPlan() UnibeeApiBeanPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlanSimplify and assigns it to the Plan field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetPlan(v UnibeeApiBeanPlanSimplify) {
	o.Plan = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetSubscription() UnibeeApiBeanSubscriptionSimplify {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscriptionSimplify
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscriptionSimplify and assigns it to the Subscription field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetSubscription(v UnibeeApiBeanSubscriptionSimplify) {
	o.Subscription = &v
}

// GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.UnfinishedSubscriptionPendingUpdate
}

// GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return nil, false
	}
	return o.UnfinishedSubscriptionPendingUpdate, true
}

// HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasUnfinishedSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetUnfinishedSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the UnfinishedSubscriptionPendingUpdate field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.UnfinishedSubscriptionPendingUpdate = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetUser() UnibeeApiBeanUserAccountSimplify {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccountSimplify
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccountSimplify, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MerchantSubscriptionDetailGet200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccountSimplify and assigns it to the User field.
func (o *MerchantSubscriptionDetailGet200ResponseData) SetUser(v UnibeeApiBeanUserAccountSimplify) {
	o.User = &v
}

func (o MerchantSubscriptionDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.LatestInvoice) {
		toSerialize["latestInvoice"] = o.LatestInvoice
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

type NullableMerchantSubscriptionDetailGet200ResponseData struct {
	value *MerchantSubscriptionDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionDetailGet200ResponseData) Get() *MerchantSubscriptionDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionDetailGet200ResponseData) Set(val *MerchantSubscriptionDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionDetailGet200ResponseData(val *MerchantSubscriptionDetailGet200ResponseData) *NullableMerchantSubscriptionDetailGet200ResponseData {
	return &NullableMerchantSubscriptionDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


