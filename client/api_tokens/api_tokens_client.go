// Code generated by go-swagger; DO NOT EDIT.

package api_tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new api tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	InsertAPIToken(params *InsertAPITokenParams, authInfo runtime.ClientAuthInfoWriter) (*InsertAPITokenCreated, error)

	ListAPITokens(params *ListAPITokensParams, authInfo runtime.ClientAuthInfoWriter) (*ListAPITokensOK, error)

	RevokeAPIToken(params *RevokeAPITokenParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeAPITokenNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InsertAPIToken Issues a new API token
*/
func (a *Client) InsertAPIToken(params *InsertAPITokenParams, authInfo runtime.ClientAuthInfoWriter) (*InsertAPITokenCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInsertAPITokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "insertApiToken",
		Method:             "POST",
		PathPattern:        "/api-tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InsertAPITokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InsertAPITokenCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for insertApiToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAPITokens List API tokens
*/
func (a *Client) ListAPITokens(params *ListAPITokensParams, authInfo runtime.ClientAuthInfoWriter) (*ListAPITokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAPITokensParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listApiTokens",
		Method:             "GET",
		PathPattern:        "/api-tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAPITokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAPITokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listApiTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RevokeAPIToken revokes an API token
*/
func (a *Client) RevokeAPIToken(params *RevokeAPITokenParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeAPITokenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeAPITokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revokeApiToken",
		Method:             "DELETE",
		PathPattern:        "/api-tokens/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevokeAPITokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeAPITokenNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revokeApiToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}