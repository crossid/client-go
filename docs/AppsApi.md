# \AppsApi

All URIs are relative to *https://acme.us.crossid.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](AppsApi.md#CreateApp) | **Post** /api/v1/apps | Create a an Application
[**Deleteapp**](AppsApi.md#Deleteapp) | **Delete** /api/v1/apps/{id} | Delete an Application



## CreateApp

> App CreateApp(ctx).Reason(reason).App(app).Execute()

Create a an Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reason := "approved by ticket 4423." // string | A descriptive reason of the change
    app := *openapiclient.NewApp() // App | An App in JSON format. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.CreateApp(context.Background()).Reason(reason).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApp`: App
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | A descriptive reason of the change | 
 **app** | [**App**](App.md) | An App in JSON format. | 

### Return type

[**App**](App.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteapp

> Deleteapp(ctx, id).Reason(reason).Execute()

Delete an Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "5eqtJP6z5bSjStMwhegqwH" // string | The App ID
    reason := "approved by ticket 4423." // string | A descriptive reason of the change

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppsApi.Deleteapp(context.Background(), id).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.Deleteapp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteappRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | A descriptive reason of the change | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

