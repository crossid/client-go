# \SchemasApi

All URIs are relative to *https://acme.us.crossid.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSCIMSchema**](SchemasApi.md#CreateSCIMSchema) | **Post** /api/v1/scim-schemas | Create a new SCIM Schema
[**GetSCIMSchema**](SchemasApi.md#GetSCIMSchema) | **Get** /api/v1/scim-schemas/{id} | Retrieve an existing SCIM Schema.
[**ReplaceSCIMSchema**](SchemasApi.md#ReplaceSCIMSchema) | **Put** /api/v1/scim-schemas/{id} | Replace an existing SCIM Schema



## CreateSCIMSchema

> ScimSchema CreateSCIMSchema(ctx).Reason(reason).ScimSchema(scimSchema).Execute()

Create a new SCIM Schema



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
    scimSchema := *openapiclient.NewScimSchema() // ScimSchema | A SCIM Schema, in JSON format. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.CreateSCIMSchema(context.Background()).Reason(reason).ScimSchema(scimSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.CreateSCIMSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSCIMSchema`: ScimSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.CreateSCIMSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSCIMSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | A descriptive reason of the change | 
 **scimSchema** | [**ScimSchema**](ScimSchema.md) | A SCIM Schema, in JSON format. | 

### Return type

[**ScimSchema**](ScimSchema.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSCIMSchema

> ScimSchema GetSCIMSchema(ctx, id).Execute()

Retrieve an existing SCIM Schema.



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
    id := "cidUser" // string | The schema ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.GetSCIMSchema(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetSCIMSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSCIMSchema`: ScimSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetSCIMSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The schema ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSCIMSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimSchema**](ScimSchema.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSCIMSchema

> ScimSchema ReplaceSCIMSchema(ctx, id).Reason(reason).ScimSchema(scimSchema).Execute()

Replace an existing SCIM Schema



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
    id := "cidUser" // string | The schema ID
    reason := "approved by ticket 4423." // string | A descriptive reason of the change
    scimSchema := *openapiclient.NewScimSchema() // ScimSchema | A SCIM Schema, in JSON format. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemasApi.ReplaceSCIMSchema(context.Background(), id).Reason(reason).ScimSchema(scimSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReplaceSCIMSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSCIMSchema`: ScimSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReplaceSCIMSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The schema ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSCIMSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | A descriptive reason of the change | 
 **scimSchema** | [**ScimSchema**](ScimSchema.md) | A SCIM Schema, in JSON format. | 

### Return type

[**ScimSchema**](ScimSchema.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

