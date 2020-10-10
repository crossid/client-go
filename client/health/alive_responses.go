// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/crossid/crossid-client-go/models"
)

// AliveReader is a Reader for the Alive structure.
type AliveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AliveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAliveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewAliveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAliveOK creates a AliveOK with default headers values
func NewAliveOK() *AliveOK {
	return &AliveOK{}
}

/*AliveOK handles this case with default header values.

healthStatus
*/
type AliveOK struct {
	Payload *models.HealthStatus
}

func (o *AliveOK) Error() string {
	return fmt.Sprintf("[GET /health/alive][%d] aliveOK  %+v", 200, o.Payload)
}

func (o *AliveOK) GetPayload() *models.HealthStatus {
	return o.Payload
}

func (o *AliveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAliveServiceUnavailable creates a AliveServiceUnavailable with default headers values
func NewAliveServiceUnavailable() *AliveServiceUnavailable {
	return &AliveServiceUnavailable{}
}

/*AliveServiceUnavailable handles this case with default header values.

healthStatus
*/
type AliveServiceUnavailable struct {
	Payload *models.HealthStatus
}

func (o *AliveServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /health/alive][%d] aliveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AliveServiceUnavailable) GetPayload() *models.HealthStatus {
	return o.Payload
}

func (o *AliveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}