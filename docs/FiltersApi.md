# \FiltersApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilter**](FiltersApi.md#CreateFilter) | **Post** /filters/ | Create a Filter
[**DeleteFilter**](FiltersApi.md#DeleteFilter) | **Delete** /filters/{id} | Delete a Filter
[**GetFilter**](FiltersApi.md#GetFilter) | **Get** /filters/{id} | Get a Filter
[**ListFilters**](FiltersApi.md#ListFilters) | **Get** /filters/ | List Filters
[**ReplaceFilter**](FiltersApi.md#ReplaceFilter) | **Put** /filters/{id} | Replace a Filter
[**ValidFilter**](FiltersApi.md#ValidFilter) | **Post** /filters/.valid | Check Filter Validity



## CreateFilter

> Filter CreateFilter(ctx).Reason(reason).Filter(filter).Execute()

Create a Filter

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
    filter := *openapiclient.NewFilter() // Filter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiltersApi.CreateFilter(context.Background()).Reason(reason).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.CreateFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFilter`: Filter
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.CreateFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **filter** | [**Filter**](Filter.md) |  | 

### Return type

[**Filter**](Filter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFilter

> DeleteFilter(ctx, id).Reason(reason).Execute()

Delete a Filter

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
    id := "id_example" // string | The id of the filter to delete.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiltersApi.DeleteFilter(context.Background(), id).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.DeleteFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilter

> Filter GetFilter(ctx, id).Execute()

Get a Filter

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
    id := "id_example" // string | The id of the filter to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiltersApi.GetFilter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.GetFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFilter`: Filter
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.GetFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Filter**](Filter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilters

> FiltersList ListFilters(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List Filters

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
    resp, r, err := api_client.FiltersApi.ListFilters(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.ListFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFilters`: FiltersList
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.ListFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFiltersRequest struct via the builder pattern


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

[**FiltersList**](FiltersList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFilter

> Filter ReplaceFilter(ctx, id).Reason(reason).Filter(filter).Execute()

Replace a Filter

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
    id := "id_example" // string | The id of the filter to replace.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    filter := *openapiclient.NewFilter() // Filter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiltersApi.ReplaceFilter(context.Background(), id).Reason(reason).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.ReplaceFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceFilter`: Filter
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.ReplaceFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the filter to replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 
 **filter** | [**Filter**](Filter.md) |  | 

### Return type

[**Filter**](Filter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidFilter

> InlineResponse2001 ValidFilter(ctx).Execute()

Check Filter Validity



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FiltersApi.ValidFilter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FiltersApi.ValidFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidFilter`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `FiltersApi.ValidFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidFilterRequest struct via the builder pattern


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

