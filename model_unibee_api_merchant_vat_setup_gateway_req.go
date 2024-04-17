/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404171839 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantVatSetupGatewayReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantVatSetupGatewayReq{}

// UnibeeApiMerchantVatSetupGatewayReq struct for UnibeeApiMerchantVatSetupGatewayReq
type UnibeeApiMerchantVatSetupGatewayReq struct {
	// IsDefault, default is true
	IsDefault *bool `json:"IsDefault,omitempty"`
	// Data
	Data string `json:"data"`
	// GatewayName, em. vatsense
	GatewayName string `json:"gatewayName"`
}

type _UnibeeApiMerchantVatSetupGatewayReq UnibeeApiMerchantVatSetupGatewayReq

// NewUnibeeApiMerchantVatSetupGatewayReq instantiates a new UnibeeApiMerchantVatSetupGatewayReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantVatSetupGatewayReq(data string, gatewayName string) *UnibeeApiMerchantVatSetupGatewayReq {
	this := UnibeeApiMerchantVatSetupGatewayReq{}
	var isDefault bool = true
	this.IsDefault = &isDefault
	this.Data = data
	this.GatewayName = gatewayName
	return &this
}

// NewUnibeeApiMerchantVatSetupGatewayReqWithDefaults instantiates a new UnibeeApiMerchantVatSetupGatewayReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantVatSetupGatewayReqWithDefaults() *UnibeeApiMerchantVatSetupGatewayReq {
	this := UnibeeApiMerchantVatSetupGatewayReq{}
	var isDefault bool = true
	this.IsDefault = &isDefault
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *UnibeeApiMerchantVatSetupGatewayReq) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantVatSetupGatewayReq) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *UnibeeApiMerchantVatSetupGatewayReq) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *UnibeeApiMerchantVatSetupGatewayReq) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetData returns the Data field value
func (o *UnibeeApiMerchantVatSetupGatewayReq) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantVatSetupGatewayReq) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UnibeeApiMerchantVatSetupGatewayReq) SetData(v string) {
	o.Data = v
}

// GetGatewayName returns the GatewayName field value
func (o *UnibeeApiMerchantVatSetupGatewayReq) GetGatewayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantVatSetupGatewayReq) GetGatewayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayName, true
}

// SetGatewayName sets field value
func (o *UnibeeApiMerchantVatSetupGatewayReq) SetGatewayName(v string) {
	o.GatewayName = v
}

func (o UnibeeApiMerchantVatSetupGatewayReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantVatSetupGatewayReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsDefault) {
		toSerialize["IsDefault"] = o.IsDefault
	}
	toSerialize["data"] = o.Data
	toSerialize["gatewayName"] = o.GatewayName
	return toSerialize, nil
}

func (o *UnibeeApiMerchantVatSetupGatewayReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"gatewayName",
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

	varUnibeeApiMerchantVatSetupGatewayReq := _UnibeeApiMerchantVatSetupGatewayReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantVatSetupGatewayReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantVatSetupGatewayReq(varUnibeeApiMerchantVatSetupGatewayReq)

	return err
}

type NullableUnibeeApiMerchantVatSetupGatewayReq struct {
	value *UnibeeApiMerchantVatSetupGatewayReq
	isSet bool
}

func (v NullableUnibeeApiMerchantVatSetupGatewayReq) Get() *UnibeeApiMerchantVatSetupGatewayReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantVatSetupGatewayReq) Set(val *UnibeeApiMerchantVatSetupGatewayReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantVatSetupGatewayReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantVatSetupGatewayReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantVatSetupGatewayReq(val *UnibeeApiMerchantVatSetupGatewayReq) *NullableUnibeeApiMerchantVatSetupGatewayReq {
	return &NullableUnibeeApiMerchantVatSetupGatewayReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantVatSetupGatewayReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantVatSetupGatewayReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


