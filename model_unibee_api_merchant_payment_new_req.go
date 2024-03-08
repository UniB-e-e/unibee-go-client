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

// checks if the UnibeeApiMerchantPaymentNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentNewReq{}

// UnibeeApiMerchantPaymentNewReq struct for UnibeeApiMerchantPaymentNewReq
type UnibeeApiMerchantPaymentNewReq struct {
	// CountryCode
	CountryCode string `json:"countryCode"`
	// Currency
	Currency string `json:"currency"`
	// Email
	Email string `json:"email"`
	// ExternalPaymentId should unique for payment
	ExternalPaymentId string `json:"externalPaymentId"`
	// ExternalUserId, should unique for user
	ExternalUserId string `json:"externalUserId"`
	// GatewayId
	GatewayId int64 `json:"gatewayId"`
	// Items
	LineItems []UnibeeApiMerchantPaymentItem `json:"lineItems,omitempty"`
	// Redirect Url
	RedirectUrl string `json:"redirectUrl"`
	// Metadata，Map
	Reference *map[string]string `json:"reference,omitempty"`
	// Total PaymentAmount, Cent
	TotalAmount int64 `json:"totalAmount"`
}

type _UnibeeApiMerchantPaymentNewReq UnibeeApiMerchantPaymentNewReq

// NewUnibeeApiMerchantPaymentNewReq instantiates a new UnibeeApiMerchantPaymentNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentNewReq(countryCode string, currency string, email string, externalPaymentId string, externalUserId string, gatewayId int64, redirectUrl string, totalAmount int64) *UnibeeApiMerchantPaymentNewReq {
	this := UnibeeApiMerchantPaymentNewReq{}
	this.CountryCode = countryCode
	this.Currency = currency
	this.Email = email
	this.ExternalPaymentId = externalPaymentId
	this.ExternalUserId = externalUserId
	this.GatewayId = gatewayId
	this.RedirectUrl = redirectUrl
	this.TotalAmount = totalAmount
	return &this
}

// NewUnibeeApiMerchantPaymentNewReqWithDefaults instantiates a new UnibeeApiMerchantPaymentNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentNewReqWithDefaults() *UnibeeApiMerchantPaymentNewReq {
	this := UnibeeApiMerchantPaymentNewReq{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *UnibeeApiMerchantPaymentNewReq) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantPaymentNewReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetCurrency(v string) {
	o.Currency = v
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantPaymentNewReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetEmail(v string) {
	o.Email = v
}

// GetExternalPaymentId returns the ExternalPaymentId field value
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalPaymentId
}

// GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalPaymentId, true
}

// SetExternalPaymentId sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetExternalPaymentId(v string) {
	o.ExternalPaymentId = v
}

// GetExternalUserId returns the ExternalUserId field value
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantPaymentNewReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetLineItems() []UnibeeApiMerchantPaymentItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []UnibeeApiMerchantPaymentItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetLineItemsOk() ([]UnibeeApiMerchantPaymentItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []UnibeeApiMerchantPaymentItem and assigns it to the LineItems field.
func (o *UnibeeApiMerchantPaymentNewReq) SetLineItems(v []UnibeeApiMerchantPaymentItem) {
	o.LineItems = v
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *UnibeeApiMerchantPaymentNewReq) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentNewReq) GetReference() map[string]string {
	if o == nil || IsNil(o.Reference) {
		var ret map[string]string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetReferenceOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentNewReq) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given map[string]string and assigns it to the Reference field.
func (o *UnibeeApiMerchantPaymentNewReq) SetReference(v map[string]string) {
	o.Reference = &v
}

// GetTotalAmount returns the TotalAmount field value
func (o *UnibeeApiMerchantPaymentNewReq) GetTotalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentNewReq) GetTotalAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *UnibeeApiMerchantPaymentNewReq) SetTotalAmount(v int64) {
	o.TotalAmount = v
}

func (o UnibeeApiMerchantPaymentNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryCode"] = o.CountryCode
	toSerialize["currency"] = o.Currency
	toSerialize["email"] = o.Email
	toSerialize["externalPaymentId"] = o.ExternalPaymentId
	toSerialize["externalUserId"] = o.ExternalUserId
	toSerialize["gatewayId"] = o.GatewayId
	if !IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["redirectUrl"] = o.RedirectUrl
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["totalAmount"] = o.TotalAmount
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPaymentNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"countryCode",
		"currency",
		"email",
		"externalPaymentId",
		"externalUserId",
		"gatewayId",
		"redirectUrl",
		"totalAmount",
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

	varUnibeeApiMerchantPaymentNewReq := _UnibeeApiMerchantPaymentNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPaymentNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPaymentNewReq(varUnibeeApiMerchantPaymentNewReq)

	return err
}

type NullableUnibeeApiMerchantPaymentNewReq struct {
	value *UnibeeApiMerchantPaymentNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentNewReq) Get() *UnibeeApiMerchantPaymentNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentNewReq) Set(val *UnibeeApiMerchantPaymentNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentNewReq(val *UnibeeApiMerchantPaymentNewReq) *NullableUnibeeApiMerchantPaymentNewReq {
	return &NullableUnibeeApiMerchantPaymentNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

