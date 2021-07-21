# \AuditApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommit**](AuditApi.md#GetCommit) | **Get** /audit/commits/{id} | Get a Commit
[**GetEventLog**](AuditApi.md#GetEventLog) | **Get** /event-logs/{id} | Get an event log
[**ListCommits**](AuditApi.md#ListCommits) | **Get** /audit/commits | List Commits



## GetCommit

> Commit GetCommit(ctx, id).Execute()

Get a Commit

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
    id := "id_example" // string | The id of the commit to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditApi.GetCommit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.GetCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommit`: Commit
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.GetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the commit to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Commit**](Commit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventLog

> EventLog GetEventLog(ctx, id).Execute()

Get an event log

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
    id := "id_example" // string | The id of the log to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditApi.GetEventLog(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.GetEventLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventLog`: EventLog
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.GetEventLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the log to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventLog**](EventLog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommits

> CommitsList ListCommits(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

List Commits

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
    filter := "displayName sw "smith"" // string | The filter string used to request a subset of models.  (optional)
    count := int64(789) // int64 | The desired maximum number of query results per page, e.g., 10. A negative value is interpreted as \"0\". A value of \"0\" indicates that no model results are to be returned except for \"totalResults\".  (optional) (default to 10)
    startIndex := int64(789) // int64 | The 1-based index of the first query result. A value less than 1 SHALL be interpreted as 1.  (optional) (default to 0)
    sortBy := "userName" // string | A string indicating the attribute whose value SHALL be used to order the returned responses. (optional)
    sortOrder := "ascending" // string | A string indicating the order in which the \"sortBy\" parameter is applied.  Allowed values are \"ascending\" and \"descending\". (optional)
    attributes := []string{"Inner_example"} // []string | A multi-valued list of strings indicating the names of resource attributes to return in the response, overriding the set of attributes that would be returned by default.  (optional)
    excludedAttributes := []string{"Inner_example"} // []string | A multi-valued list of strings indicating the names of resource attributes to be removed from the default set of attributes to return.  This parameter SHALL have no effect on attributes whose schema \"returned\" setting is \"always\".  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditApi.ListCommits(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.ListCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommits`: CommitsList
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.ListCommits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter string used to request a subset of models.  | 
 **count** | **int64** | The desired maximum number of query results per page, e.g., 10. A negative value is interpreted as \&quot;0\&quot;. A value of \&quot;0\&quot; indicates that no model results are to be returned except for \&quot;totalResults\&quot;.  | [default to 10]
 **startIndex** | **int64** | The 1-based index of the first query result. A value less than 1 SHALL be interpreted as 1.  | [default to 0]
 **sortBy** | **string** | A string indicating the attribute whose value SHALL be used to order the returned responses. | 
 **sortOrder** | **string** | A string indicating the order in which the \&quot;sortBy\&quot; parameter is applied.  Allowed values are \&quot;ascending\&quot; and \&quot;descending\&quot;. | 
 **attributes** | **[]string** | A multi-valued list of strings indicating the names of resource attributes to return in the response, overriding the set of attributes that would be returned by default.  | 
 **excludedAttributes** | **[]string** | A multi-valued list of strings indicating the names of resource attributes to be removed from the default set of attributes to return.  This parameter SHALL have no effect on attributes whose schema \&quot;returned\&quot; setting is \&quot;always\&quot;.  | 

### Return type

[**CommitsList**](CommitsList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

