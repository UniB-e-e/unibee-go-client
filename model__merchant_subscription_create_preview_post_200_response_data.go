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

// checks if the MerchantSubscriptionCreatePreviewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionCreatePreviewPost200ResponseData{}

// MerchantSubscriptionCreatePreviewPost200ResponseData struct for MerchantSubscriptionCreatePreviewPost200ResponseData
type MerchantSubscriptionCreatePreviewPost200ResponseData struct {
	AddonParams []UnibeeApiBeanPlanAddonParam `json:"addonParams,omitempty"`
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Email *string `json:"email,omitempty"`
	Gateway *UnibeeApiBeanGatewaySimplify `json:"gateway,omitempty"`
	Invoice *UnibeeApiBeanInvoiceSimplify `json:"invoice,omitempty"`
	Plan *UnibeeApiBeanPlanSimplify `json:"plan,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	TaxScale *int64 `json:"taxScale,omitempty"`
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	UserId *int64 `json:"userId,omitempty"`
	VatCountryCode *string `json:"vatCountryCode,omitempty"`
	VatCountryName *string `json:"vatCountryName,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
	VatNumberValidate *UnibeeApiBeanValidResult `json:"vatNumberValidate,omitempty"`
}

// NewMerchantSubscriptionCreatePreviewPost200ResponseData instantiates a new MerchantSubscriptionCreatePreviewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionCreatePreviewPost200ResponseData() *MerchantSubscriptionCreatePreviewPost200ResponseData {
	this := MerchantSubscriptionCreatePreviewPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionCreatePreviewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults() *MerchantSubscriptionCreatePreviewPost200ResponseData {
	this := MerchantSubscriptionCreatePreviewPost200ResponseData{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetCurrency(v string) {
	o.Currency = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetEmail(v string) {
	o.Email = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetGateway() UnibeeApiBeanGatewaySimplify {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGatewaySimplify
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGatewaySimplify and assigns it to the Gateway field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetGateway(v UnibeeApiBeanGatewaySimplify) {
	o.Gateway = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetInvoice() UnibeeApiBeanInvoiceSimplify {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoiceSimplify
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoiceSimplify and assigns it to the Invoice field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoiceSimplify) {
	o.Invoice = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetPlan() UnibeeApiBeanPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlanSimplify and assigns it to the Plan field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetPlan(v UnibeeApiBeanPlanSimplify) {
	o.Plan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTaxScale returns the TaxScale field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTaxScale() int64 {
	if o == nil || IsNil(o.TaxScale) {
		var ret int64
		return ret
	}
	return *o.TaxScale
}

// GetTaxScaleOk returns a tuple with the TaxScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTaxScaleOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxScale) {
		return nil, false
	}
	return o.TaxScale, true
}

// HasTaxScale returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTaxScale() bool {
	if o != nil && !IsNil(o.TaxScale) {
		return true
	}

	return false
}

// SetTaxScale gets a reference to the given int64 and assigns it to the TaxScale field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTaxScale(v int64) {
	o.TaxScale = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetUserId(v int64) {
	o.UserId = &v
}

// GetVatCountryCode returns the VatCountryCode field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryCode() string {
	if o == nil || IsNil(o.VatCountryCode) {
		var ret string
		return ret
	}
	return *o.VatCountryCode
}

// GetVatCountryCodeOk returns a tuple with the VatCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryCode) {
		return nil, false
	}
	return o.VatCountryCode, true
}

// HasVatCountryCode returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatCountryCode() bool {
	if o != nil && !IsNil(o.VatCountryCode) {
		return true
	}

	return false
}

// SetVatCountryCode gets a reference to the given string and assigns it to the VatCountryCode field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatCountryCode(v string) {
	o.VatCountryCode = &v
}

// GetVatCountryName returns the VatCountryName field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryName() string {
	if o == nil || IsNil(o.VatCountryName) {
		var ret string
		return ret
	}
	return *o.VatCountryName
}

// GetVatCountryNameOk returns a tuple with the VatCountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryName) {
		return nil, false
	}
	return o.VatCountryName, true
}

// HasVatCountryName returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatCountryName() bool {
	if o != nil && !IsNil(o.VatCountryName) {
		return true
	}

	return false
}

// SetVatCountryName gets a reference to the given string and assigns it to the VatCountryName field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatCountryName(v string) {
	o.VatCountryName = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetVatNumberValidate returns the VatNumberValidate field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidate() UnibeeApiBeanValidResult {
	if o == nil || IsNil(o.VatNumberValidate) {
		var ret UnibeeApiBeanValidResult
		return ret
	}
	return *o.VatNumberValidate
}

// GetVatNumberValidateOk returns a tuple with the VatNumberValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool) {
	if o == nil || IsNil(o.VatNumberValidate) {
		return nil, false
	}
	return o.VatNumberValidate, true
}

// HasVatNumberValidate returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumberValidate() bool {
	if o != nil && !IsNil(o.VatNumberValidate) {
		return true
	}

	return false
}

// SetVatNumberValidate gets a reference to the given UnibeeApiBeanValidResult and assigns it to the VatNumberValidate field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumberValidate(v UnibeeApiBeanValidResult) {
	o.VatNumberValidate = &v
}

func (o MerchantSubscriptionCreatePreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionCreatePreviewPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.TaxScale) {
		toSerialize["taxScale"] = o.TaxScale
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.VatCountryCode) {
		toSerialize["vatCountryCode"] = o.VatCountryCode
	}
	if !IsNil(o.VatCountryName) {
		toSerialize["vatCountryName"] = o.VatCountryName
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	if !IsNil(o.VatNumberValidate) {
		toSerialize["vatNumberValidate"] = o.VatNumberValidate
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionCreatePreviewPost200ResponseData struct {
	value *MerchantSubscriptionCreatePreviewPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionCreatePreviewPost200ResponseData) Get() *MerchantSubscriptionCreatePreviewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionCreatePreviewPost200ResponseData) Set(val *MerchantSubscriptionCreatePreviewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionCreatePreviewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionCreatePreviewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionCreatePreviewPost200ResponseData(val *MerchantSubscriptionCreatePreviewPost200ResponseData) *NullableMerchantSubscriptionCreatePreviewPost200ResponseData {
	return &NullableMerchantSubscriptionCreatePreviewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionCreatePreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionCreatePreviewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

