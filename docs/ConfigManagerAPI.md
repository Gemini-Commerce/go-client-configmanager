# \ConfigManagerAPI

All URIs are relative to *https://config-manager.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigManagerBulkSetConfigs**](ConfigManagerAPI.md#ConfigManagerBulkSetConfigs) | **Post** /configmanager.ConfigManager/BulkSetConfigs | BulkSetConfigs
[**ConfigManagerGetConfig**](ConfigManagerAPI.md#ConfigManagerGetConfig) | **Post** /configmanager.ConfigManager/GetConfig | GetConfig
[**ConfigManagerGetTenantIdByCode**](ConfigManagerAPI.md#ConfigManagerGetTenantIdByCode) | **Post** /configmanager.ConfigManager/GetTenantIdByCode | 



## ConfigManagerBulkSetConfigs

> map[string]interface{} ConfigManagerBulkSetConfigs(ctx).Body(body).Execute()

BulkSetConfigs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-configmanager"
)

func main() {
	body := *openapiclient.NewConfigmanagerBulkSetConfigsRequest("TenantId_example", []openapiclient.BulkSetConfigsRequestConfig{*openapiclient.NewBulkSetConfigsRequestConfig("Key_example", "Value_example")}) // ConfigmanagerBulkSetConfigsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigManagerAPI.ConfigManagerBulkSetConfigs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigManagerAPI.ConfigManagerBulkSetConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigManagerBulkSetConfigs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigManagerAPI.ConfigManagerBulkSetConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigManagerBulkSetConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigmanagerBulkSetConfigsRequest**](ConfigmanagerBulkSetConfigsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigManagerGetConfig

> ConfigmanagerConfigResponse ConfigManagerGetConfig(ctx).Body(body).Execute()

GetConfig



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-configmanager"
)

func main() {
	body := *openapiclient.NewConfigmanagerGetConfigRequest("TenantId_example", "Key_example") // ConfigmanagerGetConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigManagerAPI.ConfigManagerGetConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigManagerAPI.ConfigManagerGetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigManagerGetConfig`: ConfigmanagerConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigManagerAPI.ConfigManagerGetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigManagerGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigmanagerGetConfigRequest**](ConfigmanagerGetConfigRequest.md) |  | 

### Return type

[**ConfigmanagerConfigResponse**](ConfigmanagerConfigResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigManagerGetTenantIdByCode

> ConfigmanagerGetTenantIdByCodeResponse ConfigManagerGetTenantIdByCode(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-configmanager"
)

func main() {
	body := *openapiclient.NewConfigmanagerGetTenantIdByCodeRequest("Code_example") // ConfigmanagerGetTenantIdByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigManagerAPI.ConfigManagerGetTenantIdByCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigManagerAPI.ConfigManagerGetTenantIdByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigManagerGetTenantIdByCode`: ConfigmanagerGetTenantIdByCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigManagerAPI.ConfigManagerGetTenantIdByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigManagerGetTenantIdByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigmanagerGetTenantIdByCodeRequest**](ConfigmanagerGetTenantIdByCodeRequest.md) |  | 

### Return type

[**ConfigmanagerGetTenantIdByCodeResponse**](ConfigmanagerGetTenantIdByCodeResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

