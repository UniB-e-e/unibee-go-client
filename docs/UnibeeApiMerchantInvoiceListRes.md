# UnibeeApiMerchantInvoiceListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | Pointer to [**[]UnibeeApiBeanDetailInvoiceDetail**](UnibeeApiBeanDetailInvoiceDetail.md) | Invoice Detail Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceListRes

`func NewUnibeeApiMerchantInvoiceListRes() *UnibeeApiMerchantInvoiceListRes`

NewUnibeeApiMerchantInvoiceListRes instantiates a new UnibeeApiMerchantInvoiceListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceListResWithDefaults

`func NewUnibeeApiMerchantInvoiceListResWithDefaults() *UnibeeApiMerchantInvoiceListRes`

NewUnibeeApiMerchantInvoiceListResWithDefaults instantiates a new UnibeeApiMerchantInvoiceListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *UnibeeApiMerchantInvoiceListRes) GetInvoices() []UnibeeApiBeanDetailInvoiceDetail`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *UnibeeApiMerchantInvoiceListRes) GetInvoicesOk() (*[]UnibeeApiBeanDetailInvoiceDetail, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *UnibeeApiMerchantInvoiceListRes) SetInvoices(v []UnibeeApiBeanDetailInvoiceDetail)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *UnibeeApiMerchantInvoiceListRes) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantInvoiceListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantInvoiceListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantInvoiceListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantInvoiceListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


