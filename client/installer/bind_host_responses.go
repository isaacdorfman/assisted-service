// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// BindHostReader is a Reader for the BindHost structure.
type BindHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBindHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBindHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBindHostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBindHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewBindHostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBindHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewBindHostNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewBindHostServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBindHostOK creates a BindHostOK with default headers values
func NewBindHostOK() *BindHostOK {
	return &BindHostOK{}
}

/*BindHostOK handles this case with default header values.

Success.
*/
type BindHostOK struct {
	Payload *models.Host
}

func (o *BindHostOK) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostOK  %+v", 200, o.Payload)
}

func (o *BindHostOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *BindHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostBadRequest creates a BindHostBadRequest with default headers values
func NewBindHostBadRequest() *BindHostBadRequest {
	return &BindHostBadRequest{}
}

/*BindHostBadRequest handles this case with default header values.

Error.
*/
type BindHostBadRequest struct {
	Payload *models.Error
}

func (o *BindHostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostBadRequest  %+v", 400, o.Payload)
}

func (o *BindHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostUnauthorized creates a BindHostUnauthorized with default headers values
func NewBindHostUnauthorized() *BindHostUnauthorized {
	return &BindHostUnauthorized{}
}

/*BindHostUnauthorized handles this case with default header values.

Unauthorized.
*/
type BindHostUnauthorized struct {
	Payload *models.InfraError
}

func (o *BindHostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostUnauthorized  %+v", 401, o.Payload)
}

func (o *BindHostUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *BindHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostForbidden creates a BindHostForbidden with default headers values
func NewBindHostForbidden() *BindHostForbidden {
	return &BindHostForbidden{}
}

/*BindHostForbidden handles this case with default header values.

Forbidden.
*/
type BindHostForbidden struct {
	Payload *models.InfraError
}

func (o *BindHostForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostForbidden  %+v", 403, o.Payload)
}

func (o *BindHostForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *BindHostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostNotFound creates a BindHostNotFound with default headers values
func NewBindHostNotFound() *BindHostNotFound {
	return &BindHostNotFound{}
}

/*BindHostNotFound handles this case with default header values.

Error.
*/
type BindHostNotFound struct {
	Payload *models.Error
}

func (o *BindHostNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostNotFound  %+v", 404, o.Payload)
}

func (o *BindHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostMethodNotAllowed creates a BindHostMethodNotAllowed with default headers values
func NewBindHostMethodNotAllowed() *BindHostMethodNotAllowed {
	return &BindHostMethodNotAllowed{}
}

/*BindHostMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type BindHostMethodNotAllowed struct {
	Payload *models.Error
}

func (o *BindHostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *BindHostMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindHostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostInternalServerError creates a BindHostInternalServerError with default headers values
func NewBindHostInternalServerError() *BindHostInternalServerError {
	return &BindHostInternalServerError{}
}

/*BindHostInternalServerError handles this case with default header values.

Error.
*/
type BindHostInternalServerError struct {
	Payload *models.Error
}

func (o *BindHostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostInternalServerError  %+v", 500, o.Payload)
}

func (o *BindHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostNotImplemented creates a BindHostNotImplemented with default headers values
func NewBindHostNotImplemented() *BindHostNotImplemented {
	return &BindHostNotImplemented{}
}

/*BindHostNotImplemented handles this case with default header values.

Not implemented.
*/
type BindHostNotImplemented struct {
	Payload *models.Error
}

func (o *BindHostNotImplemented) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostNotImplemented  %+v", 501, o.Payload)
}

func (o *BindHostNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindHostNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindHostServiceUnavailable creates a BindHostServiceUnavailable with default headers values
func NewBindHostServiceUnavailable() *BindHostServiceUnavailable {
	return &BindHostServiceUnavailable{}
}

/*BindHostServiceUnavailable handles this case with default header values.

Unavailable.
*/
type BindHostServiceUnavailable struct {
	Payload *models.Error
}

func (o *BindHostServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/bind][%d] bindHostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *BindHostServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *BindHostServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
