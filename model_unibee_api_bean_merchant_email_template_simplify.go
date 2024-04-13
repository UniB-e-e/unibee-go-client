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

// checks if the UnibeeApiBeanMerchantEmailTemplateSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantEmailTemplateSimplify{}

// UnibeeApiBeanMerchantEmailTemplateSimplify struct for UnibeeApiBeanMerchantEmailTemplateSimplify
type UnibeeApiBeanMerchantEmailTemplateSimplify struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	Id *int64 `json:"id,omitempty"`
	MerchantId *int64 `json:"merchantId,omitempty"`
	Status *string `json:"status,omitempty"`
	TemplateAttachName *string `json:"templateAttachName,omitempty"`
	TemplateContent *string `json:"templateContent,omitempty"`
	TemplateDescription *string `json:"templateDescription,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	TemplateTitle *string `json:"templateTitle,omitempty"`
	// update utc time
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewUnibeeApiBeanMerchantEmailTemplateSimplify instantiates a new UnibeeApiBeanMerchantEmailTemplateSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantEmailTemplateSimplify() *UnibeeApiBeanMerchantEmailTemplateSimplify {
	this := UnibeeApiBeanMerchantEmailTemplateSimplify{}
	return &this
}

// NewUnibeeApiBeanMerchantEmailTemplateSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantEmailTemplateSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantEmailTemplateSimplifyWithDefaults() *UnibeeApiBeanMerchantEmailTemplateSimplify {
	this := UnibeeApiBeanMerchantEmailTemplateSimplify{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetStatus(v string) {
	o.Status = &v
}

// GetTemplateAttachName returns the TemplateAttachName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateAttachName() string {
	if o == nil || IsNil(o.TemplateAttachName) {
		var ret string
		return ret
	}
	return *o.TemplateAttachName
}

// GetTemplateAttachNameOk returns a tuple with the TemplateAttachName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateAttachNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateAttachName) {
		return nil, false
	}
	return o.TemplateAttachName, true
}

// HasTemplateAttachName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasTemplateAttachName() bool {
	if o != nil && !IsNil(o.TemplateAttachName) {
		return true
	}

	return false
}

// SetTemplateAttachName gets a reference to the given string and assigns it to the TemplateAttachName field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetTemplateAttachName(v string) {
	o.TemplateAttachName = &v
}

// GetTemplateContent returns the TemplateContent field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateContent() string {
	if o == nil || IsNil(o.TemplateContent) {
		var ret string
		return ret
	}
	return *o.TemplateContent
}

// GetTemplateContentOk returns a tuple with the TemplateContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateContentOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateContent) {
		return nil, false
	}
	return o.TemplateContent, true
}

// HasTemplateContent returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasTemplateContent() bool {
	if o != nil && !IsNil(o.TemplateContent) {
		return true
	}

	return false
}

// SetTemplateContent gets a reference to the given string and assigns it to the TemplateContent field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetTemplateContent(v string) {
	o.TemplateContent = &v
}

// GetTemplateDescription returns the TemplateDescription field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateDescription() string {
	if o == nil || IsNil(o.TemplateDescription) {
		var ret string
		return ret
	}
	return *o.TemplateDescription
}

// GetTemplateDescriptionOk returns a tuple with the TemplateDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateDescription) {
		return nil, false
	}
	return o.TemplateDescription, true
}

// HasTemplateDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasTemplateDescription() bool {
	if o != nil && !IsNil(o.TemplateDescription) {
		return true
	}

	return false
}

// SetTemplateDescription gets a reference to the given string and assigns it to the TemplateDescription field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetTemplateDescription(v string) {
	o.TemplateDescription = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTemplateTitle returns the TemplateTitle field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateTitle() string {
	if o == nil || IsNil(o.TemplateTitle) {
		var ret string
		return ret
	}
	return *o.TemplateTitle
}

// GetTemplateTitleOk returns a tuple with the TemplateTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetTemplateTitleOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateTitle) {
		return nil, false
	}
	return o.TemplateTitle, true
}

// HasTemplateTitle returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasTemplateTitle() bool {
	if o != nil && !IsNil(o.TemplateTitle) {
		return true
	}

	return false
}

// SetTemplateTitle gets a reference to the given string and assigns it to the TemplateTitle field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetTemplateTitle(v string) {
	o.TemplateTitle = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UnibeeApiBeanMerchantEmailTemplateSimplify) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o UnibeeApiBeanMerchantEmailTemplateSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantEmailTemplateSimplify) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TemplateAttachName) {
		toSerialize["templateAttachName"] = o.TemplateAttachName
	}
	if !IsNil(o.TemplateContent) {
		toSerialize["templateContent"] = o.TemplateContent
	}
	if !IsNil(o.TemplateDescription) {
		toSerialize["templateDescription"] = o.TemplateDescription
	}
	if !IsNil(o.TemplateName) {
		toSerialize["templateName"] = o.TemplateName
	}
	if !IsNil(o.TemplateTitle) {
		toSerialize["templateTitle"] = o.TemplateTitle
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantEmailTemplateSimplify struct {
	value *UnibeeApiBeanMerchantEmailTemplateSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantEmailTemplateSimplify) Get() *UnibeeApiBeanMerchantEmailTemplateSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantEmailTemplateSimplify) Set(val *UnibeeApiBeanMerchantEmailTemplateSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantEmailTemplateSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantEmailTemplateSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantEmailTemplateSimplify(val *UnibeeApiBeanMerchantEmailTemplateSimplify) *NullableUnibeeApiBeanMerchantEmailTemplateSimplify {
	return &NullableUnibeeApiBeanMerchantEmailTemplateSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantEmailTemplateSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantEmailTemplateSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


