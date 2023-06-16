// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/client/models"
)

// PostAccessApikeyNameReader is a Reader for the PostAccessApikeyName structure.
type PostAccessApikeyNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAccessApikeyNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAccessApikeyNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAccessApikeyNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAccessApikeyNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAccessApikeyNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAccessApikeyNameOK creates a PostAccessApikeyNameOK with default headers values
func NewPostAccessApikeyNameOK() *PostAccessApikeyNameOK {
	return &PostAccessApikeyNameOK{}
}

/*
PostAccessApikeyNameOK describes a response with status code 200, with default header values.

OK
*/
type PostAccessApikeyNameOK struct {
	Payload *models.ModelsCreateAPIKeyResponse
}

// IsSuccess returns true when this post access apikey name o k response has a 2xx status code
func (o *PostAccessApikeyNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post access apikey name o k response has a 3xx status code
func (o *PostAccessApikeyNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey name o k response has a 4xx status code
func (o *PostAccessApikeyNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post access apikey name o k response has a 5xx status code
func (o *PostAccessApikeyNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post access apikey name o k response a status code equal to that given
func (o *PostAccessApikeyNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post access apikey name o k response
func (o *PostAccessApikeyNameOK) Code() int {
	return 200
}

func (o *PostAccessApikeyNameOK) Error() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameOK  %+v", 200, o.Payload)
}

func (o *PostAccessApikeyNameOK) String() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameOK  %+v", 200, o.Payload)
}

func (o *PostAccessApikeyNameOK) GetPayload() *models.ModelsCreateAPIKeyResponse {
	return o.Payload
}

func (o *PostAccessApikeyNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCreateAPIKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessApikeyNameBadRequest creates a PostAccessApikeyNameBadRequest with default headers values
func NewPostAccessApikeyNameBadRequest() *PostAccessApikeyNameBadRequest {
	return &PostAccessApikeyNameBadRequest{}
}

/*
PostAccessApikeyNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostAccessApikeyNameBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post access apikey name bad request response has a 2xx status code
func (o *PostAccessApikeyNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access apikey name bad request response has a 3xx status code
func (o *PostAccessApikeyNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey name bad request response has a 4xx status code
func (o *PostAccessApikeyNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post access apikey name bad request response has a 5xx status code
func (o *PostAccessApikeyNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post access apikey name bad request response a status code equal to that given
func (o *PostAccessApikeyNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post access apikey name bad request response
func (o *PostAccessApikeyNameBadRequest) Code() int {
	return 400
}

func (o *PostAccessApikeyNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccessApikeyNameBadRequest) String() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccessApikeyNameBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostAccessApikeyNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessApikeyNameUnauthorized creates a PostAccessApikeyNameUnauthorized with default headers values
func NewPostAccessApikeyNameUnauthorized() *PostAccessApikeyNameUnauthorized {
	return &PostAccessApikeyNameUnauthorized{}
}

/*
PostAccessApikeyNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostAccessApikeyNameUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this post access apikey name unauthorized response has a 2xx status code
func (o *PostAccessApikeyNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access apikey name unauthorized response has a 3xx status code
func (o *PostAccessApikeyNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey name unauthorized response has a 4xx status code
func (o *PostAccessApikeyNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post access apikey name unauthorized response has a 5xx status code
func (o *PostAccessApikeyNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post access apikey name unauthorized response a status code equal to that given
func (o *PostAccessApikeyNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post access apikey name unauthorized response
func (o *PostAccessApikeyNameUnauthorized) Code() int {
	return 401
}

func (o *PostAccessApikeyNameUnauthorized) Error() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccessApikeyNameUnauthorized) String() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccessApikeyNameUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *PostAccessApikeyNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccessApikeyNameInternalServerError creates a PostAccessApikeyNameInternalServerError with default headers values
func NewPostAccessApikeyNameInternalServerError() *PostAccessApikeyNameInternalServerError {
	return &PostAccessApikeyNameInternalServerError{}
}

/*
PostAccessApikeyNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAccessApikeyNameInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post access apikey name internal server error response has a 2xx status code
func (o *PostAccessApikeyNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post access apikey name internal server error response has a 3xx status code
func (o *PostAccessApikeyNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post access apikey name internal server error response has a 4xx status code
func (o *PostAccessApikeyNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post access apikey name internal server error response has a 5xx status code
func (o *PostAccessApikeyNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post access apikey name internal server error response a status code equal to that given
func (o *PostAccessApikeyNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post access apikey name internal server error response
func (o *PostAccessApikeyNameInternalServerError) Code() int {
	return 500
}

func (o *PostAccessApikeyNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAccessApikeyNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /access/apikey/{name}][%d] postAccessApikeyNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAccessApikeyNameInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostAccessApikeyNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
