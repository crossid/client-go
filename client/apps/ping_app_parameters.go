// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewPingAppParams creates a new PingAppParams object
// with the default values initialized.
func NewPingAppParams() *PingAppParams {
	var ()
	return &PingAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPingAppParamsWithTimeout creates a new PingAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPingAppParamsWithTimeout(timeout time.Duration) *PingAppParams {
	var ()
	return &PingAppParams{

		timeout: timeout,
	}
}

// NewPingAppParamsWithContext creates a new PingAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewPingAppParamsWithContext(ctx context.Context) *PingAppParams {
	var ()
	return &PingAppParams{

		Context: ctx,
	}
}

// NewPingAppParamsWithHTTPClient creates a new PingAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPingAppParamsWithHTTPClient(client *http.Client) *PingAppParams {
	var ()
	return &PingAppParams{
		HTTPClient: client,
	}
}

/*PingAppParams contains all the parameters to send to the API endpoint
for the ping app operation typically these are written to a http.Request
*/
type PingAppParams struct {

	/*AppID*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ping app params
func (o *PingAppParams) WithTimeout(timeout time.Duration) *PingAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping app params
func (o *PingAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping app params
func (o *PingAppParams) WithContext(ctx context.Context) *PingAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping app params
func (o *PingAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping app params
func (o *PingAppParams) WithHTTPClient(client *http.Client) *PingAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping app params
func (o *PingAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the ping app params
func (o *PingAppParams) WithAppID(appID string) *PingAppParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the ping app params
func (o *PingAppParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *PingAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appID
	if err := r.SetPathParam("appID", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}