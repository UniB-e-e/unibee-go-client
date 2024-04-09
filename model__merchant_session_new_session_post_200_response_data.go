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

// checks if the MerchantSessionNewSessionPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSessionNewSessionPost200ResponseData{}

// MerchantSessionNewSessionPost200ResponseData struct for MerchantSessionNewSessionPost200ResponseData
type MerchantSessionNewSessionPost200ResponseData struct {
	// ClientSession
	ClientSession *string `json:"clientSession,omitempty"`
	// ClientToken
	ClientToken *string `json:"clientToken,omitempty"`
	// Email
	Email *string `json:"email,omitempty"`
	// ExternalUserId
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Url
	Url *string `json:"url,omitempty"`
	// UserId
	UserId *string `json:"userId,omitempty"`
}

// NewMerchantSessionNewSessionPost200ResponseData instantiates a new MerchantSessionNewSessionPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSessionNewSessionPost200ResponseData() *MerchantSessionNewSessionPost200ResponseData {
	this := MerchantSessionNewSessionPost200ResponseData{}
	return &this
}

// NewMerchantSessionNewSessionPost200ResponseDataWithDefaults instantiates a new MerchantSessionNewSessionPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSessionNewSessionPost200ResponseDataWithDefaults() *MerchantSessionNewSessionPost200ResponseData {
	this := MerchantSessionNewSessionPost200ResponseData{}
	return &this
}

// GetClientSession returns the ClientSession field value if set, zero value otherwise.
func (o *MerchantSessionNewSessionPost200ResponseData) GetClientSession() string {
	if o == nil || IsNil(o.ClientSession) {
		var ret string
		return ret
	}
	return *o.ClientSession
}

// GetClientSessionOk returns a tuple with the ClientSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) GetClientSessionOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSession) {
		return nil, false
	}
	return o.ClientSession, true
}

// HasClientSession returns a boolean if a field has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) HasClientSession() bool {
	if o != nil && !IsNil(o.ClientSession) {
		return true
	}

	return false
}

// SetClientSession gets a reference to the given string and assigns it to the ClientSession field.
func (o *MerchantSessionNewSessionPost200ResponseData) SetClientSession(v string) {
	o.ClientSession = &v
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *MerchantSessionNewSessionPost200ResponseData) GetClientToken() string {
	if o == nil || IsNil(o.ClientToken) {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) GetClientTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClientToken) {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) HasClientToken() bool {
	if o != nil && !IsNil(o.ClientToken) {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *MerchantSessionNewSessionPost200ResponseData) SetClientToken(v string) {
	o.ClientToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MerchantSessionNewSessionPost200ResponseData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MerchantSessionNewSessionPost200ResponseData) SetEmail(v string) {
	o.Email = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *MerchantSessionNewSessionPost200ResponseData) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *MerchantSessionNewSessionPost200ResponseData) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MerchantSessionNewSessionPost200ResponseData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MerchantSessionNewSessionPost200ResponseData) SetUrl(v string) {
	o.Url = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MerchantSessionNewSessionPost200ResponseData) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MerchantSessionNewSessionPost200ResponseData) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MerchantSessionNewSessionPost200ResponseData) SetUserId(v string) {
	o.UserId = &v
}

func (o MerchantSessionNewSessionPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSessionNewSessionPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientSession) {
		toSerialize["clientSession"] = o.ClientSession
	}
	if !IsNil(o.ClientToken) {
		toSerialize["clientToken"] = o.ClientToken
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableMerchantSessionNewSessionPost200ResponseData struct {
	value *MerchantSessionNewSessionPost200ResponseData
	isSet bool
}

func (v NullableMerchantSessionNewSessionPost200ResponseData) Get() *MerchantSessionNewSessionPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSessionNewSessionPost200ResponseData) Set(val *MerchantSessionNewSessionPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSessionNewSessionPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSessionNewSessionPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSessionNewSessionPost200ResponseData(val *MerchantSessionNewSessionPost200ResponseData) *NullableMerchantSessionNewSessionPost200ResponseData {
	return &NullableMerchantSessionNewSessionPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSessionNewSessionPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSessionNewSessionPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


