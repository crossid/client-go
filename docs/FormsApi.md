# \FormsApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitLoginFlow**](FormsApi.md#SubmitLoginFlow) | **Post** /auth/flows/login/{id}/.submit | Submit a login flow.
[**SubmitRegisterFlow**](FormsApi.md#SubmitRegisterFlow) | **Post** /auth/flows/register/{id}/.submit | Submit a registeration flow.



## SubmitLoginFlow

> CrossFlow SubmitLoginFlow(ctx, id).CrossFlow(crossFlow).Execute()

Submit a login flow.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the login flow to submit.
    crossFlow := *openapiclient.NewCrossFlow("DisplayName_example", "Id_example", time.Now(), "Status_example", "StepId_example", *openapiclient.NewCrossFlowSteps(), map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // CrossFlow |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.SubmitLoginFlow(context.Background(), id).CrossFlow(crossFlow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.SubmitLoginFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitLoginFlow`: CrossFlow
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.SubmitLoginFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the login flow to submit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitLoginFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crossFlow** | [**CrossFlow**](CrossFlow.md) |  | 

### Return type

[**CrossFlow**](CrossFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRegisterFlow

> CrossFlow SubmitRegisterFlow(ctx, id).CrossFlow(crossFlow).Execute()

Submit a registeration flow.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the registeration flow to submit.
    crossFlow := *openapiclient.NewCrossFlow("DisplayName_example", "Id_example", time.Now(), "Status_example", "StepId_example", *openapiclient.NewCrossFlowSteps(), map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // CrossFlow |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FormsApi.SubmitRegisterFlow(context.Background(), id).CrossFlow(crossFlow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormsApi.SubmitRegisterFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRegisterFlow`: CrossFlow
    fmt.Fprintf(os.Stdout, "Response from `FormsApi.SubmitRegisterFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the registeration flow to submit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRegisterFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crossFlow** | [**CrossFlow**](CrossFlow.md) |  | 

### Return type

[**CrossFlow**](CrossFlow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

