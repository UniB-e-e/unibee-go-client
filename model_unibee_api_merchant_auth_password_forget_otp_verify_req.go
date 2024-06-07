/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406070109 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq{}

// UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq Merchant member's password forget OTP process, verify OTP code
type UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq struct {
	// The merchant member's email address
	Email string `json:"email"`
	// The new password
	NewPassword string `json:"newPassword"`
	// OTP Code, received from email
	VerificationCode string `json:"verificationCode"`
}

type _UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq

// NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq instantiates a new UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(email string, newPassword string, verificationCode string) *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq {
	this := UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq{}
	this.Email = email
	this.NewPassword = newPassword
	this.VerificationCode = verificationCode
	return &this
}

// NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReqWithDefaults instantiates a new UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReqWithDefaults() *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq {
	this := UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) SetEmail(v string) {
	o.Email = v
}

// GetNewPassword returns the NewPassword field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) SetNewPassword(v string) {
	o.NewPassword = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) GetVerificationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) SetVerificationCode(v string) {
	o.VerificationCode = v
}

func (o UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["newPassword"] = o.NewPassword
	toSerialize["verificationCode"] = o.VerificationCode
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"newPassword",
		"verificationCode",
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

	varUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq := _UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(varUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq)

	return err
}

type NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq struct {
	value *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) Get() *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) Set(val *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(val *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) *NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq {
	return &NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


