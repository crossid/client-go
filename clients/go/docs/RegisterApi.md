# \RegisterApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegistrationFlow**](RegisterApi.md#GetRegistrationFlow) | **Get** /auth/flows/register/{id} | Info for a specific register flow.



## GetRegistrationFlow

> CrossFlow GetRegistrationFlow(ctx, id).Execute()

Info for a specific register flow.

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
    id := "id_example" // string | The id of the flow to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegisterApi.GetRegistrationFlow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegisterApi.GetRegistrationFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistrationFlow`: CrossFlow
    fmt.Fprintf(os.Stdout, "Response from `RegisterApi.GetRegistrationFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the flow to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistrationFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrossFlow**](CrossFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

