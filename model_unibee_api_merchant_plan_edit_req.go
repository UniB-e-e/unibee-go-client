/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanEditReq{}

// UnibeeApiMerchantPlanEditReq struct for UnibeeApiMerchantPlanEditReq
type UnibeeApiMerchantPlanEditReq struct {
	// Plan Ids Of Addon Type
	AddonIds []int64 `json:"addonIds,omitempty"`
	// Plan CaptureAmount
	Amount int64 `json:"amount"`
	// Plan Currency
	Currency string `json:"currency"`
	// Description
	Description *string `json:"description,omitempty"`
	// who pay the gas for crypto payment, merchant|user
	GasPayer *string `json:"gasPayer,omitempty"`
	// HomeUrl,Start With: http
	HomeUrl *string `json:"homeUrl,omitempty"`
	// ImageUrl,Start With: http
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Number Of IntervalUnit
	IntervalCount *int32 `json:"intervalCount,omitempty"`
	// Plan Interval Unit，em: day|month|year|week
	IntervalUnit string `json:"intervalUnit"`
	// Metadata，Map
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Plan's MetricLimit List
	MetricLimits []UnibeeApiBeanBulkMetricLimitPlanBindingParam `json:"metricLimits,omitempty"`
	// PlanId
	PlanId int64 `json:"planId"`
	// Plan Name
	PlanName string `json:"planName"`
	// Default Copy Description
	ProductDescription *string `json:"productDescription,omitempty"`
	// Default Copy PlanName
	ProductName *string `json:"productName,omitempty"`
}

type _UnibeeApiMerchantPlanEditReq UnibeeApiMerchantPlanEditReq

// NewUnibeeApiMerchantPlanEditReq instantiates a new UnibeeApiMerchantPlanEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanEditReq(amount int64, currency string, intervalUnit string, planId int64, planName string) *UnibeeApiMerchantPlanEditReq {
	this := UnibeeApiMerchantPlanEditReq{}
	this.Amount = amount
	this.Currency = currency
	this.IntervalUnit = intervalUnit
	this.PlanId = planId
	this.PlanName = planName
	return &this
}

// NewUnibeeApiMerchantPlanEditReqWithDefaults instantiates a new UnibeeApiMerchantPlanEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanEditReqWithDefaults() *UnibeeApiMerchantPlanEditReq {
	this := UnibeeApiMerchantPlanEditReq{}
	return &this
}

// GetAddonIds returns the AddonIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetAddonIds() []int64 {
	if o == nil || IsNil(o.AddonIds) {
		var ret []int64
		return ret
	}
	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AddonIds) {
		return nil, false
	}
	return o.AddonIds, true
}

// HasAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasAddonIds() bool {
	if o != nil && !IsNil(o.AddonIds) {
		return true
	}

	return false
}

// SetAddonIds gets a reference to the given []int64 and assigns it to the AddonIds field.
func (o *UnibeeApiMerchantPlanEditReq) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetAmount returns the Amount field value
func (o *UnibeeApiMerchantPlanEditReq) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UnibeeApiMerchantPlanEditReq) SetAmount(v int64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantPlanEditReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantPlanEditReq) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantPlanEditReq) SetDescription(v string) {
	o.Description = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiMerchantPlanEditReq) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiMerchantPlanEditReq) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiMerchantPlanEditReq) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *UnibeeApiMerchantPlanEditReq) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetIntervalUnit returns the IntervalUnit field value
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntervalUnit
}

// GetIntervalUnitOk returns a tuple with the IntervalUnit field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntervalUnit, true
}

// SetIntervalUnit sets field value
func (o *UnibeeApiMerchantPlanEditReq) SetIntervalUnit(v string) {
	o.IntervalUnit = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiMerchantPlanEditReq) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMetricLimits returns the MetricLimits field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimits() []UnibeeApiBeanBulkMetricLimitPlanBindingParam {
	if o == nil || IsNil(o.MetricLimits) {
		var ret []UnibeeApiBeanBulkMetricLimitPlanBindingParam
		return ret
	}
	return o.MetricLimits
}

// GetMetricLimitsOk returns a tuple with the MetricLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimitsOk() ([]UnibeeApiBeanBulkMetricLimitPlanBindingParam, bool) {
	if o == nil || IsNil(o.MetricLimits) {
		return nil, false
	}
	return o.MetricLimits, true
}

// HasMetricLimits returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasMetricLimits() bool {
	if o != nil && !IsNil(o.MetricLimits) {
		return true
	}

	return false
}

// SetMetricLimits gets a reference to the given []UnibeeApiBeanBulkMetricLimitPlanBindingParam and assigns it to the MetricLimits field.
func (o *UnibeeApiMerchantPlanEditReq) SetMetricLimits(v []UnibeeApiBeanBulkMetricLimitPlanBindingParam) {
	o.MetricLimits = v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanEditReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanEditReq) SetPlanId(v int64) {
	o.PlanId = v
}

// GetPlanName returns the PlanName field value
func (o *UnibeeApiMerchantPlanEditReq) GetPlanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanName, true
}

// SetPlanName sets field value
func (o *UnibeeApiMerchantPlanEditReq) SetPlanName(v string) {
	o.PlanName = v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetProductDescription() string {
	if o == nil || IsNil(o.ProductDescription) {
		var ret string
		return ret
	}
	return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDescription) {
		return nil, false
	}
	return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasProductDescription() bool {
	if o != nil && !IsNil(o.ProductDescription) {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *UnibeeApiMerchantPlanEditReq) SetProductDescription(v string) {
	o.ProductDescription = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *UnibeeApiMerchantPlanEditReq) SetProductName(v string) {
	o.ProductName = &v
}

func (o UnibeeApiMerchantPlanEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonIds) {
		toSerialize["addonIds"] = o.AddonIds
	}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GasPayer) {
		toSerialize["gasPayer"] = o.GasPayer
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.IntervalCount) {
		toSerialize["intervalCount"] = o.IntervalCount
	}
	toSerialize["intervalUnit"] = o.IntervalUnit
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MetricLimits) {
		toSerialize["metricLimits"] = o.MetricLimits
	}
	toSerialize["planId"] = o.PlanId
	toSerialize["planName"] = o.PlanName
	if !IsNil(o.ProductDescription) {
		toSerialize["productDescription"] = o.ProductDescription
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"intervalUnit",
		"planId",
		"planName",
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

	varUnibeeApiMerchantPlanEditReq := _UnibeeApiMerchantPlanEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanEditReq(varUnibeeApiMerchantPlanEditReq)

	return err
}

type NullableUnibeeApiMerchantPlanEditReq struct {
	value *UnibeeApiMerchantPlanEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanEditReq) Get() *UnibeeApiMerchantPlanEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanEditReq) Set(val *UnibeeApiMerchantPlanEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanEditReq(val *UnibeeApiMerchantPlanEditReq) *NullableUnibeeApiMerchantPlanEditReq {
	return &NullableUnibeeApiMerchantPlanEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


