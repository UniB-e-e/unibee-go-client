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

// checks if the MerchantGetGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGetGet200ResponseData{}

// MerchantGetGet200ResponseData struct for MerchantGetGet200ResponseData
type MerchantGetGet200ResponseData struct {
	// Currency List
	Currency []UnibeeApiBeanCurrency `json:"Currency,omitempty"`
	// TimeZone List
	TimeZone []string `json:"TimeZone,omitempty"`
	// System Env, em: daily|stage|local|prod
	Env *string `json:"env,omitempty"`
	// Gateway List
	Gateway []UnibeeApiBeanGatewaySimplify `json:"gateway,omitempty"`
	// Check System Env Is Prod, true|false
	IsProd *bool `json:"isProd,omitempty"`
	Merchant *UnibeeApiBeanMerchantSimplify `json:"merchant,omitempty"`
	MerchantMember *UnibeeApiBeanDetailMerchantMemberDetail `json:"merchantMember,omitempty"`
}

// NewMerchantGetGet200ResponseData instantiates a new MerchantGetGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGetGet200ResponseData() *MerchantGetGet200ResponseData {
	this := MerchantGetGet200ResponseData{}
	return &this
}

// NewMerchantGetGet200ResponseDataWithDefaults instantiates a new MerchantGetGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGetGet200ResponseDataWithDefaults() *MerchantGetGet200ResponseData {
	this := MerchantGetGet200ResponseData{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetCurrency() []UnibeeApiBeanCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret []UnibeeApiBeanCurrency
		return ret
	}
	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetCurrencyOk() ([]UnibeeApiBeanCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given []UnibeeApiBeanCurrency and assigns it to the Currency field.
func (o *MerchantGetGet200ResponseData) SetCurrency(v []UnibeeApiBeanCurrency) {
	o.Currency = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetTimeZone() []string {
	if o == nil || IsNil(o.TimeZone) {
		var ret []string
		return ret
	}
	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetTimeZoneOk() ([]string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given []string and assigns it to the TimeZone field.
func (o *MerchantGetGet200ResponseData) SetTimeZone(v []string) {
	o.TimeZone = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetEnv() string {
	if o == nil || IsNil(o.Env) {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetEnvOk() (*string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *MerchantGetGet200ResponseData) SetEnv(v string) {
	o.Env = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetGateway() []UnibeeApiBeanGatewaySimplify {
	if o == nil || IsNil(o.Gateway) {
		var ret []UnibeeApiBeanGatewaySimplify
		return ret
	}
	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetGatewayOk() ([]UnibeeApiBeanGatewaySimplify, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given []UnibeeApiBeanGatewaySimplify and assigns it to the Gateway field.
func (o *MerchantGetGet200ResponseData) SetGateway(v []UnibeeApiBeanGatewaySimplify) {
	o.Gateway = v
}

// GetIsProd returns the IsProd field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetIsProd() bool {
	if o == nil || IsNil(o.IsProd) {
		var ret bool
		return ret
	}
	return *o.IsProd
}

// GetIsProdOk returns a tuple with the IsProd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetIsProdOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProd) {
		return nil, false
	}
	return o.IsProd, true
}

// HasIsProd returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasIsProd() bool {
	if o != nil && !IsNil(o.IsProd) {
		return true
	}

	return false
}

// SetIsProd gets a reference to the given bool and assigns it to the IsProd field.
func (o *MerchantGetGet200ResponseData) SetIsProd(v bool) {
	o.IsProd = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetMerchant() UnibeeApiBeanMerchantSimplify {
	if o == nil || IsNil(o.Merchant) {
		var ret UnibeeApiBeanMerchantSimplify
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchantSimplify, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given UnibeeApiBeanMerchantSimplify and assigns it to the Merchant field.
func (o *MerchantGetGet200ResponseData) SetMerchant(v UnibeeApiBeanMerchantSimplify) {
	o.Merchant = &v
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *MerchantGetGet200ResponseData) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanDetailMerchantMemberDetail
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGetGet200ResponseData) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *MerchantGetGet200ResponseData) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanDetailMerchantMemberDetail and assigns it to the MerchantMember field.
func (o *MerchantGetGet200ResponseData) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail) {
	o.MerchantMember = &v
}

func (o MerchantGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGetGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["Currency"] = o.Currency
	}
	if !IsNil(o.TimeZone) {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.IsProd) {
		toSerialize["isProd"] = o.IsProd
	}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	return toSerialize, nil
}

type NullableMerchantGetGet200ResponseData struct {
	value *MerchantGetGet200ResponseData
	isSet bool
}

func (v NullableMerchantGetGet200ResponseData) Get() *MerchantGetGet200ResponseData {
	return v.value
}

func (v *NullableMerchantGetGet200ResponseData) Set(val *MerchantGetGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGetGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGetGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGetGet200ResponseData(val *MerchantGetGet200ResponseData) *NullableMerchantGetGet200ResponseData {
	return &NullableMerchantGetGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantGetGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGetGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


