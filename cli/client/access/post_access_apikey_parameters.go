// Code generated by go-swagger; DO NOT EDIT.

package access

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

	"github.com/loopholelabs/scale/cli/client/models"
)

// NewPostAccessApikeyParams creates a new PostAccessApikeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAccessApikeyParams() *PostAccessApikeyParams {
	return &PostAccessApikeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccessApikeyParamsWithTimeout creates a new PostAccessApikeyParams object
// with the ability to set a timeout on a request.
func NewPostAccessApikeyParamsWithTimeout(timeout time.Duration) *PostAccessApikeyParams {
	return &PostAccessApikeyParams{
		timeout: timeout,
	}
}

// NewPostAccessApikeyParamsWithContext creates a new PostAccessApikeyParams object
// with the ability to set a context for a request.
func NewPostAccessApikeyParamsWithContext(ctx context.Context) *PostAccessApikeyParams {
	return &PostAccessApikeyParams{
		Context: ctx,
	}
}

// NewPostAccessApikeyParamsWithHTTPClient creates a new PostAccessApikeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAccessApikeyParamsWithHTTPClient(client *http.Client) *PostAccessApikeyParams {
	return &PostAccessApikeyParams{
		HTTPClient: client,
	}
}

/*
PostAccessApikeyParams contains all the parameters to send to the API endpoint

	for the post access apikey operation.

	Typically these are written to a http.Request.
*/
type PostAccessApikeyParams struct {

	/* Request.

	   Create API Key Request
	*/
	Request *models.ModelsCreateAPIKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post access apikey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccessApikeyParams) WithDefaults() *PostAccessApikeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post access apikey params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAccessApikeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post access apikey params
func (o *PostAccessApikeyParams) WithTimeout(timeout time.Duration) *PostAccessApikeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post access apikey params
func (o *PostAccessApikeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post access apikey params
func (o *PostAccessApikeyParams) WithContext(ctx context.Context) *PostAccessApikeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post access apikey params
func (o *PostAccessApikeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post access apikey params
func (o *PostAccessApikeyParams) WithHTTPClient(client *http.Client) *PostAccessApikeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post access apikey params
func (o *PostAccessApikeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post access apikey params
func (o *PostAccessApikeyParams) WithRequest(request *models.ModelsCreateAPIKeyRequest) *PostAccessApikeyParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post access apikey params
func (o *PostAccessApikeyParams) SetRequest(request *models.ModelsCreateAPIKeyRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccessApikeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
