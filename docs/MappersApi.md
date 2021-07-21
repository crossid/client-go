# \MappersApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMapper**](MappersApi.md#CreateMapper) | **Post** /mappers/ | Create a mapper
[**GetMapper**](MappersApi.md#GetMapper) | **Get** /mappers/{id} | Get a Mapper
[**ListMappers**](MappersApi.md#ListMappers) | **Get** /mappers/ | List Mappers
[**PatchMapper**](MappersApi.md#PatchMapper) | **Patch** /mappers/{id} | Update a Mapper
[**ReplaceMapper**](MappersApi.md#ReplaceMapper) | **Put** /mappers/{id} | Replace a Mapper



## CreateMapper

> Mapper CreateMapper(ctx).Reason(reason).Mapper(mapper).Execute()

Create a mapper

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
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    mapper := *openapiclient.NewMapper() // Mapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MappersApi.CreateMapper(context.Background()).Reason(reason).Mapper(mapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappersApi.CreateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMapper`: Mapper
    fmt.Fprintf(os.Stdout, "Response from `MappersApi.CreateMapper`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **mapper** | [**Mapper**](Mapper.md) |  | 

### Return type

[**Mapper**](Mapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapper

> Mapper GetMapper(ctx, id).Execute()

Get a Mapper

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
    id := "id_example" // string | The id of the mappers to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MappersApi.GetMapper(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappersApi.GetMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapper`: Mapper
    fmt.Fprintf(os.Stdout, "Response from `MappersApi.GetMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the mappers to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mapper**](Mapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappers

> MappersList ListMappers(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List Mappers

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
    filter := "displayName sw "smith"" // string | The filter string used to request a subset of models.  (optional)
    count := int64(789) // int64 | The desired maximum number of query results per page, e.g., 10. A negative value is interpreted as \"0\". A value of \"0\" indicates that no model results are to be returned except for \"totalResults\".  (optional) (default to 10)
    startIndex := int64(789) // int64 | The 1-based index of the first query result. A value less than 1 SHALL be interpreted as 1.  (optional) (default to 0)
    sortBy := "userName" // string | A string indicating the attribute whose value SHALL be used to order the returned responses. (optional)
    sortOrder := "ascending" // string | A string indicating the order in which the \"sortBy\" parameter is applied.  Allowed values are \"ascending\" and \"descending\". (optional)
    attributes := []string{"Inner_example"} // []string | A multi-valued list of strings indicating the names of resource attributes to return in the response, overriding the set of attributes that would be returned by default.  (optional)
    excludedAttributes := []string{"Inner_example"} // []string | A multi-valued list of strings indicating the names of resource attributes to be removed from the default set of attributes to return.  This parameter SHALL have no effect on attributes whose schema \"returned\" setting is \"always\".  (optional)
    forTime := time.Now() // time.Time | a date time indicating that the requested resources should be retrieved from history as how they looked for the specified time.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MappersApi.ListMappers(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappersApi.ListMappers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappers`: MappersList
    fmt.Fprintf(os.Stdout, "Response from `MappersApi.ListMappers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter string used to request a subset of models.  | 
 **count** | **int64** | The desired maximum number of query results per page, e.g., 10. A negative value is interpreted as \&quot;0\&quot;. A value of \&quot;0\&quot; indicates that no model results are to be returned except for \&quot;totalResults\&quot;.  | [default to 10]
 **startIndex** | **int64** | The 1-based index of the first query result. A value less than 1 SHALL be interpreted as 1.  | [default to 0]
 **sortBy** | **string** | A string indicating the attribute whose value SHALL be used to order the returned responses. | 
 **sortOrder** | **string** | A string indicating the order in which the \&quot;sortBy\&quot; parameter is applied.  Allowed values are \&quot;ascending\&quot; and \&quot;descending\&quot;. | 
 **attributes** | **[]string** | A multi-valued list of strings indicating the names of resource attributes to return in the response, overriding the set of attributes that would be returned by default.  | 
 **excludedAttributes** | **[]string** | A multi-valued list of strings indicating the names of resource attributes to be removed from the default set of attributes to return.  This parameter SHALL have no effect on attributes whose schema \&quot;returned\&quot; setting is \&quot;always\&quot;.  | 
 **forTime** | **time.Time** | a date time indicating that the requested resources should be retrieved from history as how they looked for the specified time.  | 

### Return type

[**MappersList**](MappersList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMapper

> Mapper PatchMapper(ctx, id).Reason(reason).Mapper(mapper).Execute()

Update a Mapper

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
    id := "id_example" // string | The id of the mapper to patch.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    mapper := *openapiclient.NewMapper() // Mapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MappersApi.PatchMapper(context.Background(), id).Reason(reason).Mapper(mapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappersApi.PatchMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMapper`: Mapper
    fmt.Fprintf(os.Stdout, "Response from `MappersApi.PatchMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the mapper to patch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 
 **mapper** | [**Mapper**](Mapper.md) |  | 

### Return type

[**Mapper**](Mapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMapper

> Mapper ReplaceMapper(ctx, id).Reason(reason).Mapper(mapper).Execute()

Replace a Mapper

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
    id := "id_example" // string | The id of the mappers to replace.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    mapper := *openapiclient.NewMapper() // Mapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MappersApi.ReplaceMapper(context.Background(), id).Reason(reason).Mapper(mapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MappersApi.ReplaceMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMapper`: Mapper
    fmt.Fprintf(os.Stdout, "Response from `MappersApi.ReplaceMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the mappers to replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 
 **mapper** | [**Mapper**](Mapper.md) |  | 

### Return type

[**Mapper**](Mapper.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

