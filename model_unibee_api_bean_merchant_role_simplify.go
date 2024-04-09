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

// checks if the UnibeeApiBeanMerchantRoleSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantRoleSimplify{}

// UnibeeApiBeanMerchantRoleSimplify struct for UnibeeApiBeanMerchantRoleSimplify
type UnibeeApiBeanMerchantRoleSimplify struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// userId
	Id *int64 `json:"id,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// permissions
	Permissions []UnibeeApiBeanMerchantRolePermission `json:"permissions,omitempty"`
	// role
	Role *string `json:"role,omitempty"`
}

// NewUnibeeApiBeanMerchantRoleSimplify instantiates a new UnibeeApiBeanMerchantRoleSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantRoleSimplify() *UnibeeApiBeanMerchantRoleSimplify {
	this := UnibeeApiBeanMerchantRoleSimplify{}
	return &this
}

// NewUnibeeApiBeanMerchantRoleSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantRoleSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantRoleSimplifyWithDefaults() *UnibeeApiBeanMerchantRoleSimplify {
	this := UnibeeApiBeanMerchantRoleSimplify{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantRoleSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantRoleSimplify) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantRoleSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetPermissions() []UnibeeApiBeanMerchantRolePermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []UnibeeApiBeanMerchantRolePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetPermissionsOk() ([]UnibeeApiBeanMerchantRolePermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UnibeeApiBeanMerchantRolePermission and assigns it to the Permissions field.
func (o *UnibeeApiBeanMerchantRoleSimplify) SetPermissions(v []UnibeeApiBeanMerchantRolePermission) {
	o.Permissions = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRoleSimplify) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UnibeeApiBeanMerchantRoleSimplify) SetRole(v string) {
	o.Role = &v
}

func (o UnibeeApiBeanMerchantRoleSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantRoleSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantRoleSimplify struct {
	value *UnibeeApiBeanMerchantRoleSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantRoleSimplify) Get() *UnibeeApiBeanMerchantRoleSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantRoleSimplify) Set(val *UnibeeApiBeanMerchantRoleSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantRoleSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantRoleSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantRoleSimplify(val *UnibeeApiBeanMerchantRoleSimplify) *NullableUnibeeApiBeanMerchantRoleSimplify {
	return &NullableUnibeeApiBeanMerchantRoleSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantRoleSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantRoleSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


