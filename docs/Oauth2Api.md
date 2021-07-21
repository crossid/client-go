# \Oauth2Api

All URIs are relative to *http://dev.local.crossid.io:8000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptConsentRequest**](Oauth2Api.md#AcceptConsentRequest) | **Put** /oauth2/auth/requests/consent/accept | 
[**AcceptLoginRequest**](Oauth2Api.md#AcceptLoginRequest) | **Put** /oauth2/auth/requests/login/accept | 
[**GetConsentRequest**](Oauth2Api.md#GetConsentRequest) | **Get** /oauth2/auth/requests/consent | 
[**GetLoginRequest**](Oauth2Api.md#GetLoginRequest) | **Get** /oauth2/auth/requests/login | 
[**ListOAUTH2AuthorizationServers**](Oauth2Api.md#ListOAUTH2AuthorizationServers) | **Get** /oauth2/auhtorization-servers/ | List OAUTH 2.0 Authorization Servers.
[**ListOAUTH2Clients**](Oauth2Api.md#ListOAUTH2Clients) | **Get** /oauth2/clients/ | List OAUTH 2.0 Clients.
[**RejectConsentRequest**](Oauth2Api.md#RejectConsentRequest) | **Put** /oauth2/auth/requests/consent/reject | 
[**RejectLoginRequest**](Oauth2Api.md#RejectLoginRequest) | **Put** /oauth2/auth/requests/login/reject | 



## AcceptConsentRequest

> OAuth2RequestHandlerResponse AcceptConsentRequest(ctx).OAuth2HandledConsentRequest(oAuth2HandledConsentRequest).Challenge(challenge).Execute()



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
    oAuth2HandledConsentRequest := *openapiclient.NewOAuth2HandledConsentRequest() // OAuth2HandledConsentRequest | 
    challenge := "challenge_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Oauth2Api.AcceptConsentRequest(context.Background()).OAuth2HandledConsentRequest(oAuth2HandledConsentRequest).Challenge(challenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.AcceptConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptConsentRequest`: OAuth2RequestHandlerResponse
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.AcceptConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2HandledConsentRequest** | [**OAuth2HandledConsentRequest**](OAuth2HandledConsentRequest.md) |  | 
 **challenge** | **string** |  | 

### Return type

[**OAuth2RequestHandlerResponse**](OAuth2RequestHandlerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcceptLoginRequest

> OAuth2RequestHandlerResponse AcceptLoginRequest(ctx).OAuth2HandledLoginRequest(oAuth2HandledLoginRequest).Challenge(challenge).Execute()



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
    oAuth2HandledLoginRequest := *openapiclient.NewOAuth2HandledLoginRequest() // OAuth2HandledLoginRequest | 
    challenge := "challenge_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Oauth2Api.AcceptLoginRequest(context.Background()).OAuth2HandledLoginRequest(oAuth2HandledLoginRequest).Challenge(challenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.AcceptLoginRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptLoginRequest`: OAuth2RequestHandlerResponse
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.AcceptLoginRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptLoginRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2HandledLoginRequest** | [**OAuth2HandledLoginRequest**](OAuth2HandledLoginRequest.md) |  | 
 **challenge** | **string** |  | 

### Return type

[**OAuth2RequestHandlerResponse**](OAuth2RequestHandlerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsentRequest

> OAuth2ConsentRequest GetConsentRequest(ctx).Challenge(challenge).Execute()



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
    challenge := "challenge_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Oauth2Api.GetConsentRequest(context.Background()).Challenge(challenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.GetConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentRequest`: OAuth2ConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.GetConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challenge** | **string** |  | 

### Return type

[**OAuth2ConsentRequest**](OAuth2ConsentRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginRequest

> OAuth2LoginRequest GetLoginRequest(ctx).Challenge(challenge).Execute()



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
    challenge := "challenge_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Oauth2Api.GetLoginRequest(context.Background()).Challenge(challenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.GetLoginRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginRequest`: OAuth2LoginRequest
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.GetLoginRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challenge** | **string** |  | 

### Return type

[**OAuth2LoginRequest**](OAuth2LoginRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAUTH2AuthorizationServers

> ListOAUTH2AuthorizationServers(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List OAUTH 2.0 Authorization Servers.

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
    resp, r, err := api_client.Oauth2Api.ListOAUTH2AuthorizationServers(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.ListOAUTH2AuthorizationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOAUTH2AuthorizationServersRequest struct via the builder pattern


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

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAUTH2Clients

> Oauth2ClientList ListOAUTH2Clients(ctx).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()

List OAUTH 2.0 Clients.

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
    resp, r, err := api_client.Oauth2Api.ListOAUTH2Clients(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).Attributes(attributes).ExcludedAttributes(excludedAttributes).ForTime(forTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.ListOAUTH2Clients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAUTH2Clients`: Oauth2ClientList
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.ListOAUTH2Clients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOAUTH2ClientsRequest struct via the builder pattern


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

[**Oauth2ClientList**](Oauth2ClientList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectConsentRequest

> OAuth2RequestHandlerResponse RejectConsentRequest(ctx).OAuth2RequestDeniedError(oAuth2RequestDeniedError).Challenge(challenge).Execute()



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
    oAuth2RequestDeniedError := *openapiclient.NewOAuth2RequestDeniedError() // OAuth2RequestDeniedError | 
    challenge := "challenge_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Oauth2Api.RejectConsentRequest(context.Background()).OAuth2RequestDeniedError(oAuth2RequestDeniedError).Challenge(challenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.RejectConsentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectConsentRequest`: OAuth2RequestHandlerResponse
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.RejectConsentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectConsentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2RequestDeniedError** | [**OAuth2RequestDeniedError**](OAuth2RequestDeniedError.md) |  | 
 **challenge** | **string** |  | 

### Return type

[**OAuth2RequestHandlerResponse**](OAuth2RequestHandlerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectLoginRequest

> OAuth2RequestHandlerResponse RejectLoginRequest(ctx).OAuth2RequestDeniedError(oAuth2RequestDeniedError).Challenge(challenge).Execute()



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
    oAuth2RequestDeniedError := *openapiclient.NewOAuth2RequestDeniedError() // OAuth2RequestDeniedError | 
    challenge := "challenge_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Oauth2Api.RejectLoginRequest(context.Background()).OAuth2RequestDeniedError(oAuth2RequestDeniedError).Challenge(challenge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.RejectLoginRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectLoginRequest`: OAuth2RequestHandlerResponse
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.RejectLoginRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectLoginRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2RequestDeniedError** | [**OAuth2RequestDeniedError**](OAuth2RequestDeniedError.md) |  | 
 **challenge** | **string** |  | 

### Return type

[**OAuth2RequestHandlerResponse**](OAuth2RequestHandlerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

