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

// V2UpdateClusterReader is a Reader for the V2UpdateCluster structure.
type V2UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2UpdateClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2UpdateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2UpdateClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UpdateClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UpdateClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2UpdateClusterMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2UpdateClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UpdateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UpdateClusterCreated creates a V2UpdateClusterCreated with default headers values
func NewV2UpdateClusterCreated() *V2UpdateClusterCreated {
	return &V2UpdateClusterCreated{}
}

/*V2UpdateClusterCreated handles this case with default header values.

Success.
*/
type V2UpdateClusterCreated struct {
	Payload *models.Cluster
}

func (o *V2UpdateClusterCreated) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterCreated  %+v", 201, o.Payload)
}

func (o *V2UpdateClusterCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *V2UpdateClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterBadRequest creates a V2UpdateClusterBadRequest with default headers values
func NewV2UpdateClusterBadRequest() *V2UpdateClusterBadRequest {
	return &V2UpdateClusterBadRequest{}
}

/*V2UpdateClusterBadRequest handles this case with default header values.

Error.
*/
type V2UpdateClusterBadRequest struct {
	Payload *models.Error
}

func (o *V2UpdateClusterBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterBadRequest  %+v", 400, o.Payload)
}

func (o *V2UpdateClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterUnauthorized creates a V2UpdateClusterUnauthorized with default headers values
func NewV2UpdateClusterUnauthorized() *V2UpdateClusterUnauthorized {
	return &V2UpdateClusterUnauthorized{}
}

/*V2UpdateClusterUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2UpdateClusterUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2UpdateClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UpdateClusterUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterForbidden creates a V2UpdateClusterForbidden with default headers values
func NewV2UpdateClusterForbidden() *V2UpdateClusterForbidden {
	return &V2UpdateClusterForbidden{}
}

/*V2UpdateClusterForbidden handles this case with default header values.

Forbidden.
*/
type V2UpdateClusterForbidden struct {
	Payload *models.InfraError
}

func (o *V2UpdateClusterForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterForbidden  %+v", 403, o.Payload)
}

func (o *V2UpdateClusterForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterNotFound creates a V2UpdateClusterNotFound with default headers values
func NewV2UpdateClusterNotFound() *V2UpdateClusterNotFound {
	return &V2UpdateClusterNotFound{}
}

/*V2UpdateClusterNotFound handles this case with default header values.

Error.
*/
type V2UpdateClusterNotFound struct {
	Payload *models.Error
}

func (o *V2UpdateClusterNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterNotFound  %+v", 404, o.Payload)
}

func (o *V2UpdateClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterMethodNotAllowed creates a V2UpdateClusterMethodNotAllowed with default headers values
func NewV2UpdateClusterMethodNotAllowed() *V2UpdateClusterMethodNotAllowed {
	return &V2UpdateClusterMethodNotAllowed{}
}

/*V2UpdateClusterMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2UpdateClusterMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2UpdateClusterMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2UpdateClusterMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterConflict creates a V2UpdateClusterConflict with default headers values
func NewV2UpdateClusterConflict() *V2UpdateClusterConflict {
	return &V2UpdateClusterConflict{}
}

/*V2UpdateClusterConflict handles this case with default header values.

Error.
*/
type V2UpdateClusterConflict struct {
	Payload *models.Error
}

func (o *V2UpdateClusterConflict) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterConflict  %+v", 409, o.Payload)
}

func (o *V2UpdateClusterConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterInternalServerError creates a V2UpdateClusterInternalServerError with default headers values
func NewV2UpdateClusterInternalServerError() *V2UpdateClusterInternalServerError {
	return &V2UpdateClusterInternalServerError{}
}

/*V2UpdateClusterInternalServerError handles this case with default header values.

Error.
*/
type V2UpdateClusterInternalServerError struct {
	Payload *models.Error
}

func (o *V2UpdateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v2/clusters/{cluster_id}][%d] v2UpdateClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UpdateClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
