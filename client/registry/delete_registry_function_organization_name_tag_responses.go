// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRegistryFunctionOrganizationNameTagReader is a Reader for the DeleteRegistryFunctionOrganizationNameTag structure.
type DeleteRegistryFunctionOrganizationNameTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistryFunctionOrganizationNameTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistryFunctionOrganizationNameTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRegistryFunctionOrganizationNameTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRegistryFunctionOrganizationNameTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegistryFunctionOrganizationNameTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRegistryFunctionOrganizationNameTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRegistryFunctionOrganizationNameTagOK creates a DeleteRegistryFunctionOrganizationNameTagOK with default headers values
func NewDeleteRegistryFunctionOrganizationNameTagOK() *DeleteRegistryFunctionOrganizationNameTagOK {
	return &DeleteRegistryFunctionOrganizationNameTagOK{}
}

/*
DeleteRegistryFunctionOrganizationNameTagOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRegistryFunctionOrganizationNameTagOK struct {
	Payload string
}

// IsSuccess returns true when this delete registry function organization name tag o k response has a 2xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete registry function organization name tag o k response has a 3xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry function organization name tag o k response has a 4xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry function organization name tag o k response has a 5xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry function organization name tag o k response a status code equal to that given
func (o *DeleteRegistryFunctionOrganizationNameTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete registry function organization name tag o k response
func (o *DeleteRegistryFunctionOrganizationNameTagOK) Code() int {
	return 200
}

func (o *DeleteRegistryFunctionOrganizationNameTagOK) Error() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagOK) String() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistryFunctionOrganizationNameTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryFunctionOrganizationNameTagBadRequest creates a DeleteRegistryFunctionOrganizationNameTagBadRequest with default headers values
func NewDeleteRegistryFunctionOrganizationNameTagBadRequest() *DeleteRegistryFunctionOrganizationNameTagBadRequest {
	return &DeleteRegistryFunctionOrganizationNameTagBadRequest{}
}

/*
DeleteRegistryFunctionOrganizationNameTagBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteRegistryFunctionOrganizationNameTagBadRequest struct {
	Payload string
}

// IsSuccess returns true when this delete registry function organization name tag bad request response has a 2xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry function organization name tag bad request response has a 3xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry function organization name tag bad request response has a 4xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry function organization name tag bad request response has a 5xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry function organization name tag bad request response a status code equal to that given
func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete registry function organization name tag bad request response
func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) Code() int {
	return 400
}

func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) String() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistryFunctionOrganizationNameTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryFunctionOrganizationNameTagUnauthorized creates a DeleteRegistryFunctionOrganizationNameTagUnauthorized with default headers values
func NewDeleteRegistryFunctionOrganizationNameTagUnauthorized() *DeleteRegistryFunctionOrganizationNameTagUnauthorized {
	return &DeleteRegistryFunctionOrganizationNameTagUnauthorized{}
}

/*
DeleteRegistryFunctionOrganizationNameTagUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRegistryFunctionOrganizationNameTagUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this delete registry function organization name tag unauthorized response has a 2xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry function organization name tag unauthorized response has a 3xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry function organization name tag unauthorized response has a 4xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry function organization name tag unauthorized response has a 5xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry function organization name tag unauthorized response a status code equal to that given
func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete registry function organization name tag unauthorized response
func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) Code() int {
	return 401
}

func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistryFunctionOrganizationNameTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryFunctionOrganizationNameTagNotFound creates a DeleteRegistryFunctionOrganizationNameTagNotFound with default headers values
func NewDeleteRegistryFunctionOrganizationNameTagNotFound() *DeleteRegistryFunctionOrganizationNameTagNotFound {
	return &DeleteRegistryFunctionOrganizationNameTagNotFound{}
}

/*
DeleteRegistryFunctionOrganizationNameTagNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRegistryFunctionOrganizationNameTagNotFound struct {
	Payload string
}

// IsSuccess returns true when this delete registry function organization name tag not found response has a 2xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry function organization name tag not found response has a 3xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry function organization name tag not found response has a 4xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry function organization name tag not found response has a 5xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry function organization name tag not found response a status code equal to that given
func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete registry function organization name tag not found response
func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) Code() int {
	return 404
}

func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) String() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistryFunctionOrganizationNameTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryFunctionOrganizationNameTagInternalServerError creates a DeleteRegistryFunctionOrganizationNameTagInternalServerError with default headers values
func NewDeleteRegistryFunctionOrganizationNameTagInternalServerError() *DeleteRegistryFunctionOrganizationNameTagInternalServerError {
	return &DeleteRegistryFunctionOrganizationNameTagInternalServerError{}
}

/*
DeleteRegistryFunctionOrganizationNameTagInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteRegistryFunctionOrganizationNameTagInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this delete registry function organization name tag internal server error response has a 2xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry function organization name tag internal server error response has a 3xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry function organization name tag internal server error response has a 4xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry function organization name tag internal server error response has a 5xx status code
func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete registry function organization name tag internal server error response a status code equal to that given
func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete registry function organization name tag internal server error response
func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) Code() int {
	return 500
}

func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /registry/function/{organization}/{name}/{tag}][%d] deleteRegistryFunctionOrganizationNameTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *DeleteRegistryFunctionOrganizationNameTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
