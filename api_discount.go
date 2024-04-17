/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202404171839 
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


// DiscountService Discount service
type DiscountService service

type DiscountDiscountActivatePostRequest struct {
	ctx context.Context
	ApiService *DiscountService
	unibeeApiMerchantDiscountActivateReq *UnibeeApiMerchantDiscountActivateReq
}

func (r DiscountDiscountActivatePostRequest) UnibeeApiMerchantDiscountActivateReq(unibeeApiMerchantDiscountActivateReq UnibeeApiMerchantDiscountActivateReq) DiscountDiscountActivatePostRequest {
	r.unibeeApiMerchantDiscountActivateReq = &unibeeApiMerchantDiscountActivateReq
	return r
}

func (r DiscountDiscountActivatePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.DiscountActivatePostExecute(r)
}

/*
DiscountActivatePost ActivateDiscountCode

Activate discount code, the discount code can only effect to payment or subscription after activated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DiscountDiscountActivatePostRequest
*/
func (a *DiscountService) DiscountActivatePost(ctx context.Context) DiscountDiscountActivatePostRequest {
	return DiscountDiscountActivatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *DiscountService) DiscountActivatePostExecute(r DiscountDiscountActivatePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountService.DiscountActivatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/discount/activate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantDiscountActivateReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantDiscountActivateReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantDiscountActivateReq
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

type DiscountDiscountDeactivatePostRequest struct {
	ctx context.Context
	ApiService *DiscountService
	unibeeApiMerchantDiscountDeactivateReq *UnibeeApiMerchantDiscountDeactivateReq
}

func (r DiscountDiscountDeactivatePostRequest) UnibeeApiMerchantDiscountDeactivateReq(unibeeApiMerchantDiscountDeactivateReq UnibeeApiMerchantDiscountDeactivateReq) DiscountDiscountDeactivatePostRequest {
	r.unibeeApiMerchantDiscountDeactivateReq = &unibeeApiMerchantDiscountDeactivateReq
	return r
}

func (r DiscountDiscountDeactivatePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.DiscountDeactivatePostExecute(r)
}

/*
DiscountDeactivatePost DeactivateDiscountCode

Deactivate discount code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DiscountDiscountDeactivatePostRequest
*/
func (a *DiscountService) DiscountDeactivatePost(ctx context.Context) DiscountDiscountDeactivatePostRequest {
	return DiscountDiscountDeactivatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *DiscountService) DiscountDeactivatePostExecute(r DiscountDiscountDeactivatePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountService.DiscountDeactivatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/discount/deactivate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantDiscountDeactivateReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantDiscountDeactivateReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantDiscountDeactivateReq
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

type DiscountDiscountDeletePostRequest struct {
	ctx context.Context
	ApiService *DiscountService
	unibeeApiMerchantDiscountDeleteReq *UnibeeApiMerchantDiscountDeleteReq
}

func (r DiscountDiscountDeletePostRequest) UnibeeApiMerchantDiscountDeleteReq(unibeeApiMerchantDiscountDeleteReq UnibeeApiMerchantDiscountDeleteReq) DiscountDiscountDeletePostRequest {
	r.unibeeApiMerchantDiscountDeleteReq = &unibeeApiMerchantDiscountDeleteReq
	return r
}

func (r DiscountDiscountDeletePostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.DiscountDeletePostExecute(r)
}

/*
DiscountDeletePost DeleteDiscountCode

Delete discount code before activate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DiscountDiscountDeletePostRequest
*/
func (a *DiscountService) DiscountDeletePost(ctx context.Context) DiscountDiscountDeletePostRequest {
	return DiscountDiscountDeletePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *DiscountService) DiscountDeletePostExecute(r DiscountDiscountDeletePostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountService.DiscountDeletePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/discount/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantDiscountDeleteReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantDiscountDeleteReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantDiscountDeleteReq
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

type DiscountDiscountEditPostRequest struct {
	ctx context.Context
	ApiService *DiscountService
	unibeeApiMerchantDiscountEditReq *UnibeeApiMerchantDiscountEditReq
}

func (r DiscountDiscountEditPostRequest) UnibeeApiMerchantDiscountEditReq(unibeeApiMerchantDiscountEditReq UnibeeApiMerchantDiscountEditReq) DiscountDiscountEditPostRequest {
	r.unibeeApiMerchantDiscountEditReq = &unibeeApiMerchantDiscountEditReq
	return r
}

func (r DiscountDiscountEditPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.DiscountEditPostExecute(r)
}

/*
DiscountEditPost EditDiscountCode

Edit the discount code before activate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DiscountDiscountEditPostRequest
*/
func (a *DiscountService) DiscountEditPost(ctx context.Context) DiscountDiscountEditPostRequest {
	return DiscountDiscountEditPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *DiscountService) DiscountEditPostExecute(r DiscountDiscountEditPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountService.DiscountEditPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/discount/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantDiscountEditReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantDiscountEditReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantDiscountEditReq
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

type DiscountDiscountListGetRequest struct {
	ctx context.Context
	ApiService *DiscountService
}

func (r DiscountDiscountListGetRequest) Execute() (*MerchantDiscountListGet200Response, *http.Response, error) {
	return r.ApiService.DiscountListGetExecute(r)
}

/*
DiscountListGet DiscountCodeList

Get discountCode list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DiscountDiscountListGetRequest
*/
func (a *DiscountService) DiscountListGet(ctx context.Context) DiscountDiscountListGetRequest {
	return DiscountDiscountListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantDiscountListGet200Response
func (a *DiscountService) DiscountListGetExecute(r DiscountDiscountListGetRequest) (*MerchantDiscountListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantDiscountListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountService.DiscountListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/discount/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DiscountDiscountNewPostRequest struct {
	ctx context.Context
	ApiService *DiscountService
	unibeeApiMerchantDiscountNewReq *UnibeeApiMerchantDiscountNewReq
}

func (r DiscountDiscountNewPostRequest) UnibeeApiMerchantDiscountNewReq(unibeeApiMerchantDiscountNewReq UnibeeApiMerchantDiscountNewReq) DiscountDiscountNewPostRequest {
	r.unibeeApiMerchantDiscountNewReq = &unibeeApiMerchantDiscountNewReq
	return r
}

func (r DiscountDiscountNewPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.DiscountNewPostExecute(r)
}

/*
DiscountNewPost NewDiscountCode

Create a new discount code, code can used in onetime or subscription purchase to make discount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DiscountDiscountNewPostRequest
*/
func (a *DiscountService) DiscountNewPost(ctx context.Context) DiscountDiscountNewPostRequest {
	return DiscountDiscountNewPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *DiscountService) DiscountNewPostExecute(r DiscountDiscountNewPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscountService.DiscountNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/discount/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantDiscountNewReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantDiscountNewReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantDiscountNewReq
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
