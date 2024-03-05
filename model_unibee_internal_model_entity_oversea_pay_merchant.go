/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalModelEntityOverseaPayMerchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalModelEntityOverseaPayMerchant{}

// UnibeeInternalModelEntityOverseaPayMerchant struct for UnibeeInternalModelEntityOverseaPayMerchant
type UnibeeInternalModelEntityOverseaPayMerchant struct {
	// address
	Address *string `json:"address,omitempty"`
	// merchant open api key
	ApiKey *string `json:"apiKey,omitempty"`
	// business_num
	BusinessNum *string `json:"businessNum,omitempty"`
	// company_id
	CompanyId *int64 `json:"companyId,omitempty"`
	// company_logo
	CompanyLogo *string `json:"companyLogo,omitempty"`
	// company_name
	CompanyName *string `json:"companyName,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// email
	Email *string `json:"email,omitempty"`
	// create time
	GmtCreate *string `json:"gmtCreate,omitempty"`
	// update_time
	GmtModify *string `json:"gmtModify,omitempty"`
	HomeUrl *string `json:"homeUrl,omitempty"`
	// merchant user portal host
	Host *string `json:"host,omitempty"`
	// merchant_id
	Id *int32 `json:"id,omitempty"`
	// idcard
	Idcard *string `json:"idcard,omitempty"`
	// 0-UnDeleted，1-Deleted
	IsDeleted *int32 `json:"isDeleted,omitempty"`
	// location
	Location *string `json:"location,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// phone
	Phone *string `json:"phone,omitempty"`
	// merchant default time zone
	TimeZone *string `json:"timeZone,omitempty"`
	// type
	Type *int32 `json:"type,omitempty"`
	// create_user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeInternalModelEntityOverseaPayMerchant instantiates a new UnibeeInternalModelEntityOverseaPayMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalModelEntityOverseaPayMerchant() *UnibeeInternalModelEntityOverseaPayMerchant {
	this := UnibeeInternalModelEntityOverseaPayMerchant{}
	return &this
}

// NewUnibeeInternalModelEntityOverseaPayMerchantWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalModelEntityOverseaPayMerchantWithDefaults() *UnibeeInternalModelEntityOverseaPayMerchant {
	this := UnibeeInternalModelEntityOverseaPayMerchant{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetAddress(v string) {
	o.Address = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetBusinessNum returns the BusinessNum field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetBusinessNum() string {
	if o == nil || IsNil(o.BusinessNum) {
		var ret string
		return ret
	}
	return *o.BusinessNum
}

// GetBusinessNumOk returns a tuple with the BusinessNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetBusinessNumOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessNum) {
		return nil, false
	}
	return o.BusinessNum, true
}

// HasBusinessNum returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasBusinessNum() bool {
	if o != nil && !IsNil(o.BusinessNum) {
		return true
	}

	return false
}

// SetBusinessNum gets a reference to the given string and assigns it to the BusinessNum field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetBusinessNum(v string) {
	o.BusinessNum = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCompanyId() int64 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int64
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCompanyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int64 and assigns it to the CompanyId field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetCompanyId(v int64) {
	o.CompanyId = &v
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCompanyLogo() string {
	if o == nil || IsNil(o.CompanyLogo) {
		var ret string
		return ret
	}
	return *o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCompanyLogoOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyLogo) {
		return nil, false
	}
	return o.CompanyLogo, true
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasCompanyLogo() bool {
	if o != nil && !IsNil(o.CompanyLogo) {
		return true
	}

	return false
}

// SetCompanyLogo gets a reference to the given string and assigns it to the CompanyLogo field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetCompanyLogo(v string) {
	o.CompanyLogo = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetEmail(v string) {
	o.Email = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModify returns the GmtModify field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetGmtModify() string {
	if o == nil || IsNil(o.GmtModify) {
		var ret string
		return ret
	}
	return *o.GmtModify
}

// GetGmtModifyOk returns a tuple with the GmtModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetGmtModifyOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModify) {
		return nil, false
	}
	return o.GmtModify, true
}

// HasGmtModify returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasGmtModify() bool {
	if o != nil && !IsNil(o.GmtModify) {
		return true
	}

	return false
}

// SetGmtModify gets a reference to the given string and assigns it to the GmtModify field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetGmtModify(v string) {
	o.GmtModify = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetId(v int32) {
	o.Id = &v
}

// GetIdcard returns the Idcard field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetIdcard() string {
	if o == nil || IsNil(o.Idcard) {
		var ret string
		return ret
	}
	return *o.Idcard
}

// GetIdcardOk returns a tuple with the Idcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetIdcardOk() (*string, bool) {
	if o == nil || IsNil(o.Idcard) {
		return nil, false
	}
	return o.Idcard, true
}

// HasIdcard returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasIdcard() bool {
	if o != nil && !IsNil(o.Idcard) {
		return true
	}

	return false
}

// SetIdcard gets a reference to the given string and assigns it to the Idcard field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetIdcard(v string) {
	o.Idcard = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetIsDeleted() int32 {
	if o == nil || IsNil(o.IsDeleted) {
		var ret int32
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetIsDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given int32 and assigns it to the IsDeleted field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetIsDeleted(v int32) {
	o.IsDeleted = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetPhone(v string) {
	o.Phone = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetType(v int32) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeInternalModelEntityOverseaPayMerchant) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeInternalModelEntityOverseaPayMerchant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalModelEntityOverseaPayMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.BusinessNum) {
		toSerialize["businessNum"] = o.BusinessNum
	}
	if !IsNil(o.CompanyId) {
		toSerialize["companyId"] = o.CompanyId
	}
	if !IsNil(o.CompanyLogo) {
		toSerialize["companyLogo"] = o.CompanyLogo
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmtCreate"] = o.GmtCreate
	}
	if !IsNil(o.GmtModify) {
		toSerialize["gmtModify"] = o.GmtModify
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Idcard) {
		toSerialize["idcard"] = o.Idcard
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeInternalModelEntityOverseaPayMerchant struct {
	value *UnibeeInternalModelEntityOverseaPayMerchant
	isSet bool
}

func (v NullableUnibeeInternalModelEntityOverseaPayMerchant) Get() *UnibeeInternalModelEntityOverseaPayMerchant {
	return v.value
}

func (v *NullableUnibeeInternalModelEntityOverseaPayMerchant) Set(val *UnibeeInternalModelEntityOverseaPayMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalModelEntityOverseaPayMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalModelEntityOverseaPayMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalModelEntityOverseaPayMerchant(val *UnibeeInternalModelEntityOverseaPayMerchant) *NullableUnibeeInternalModelEntityOverseaPayMerchant {
	return &NullableUnibeeInternalModelEntityOverseaPayMerchant{value: val, isSet: true}
}

func (v NullableUnibeeInternalModelEntityOverseaPayMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalModelEntityOverseaPayMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

