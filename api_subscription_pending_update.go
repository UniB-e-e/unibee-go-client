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


// SubscriptionPendingUpdateService SubscriptionPendingUpdate service
type SubscriptionPendingUpdateService service

type SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest struct {
	ctx context.Context
	ApiService *SubscriptionPendingUpdateService
	subscriptionId *string
	sortField *string
	sortType *string
	page *int32
	count *int32
}

// SubscriptionId
func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) SubscriptionId(subscriptionId string) SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest {
	r.subscriptionId = &subscriptionId
	return r
}

// Sort Field，gmt_create|gmt_modify，Default gmt_modify
func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) SortField(sortField string) SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest {
	r.sortField = &sortField
	return r
}

// Sort Type，asc|desc，Default desc
func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) SortType(sortType string) SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest {
	r.sortType = &sortType
	return r
}

// Page, Start WIth 0
func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) Page(page int32) SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest {
	r.page = &page
	return r
}

// Count Of Page
func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) Count(count int32) SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest {
	r.count = &count
	return r
}

func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) Execute() (*MerchantSubscriptionPendingUpdateListGet200Response, *http.Response, error) {
	return r.ApiService.SubscriptionPendingUpdateListGetExecute(r)
}

/*
SubscriptionPendingUpdateListGet SubscriptionPendingUpdateList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest
*/
func (a *SubscriptionPendingUpdateService) SubscriptionPendingUpdateListGet(ctx context.Context) SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest {
	return SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSubscriptionPendingUpdateListGet200Response
func (a *SubscriptionPendingUpdateService) SubscriptionPendingUpdateListGetExecute(r SubscriptionPendingUpdateSubscriptionPendingUpdateListGetRequest) (*MerchantSubscriptionPendingUpdateListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSubscriptionPendingUpdateListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPendingUpdateService.SubscriptionPendingUpdateListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/pending_update_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionId == nil {
		return localVarReturnValue, nil, reportError("subscriptionId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionId", r.subscriptionId, "")
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

type SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest struct {
	ctx context.Context
	ApiService *SubscriptionPendingUpdateService
	unibeeApiMerchantSubscriptionPendingUpdateListReq *UnibeeApiMerchantSubscriptionPendingUpdateListReq
}

func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest) UnibeeApiMerchantSubscriptionPendingUpdateListReq(unibeeApiMerchantSubscriptionPendingUpdateListReq UnibeeApiMerchantSubscriptionPendingUpdateListReq) SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest {
	r.unibeeApiMerchantSubscriptionPendingUpdateListReq = &unibeeApiMerchantSubscriptionPendingUpdateListReq
	return r
}

func (r SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest) Execute() (*MerchantSubscriptionPendingUpdateListGet200Response, *http.Response, error) {
	return r.ApiService.SubscriptionPendingUpdateListPostExecute(r)
}

/*
SubscriptionPendingUpdateListPost SubscriptionPendingUpdateList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest
*/
func (a *SubscriptionPendingUpdateService) SubscriptionPendingUpdateListPost(ctx context.Context) SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest {
	return SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSubscriptionPendingUpdateListGet200Response
func (a *SubscriptionPendingUpdateService) SubscriptionPendingUpdateListPostExecute(r SubscriptionPendingUpdateSubscriptionPendingUpdateListPostRequest) (*MerchantSubscriptionPendingUpdateListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSubscriptionPendingUpdateListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPendingUpdateService.SubscriptionPendingUpdateListPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/subscription/pending_update_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantSubscriptionPendingUpdateListReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantSubscriptionPendingUpdateListReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantSubscriptionPendingUpdateListReq
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
