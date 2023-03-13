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
)

// NewGetAccessApikeyNameoridParams creates a new GetAccessApikeyNameoridParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccessApikeyNameoridParams() *GetAccessApikeyNameoridParams {
	return &GetAccessApikeyNameoridParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessApikeyNameoridParamsWithTimeout creates a new GetAccessApikeyNameoridParams object
// with the ability to set a timeout on a request.
func NewGetAccessApikeyNameoridParamsWithTimeout(timeout time.Duration) *GetAccessApikeyNameoridParams {
	return &GetAccessApikeyNameoridParams{
		timeout: timeout,
	}
}

// NewGetAccessApikeyNameoridParamsWithContext creates a new GetAccessApikeyNameoridParams object
// with the ability to set a context for a request.
func NewGetAccessApikeyNameoridParamsWithContext(ctx context.Context) *GetAccessApikeyNameoridParams {
	return &GetAccessApikeyNameoridParams{
		Context: ctx,
	}
}

// NewGetAccessApikeyNameoridParamsWithHTTPClient creates a new GetAccessApikeyNameoridParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccessApikeyNameoridParamsWithHTTPClient(client *http.Client) *GetAccessApikeyNameoridParams {
	return &GetAccessApikeyNameoridParams{
		HTTPClient: client,
	}
}

/*
GetAccessApikeyNameoridParams contains all the parameters to send to the API endpoint

	for the get access apikey nameorid operation.

	Typically these are written to a http.Request.
*/
type GetAccessApikeyNameoridParams struct {

	/* Nameorid.

	   API Key Name or ID
	*/
	Nameorid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get access apikey nameorid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessApikeyNameoridParams) WithDefaults() *GetAccessApikeyNameoridParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get access apikey nameorid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccessApikeyNameoridParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) WithTimeout(timeout time.Duration) *GetAccessApikeyNameoridParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) WithContext(ctx context.Context) *GetAccessApikeyNameoridParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) WithHTTPClient(client *http.Client) *GetAccessApikeyNameoridParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNameorid adds the nameorid to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) WithNameorid(nameorid string) *GetAccessApikeyNameoridParams {
	o.SetNameorid(nameorid)
	return o
}

// SetNameorid adds the nameorid to the get access apikey nameorid params
func (o *GetAccessApikeyNameoridParams) SetNameorid(nameorid string) {
	o.Nameorid = nameorid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessApikeyNameoridParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nameorid
	if err := r.SetPathParam("nameorid", o.Nameorid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}