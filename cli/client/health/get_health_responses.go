// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/cli/client/models"
)

// GetHealthReader is a Reader for the GetHealth structure.
type GetHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /health] GetHealth", response, response.Code())
	}
}

// NewGetHealthOK creates a GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {
	return &GetHealthOK{}
}

/*
GetHealthOK describes a response with status code 200, with default header values.

OK
*/
type GetHealthOK struct {
	Payload *models.ModelsHealthResponse
}

// IsSuccess returns true when this get health o k response has a 2xx status code
func (o *GetHealthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get health o k response has a 3xx status code
func (o *GetHealthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health o k response has a 4xx status code
func (o *GetHealthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health o k response has a 5xx status code
func (o *GetHealthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get health o k response a status code equal to that given
func (o *GetHealthOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get health o k response
func (o *GetHealthOK) Code() int {
	return 200
}

func (o *GetHealthOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthOK  %+v", 200, o.Payload)
}

func (o *GetHealthOK) String() string {
	return fmt.Sprintf("[GET /health][%d] getHealthOK  %+v", 200, o.Payload)
}

func (o *GetHealthOK) GetPayload() *models.ModelsHealthResponse {
	return o.Payload
}

func (o *GetHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsHealthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthInternalServerError creates a GetHealthInternalServerError with default headers values
func NewGetHealthInternalServerError() *GetHealthInternalServerError {
	return &GetHealthInternalServerError{}
}

/*
GetHealthInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHealthInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get health internal server error response has a 2xx status code
func (o *GetHealthInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health internal server error response has a 3xx status code
func (o *GetHealthInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health internal server error response has a 4xx status code
func (o *GetHealthInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health internal server error response has a 5xx status code
func (o *GetHealthInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get health internal server error response a status code equal to that given
func (o *GetHealthInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get health internal server error response
func (o *GetHealthInternalServerError) Code() int {
	return 500
}

func (o *GetHealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /health][%d] getHealthInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHealthInternalServerError) String() string {
	return fmt.Sprintf("[GET /health][%d] getHealthInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHealthInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
