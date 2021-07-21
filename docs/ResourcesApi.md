# \ResourcesApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorrelateResources**](ResourcesApi.md#CorrelateResources) | **Post** /resources/.correlate | Correlate Resources
[**CreateResource**](ResourcesApi.md#CreateResource) | **Post** /resources/{appID}/{resourceTypes} | Create a Resource
[**CreateResourceAsync**](ResourcesApi.md#CreateResourceAsync) | **Post** /resources/{appID}/{resourceTypes}/.async | Create a Resource Asynchonously
[**DeleteResourceById**](ResourcesApi.md#DeleteResourceById) | **Delete** /resources/{appID}/{resourceTypes}/{id} | Delete a Resource
[**DeleteResourceByIdAsync**](ResourcesApi.md#DeleteResourceByIdAsync) | **Delete** /resources/{appID}/{resourceTypes}/{id}/.async | Delete a Resource Asynchonously
[**DesiredResource**](ResourcesApi.md#DesiredResource) | **Get** /resources/{appID}/{resourceTypes}/{id}/.desired/ | Get Desired Resource
[**DiffResources**](ResourcesApi.md#DiffResources) | **Post** /resources/ | Diffing Resources
[**GetResourceById**](ResourcesApi.md#GetResourceById) | **Get** /resources/{appID}/{resourceTypes}/{id} | Get a Resource
[**ListResources**](ResourcesApi.md#ListResources) | **Get** /resources/ | List Resources
[**ListResourcesByAppAndType**](ResourcesApi.md#ListResourcesByAppAndType) | **Get** /resources/{appID}/{resourceTypes} | List Resources for a Type
[**MapResources**](ResourcesApi.md#MapResources) | **Post** /resources/.map | Map Resources
[**PatchResource**](ResourcesApi.md#PatchResource) | **Patch** /resources/{appID}/{resourceTypes}/{id} | Update a Resource
[**PatchResourceAsync**](ResourcesApi.md#PatchResourceAsync) | **Patch** /resources/{appID}/{resourceTypes}/{id}/.async | Update a Resource Asynchonously
[**ReplaceResource**](ResourcesApi.md#ReplaceResource) | **Put** /resources/{appID}/{resourceTypes}/{id} | Replace a Resource
[**ReplaceResourceAsync**](ResourcesApi.md#ReplaceResourceAsync) | **Put** /resources/{appID}/{resourceTypes}/{id}/.async | Replace a Resource Asynchonously
[**ResourceRules**](ResourcesApi.md#ResourceRules) | **Post** /resources/{appID}/{resourceTypes}/{id}/.rules | List rules applying on a Resource
[**TestResource**](ResourcesApi.md#TestResource) | **Get** /resources/{appID}/{resourceTypes}/{id}/.test | Test a Resource Against a Filter



## CorrelateResources

> ChangeLog CorrelateResources(ctx).Reason(reason).Correlation(correlation).CorrelateResource(correlateResource).Execute()

Correlate Resources



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
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)
    correlateResource := *openapiclient.NewCorrelateResource("CorrId_example", "Filter_example", "RefPath_example") // CorrelateResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.CorrelateResources(context.Background()).Reason(reason).Correlation(correlation).CorrelateResource(correlateResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.CorrelateResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorrelateResources`: ChangeLog
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.CorrelateResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorrelateResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 
 **correlateResource** | [**CorrelateResource**](CorrelateResource.md) |  | 

### Return type

[**ChangeLog**](ChangeLog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResource

> Resource CreateResource(ctx, appID, resourceTypes).Reason(reason).Correlation(correlation).Resource(resource).Execute()

Create a Resource



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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)
    resource := *openapiclient.NewResource() // Resource |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.CreateResource(context.Background(), appID, resourceTypes).Reason(reason).Correlation(correlation).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.CreateResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.CreateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 
 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Resource**](Resource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResourceAsync

> Job CreateResourceAsync(ctx, appID, resourceTypes).NotBefore(notBefore).Reason(reason).Correlation(correlation).Resource(resource).Execute()

Create a Resource Asynchonously

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    notBefore := time.Now() // time.Time | the time where the operation should be executed. (optional)
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)
    resource := *openapiclient.NewResource() // Resource |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.CreateResourceAsync(context.Background(), appID, resourceTypes).NotBefore(notBefore).Reason(reason).Correlation(correlation).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.CreateResourceAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceAsync`: Job
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.CreateResourceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notBefore** | **time.Time** | the time where the operation should be executed. | 
 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 
 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceById

> DeleteResourceById(ctx, appID, resourceTypes, id).Reason(reason).Correlation(correlation).Execute()

Delete a Resource

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.DeleteResourceById(context.Background(), appID, resourceTypes, id).Reason(reason).Correlation(correlation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DeleteResourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 

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


## DeleteResourceByIdAsync

> Job DeleteResourceByIdAsync(ctx, appID, resourceTypes, id).Reason(reason).NotBefore(notBefore).Correlation(correlation).Execute()

Delete a Resource Asynchonously

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    notBefore := time.Now() // time.Time | the time where the operation should be executed. (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.DeleteResourceByIdAsync(context.Background(), appID, resourceTypes, id).Reason(reason).NotBefore(notBefore).Correlation(correlation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DeleteResourceByIdAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteResourceByIdAsync`: Job
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.DeleteResourceByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **string** | a descriptive reason of the change  | 
 **notBefore** | **time.Time** | the time where the operation should be executed. | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 

### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DesiredResource

> Resource DesiredResource(ctx, appID, resourceTypes, id).Execute()

Get Desired Resource



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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.DesiredResource(context.Background(), appID, resourceTypes, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DesiredResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DesiredResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.DesiredResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDesiredResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Resource**](Resource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiffResources

> ResourceDiffRes DiffResources(ctx).ResourceDiffReq(resourceDiffReq).Execute()

Diffing Resources



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
    resourceDiffReq := *openapiclient.NewResourceDiffReq(*openapiclient.NewResource(), *openapiclient.NewResource(), "Location_example") // ResourceDiffReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.DiffResources(context.Background()).ResourceDiffReq(resourceDiffReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DiffResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiffResources`: ResourceDiffRes
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.DiffResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiffResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceDiffReq** | [**ResourceDiffReq**](ResourceDiffReq.md) |  | 

### Return type

[**ResourceDiffRes**](ResourceDiffRes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceById

> Resource GetResourceById(ctx, appID, resourceTypes, id).ForTime(forTime).Live(live).Execute()

Get a Resource

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    forTime := time.Now() // time.Time | a date time indicating that the requested resources should be retrieved from history as how they looked for the specified time.  (optional)
    live := true // bool | true if the resource should be fetched directly from the application rather from store. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.GetResourceById(context.Background(), appID, resourceTypes, id).ForTime(forTime).Live(live).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.GetResourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceById`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.GetResourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forTime** | **time.Time** | a date time indicating that the requested resources should be retrieved from history as how they looked for the specified time.  | 
 **live** | **bool** | true if the resource should be fetched directly from the application rather from store. | 

### Return type

[**Resource**](Resource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources

> ResourcesList ListResources(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List Resources



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
    resp, r, err := api_client.ResourcesApi.ListResources(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ListResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResources`: ResourcesList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ListResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesRequest struct via the builder pattern


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

[**ResourcesList**](ResourcesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourcesByAppAndType

> ResourcesList ListResourcesByAppAndType(ctx, appID, resourceTypes).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List Resources for a Type



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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
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
    resp, r, err := api_client.ResourcesApi.ListResourcesByAppAndType(context.Background(), appID, resourceTypes).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ListResourcesByAppAndType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResourcesByAppAndType`: ResourcesList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ListResourcesByAppAndType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourcesByAppAndTypeRequest struct via the builder pattern


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

[**ResourcesList**](ResourcesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapResources

> ChangeLog MapResources(ctx).Reason(reason).Correlation(correlation).MapResource(mapResource).Execute()

Map Resources



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
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)
    mapResource := *openapiclient.NewMapResource("CorrId_example", "Filter_example", "MapperId_example") // MapResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.MapResources(context.Background()).Reason(reason).Correlation(correlation).MapResource(mapResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.MapResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MapResources`: ChangeLog
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.MapResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMapResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 
 **mapResource** | [**MapResource**](MapResource.md) |  | 

### Return type

[**ChangeLog**](ChangeLog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchResource

> Resource PatchResource(ctx, appID, resourceTypes, id).PatchRequest(patchRequest).Reason(reason).Correlation(correlation).Execute()

Update a Resource



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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    patchRequest := *openapiclient.NewPatchRequest() // PatchRequest | a request to patch (update) a single model
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.PatchResource(context.Background(), appID, resourceTypes, id).PatchRequest(patchRequest).Reason(reason).Correlation(correlation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.PatchResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.PatchResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequest** | [**PatchRequest**](PatchRequest.md) | a request to patch (update) a single model | 
 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 

### Return type

[**Resource**](Resource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchResourceAsync

> Job PatchResourceAsync(ctx, appID, resourceTypes, id).PatchRequest(patchRequest).Reason(reason).NotBefore(notBefore).Correlation(correlation).Execute()

Update a Resource Asynchonously

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    patchRequest := *openapiclient.NewPatchRequest() // PatchRequest | a request to patch (update) a single model
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    notBefore := time.Now() // time.Time | the time where the operation should be executed. (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.PatchResourceAsync(context.Background(), appID, resourceTypes, id).PatchRequest(patchRequest).Reason(reason).NotBefore(notBefore).Correlation(correlation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.PatchResourceAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchResourceAsync`: Job
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.PatchResourceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchResourceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchRequest** | [**PatchRequest**](PatchRequest.md) | a request to patch (update) a single model | 
 **reason** | **string** | a descriptive reason of the change  | 
 **notBefore** | **time.Time** | the time where the operation should be executed. | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 

### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceResource

> Resource ReplaceResource(ctx, appID, resourceTypes, id).Reason(reason).Correlation(correlation).Resource(resource).Execute()

Replace a Resource



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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)
    resource := *openapiclient.NewResource() // Resource |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.ReplaceResource(context.Background(), appID, resourceTypes, id).Reason(reason).Correlation(correlation).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ReplaceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ReplaceResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **string** | a descriptive reason of the change  | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 
 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Resource**](Resource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceResourceAsync

> Job ReplaceResourceAsync(ctx, appID, resourceTypes, id).Reason(reason).NotBefore(notBefore).Correlation(correlation).Resource(resource).Execute()

Replace a Resource Asynchonously

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    notBefore := time.Now() // time.Time | the time where the operation should be executed. (optional)
    correlation := "/flows/1234" // string | a unique identifier of an external resource where this change should be correlated to  (optional)
    resource := *openapiclient.NewResource() // Resource |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.ReplaceResourceAsync(context.Background(), appID, resourceTypes, id).Reason(reason).NotBefore(notBefore).Correlation(correlation).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ReplaceResourceAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceResourceAsync`: Job
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ReplaceResourceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceResourceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **string** | a descriptive reason of the change  | 
 **notBefore** | **time.Time** | the time where the operation should be executed. | 
 **correlation** | **string** | a unique identifier of an external resource where this change should be correlated to  | 
 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceRules

> RulesList ResourceRules(ctx, appID, resourceTypes, id).ResourceRules(resourceRules).Execute()

List rules applying on a Resource



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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    resourceRules := *openapiclient.NewResourceRules() // ResourceRules |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.ResourceRules(context.Background(), appID, resourceTypes, id).ResourceRules(resourceRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ResourceRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourceRules`: RulesList
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ResourceRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceRules** | [**ResourceRules**](ResourceRules.md) |  | 

### Return type

[**RulesList**](RulesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestResource

> InlineResponse2002 TestResource(ctx, appID, resourceTypes, id).Filter(filter).Execute()

Test a Resource Against a Filter

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
    appID := "azuread" // string | The id of the app the resource belongs to.
    resourceTypes := "Users" // string | the type of the resource.
    id := "111111" // string | the unique identifier of the resource.
    filter := "displayName sw "smith"" // string | The filter string used to request a subset of models.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourcesApi.TestResource(context.Background(), appID, resourceTypes, id).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.TestResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestResource`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.TestResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appID** | **string** | The id of the app the resource belongs to. | 
**resourceTypes** | **string** | the type of the resource. | 
**id** | **string** | the unique identifier of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **string** | The filter string used to request a subset of models.  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

