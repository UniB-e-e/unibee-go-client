# UnibeeApiMerchantUserListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count OF Page | [optional] 
**DeleteInclude** | Pointer to **bool** | Deleted Involved，Need Admin | [optional] 
**Email** | Pointer to **string** | Search Filter Email | [optional] 
**FirstName** | Pointer to **string** | Search FirstName | [optional] 
**LastName** | Pointer to **string** | Search LastName | [optional] 
**Page** | Pointer to **int32** | Page,Start 0 | [optional] 
**SortField** | Pointer to **string** | Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | Status, 0-Active｜2-Frozen | [optional] 
**UserId** | Pointer to **int32** | Filter UserId | [optional] 

## Methods

### NewUnibeeApiMerchantUserListReq

`func NewUnibeeApiMerchantUserListReq() *UnibeeApiMerchantUserListReq`

NewUnibeeApiMerchantUserListReq instantiates a new UnibeeApiMerchantUserListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserListReqWithDefaults

`func NewUnibeeApiMerchantUserListReqWithDefaults() *UnibeeApiMerchantUserListReq`

NewUnibeeApiMerchantUserListReqWithDefaults instantiates a new UnibeeApiMerchantUserListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantUserListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantUserListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantUserListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantUserListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDeleteInclude

`func (o *UnibeeApiMerchantUserListReq) GetDeleteInclude() bool`

GetDeleteInclude returns the DeleteInclude field if non-nil, zero value otherwise.

### GetDeleteIncludeOk

`func (o *UnibeeApiMerchantUserListReq) GetDeleteIncludeOk() (*bool, bool)`

GetDeleteIncludeOk returns a tuple with the DeleteInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInclude

`func (o *UnibeeApiMerchantUserListReq) SetDeleteInclude(v bool)`

SetDeleteInclude sets DeleteInclude field to given value.

### HasDeleteInclude

`func (o *UnibeeApiMerchantUserListReq) HasDeleteInclude() bool`

HasDeleteInclude returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantUserListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantUserListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantUserListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantUserListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiMerchantUserListReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantUserListReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantUserListReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantUserListReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantUserListReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantUserListReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantUserListReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantUserListReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantUserListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantUserListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantUserListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantUserListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantUserListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantUserListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantUserListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantUserListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantUserListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantUserListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantUserListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantUserListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantUserListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantUserListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantUserListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantUserListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserListReq) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserListReq) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserListReq) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantUserListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


