/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404191455 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantInvoiceMarkRefundReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceMarkRefundReq{}

// UnibeeApiMerchantInvoiceMarkRefundReq Mark invoice as refund
type UnibeeApiMerchantInvoiceMarkRefundReq struct {
	// The unique id of invoice
	InvoiceId string `json:"invoiceId"`
	// The reason of refund
	Reason string `json:"reason"`
	// The amount of refund
	RefundAmount int64 `json:"refundAmount"`
	// The out refund number
	RefundNo *string `json:"refundNo,omitempty"`
}

type _UnibeeApiMerchantInvoiceMarkRefundReq UnibeeApiMerchantInvoiceMarkRefundReq

// NewUnibeeApiMerchantInvoiceMarkRefundReq instantiates a new UnibeeApiMerchantInvoiceMarkRefundReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceMarkRefundReq(invoiceId string, reason string, refundAmount int64) *UnibeeApiMerchantInvoiceMarkRefundReq {
	this := UnibeeApiMerchantInvoiceMarkRefundReq{}
	this.InvoiceId = invoiceId
	this.Reason = reason
	this.RefundAmount = refundAmount
	return &this
}

// NewUnibeeApiMerchantInvoiceMarkRefundReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkRefundReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceMarkRefundReqWithDefaults() *UnibeeApiMerchantInvoiceMarkRefundReq {
	this := UnibeeApiMerchantInvoiceMarkRefundReq{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetReason returns the Reason field value
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) SetReason(v string) {
	o.Reason = v
}

// GetRefundAmount returns the RefundAmount field value
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetRefundAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetRefundAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundAmount, true
}

// SetRefundAmount sets field value
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) SetRefundAmount(v int64) {
	o.RefundAmount = v
}

// GetRefundNo returns the RefundNo field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetRefundNo() string {
	if o == nil || IsNil(o.RefundNo) {
		var ret string
		return ret
	}
	return *o.RefundNo
}

// GetRefundNoOk returns a tuple with the RefundNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) GetRefundNoOk() (*string, bool) {
	if o == nil || IsNil(o.RefundNo) {
		return nil, false
	}
	return o.RefundNo, true
}

// HasRefundNo returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) HasRefundNo() bool {
	if o != nil && !IsNil(o.RefundNo) {
		return true
	}

	return false
}

// SetRefundNo gets a reference to the given string and assigns it to the RefundNo field.
func (o *UnibeeApiMerchantInvoiceMarkRefundReq) SetRefundNo(v string) {
	o.RefundNo = &v
}

func (o UnibeeApiMerchantInvoiceMarkRefundReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceMarkRefundReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoiceId"] = o.InvoiceId
	toSerialize["reason"] = o.Reason
	toSerialize["refundAmount"] = o.RefundAmount
	if !IsNil(o.RefundNo) {
		toSerialize["refundNo"] = o.RefundNo
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantInvoiceMarkRefundReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoiceId",
		"reason",
		"refundAmount",
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

	varUnibeeApiMerchantInvoiceMarkRefundReq := _UnibeeApiMerchantInvoiceMarkRefundReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantInvoiceMarkRefundReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantInvoiceMarkRefundReq(varUnibeeApiMerchantInvoiceMarkRefundReq)

	return err
}

type NullableUnibeeApiMerchantInvoiceMarkRefundReq struct {
	value *UnibeeApiMerchantInvoiceMarkRefundReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundReq) Get() *UnibeeApiMerchantInvoiceMarkRefundReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundReq) Set(val *UnibeeApiMerchantInvoiceMarkRefundReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceMarkRefundReq(val *UnibeeApiMerchantInvoiceMarkRefundReq) *NullableUnibeeApiMerchantInvoiceMarkRefundReq {
	return &NullableUnibeeApiMerchantInvoiceMarkRefundReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceMarkRefundReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceMarkRefundReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

