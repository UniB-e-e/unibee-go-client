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

// checks if the UnibeeApiMerchantUserListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserListReq{}

// UnibeeApiMerchantUserListReq struct for UnibeeApiMerchantUserListReq
type UnibeeApiMerchantUserListReq struct {
	// Count OF Page
	Count *int32 `json:"count,omitempty"`
	// Deleted Involved，Need Admin
	DeleteInclude *bool `json:"deleteInclude,omitempty"`
	// Search Filter Email
	Email *string `json:"email,omitempty"`
	// Search FirstName
	FirstName *string `json:"firstName,omitempty"`
	// Search LastName
	LastName *string `json:"lastName,omitempty"`
	// Page,Start 0
	Page *int32 `json:"page,omitempty"`
	// Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// Status, 0-Active｜2-Frozen
	Status []int32 `json:"status,omitempty"`
	// Filter UserId
	UserId *int32 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantUserListReq instantiates a new UnibeeApiMerchantUserListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserListReq() *UnibeeApiMerchantUserListReq {
	this := UnibeeApiMerchantUserListReq{}
	return &this
}

// NewUnibeeApiMerchantUserListReqWithDefaults instantiates a new UnibeeApiMerchantUserListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserListReqWithDefaults() *UnibeeApiMerchantUserListReq {
	this := UnibeeApiMerchantUserListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantUserListReq) SetCount(v int32) {
	o.Count = &v
}

// GetDeleteInclude returns the DeleteInclude field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetDeleteInclude() bool {
	if o == nil || IsNil(o.DeleteInclude) {
		var ret bool
		return ret
	}
	return *o.DeleteInclude
}

// GetDeleteIncludeOk returns a tuple with the DeleteInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetDeleteIncludeOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteInclude) {
		return nil, false
	}
	return o.DeleteInclude, true
}

// HasDeleteInclude returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasDeleteInclude() bool {
	if o != nil && !IsNil(o.DeleteInclude) {
		return true
	}

	return false
}

// SetDeleteInclude gets a reference to the given bool and assigns it to the DeleteInclude field.
func (o *UnibeeApiMerchantUserListReq) SetDeleteInclude(v bool) {
	o.DeleteInclude = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantUserListReq) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiMerchantUserListReq) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiMerchantUserListReq) SetLastName(v string) {
	o.LastName = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantUserListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantUserListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantUserListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetStatus() []int32 {
	if o == nil || IsNil(o.Status) {
		var ret []int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetStatusOk() ([]int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantUserListReq) SetStatus(v []int32) {
	o.Status = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserListReq) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserListReq) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *UnibeeApiMerchantUserListReq) SetUserId(v int32) {
	o.UserId = &v
}

func (o UnibeeApiMerchantUserListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.DeleteInclude) {
		toSerialize["deleteInclude"] = o.DeleteInclude
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.SortField) {
		toSerialize["sortField"] = o.SortField
	}
	if !IsNil(o.SortType) {
		toSerialize["sortType"] = o.SortType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserListReq struct {
	value *UnibeeApiMerchantUserListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserListReq) Get() *UnibeeApiMerchantUserListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserListReq) Set(val *UnibeeApiMerchantUserListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserListReq(val *UnibeeApiMerchantUserListReq) *NullableUnibeeApiMerchantUserListReq {
	return &NullableUnibeeApiMerchantUserListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


