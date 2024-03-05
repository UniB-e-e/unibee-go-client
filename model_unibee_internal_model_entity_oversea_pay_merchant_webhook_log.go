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

// checks if the UnibeeInternalModelEntityOverseaPayMerchantWebhookLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalModelEntityOverseaPayMerchantWebhookLog{}

// UnibeeInternalModelEntityOverseaPayMerchantWebhookLog struct for UnibeeInternalModelEntityOverseaPayMerchantWebhookLog
type UnibeeInternalModelEntityOverseaPayMerchantWebhookLog struct {
	// body(json)
	Body *string `json:"body,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	EndpointId *int64 `json:"endpointId,omitempty"`
	// create time
	GmtCreate *string `json:"gmtCreate,omitempty"`
	// update time
	GmtModify *string `json:"gmtModify,omitempty"`
	// id
	Id *int32 `json:"id,omitempty"`
	// mamo
	Mamo *string `json:"mamo,omitempty"`
	// webhook url
	MerchantId *int32 `json:"merchantId,omitempty"`
	ReconsumeCount *int32 `json:"reconsumeCount,omitempty"`
	// request_id
	RequestId *string `json:"requestId,omitempty"`
	// response
	Response *string `json:"response,omitempty"`
	// webhook_event
	WebhookEvent *string `json:"webhookEvent,omitempty"`
	// webhook url
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

// NewUnibeeInternalModelEntityOverseaPayMerchantWebhookLog instantiates a new UnibeeInternalModelEntityOverseaPayMerchantWebhookLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalModelEntityOverseaPayMerchantWebhookLog() *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog {
	this := UnibeeInternalModelEntityOverseaPayMerchantWebhookLog{}
	return &this
}

// NewUnibeeInternalModelEntityOverseaPayMerchantWebhookLogWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayMerchantWebhookLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalModelEntityOverseaPayMerchantWebhookLogWithDefaults() *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog {
	this := UnibeeInternalModelEntityOverseaPayMerchantWebhookLog{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetBody(v string) {
	o.Body = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetEndpointId() int64 {
	if o == nil || IsNil(o.EndpointId) {
		var ret int64
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetEndpointIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given int64 and assigns it to the EndpointId field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetEndpointId(v int64) {
	o.EndpointId = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModify returns the GmtModify field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetGmtModify() string {
	if o == nil || IsNil(o.GmtModify) {
		var ret string
		return ret
	}
	return *o.GmtModify
}

// GetGmtModifyOk returns a tuple with the GmtModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetGmtModifyOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModify) {
		return nil, false
	}
	return o.GmtModify, true
}

// HasGmtModify returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasGmtModify() bool {
	if o != nil && !IsNil(o.GmtModify) {
		return true
	}

	return false
}

// SetGmtModify gets a reference to the given string and assigns it to the GmtModify field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetGmtModify(v string) {
	o.GmtModify = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetId(v int32) {
	o.Id = &v
}

// GetMamo returns the Mamo field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetMamo() string {
	if o == nil || IsNil(o.Mamo) {
		var ret string
		return ret
	}
	return *o.Mamo
}

// GetMamoOk returns a tuple with the Mamo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetMamoOk() (*string, bool) {
	if o == nil || IsNil(o.Mamo) {
		return nil, false
	}
	return o.Mamo, true
}

// HasMamo returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasMamo() bool {
	if o != nil && !IsNil(o.Mamo) {
		return true
	}

	return false
}

// SetMamo gets a reference to the given string and assigns it to the Mamo field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetMamo(v string) {
	o.Mamo = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetMerchantId() int32 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int32
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetMerchantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int32 and assigns it to the MerchantId field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetMerchantId(v int32) {
	o.MerchantId = &v
}

// GetReconsumeCount returns the ReconsumeCount field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetReconsumeCount() int32 {
	if o == nil || IsNil(o.ReconsumeCount) {
		var ret int32
		return ret
	}
	return *o.ReconsumeCount
}

// GetReconsumeCountOk returns a tuple with the ReconsumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetReconsumeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReconsumeCount) {
		return nil, false
	}
	return o.ReconsumeCount, true
}

// HasReconsumeCount returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasReconsumeCount() bool {
	if o != nil && !IsNil(o.ReconsumeCount) {
		return true
	}

	return false
}

// SetReconsumeCount gets a reference to the given int32 and assigns it to the ReconsumeCount field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetReconsumeCount(v int32) {
	o.ReconsumeCount = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetResponse(v string) {
	o.Response = &v
}

// GetWebhookEvent returns the WebhookEvent field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetWebhookEvent() string {
	if o == nil || IsNil(o.WebhookEvent) {
		var ret string
		return ret
	}
	return *o.WebhookEvent
}

// GetWebhookEventOk returns a tuple with the WebhookEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetWebhookEventOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEvent) {
		return nil, false
	}
	return o.WebhookEvent, true
}

// HasWebhookEvent returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasWebhookEvent() bool {
	if o != nil && !IsNil(o.WebhookEvent) {
		return true
	}

	return false
}

// SetWebhookEvent gets a reference to the given string and assigns it to the WebhookEvent field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetWebhookEvent(v string) {
	o.WebhookEvent = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.EndpointId) {
		toSerialize["endpointId"] = o.EndpointId
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmtCreate"] = o.GmtCreate
	}
	if !IsNil(o.GmtModify) {
		toSerialize["gmtModify"] = o.GmtModify
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mamo) {
		toSerialize["mamo"] = o.Mamo
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.ReconsumeCount) {
		toSerialize["reconsumeCount"] = o.ReconsumeCount
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.WebhookEvent) {
		toSerialize["webhookEvent"] = o.WebhookEvent
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return toSerialize, nil
}

type NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog struct {
	value *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog
	isSet bool
}

func (v NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog) Get() *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog {
	return v.value
}

func (v *NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog) Set(val *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog(val *UnibeeInternalModelEntityOverseaPayMerchantWebhookLog) *NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog {
	return &NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog{value: val, isSet: true}
}

func (v NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalModelEntityOverseaPayMerchantWebhookLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

