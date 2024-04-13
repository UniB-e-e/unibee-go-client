/*
OpenAPI UniBee

This is UniBee api server

API version: buildtime:202404131246 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanCurrency{}

// UnibeeApiBeanCurrency struct for UnibeeApiBeanCurrency
type UnibeeApiBeanCurrency struct {
	Currency *string `json:"Currency,omitempty"`
	Scale *int32 `json:"Scale,omitempty"`
	Symbol *string `json:"Symbol,omitempty"`
}

// NewUnibeeApiBeanCurrency instantiates a new UnibeeApiBeanCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanCurrency() *UnibeeApiBeanCurrency {
	this := UnibeeApiBeanCurrency{}
	return &this
}

// NewUnibeeApiBeanCurrencyWithDefaults instantiates a new UnibeeApiBeanCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanCurrencyWithDefaults() *UnibeeApiBeanCurrency {
	this := UnibeeApiBeanCurrency{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanCurrency) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCurrency) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanCurrency) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanCurrency) SetCurrency(v string) {
	o.Currency = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *UnibeeApiBeanCurrency) GetScale() int32 {
	if o == nil || IsNil(o.Scale) {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCurrency) GetScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.Scale) {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *UnibeeApiBeanCurrency) HasScale() bool {
	if o != nil && !IsNil(o.Scale) {
		return true
	}

	return false
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *UnibeeApiBeanCurrency) SetScale(v int32) {
	o.Scale = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UnibeeApiBeanCurrency) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCurrency) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UnibeeApiBeanCurrency) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UnibeeApiBeanCurrency) SetSymbol(v string) {
	o.Symbol = &v
}

func (o UnibeeApiBeanCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["Currency"] = o.Currency
	}
	if !IsNil(o.Scale) {
		toSerialize["Scale"] = o.Scale
	}
	if !IsNil(o.Symbol) {
		toSerialize["Symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanCurrency struct {
	value *UnibeeApiBeanCurrency
	isSet bool
}

func (v NullableUnibeeApiBeanCurrency) Get() *UnibeeApiBeanCurrency {
	return v.value
}

func (v *NullableUnibeeApiBeanCurrency) Set(val *UnibeeApiBeanCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanCurrency(val *UnibeeApiBeanCurrency) *NullableUnibeeApiBeanCurrency {
	return &NullableUnibeeApiBeanCurrency{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


