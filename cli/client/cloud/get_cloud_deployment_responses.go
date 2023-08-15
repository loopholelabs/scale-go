// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/cli/client/models"
)

// GetCloudDeploymentReader is a Reader for the GetCloudDeployment structure.
type GetCloudDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCloudDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud/deployment] GetCloudDeployment", response, response.Code())
	}
}

// NewGetCloudDeploymentOK creates a GetCloudDeploymentOK with default headers values
func NewGetCloudDeploymentOK() *GetCloudDeploymentOK {
	return &GetCloudDeploymentOK{}
}

/*
GetCloudDeploymentOK describes a response with status code 200, with default header values.

OK
*/
type GetCloudDeploymentOK struct {
	Payload []*models.ModelsDeploymentResponse
}

// IsSuccess returns true when this get cloud deployment o k response has a 2xx status code
func (o *GetCloudDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud deployment o k response has a 3xx status code
func (o *GetCloudDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud deployment o k response has a 4xx status code
func (o *GetCloudDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud deployment o k response has a 5xx status code
func (o *GetCloudDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud deployment o k response a status code equal to that given
func (o *GetCloudDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cloud deployment o k response
func (o *GetCloudDeploymentOK) Code() int {
	return 200
}

func (o *GetCloudDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /cloud/deployment][%d] getCloudDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetCloudDeploymentOK) String() string {
	return fmt.Sprintf("[GET /cloud/deployment][%d] getCloudDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetCloudDeploymentOK) GetPayload() []*models.ModelsDeploymentResponse {
	return o.Payload
}

func (o *GetCloudDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudDeploymentUnauthorized creates a GetCloudDeploymentUnauthorized with default headers values
func NewGetCloudDeploymentUnauthorized() *GetCloudDeploymentUnauthorized {
	return &GetCloudDeploymentUnauthorized{}
}

/*
GetCloudDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCloudDeploymentUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get cloud deployment unauthorized response has a 2xx status code
func (o *GetCloudDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud deployment unauthorized response has a 3xx status code
func (o *GetCloudDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud deployment unauthorized response has a 4xx status code
func (o *GetCloudDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud deployment unauthorized response has a 5xx status code
func (o *GetCloudDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud deployment unauthorized response a status code equal to that given
func (o *GetCloudDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cloud deployment unauthorized response
func (o *GetCloudDeploymentUnauthorized) Code() int {
	return 401
}

func (o *GetCloudDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloud/deployment][%d] getCloudDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCloudDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[GET /cloud/deployment][%d] getCloudDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCloudDeploymentUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetCloudDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudDeploymentInternalServerError creates a GetCloudDeploymentInternalServerError with default headers values
func NewGetCloudDeploymentInternalServerError() *GetCloudDeploymentInternalServerError {
	return &GetCloudDeploymentInternalServerError{}
}

/*
GetCloudDeploymentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCloudDeploymentInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get cloud deployment internal server error response has a 2xx status code
func (o *GetCloudDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud deployment internal server error response has a 3xx status code
func (o *GetCloudDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud deployment internal server error response has a 4xx status code
func (o *GetCloudDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud deployment internal server error response has a 5xx status code
func (o *GetCloudDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cloud deployment internal server error response a status code equal to that given
func (o *GetCloudDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cloud deployment internal server error response
func (o *GetCloudDeploymentInternalServerError) Code() int {
	return 500
}

func (o *GetCloudDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud/deployment][%d] getCloudDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCloudDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud/deployment][%d] getCloudDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCloudDeploymentInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetCloudDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
