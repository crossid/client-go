// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAliveParams creates a new AliveParams object
// with the default values initialized.
func NewAliveParams() *AliveParams {

	return &AliveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAliveParamsWithTimeout creates a new AliveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAliveParamsWithTimeout(timeout time.Duration) *AliveParams {

	return &AliveParams{

		timeout: timeout,
	}
}

// NewAliveParamsWithContext creates a new AliveParams object
// with the default values initialized, and the ability to set a context for a request
func NewAliveParamsWithContext(ctx context.Context) *AliveParams {

	return &AliveParams{

		Context: ctx,
	}
}

// NewAliveParamsWithHTTPClient creates a new AliveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAliveParamsWithHTTPClient(client *http.Client) *AliveParams {

	return &AliveParams{
		HTTPClient: client,
	}
}

/*AliveParams contains all the parameters to send to the API endpoint
for the alive operation typically these are written to a http.Request
*/
type AliveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alive params
func (o *AliveParams) WithTimeout(timeout time.Duration) *AliveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alive params
func (o *AliveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alive params
func (o *AliveParams) WithContext(ctx context.Context) *AliveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alive params
func (o *AliveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alive params
func (o *AliveParams) WithHTTPClient(client *http.Client) *AliveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alive params
func (o *AliveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AliveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
