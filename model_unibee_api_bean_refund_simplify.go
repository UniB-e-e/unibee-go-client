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

// checks if the UnibeeApiBeanRefundSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanRefundSimplify{}

// UnibeeApiBeanRefundSimplify struct for UnibeeApiBeanRefundSimplify
type UnibeeApiBeanRefundSimplify struct {
	// country code
	CountryCode *string `json:"countryCode,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// external_refund_id
	ExternalRefundId *string `json:"externalRefundId,omitempty"`
	// gateway_id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// gateway refund id
	GatewayRefundId *string `json:"gatewayRefundId,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	// relative payment id
	PaymentId *string `json:"paymentId,omitempty"`
	// refund amount, cent
	RefundAmount *int64 `json:"refundAmount,omitempty"`
	// refund comment
	RefundComment *string `json:"refundComment,omitempty"`
	// refund id (system generate)
	RefundId *string `json:"refundId,omitempty"`
	// refund success time
	RefundTime *int64 `json:"refundTime,omitempty"`
	// return url after refund success
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// status。10-pending，20-success，30-failure, 40-cancel
	Status *int32 `json:"status,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanRefundSimplify instantiates a new UnibeeApiBeanRefundSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanRefundSimplify() *UnibeeApiBeanRefundSimplify {
	this := UnibeeApiBeanRefundSimplify{}
	return &this
}

// NewUnibeeApiBeanRefundSimplifyWithDefaults instantiates a new UnibeeApiBeanRefundSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanRefundSimplifyWithDefaults() *UnibeeApiBeanRefundSimplify {
	this := UnibeeApiBeanRefundSimplify{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanRefundSimplify) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanRefundSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanRefundSimplify) SetCurrency(v string) {
	o.Currency = &v
}

// GetExternalRefundId returns the ExternalRefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetExternalRefundId() string {
	if o == nil || IsNil(o.ExternalRefundId) {
		var ret string
		return ret
	}
	return *o.ExternalRefundId
}

// GetExternalRefundIdOk returns a tuple with the ExternalRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetExternalRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalRefundId) {
		return nil, false
	}
	return o.ExternalRefundId, true
}

// HasExternalRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasExternalRefundId() bool {
	if o != nil && !IsNil(o.ExternalRefundId) {
		return true
	}

	return false
}

// SetExternalRefundId gets a reference to the given string and assigns it to the ExternalRefundId field.
func (o *UnibeeApiBeanRefundSimplify) SetExternalRefundId(v string) {
	o.ExternalRefundId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanRefundSimplify) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayRefundId returns the GatewayRefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetGatewayRefundId() string {
	if o == nil || IsNil(o.GatewayRefundId) {
		var ret string
		return ret
	}
	return *o.GatewayRefundId
}

// GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetGatewayRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayRefundId) {
		return nil, false
	}
	return o.GatewayRefundId, true
}

// HasGatewayRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasGatewayRefundId() bool {
	if o != nil && !IsNil(o.GatewayRefundId) {
		return true
	}

	return false
}

// SetGatewayRefundId gets a reference to the given string and assigns it to the GatewayRefundId field.
func (o *UnibeeApiBeanRefundSimplify) SetGatewayRefundId(v string) {
	o.GatewayRefundId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanRefundSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiBeanRefundSimplify) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanRefundSimplify) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetRefundAmount() int64 {
	if o == nil || IsNil(o.RefundAmount) {
		var ret int64
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetRefundAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given int64 and assigns it to the RefundAmount field.
func (o *UnibeeApiBeanRefundSimplify) SetRefundAmount(v int64) {
	o.RefundAmount = &v
}

// GetRefundComment returns the RefundComment field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetRefundComment() string {
	if o == nil || IsNil(o.RefundComment) {
		var ret string
		return ret
	}
	return *o.RefundComment
}

// GetRefundCommentOk returns a tuple with the RefundComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetRefundCommentOk() (*string, bool) {
	if o == nil || IsNil(o.RefundComment) {
		return nil, false
	}
	return o.RefundComment, true
}

// HasRefundComment returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasRefundComment() bool {
	if o != nil && !IsNil(o.RefundComment) {
		return true
	}

	return false
}

// SetRefundComment gets a reference to the given string and assigns it to the RefundComment field.
func (o *UnibeeApiBeanRefundSimplify) SetRefundComment(v string) {
	o.RefundComment = &v
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *UnibeeApiBeanRefundSimplify) SetRefundId(v string) {
	o.RefundId = &v
}

// GetRefundTime returns the RefundTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetRefundTime() int64 {
	if o == nil || IsNil(o.RefundTime) {
		var ret int64
		return ret
	}
	return *o.RefundTime
}

// GetRefundTimeOk returns a tuple with the RefundTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetRefundTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RefundTime) {
		return nil, false
	}
	return o.RefundTime, true
}

// HasRefundTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasRefundTime() bool {
	if o != nil && !IsNil(o.RefundTime) {
		return true
	}

	return false
}

// SetRefundTime gets a reference to the given int64 and assigns it to the RefundTime field.
func (o *UnibeeApiBeanRefundSimplify) SetRefundTime(v int64) {
	o.RefundTime = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiBeanRefundSimplify) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanRefundSimplify) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanRefundSimplify) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefundSimplify) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefundSimplify) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefundSimplify) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanRefundSimplify) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanRefundSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanRefundSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExternalRefundId) {
		toSerialize["externalRefundId"] = o.ExternalRefundId
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayRefundId) {
		toSerialize["gatewayRefundId"] = o.GatewayRefundId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refundAmount"] = o.RefundAmount
	}
	if !IsNil(o.RefundComment) {
		toSerialize["refundComment"] = o.RefundComment
	}
	if !IsNil(o.RefundId) {
		toSerialize["refundId"] = o.RefundId
	}
	if !IsNil(o.RefundTime) {
		toSerialize["refundTime"] = o.RefundTime
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanRefundSimplify struct {
	value *UnibeeApiBeanRefundSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanRefundSimplify) Get() *UnibeeApiBeanRefundSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanRefundSimplify) Set(val *UnibeeApiBeanRefundSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanRefundSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanRefundSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanRefundSimplify(val *UnibeeApiBeanRefundSimplify) *NullableUnibeeApiBeanRefundSimplify {
	return &NullableUnibeeApiBeanRefundSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanRefundSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanRefundSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

