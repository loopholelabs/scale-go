// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

	"github.com/loopholelabs/scale/client/models"
)

// NewPostCloudDomainParams creates a new PostCloudDomainParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCloudDomainParams() *PostCloudDomainParams {
	return &PostCloudDomainParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCloudDomainParamsWithTimeout creates a new PostCloudDomainParams object
// with the ability to set a timeout on a request.
func NewPostCloudDomainParamsWithTimeout(timeout time.Duration) *PostCloudDomainParams {
	return &PostCloudDomainParams{
		timeout: timeout,
	}
}

// NewPostCloudDomainParamsWithContext creates a new PostCloudDomainParams object
// with the ability to set a context for a request.
func NewPostCloudDomainParamsWithContext(ctx context.Context) *PostCloudDomainParams {
	return &PostCloudDomainParams{
		Context: ctx,
	}
}

// NewPostCloudDomainParamsWithHTTPClient creates a new PostCloudDomainParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCloudDomainParamsWithHTTPClient(client *http.Client) *PostCloudDomainParams {
	return &PostCloudDomainParams{
		HTTPClient: client,
	}
}

/*
PostCloudDomainParams contains all the parameters to send to the API endpoint

	for the post cloud domain operation.

	Typically these are written to a http.Request.
*/
type PostCloudDomainParams struct {

	/* Request.

	   Create Domain Request
	*/
	Request *models.ModelsCreateDomainRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post cloud domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCloudDomainParams) WithDefaults() *PostCloudDomainParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post cloud domain params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCloudDomainParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post cloud domain params
func (o *PostCloudDomainParams) WithTimeout(timeout time.Duration) *PostCloudDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cloud domain params
func (o *PostCloudDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cloud domain params
func (o *PostCloudDomainParams) WithContext(ctx context.Context) *PostCloudDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cloud domain params
func (o *PostCloudDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cloud domain params
func (o *PostCloudDomainParams) WithHTTPClient(client *http.Client) *PostCloudDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cloud domain params
func (o *PostCloudDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post cloud domain params
func (o *PostCloudDomainParams) WithRequest(request *models.ModelsCreateDomainRequest) *PostCloudDomainParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post cloud domain params
func (o *PostCloudDomainParams) SetRequest(request *models.ModelsCreateDomainRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostCloudDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
