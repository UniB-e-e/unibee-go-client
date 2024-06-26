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

// checks if the UnibeeApiBeanMerchantMemberSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantMemberSimplify{}

// UnibeeApiBeanMerchantMemberSimplify struct for UnibeeApiBeanMerchantMemberSimplify
type UnibeeApiBeanMerchantMemberSimplify struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// email
	Email *string `json:"email,omitempty"`
	// first name
	FirstName *string `json:"firstName,omitempty"`
	// userId
	Id *int64 `json:"id,omitempty"`
	// last name
	LastName *string `json:"lastName,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// mobile
	Mobile *string `json:"mobile,omitempty"`
	// role
	Role *string `json:"role,omitempty"`
}

// NewUnibeeApiBeanMerchantMemberSimplify instantiates a new UnibeeApiBeanMerchantMemberSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantMemberSimplify() *UnibeeApiBeanMerchantMemberSimplify {
	this := UnibeeApiBeanMerchantMemberSimplify{}
	return &this
}

// NewUnibeeApiBeanMerchantMemberSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantMemberSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantMemberSimplifyWithDefaults() *UnibeeApiBeanMerchantMemberSimplify {
	this := UnibeeApiBeanMerchantMemberSimplify{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetFirstName(v string) {
	o.FirstName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetId(v int64) {
	o.Id = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetLastName(v string) {
	o.LastName = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetMobile(v string) {
	o.Mobile = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMemberSimplify) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UnibeeApiBeanMerchantMemberSimplify) SetRole(v string) {
	o.Role = &v
}

func (o UnibeeApiBeanMerchantMemberSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantMemberSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantMemberSimplify struct {
	value *UnibeeApiBeanMerchantMemberSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantMemberSimplify) Get() *UnibeeApiBeanMerchantMemberSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantMemberSimplify) Set(val *UnibeeApiBeanMerchantMemberSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantMemberSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantMemberSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantMemberSimplify(val *UnibeeApiBeanMerchantMemberSimplify) *NullableUnibeeApiBeanMerchantMemberSimplify {
	return &NullableUnibeeApiBeanMerchantMemberSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantMemberSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantMemberSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


