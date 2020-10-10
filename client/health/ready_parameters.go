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

// NewReadyParams creates a new ReadyParams object
// with the default values initialized.
func NewReadyParams() *ReadyParams {

	return &ReadyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadyParamsWithTimeout creates a new ReadyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadyParamsWithTimeout(timeout time.Duration) *ReadyParams {

	return &ReadyParams{

		timeout: timeout,
	}
}

// NewReadyParamsWithContext creates a new ReadyParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadyParamsWithContext(ctx context.Context) *ReadyParams {

	return &ReadyParams{

		Context: ctx,
	}
}

// NewReadyParamsWithHTTPClient creates a new ReadyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadyParamsWithHTTPClient(client *http.Client) *ReadyParams {

	return &ReadyParams{
		HTTPClient: client,
	}
}

/*ReadyParams contains all the parameters to send to the API endpoint
for the ready operation typically these are written to a http.Request
*/
type ReadyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ready params
func (o *ReadyParams) WithTimeout(timeout time.Duration) *ReadyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ready params
func (o *ReadyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ready params
func (o *ReadyParams) WithContext(ctx context.Context) *ReadyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ready params
func (o *ReadyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ready params
func (o *ReadyParams) WithHTTPClient(client *http.Client) *ReadyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ready params
func (o *ReadyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
