/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406070109 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiSystemInformationGetRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemInformationGetRes{}

// UnibeeApiSystemInformationGetRes struct for UnibeeApiSystemInformationGetRes
type UnibeeApiSystemInformationGetRes struct {
	// System Env, em: daily|stage|local|prod
	Env *string `json:"env,omitempty"`
	// Support Currency List
	Gateway []UnibeeApiBeanGatewaySimplify `json:"gateway,omitempty"`
	// Check System Env Is Prod, true|false
	IsProd *bool `json:"isProd,omitempty"`
	// Support Currency List
	SupportCurrency []UnibeeApiBeanCurrency `json:"supportCurrency,omitempty"`
	// Support TimeZone List
	SupportTimeZone []string `json:"supportTimeZone,omitempty"`
}

// NewUnibeeApiSystemInformationGetRes instantiates a new UnibeeApiSystemInformationGetRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemInformationGetRes() *UnibeeApiSystemInformationGetRes {
	this := UnibeeApiSystemInformationGetRes{}
	return &this
}

// NewUnibeeApiSystemInformationGetResWithDefaults instantiates a new UnibeeApiSystemInformationGetRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemInformationGetResWithDefaults() *UnibeeApiSystemInformationGetRes {
	this := UnibeeApiSystemInformationGetRes{}
	return &this
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *UnibeeApiSystemInformationGetRes) GetEnv() string {
	if o == nil || IsNil(o.Env) {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInformationGetRes) GetEnvOk() (*string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *UnibeeApiSystemInformationGetRes) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *UnibeeApiSystemInformationGetRes) SetEnv(v string) {
	o.Env = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiSystemInformationGetRes) GetGateway() []UnibeeApiBeanGatewaySimplify {
	if o == nil || IsNil(o.Gateway) {
		var ret []UnibeeApiBeanGatewaySimplify
		return ret
	}
	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInformationGetRes) GetGatewayOk() ([]UnibeeApiBeanGatewaySimplify, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiSystemInformationGetRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given []UnibeeApiBeanGatewaySimplify and assigns it to the Gateway field.
func (o *UnibeeApiSystemInformationGetRes) SetGateway(v []UnibeeApiBeanGatewaySimplify) {
	o.Gateway = v
}

// GetIsProd returns the IsProd field value if set, zero value otherwise.
func (o *UnibeeApiSystemInformationGetRes) GetIsProd() bool {
	if o == nil || IsNil(o.IsProd) {
		var ret bool
		return ret
	}
	return *o.IsProd
}

// GetIsProdOk returns a tuple with the IsProd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInformationGetRes) GetIsProdOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProd) {
		return nil, false
	}
	return o.IsProd, true
}

// HasIsProd returns a boolean if a field has been set.
func (o *UnibeeApiSystemInformationGetRes) HasIsProd() bool {
	if o != nil && !IsNil(o.IsProd) {
		return true
	}

	return false
}

// SetIsProd gets a reference to the given bool and assigns it to the IsProd field.
func (o *UnibeeApiSystemInformationGetRes) SetIsProd(v bool) {
	o.IsProd = &v
}

// GetSupportCurrency returns the SupportCurrency field value if set, zero value otherwise.
func (o *UnibeeApiSystemInformationGetRes) GetSupportCurrency() []UnibeeApiBeanCurrency {
	if o == nil || IsNil(o.SupportCurrency) {
		var ret []UnibeeApiBeanCurrency
		return ret
	}
	return o.SupportCurrency
}

// GetSupportCurrencyOk returns a tuple with the SupportCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInformationGetRes) GetSupportCurrencyOk() ([]UnibeeApiBeanCurrency, bool) {
	if o == nil || IsNil(o.SupportCurrency) {
		return nil, false
	}
	return o.SupportCurrency, true
}

// HasSupportCurrency returns a boolean if a field has been set.
func (o *UnibeeApiSystemInformationGetRes) HasSupportCurrency() bool {
	if o != nil && !IsNil(o.SupportCurrency) {
		return true
	}

	return false
}

// SetSupportCurrency gets a reference to the given []UnibeeApiBeanCurrency and assigns it to the SupportCurrency field.
func (o *UnibeeApiSystemInformationGetRes) SetSupportCurrency(v []UnibeeApiBeanCurrency) {
	o.SupportCurrency = v
}

// GetSupportTimeZone returns the SupportTimeZone field value if set, zero value otherwise.
func (o *UnibeeApiSystemInformationGetRes) GetSupportTimeZone() []string {
	if o == nil || IsNil(o.SupportTimeZone) {
		var ret []string
		return ret
	}
	return o.SupportTimeZone
}

// GetSupportTimeZoneOk returns a tuple with the SupportTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInformationGetRes) GetSupportTimeZoneOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportTimeZone) {
		return nil, false
	}
	return o.SupportTimeZone, true
}

// HasSupportTimeZone returns a boolean if a field has been set.
func (o *UnibeeApiSystemInformationGetRes) HasSupportTimeZone() bool {
	if o != nil && !IsNil(o.SupportTimeZone) {
		return true
	}

	return false
}

// SetSupportTimeZone gets a reference to the given []string and assigns it to the SupportTimeZone field.
func (o *UnibeeApiSystemInformationGetRes) SetSupportTimeZone(v []string) {
	o.SupportTimeZone = v
}

func (o UnibeeApiSystemInformationGetRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemInformationGetRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.IsProd) {
		toSerialize["isProd"] = o.IsProd
	}
	if !IsNil(o.SupportCurrency) {
		toSerialize["supportCurrency"] = o.SupportCurrency
	}
	if !IsNil(o.SupportTimeZone) {
		toSerialize["supportTimeZone"] = o.SupportTimeZone
	}
	return toSerialize, nil
}

type NullableUnibeeApiSystemInformationGetRes struct {
	value *UnibeeApiSystemInformationGetRes
	isSet bool
}

func (v NullableUnibeeApiSystemInformationGetRes) Get() *UnibeeApiSystemInformationGetRes {
	return v.value
}

func (v *NullableUnibeeApiSystemInformationGetRes) Set(val *UnibeeApiSystemInformationGetRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemInformationGetRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemInformationGetRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemInformationGetRes(val *UnibeeApiSystemInformationGetRes) *NullableUnibeeApiSystemInformationGetRes {
	return &NullableUnibeeApiSystemInformationGetRes{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemInformationGetRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemInformationGetRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


