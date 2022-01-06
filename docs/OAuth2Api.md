# \OAuth2Api

All URIs are relative to *https://acme.us.crossid.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2Token**](OAuth2Api.md#Oauth2Token) | **Post** /oauth2/token | Request a token from the OAuth2 authorization server.



## Oauth2Token

> Token Oauth2Token(ctx).TokenRequest(tokenRequest).Execute()

Request a token from the OAuth2 authorization server.



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
    tokenRequest := TODO // TokenRequest | Token Request, as form data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OAuth2Api.Oauth2Token(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2Api.Oauth2Token``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2Token`: Token
    fmt.Fprintf(os.Stdout, "Response from `OAuth2Api.Oauth2Token`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2TokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) | Token Request, as form data. | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

