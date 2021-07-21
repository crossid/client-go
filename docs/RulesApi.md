# \RulesApi

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](RulesApi.md#CreateRule) | **Post** /rules/ | Create a Rule
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /rules/{id} | Delete a Rule
[**GetRule**](RulesApi.md#GetRule) | **Get** /rules/{id} | Get a Rule
[**ListRules**](RulesApi.md#ListRules) | **Get** /rules/ | List Rules
[**ReplaceRule**](RulesApi.md#ReplaceRule) | **Put** /rules/{id} | Replace a Rule



## CreateRule

> Rule CreateRule(ctx).Reason(reason).Rule(rule).Execute()

Create a Rule

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
    rule := *openapiclient.NewRule("DisplayName_example", "Id_example", *openapiclient.NewRuleResult()) // Rule |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.CreateRule(context.Background()).Reason(reason).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: Rule
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reason** | **string** | a descriptive reason of the change  | 
 **rule** | [**Rule**](Rule.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, id).Reason(reason).Execute()

Delete a Rule

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
    id := "id_example" // string | The id of the rule to delete.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.DeleteRule(context.Background(), id).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetRule

> Rule GetRule(ctx, id).Execute()

Get a Rule

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
    id := "id_example" // string | The id of the rules to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.GetRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: Rule
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the rules to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rule**](Rule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> RulesList ListRules(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List Rules

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
    resp, r, err := api_client.RulesApi.ListRules(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.ListRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRules`: RulesList
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.ListRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


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

[**RulesList**](RulesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRule

> Rule ReplaceRule(ctx, id).Reason(reason).Rule(rule).Execute()

Replace a Rule

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
    id := "id_example" // string | The id of the rule to replace.
    reason := "required due to..." // string | a descriptive reason of the change  (optional)
    rule := *openapiclient.NewRule("DisplayName_example", "Id_example", *openapiclient.NewRuleResult()) // Rule |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RulesApi.ReplaceRule(context.Background(), id).Reason(reason).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.ReplaceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRule`: Rule
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.ReplaceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the rule to replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | a descriptive reason of the change  | 
 **rule** | [**Rule**](Rule.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

