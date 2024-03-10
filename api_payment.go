/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// PaymentService Payment service
type PaymentService service

type PaymentPaymentCancelPostRequest struct {
	ctx context.Context
	ApiService *PaymentService
	unibeeApiMerchantPaymentCancelReq *UnibeeApiMerchantPaymentCancelReq
}

func (r PaymentPaymentCancelPostRequest) UnibeeApiMerchantPaymentCancelReq(unibeeApiMerchantPaymentCancelReq UnibeeApiMerchantPaymentCancelReq) PaymentPaymentCancelPostRequest {
	r.unibeeApiMerchantPaymentCancelReq = &unibeeApiMerchantPaymentCancelReq
	return r
}

func (r PaymentPaymentCancelPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.PaymentCancelPostExecute(r)
}

/*
PaymentCancelPost Cancel Payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentCancelPostRequest
*/
func (a *PaymentService) PaymentCancelPost(ctx context.Context) PaymentPaymentCancelPostRequest {
	return PaymentPaymentCancelPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *PaymentService) PaymentCancelPostExecute(r PaymentPaymentCancelPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentCancelPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantPaymentCancelReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantPaymentCancelReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantPaymentCancelReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentCapturePostRequest struct {
	ctx context.Context
	ApiService *PaymentService
	captureAmount *int64
	currency *string
	merchantPaymentCapturePostRequest *MerchantPaymentCapturePostRequest
}

// CaptureAmount, Cent
func (r PaymentPaymentCapturePostRequest) CaptureAmount(captureAmount int64) PaymentPaymentCapturePostRequest {
	r.captureAmount = &captureAmount
	return r
}

// Currency
func (r PaymentPaymentCapturePostRequest) Currency(currency string) PaymentPaymentCapturePostRequest {
	r.currency = &currency
	return r
}

func (r PaymentPaymentCapturePostRequest) MerchantPaymentCapturePostRequest(merchantPaymentCapturePostRequest MerchantPaymentCapturePostRequest) PaymentPaymentCapturePostRequest {
	r.merchantPaymentCapturePostRequest = &merchantPaymentCapturePostRequest
	return r
}

func (r PaymentPaymentCapturePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.PaymentCapturePostExecute(r)
}

/*
PaymentCapturePost Capture Payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentCapturePostRequest
*/
func (a *PaymentService) PaymentCapturePost(ctx context.Context) PaymentPaymentCapturePostRequest {
	return PaymentPaymentCapturePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *PaymentService) PaymentCapturePostExecute(r PaymentPaymentCapturePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentCapturePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/capture"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.captureAmount == nil {
		return localVarReturnValue, nil, reportError("captureAmount is required and must be specified")
	}
	if r.currency == nil {
		return localVarReturnValue, nil, reportError("currency is required and must be specified")
	}
	if r.merchantPaymentCapturePostRequest == nil {
		return localVarReturnValue, nil, reportError("merchantPaymentCapturePostRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "captureAmount", r.captureAmount, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.merchantPaymentCapturePostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentDetailGetRequest struct {
	ctx context.Context
	ApiService *PaymentService
	paymentId *string
}

// PaymentId
func (r PaymentPaymentDetailGetRequest) PaymentId(paymentId string) PaymentPaymentDetailGetRequest {
	r.paymentId = &paymentId
	return r
}

func (r PaymentPaymentDetailGetRequest) Execute() (*MerchantPaymentDetailGet200Response, *http.Response, error) {
	return r.ApiService.PaymentDetailGetExecute(r)
}

/*
PaymentDetailGet Query Payment Detail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentDetailGetRequest
*/
func (a *PaymentService) PaymentDetailGet(ctx context.Context) PaymentPaymentDetailGetRequest {
	return PaymentPaymentDetailGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentDetailGet200Response
func (a *PaymentService) PaymentDetailGetExecute(r PaymentPaymentDetailGetRequest) (*MerchantPaymentDetailGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentDetailGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentDetailGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paymentId == nil {
		return localVarReturnValue, nil, reportError("paymentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "paymentId", r.paymentId, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentListGetRequest struct {
	ctx context.Context
	ApiService *PaymentService
	gatewayId *int64
	userId *int64
	email *string
	status *int32
	currency *string
	countryCode *string
	sortField *string
	sortType *string
	page *int32
	count *int32
}

// GatewayId
func (r PaymentPaymentListGetRequest) GatewayId(gatewayId int64) PaymentPaymentListGetRequest {
	r.gatewayId = &gatewayId
	return r
}

// UserId 
func (r PaymentPaymentListGetRequest) UserId(userId int64) PaymentPaymentListGetRequest {
	r.userId = &userId
	return r
}

// Email
func (r PaymentPaymentListGetRequest) Email(email string) PaymentPaymentListGetRequest {
	r.email = &email
	return r
}

// Status, 10-Created|20-Success|30-Failed|40-Cancelled
func (r PaymentPaymentListGetRequest) Status(status int32) PaymentPaymentListGetRequest {
	r.status = &status
	return r
}

// Currency
func (r PaymentPaymentListGetRequest) Currency(currency string) PaymentPaymentListGetRequest {
	r.currency = &currency
	return r
}

// CountryCode
func (r PaymentPaymentListGetRequest) CountryCode(countryCode string) PaymentPaymentListGetRequest {
	r.countryCode = &countryCode
	return r
}

// Sort Field，user_id|create_time|status
func (r PaymentPaymentListGetRequest) SortField(sortField string) PaymentPaymentListGetRequest {
	r.sortField = &sortField
	return r
}

// Sort Type，asc|desc
func (r PaymentPaymentListGetRequest) SortType(sortType string) PaymentPaymentListGetRequest {
	r.sortType = &sortType
	return r
}

// Page, Start With 0
func (r PaymentPaymentListGetRequest) Page(page int32) PaymentPaymentListGetRequest {
	r.page = &page
	return r
}

// Count Of Page
func (r PaymentPaymentListGetRequest) Count(count int32) PaymentPaymentListGetRequest {
	r.count = &count
	return r
}

func (r PaymentPaymentListGetRequest) Execute() (*MerchantPaymentListGet200Response, *http.Response, error) {
	return r.ApiService.PaymentListGetExecute(r)
}

/*
PaymentListGet Query Payment List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentListGetRequest
*/
func (a *PaymentService) PaymentListGet(ctx context.Context) PaymentPaymentListGetRequest {
	return PaymentPaymentListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentListGet200Response
func (a *PaymentService) PaymentListGetExecute(r PaymentPaymentListGetRequest) (*MerchantPaymentListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.gatewayId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayId", r.gatewayId, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.currency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	}
	if r.countryCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "countryCode", r.countryCode, "")
	}
	if r.sortField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortField", r.sortField, "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortType", r.sortType, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentNewPostRequest struct {
	ctx context.Context
	ApiService *PaymentService
	unibeeApiMerchantPaymentNewReq *UnibeeApiMerchantPaymentNewReq
}

func (r PaymentPaymentNewPostRequest) UnibeeApiMerchantPaymentNewReq(unibeeApiMerchantPaymentNewReq UnibeeApiMerchantPaymentNewReq) PaymentPaymentNewPostRequest {
	r.unibeeApiMerchantPaymentNewReq = &unibeeApiMerchantPaymentNewReq
	return r
}

func (r PaymentPaymentNewPostRequest) Execute() (*MerchantPaymentNewPost200Response, *http.Response, error) {
	return r.ApiService.PaymentNewPostExecute(r)
}

/*
PaymentNewPost New Payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentNewPostRequest
*/
func (a *PaymentService) PaymentNewPost(ctx context.Context) PaymentPaymentNewPostRequest {
	return PaymentPaymentNewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentNewPost200Response
func (a *PaymentService) PaymentNewPostExecute(r PaymentPaymentNewPostRequest) (*MerchantPaymentNewPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentNewPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantPaymentNewReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantPaymentNewReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantPaymentNewReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentRefundDetailGetRequest struct {
	ctx context.Context
	ApiService *PaymentService
	refundId *string
}

// RefundId
func (r PaymentPaymentRefundDetailGetRequest) RefundId(refundId string) PaymentPaymentRefundDetailGetRequest {
	r.refundId = &refundId
	return r
}

func (r PaymentPaymentRefundDetailGetRequest) Execute() (*MerchantPaymentRefundDetailGet200Response, *http.Response, error) {
	return r.ApiService.PaymentRefundDetailGetExecute(r)
}

/*
PaymentRefundDetailGet Query Payment Refund Detail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentRefundDetailGetRequest
*/
func (a *PaymentService) PaymentRefundDetailGet(ctx context.Context) PaymentPaymentRefundDetailGetRequest {
	return PaymentPaymentRefundDetailGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentRefundDetailGet200Response
func (a *PaymentService) PaymentRefundDetailGetExecute(r PaymentPaymentRefundDetailGetRequest) (*MerchantPaymentRefundDetailGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentRefundDetailGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentRefundDetailGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/refund/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.refundId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refundId", r.refundId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentRefundListGetRequest struct {
	ctx context.Context
	ApiService *PaymentService
	paymentId *string
	status *int32
	gatewayId *int64
	userId *int64
	email *string
	currency *string
}

// PaymentId
func (r PaymentPaymentRefundListGetRequest) PaymentId(paymentId string) PaymentPaymentRefundListGetRequest {
	r.paymentId = &paymentId
	return r
}

// Status,10-create|20-success|30-Failed|40-Reverse
func (r PaymentPaymentRefundListGetRequest) Status(status int32) PaymentPaymentRefundListGetRequest {
	r.status = &status
	return r
}

// GatewayId
func (r PaymentPaymentRefundListGetRequest) GatewayId(gatewayId int64) PaymentPaymentRefundListGetRequest {
	r.gatewayId = &gatewayId
	return r
}

// UserId
func (r PaymentPaymentRefundListGetRequest) UserId(userId int64) PaymentPaymentRefundListGetRequest {
	r.userId = &userId
	return r
}

// Email
func (r PaymentPaymentRefundListGetRequest) Email(email string) PaymentPaymentRefundListGetRequest {
	r.email = &email
	return r
}

// Currency
func (r PaymentPaymentRefundListGetRequest) Currency(currency string) PaymentPaymentRefundListGetRequest {
	r.currency = &currency
	return r
}

func (r PaymentPaymentRefundListGetRequest) Execute() (*MerchantPaymentRefundListGet200Response, *http.Response, error) {
	return r.ApiService.PaymentRefundListGetExecute(r)
}

/*
PaymentRefundListGet Query Payment Refund List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentRefundListGetRequest
*/
func (a *PaymentService) PaymentRefundListGet(ctx context.Context) PaymentPaymentRefundListGetRequest {
	return PaymentPaymentRefundListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentRefundListGet200Response
func (a *PaymentService) PaymentRefundListGetExecute(r PaymentPaymentRefundListGetRequest) (*MerchantPaymentRefundListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentRefundListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentRefundListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/refund/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paymentId == nil {
		return localVarReturnValue, nil, reportError("paymentId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "paymentId", r.paymentId, "")
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.gatewayId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayId", r.gatewayId, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "")
	}
	if r.currency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentRefundNewPostRequest struct {
	ctx context.Context
	ApiService *PaymentService
	unibeeApiMerchantPaymentNewPaymentRefundReq *UnibeeApiMerchantPaymentNewPaymentRefundReq
}

func (r PaymentPaymentRefundNewPostRequest) UnibeeApiMerchantPaymentNewPaymentRefundReq(unibeeApiMerchantPaymentNewPaymentRefundReq UnibeeApiMerchantPaymentNewPaymentRefundReq) PaymentPaymentRefundNewPostRequest {
	r.unibeeApiMerchantPaymentNewPaymentRefundReq = &unibeeApiMerchantPaymentNewPaymentRefundReq
	return r
}

func (r PaymentPaymentRefundNewPostRequest) Execute() (*MerchantPaymentRefundNewPost200Response, *http.Response, error) {
	return r.ApiService.PaymentRefundNewPostExecute(r)
}

/*
PaymentRefundNewPost New Payment Refund

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentRefundNewPostRequest
*/
func (a *PaymentService) PaymentRefundNewPost(ctx context.Context) PaymentPaymentRefundNewPostRequest {
	return PaymentPaymentRefundNewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentRefundNewPost200Response
func (a *PaymentService) PaymentRefundNewPostExecute(r PaymentPaymentRefundNewPostRequest) (*MerchantPaymentRefundNewPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentRefundNewPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentRefundNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/refund/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantPaymentNewPaymentRefundReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantPaymentNewPaymentRefundReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantPaymentNewPaymentRefundReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentPaymentTimelineListGetRequest struct {
	ctx context.Context
	ApiService *PaymentService
	userId *int64
	sortField *string
	sortType *string
	page *int32
	count *int32
}

// Filter UserId, Default All
func (r PaymentPaymentTimelineListGetRequest) UserId(userId int64) PaymentPaymentTimelineListGetRequest {
	r.userId = &userId
	return r
}

// Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify
func (r PaymentPaymentTimelineListGetRequest) SortField(sortField string) PaymentPaymentTimelineListGetRequest {
	r.sortField = &sortField
	return r
}

// Sort Type，asc|desc，Default desc
func (r PaymentPaymentTimelineListGetRequest) SortType(sortType string) PaymentPaymentTimelineListGetRequest {
	r.sortType = &sortType
	return r
}

// Page,Start 0
func (r PaymentPaymentTimelineListGetRequest) Page(page int32) PaymentPaymentTimelineListGetRequest {
	r.page = &page
	return r
}

// Count Of Page
func (r PaymentPaymentTimelineListGetRequest) Count(count int32) PaymentPaymentTimelineListGetRequest {
	r.count = &count
	return r
}

func (r PaymentPaymentTimelineListGetRequest) Execute() (*MerchantPaymentTimelineListGet200Response, *http.Response, error) {
	return r.ApiService.PaymentTimelineListGetExecute(r)
}

/*
PaymentTimelineListGet Payment TimeLine List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentPaymentTimelineListGetRequest
*/
func (a *PaymentService) PaymentTimelineListGet(ctx context.Context) PaymentPaymentTimelineListGetRequest {
	return PaymentPaymentTimelineListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantPaymentTimelineListGet200Response
func (a *PaymentService) PaymentTimelineListGetExecute(r PaymentPaymentTimelineListGetRequest) (*MerchantPaymentTimelineListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantPaymentTimelineListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentService.PaymentTimelineListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/payment/timeline/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.sortField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortField", r.sortField, "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortType", r.sortType, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
