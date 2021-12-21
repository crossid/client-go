/*
Crossid API

# Introduction  Crossid API allows you to manage resources in a simple, programmatic way using conventional HTTP requests.  All of the functionality that you are familiar with in the Crossid UI is also available through this API, allowing you to code actions that suites your requirements.  The rest of this page provides a general overview about the API design and technology that has been implemented.  ## Requests  ## HTTP Statuses  ## Meta  ### Sample Links Object  ## Rate Limit  ## Curl Examples  ## OpenAPI Specification  Crossid API conforms the OpenAPI V3 specification.  The goal of The OpenAPI Specification is to define a standard, language-agnostic interface to REST APIs which  allows both humans and computers to discover and understand the capabilities of the service without access to source  code, documentation, or through network traffic inspection. When properly defined via OpenAPI, a consumer can  understand and interact with the remote service with a minimal amount of implementation logic. Similar to what  interfaces have done for lower-level programming, OpenAPI removes the guesswork in calling the service. 

API version: 2.0.0
Contact: api-engineering@crossid.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cidclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SchemasApiService SchemasApi service
type SchemasApiService service

type ApiCreateSCIMSchemaRequest struct {
	ctx _context.Context
	ApiService *SchemasApiService
	reason *string
	scimSchema *ScimSchema
}

// A descriptive reason of the change
func (r ApiCreateSCIMSchemaRequest) Reason(reason string) ApiCreateSCIMSchemaRequest {
	r.reason = &reason
	return r
}
// A SCIM Schema, in JSON format.
func (r ApiCreateSCIMSchemaRequest) ScimSchema(scimSchema ScimSchema) ApiCreateSCIMSchemaRequest {
	r.scimSchema = &scimSchema
	return r
}

func (r ApiCreateSCIMSchemaRequest) Execute() (ScimSchema, *_nethttp.Response, error) {
	return r.ApiService.CreateSCIMSchemaExecute(r)
}

/*
CreateSCIMSchema Create a new SCIM Schema

Create a new SCIM Schema with the given payload.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateSCIMSchemaRequest
*/
func (a *SchemasApiService) CreateSCIMSchema(ctx _context.Context) ApiCreateSCIMSchemaRequest {
	return ApiCreateSCIMSchemaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ScimSchema
func (a *SchemasApiService) CreateSCIMSchemaExecute(r ApiCreateSCIMSchemaRequest) (ScimSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScimSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemasApiService.CreateSCIMSchema")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/scim-schemas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.reason == nil {
		return localVarReturnValue, nil, reportError("reason is required and must be specified")
	}

	localVarQueryParams.Add("reason", parameterToString(*r.reason, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.scimSchema
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSCIMSchemaRequest struct {
	ctx _context.Context
	ApiService *SchemasApiService
	id string
}


func (r ApiGetSCIMSchemaRequest) Execute() (ScimSchema, *_nethttp.Response, error) {
	return r.ApiService.GetSCIMSchemaExecute(r)
}

/*
GetSCIMSchema Retrieve an existing SCIM Schema.

Retrieve details about an existing SCIM Schema by its id.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The schema ID
 @return ApiGetSCIMSchemaRequest
*/
func (a *SchemasApiService) GetSCIMSchema(ctx _context.Context, id string) ApiGetSCIMSchemaRequest {
	return ApiGetSCIMSchemaRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ScimSchema
func (a *SchemasApiService) GetSCIMSchemaExecute(r ApiGetSCIMSchemaRequest) (ScimSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScimSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemasApiService.GetSCIMSchema")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/scim-schemas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReplaceSCIMSchemaRequest struct {
	ctx _context.Context
	ApiService *SchemasApiService
	id string
	reason *string
	scimSchema *ScimSchema
}

// A descriptive reason of the change
func (r ApiReplaceSCIMSchemaRequest) Reason(reason string) ApiReplaceSCIMSchemaRequest {
	r.reason = &reason
	return r
}
// A SCIM Schema, in JSON format.
func (r ApiReplaceSCIMSchemaRequest) ScimSchema(scimSchema ScimSchema) ApiReplaceSCIMSchemaRequest {
	r.scimSchema = &scimSchema
	return r
}

func (r ApiReplaceSCIMSchemaRequest) Execute() (ScimSchema, *_nethttp.Response, error) {
	return r.ApiService.ReplaceSCIMSchemaExecute(r)
}

/*
ReplaceSCIMSchema Replace an existing SCIM Schema

Create an existing SCIM Schema with the given payload.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The schema ID
 @return ApiReplaceSCIMSchemaRequest
*/
func (a *SchemasApiService) ReplaceSCIMSchema(ctx _context.Context, id string) ApiReplaceSCIMSchemaRequest {
	return ApiReplaceSCIMSchemaRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ScimSchema
func (a *SchemasApiService) ReplaceSCIMSchemaExecute(r ApiReplaceSCIMSchemaRequest) (ScimSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScimSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemasApiService.ReplaceSCIMSchema")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/scim-schemas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.reason == nil {
		return localVarReturnValue, nil, reportError("reason is required and must be specified")
	}

	localVarQueryParams.Add("reason", parameterToString(*r.reason, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.scimSchema
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
