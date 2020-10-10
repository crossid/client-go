// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crossid/crossid-client-go/models"
)

// PingAppReader is a Reader for the PingApp structure.
type PingAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPingAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPingAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPingAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPingAppNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPingAppOK creates a PingAppOK with default headers values
func NewPingAppOK() *PingAppOK {
	return &PingAppOK{}
}

/*PingAppOK handles this case with default header values.

Application ping response.
*/
type PingAppOK struct {
	Payload *PingAppOKBody
}

func (o *PingAppOK) Error() string {
	return fmt.Sprintf("[POST /apps/{appID}/.ping][%d] pingAppOK  %+v", 200, o.Payload)
}

func (o *PingAppOK) GetPayload() *PingAppOKBody {
	return o.Payload
}

func (o *PingAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PingAppOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingAppBadRequest creates a PingAppBadRequest with default headers values
func NewPingAppBadRequest() *PingAppBadRequest {
	return &PingAppBadRequest{}
}

/*PingAppBadRequest handles this case with default header values.

genericError
*/
type PingAppBadRequest struct {
	Payload models.GenericError
}

func (o *PingAppBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appID}/.ping][%d] pingAppBadRequest  %+v", 400, o.Payload)
}

func (o *PingAppBadRequest) GetPayload() models.GenericError {
	return o.Payload
}

func (o *PingAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingAppNotFound creates a PingAppNotFound with default headers values
func NewPingAppNotFound() *PingAppNotFound {
	return &PingAppNotFound{}
}

/*PingAppNotFound handles this case with default header values.

genericError
*/
type PingAppNotFound struct {
	Payload models.GenericError
}

func (o *PingAppNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{appID}/.ping][%d] pingAppNotFound  %+v", 404, o.Payload)
}

func (o *PingAppNotFound) GetPayload() models.GenericError {
	return o.Payload
}

func (o *PingAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingAppInternalServerError creates a PingAppInternalServerError with default headers values
func NewPingAppInternalServerError() *PingAppInternalServerError {
	return &PingAppInternalServerError{}
}

/*PingAppInternalServerError handles this case with default header values.

genericError
*/
type PingAppInternalServerError struct {
	Payload models.GenericError
}

func (o *PingAppInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{appID}/.ping][%d] pingAppInternalServerError  %+v", 500, o.Payload)
}

func (o *PingAppInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *PingAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingAppNotImplemented creates a PingAppNotImplemented with default headers values
func NewPingAppNotImplemented() *PingAppNotImplemented {
	return &PingAppNotImplemented{}
}

/*PingAppNotImplemented handles this case with default header values.

genericError
*/
type PingAppNotImplemented struct {
	Payload models.GenericError
}

func (o *PingAppNotImplemented) Error() string {
	return fmt.Sprintf("[POST /apps/{appID}/.ping][%d] pingAppNotImplemented  %+v", 501, o.Payload)
}

func (o *PingAppNotImplemented) GetPayload() models.GenericError {
	return o.Payload
}

func (o *PingAppNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PingAppOKBody ping app o k body
swagger:model PingAppOKBody
*/
type PingAppOKBody struct {

	// latency
	Latency int64 `json:"latency,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`
}

// Validate validates this ping app o k body
func (o *PingAppOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PingAppOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PingAppOKBody) UnmarshalBinary(b []byte) error {
	var res PingAppOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}