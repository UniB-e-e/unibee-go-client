# \Subscription

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionAddNewTrialStartPost**](Subscription.md#SubscriptionAddNewTrialStartPost) | **Post** /merchant/subscription/add_new_trial_start | Merchant Edit Subscription-add appendTrialEndHour For Free
[**SubscriptionCancelAtPeriodEndPost**](Subscription.md#SubscriptionCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_at_period_end | Merchant Edit Subscription-Set Cancel Ad Period End
[**SubscriptionCancelLastCancelAtPeriodEndPost**](Subscription.md#SubscriptionCancelLastCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_last_cancel_at_period_end | Merchant Edit Subscription-Cancel Last CancelAtPeriod
[**SubscriptionCancelPost**](Subscription.md#SubscriptionCancelPost) | **Post** /merchant/subscription/cancel | Merchant Cancel Subscription Immediately (Will Not Generate Proration Invoice)
[**SubscriptionChangeGatewayPost**](Subscription.md#SubscriptionChangeGatewayPost) | **Post** /merchant/subscription/change_gateway | Change Subscription Gateway
[**SubscriptionConfigGet**](Subscription.md#SubscriptionConfigGet) | **Get** /merchant/subscription/config | Get Merchant Subscription Config
[**SubscriptionConfigUpdateGet**](Subscription.md#SubscriptionConfigUpdateGet) | **Get** /merchant/subscription/config/update | Update Merchant Subscription Config
[**SubscriptionCreatePreviewPost**](Subscription.md#SubscriptionCreatePreviewPost) | **Post** /merchant/subscription/create_preview | Create Subscription Preview
[**SubscriptionCreateSubmitPost**](Subscription.md#SubscriptionCreateSubmitPost) | **Post** /merchant/subscription/create_submit | Create Subscription
[**SubscriptionDetailGet**](Subscription.md#SubscriptionDetailGet) | **Get** /merchant/subscription/detail | Subscription Detail
[**SubscriptionDetailPost**](Subscription.md#SubscriptionDetailPost) | **Post** /merchant/subscription/detail | Subscription Detail
[**SubscriptionListGet**](Subscription.md#SubscriptionListGet) | **Get** /merchant/subscription/list | Subscription List
[**SubscriptionListPost**](Subscription.md#SubscriptionListPost) | **Post** /merchant/subscription/list | Subscription List
[**SubscriptionRenewPost**](Subscription.md#SubscriptionRenewPost) | **Post** /merchant/subscription/renew | Renew Subscription, will create new subscription based on one provided 
[**SubscriptionResumePost**](Subscription.md#SubscriptionResumePost) | **Post** /merchant/subscription/resume | Merchant Edit Subscription-Resume
[**SubscriptionSuspendPost**](Subscription.md#SubscriptionSuspendPost) | **Post** /merchant/subscription/suspend | Merchant Edit Subscription-Stop
[**SubscriptionUpdatePreviewPost**](Subscription.md#SubscriptionUpdatePreviewPost) | **Post** /merchant/subscription/update_preview | Merchant Update Subscription Preview
[**SubscriptionUpdateSubmitPost**](Subscription.md#SubscriptionUpdateSubmitPost) | **Post** /merchant/subscription/update_submit | Merchant Update Subscription Submit
[**SubscriptionUserSubscriptionDetailGet**](Subscription.md#SubscriptionUserSubscriptionDetailGet) | **Get** /merchant/subscription/user_subscription_detail | Subscription Detail
[**SubscriptionUserSubscriptionDetailPost**](Subscription.md#SubscriptionUserSubscriptionDetailPost) | **Post** /merchant/subscription/user_subscription_detail | Subscription Detail



## SubscriptionAddNewTrialStartPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionAddNewTrialStartPost(ctx).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()

Merchant Edit Subscription-add appendTrialEndHour For Free

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionAddNewTrialStartReq := *openapiclient.NewUnibeeApiMerchantSubscriptionAddNewTrialStartReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionAddNewTrialStartReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionAddNewTrialStartPost(context.Background()).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionAddNewTrialStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAddNewTrialStartPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionAddNewTrialStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAddNewTrialStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionAddNewTrialStartReq** | [**UnibeeApiMerchantSubscriptionAddNewTrialStartReq**](UnibeeApiMerchantSubscriptionAddNewTrialStartReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelAtPeriodEndPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()

Merchant Edit Subscription-Set Cancel Ad Period End

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionCancelAtPeriodEndReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCancelAtPeriodEndPost(context.Background()).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCancelAtPeriodEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancelAtPeriodEndPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCancelAtPeriodEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelAtPeriodEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelAtPeriodEndReq** | [**UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq**](UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelLastCancelAtPeriodEndPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionCancelLastCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()

Merchant Edit Subscription-Cancel Last CancelAtPeriod

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCancelLastCancelAtPeriodEndPost(context.Background()).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCancelLastCancelAtPeriodEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancelLastCancelAtPeriodEndPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCancelLastCancelAtPeriodEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelLastCancelAtPeriodEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq** | [**UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq**](UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancelPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionCancelPost(ctx).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()

Merchant Cancel Subscription Immediately (Will Not Generate Proration Invoice)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionCancelReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCancelPost(context.Background()).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancelPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelReq** | [**UnibeeApiMerchantSubscriptionCancelReq**](UnibeeApiMerchantSubscriptionCancelReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionChangeGatewayPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionChangeGatewayPost(ctx).UnibeeApiMerchantSubscriptionChangeGatewayReq(unibeeApiMerchantSubscriptionChangeGatewayReq).Execute()

Change Subscription Gateway

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionChangeGatewayReq := *openapiclient.NewUnibeeApiMerchantSubscriptionChangeGatewayReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionChangeGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionChangeGatewayPost(context.Background()).UnibeeApiMerchantSubscriptionChangeGatewayReq(unibeeApiMerchantSubscriptionChangeGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionChangeGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionChangeGatewayPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionChangeGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionChangeGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionChangeGatewayReq** | [**UnibeeApiMerchantSubscriptionChangeGatewayReq**](UnibeeApiMerchantSubscriptionChangeGatewayReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionConfigGet

> MerchantSubscriptionConfigGet200Response SubscriptionConfigGet(ctx).Execute()

Get Merchant Subscription Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionConfigGet`: MerchantSubscriptionConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionConfigGetRequest struct via the builder pattern


### Return type

[**MerchantSubscriptionConfigGet200Response**](MerchantSubscriptionConfigGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionConfigUpdateGet

> MerchantSubscriptionConfigGet200Response SubscriptionConfigUpdateGet(ctx).DowngradeEffectImmediately(downgradeEffectImmediately).UpgradeProration(upgradeProration).IncompleteExpireTime(incompleteExpireTime).InvoiceEmail(invoiceEmail).Execute()

Update Merchant Subscription Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	downgradeEffectImmediately := true // bool | DowngradeEffectImmediately, whether subscription downgrade should effect immediately or at period end, default at period end (optional)
	upgradeProration := true // bool | UpgradeProration, whether subscription update generation proration invoice or not, default yes (optional)
	incompleteExpireTime := int32(56) // int32 | IncompleteExpireTime, em.. default 1day for plan of month type (optional)
	invoiceEmail := true // bool | InvoiceEmail, whether to send invoice email to user, default yes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionConfigUpdateGet(context.Background()).DowngradeEffectImmediately(downgradeEffectImmediately).UpgradeProration(upgradeProration).IncompleteExpireTime(incompleteExpireTime).InvoiceEmail(invoiceEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionConfigUpdateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionConfigUpdateGet`: MerchantSubscriptionConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionConfigUpdateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionConfigUpdateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downgradeEffectImmediately** | **bool** | DowngradeEffectImmediately, whether subscription downgrade should effect immediately or at period end, default at period end | 
 **upgradeProration** | **bool** | UpgradeProration, whether subscription update generation proration invoice or not, default yes | 
 **incompleteExpireTime** | **int32** | IncompleteExpireTime, em.. default 1day for plan of month type | 
 **invoiceEmail** | **bool** | InvoiceEmail, whether to send invoice email to user, default yes | 

### Return type

[**MerchantSubscriptionConfigGet200Response**](MerchantSubscriptionConfigGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCreatePreviewPost

> MerchantSubscriptionCreatePreviewPost200Response SubscriptionCreatePreviewPost(ctx).UnibeeApiMerchantSubscriptionCreatePreviewReq(unibeeApiMerchantSubscriptionCreatePreviewReq).Execute()

Create Subscription Preview

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionCreatePreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCreatePreviewReq(int64(123), int64(123), int64(123)) // UnibeeApiMerchantSubscriptionCreatePreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCreatePreviewPost(context.Background()).UnibeeApiMerchantSubscriptionCreatePreviewReq(unibeeApiMerchantSubscriptionCreatePreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCreatePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCreatePreviewPost`: MerchantSubscriptionCreatePreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCreatePreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCreatePreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCreatePreviewReq** | [**UnibeeApiMerchantSubscriptionCreatePreviewReq**](UnibeeApiMerchantSubscriptionCreatePreviewReq.md) |  | 

### Return type

[**MerchantSubscriptionCreatePreviewPost200Response**](MerchantSubscriptionCreatePreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCreateSubmitPost

> MerchantSubscriptionCreateSubmitPost200Response SubscriptionCreateSubmitPost(ctx).UnibeeApiMerchantSubscriptionCreateReq(unibeeApiMerchantSubscriptionCreateReq).Execute()

Create Subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionCreateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCreateReq(int64(123), int64(123), int64(123)) // UnibeeApiMerchantSubscriptionCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionCreateSubmitPost(context.Background()).UnibeeApiMerchantSubscriptionCreateReq(unibeeApiMerchantSubscriptionCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionCreateSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCreateSubmitPost`: MerchantSubscriptionCreateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionCreateSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCreateSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCreateReq** | [**UnibeeApiMerchantSubscriptionCreateReq**](UnibeeApiMerchantSubscriptionCreateReq.md) |  | 

### Return type

[**MerchantSubscriptionCreateSubmitPost200Response**](MerchantSubscriptionCreateSubmitPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDetailGet

> MerchantSubscriptionDetailGet200Response SubscriptionDetailGet(ctx).SubscriptionId(subscriptionId).Execute()

Subscription Detail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | SubscriptionId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionDetailGet(context.Background()).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionDetailGet`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDetailPost

> MerchantSubscriptionDetailGet200Response SubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionDetailReq(unibeeApiMerchantSubscriptionDetailReq).Execute()

Subscription Detail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionDetailReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionDetailReq(unibeeApiMerchantSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionDetailPost`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionDetailReq.md) |  | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionListGet

> MerchantSubscriptionListGet200Response SubscriptionListGet(ctx).UserId(userId).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Subscription List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	userId := int64(789) // int64 | UserId (optional)
	status := []int32{int32(123)} // []int32 | Filter, Default All，Status，0-Init | 1-Create｜2-Active｜3-Suspend | 4-Cancel | 5-Expire (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionListGet(context.Background()).UserId(userId).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionListGet`: MerchantSubscriptionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **status** | **[]int32** | Filter, Default All，Status，0-Init | 1-Create｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start WIth 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionListGet200Response**](MerchantSubscriptionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionListPost

> MerchantSubscriptionListGet200Response SubscriptionListPost(ctx).UnibeeApiMerchantSubscriptionListReq(unibeeApiMerchantSubscriptionListReq).Execute()

Subscription List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionListReq() // UnibeeApiMerchantSubscriptionListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionListPost(context.Background()).UnibeeApiMerchantSubscriptionListReq(unibeeApiMerchantSubscriptionListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionListPost`: MerchantSubscriptionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionListReq** | [**UnibeeApiMerchantSubscriptionListReq**](UnibeeApiMerchantSubscriptionListReq.md) |  | 

### Return type

[**MerchantSubscriptionListGet200Response**](MerchantSubscriptionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionRenewPost

> MerchantSubscriptionCreateSubmitPost200Response SubscriptionRenewPost(ctx).UnibeeApiMerchantSubscriptionRenewReq(unibeeApiMerchantSubscriptionRenewReq).Execute()

Renew Subscription, will create new subscription based on one provided 

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionRenewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionRenewReq("SubscriptionId_example", int64(123)) // UnibeeApiMerchantSubscriptionRenewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionRenewPost(context.Background()).UnibeeApiMerchantSubscriptionRenewReq(unibeeApiMerchantSubscriptionRenewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionRenewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRenewPost`: MerchantSubscriptionCreateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionRenewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRenewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionRenewReq** | [**UnibeeApiMerchantSubscriptionRenewReq**](UnibeeApiMerchantSubscriptionRenewReq.md) |  | 

### Return type

[**MerchantSubscriptionCreateSubmitPost200Response**](MerchantSubscriptionCreateSubmitPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionResumePost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionResumePost(ctx).UnibeeApiMerchantSubscriptionResumeReq(unibeeApiMerchantSubscriptionResumeReq).Execute()

Merchant Edit Subscription-Resume

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionResumeReq := *openapiclient.NewUnibeeApiMerchantSubscriptionResumeReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionResumeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionResumePost(context.Background()).UnibeeApiMerchantSubscriptionResumeReq(unibeeApiMerchantSubscriptionResumeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionResumePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionResumePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionResumeReq** | [**UnibeeApiMerchantSubscriptionResumeReq**](UnibeeApiMerchantSubscriptionResumeReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionSuspendPost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionSuspendPost(ctx).UnibeeApiMerchantSubscriptionSuspendReq(unibeeApiMerchantSubscriptionSuspendReq).Execute()

Merchant Edit Subscription-Stop

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionSuspendReq := *openapiclient.NewUnibeeApiMerchantSubscriptionSuspendReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionSuspendReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionSuspendPost(context.Background()).UnibeeApiMerchantSubscriptionSuspendReq(unibeeApiMerchantSubscriptionSuspendReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionSuspendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionSuspendPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionSuspendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionSuspendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionSuspendReq** | [**UnibeeApiMerchantSubscriptionSuspendReq**](UnibeeApiMerchantSubscriptionSuspendReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdatePreviewPost

> MerchantSubscriptionUpdatePreviewPost200Response SubscriptionUpdatePreviewPost(ctx).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()

Merchant Update Subscription Preview

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionUpdatePreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionUpdatePreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUpdatePreviewPost(context.Background()).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUpdatePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdatePreviewPost`: MerchantSubscriptionUpdatePreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUpdatePreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdatePreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUpdatePreviewReq** | [**UnibeeApiMerchantSubscriptionUpdatePreviewReq**](UnibeeApiMerchantSubscriptionUpdatePreviewReq.md) |  | 

### Return type

[**MerchantSubscriptionUpdatePreviewPost200Response**](MerchantSubscriptionUpdatePreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdateSubmitPost

> MerchantSubscriptionUpdateSubmitPost200Response SubscriptionUpdateSubmitPost(ctx).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()

Merchant Update Subscription Submit

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionUpdateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdateReq(int64(123), int64(123), int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUpdateSubmitPost(context.Background()).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUpdateSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdateSubmitPost`: MerchantSubscriptionUpdateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUpdateSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUpdateReq** | [**UnibeeApiMerchantSubscriptionUpdateReq**](UnibeeApiMerchantSubscriptionUpdateReq.md) |  | 

### Return type

[**MerchantSubscriptionUpdateSubmitPost200Response**](MerchantSubscriptionUpdateSubmitPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserSubscriptionDetailGet

> MerchantSubscriptionUserSubscriptionDetailGet200Response SubscriptionUserSubscriptionDetailGet(ctx).UserId(userId).Execute()

Subscription Detail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	userId := int64(789) // int64 | UserId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserSubscriptionDetailGet(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserSubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserSubscriptionDetailGet`: MerchantSubscriptionUserSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUserSubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUserSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 

### Return type

[**MerchantSubscriptionUserSubscriptionDetailGet200Response**](MerchantSubscriptionUserSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUserSubscriptionDetailPost

> MerchantSubscriptionUserSubscriptionDetailGet200Response SubscriptionUserSubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserSubscriptionDetailReq).Execute()

Subscription Detail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionUserSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(int64(123)) // UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Subscription.SubscriptionUserSubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Subscription.SubscriptionUserSubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUserSubscriptionDetailPost`: MerchantSubscriptionUserSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Subscription.SubscriptionUserSubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUserSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUserSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq.md) |  | 

### Return type

[**MerchantSubscriptionUserSubscriptionDetailGet200Response**](MerchantSubscriptionUserSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

