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

// NewDeleteAccessApikeyNameoridParams creates a new DeleteAccessApikeyNameoridParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAccessApikeyNameoridParams() *DeleteAccessApikeyNameoridParams {
	return &DeleteAccessApikeyNameoridParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccessApikeyNameoridParamsWithTimeout creates a new DeleteAccessApikeyNameoridParams object
// with the ability to set a timeout on a request.
func NewDeleteAccessApikeyNameoridParamsWithTimeout(timeout time.Duration) *DeleteAccessApikeyNameoridParams {
	return &DeleteAccessApikeyNameoridParams{
		timeout: timeout,
	}
}

// NewDeleteAccessApikeyNameoridParamsWithContext creates a new DeleteAccessApikeyNameoridParams object
// with the ability to set a context for a request.
func NewDeleteAccessApikeyNameoridParamsWithContext(ctx context.Context) *DeleteAccessApikeyNameoridParams {
	return &DeleteAccessApikeyNameoridParams{
		Context: ctx,
	}
}

// NewDeleteAccessApikeyNameoridParamsWithHTTPClient creates a new DeleteAccessApikeyNameoridParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAccessApikeyNameoridParamsWithHTTPClient(client *http.Client) *DeleteAccessApikeyNameoridParams {
	return &DeleteAccessApikeyNameoridParams{
		HTTPClient: client,
	}
}

/*
DeleteAccessApikeyNameoridParams contains all the parameters to send to the API endpoint

	for the delete access apikey nameorid operation.

	Typically these are written to a http.Request.
*/
type DeleteAccessApikeyNameoridParams struct {

	/* Nameorid.

	   API Key Name or ID
	*/
	Nameorid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete access apikey nameorid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccessApikeyNameoridParams) WithDefaults() *DeleteAccessApikeyNameoridParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete access apikey nameorid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAccessApikeyNameoridParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) WithTimeout(timeout time.Duration) *DeleteAccessApikeyNameoridParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) WithContext(ctx context.Context) *DeleteAccessApikeyNameoridParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) WithHTTPClient(client *http.Client) *DeleteAccessApikeyNameoridParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNameorid adds the nameorid to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) WithNameorid(nameorid string) *DeleteAccessApikeyNameoridParams {
	o.SetNameorid(nameorid)
	return o
}

// SetNameorid adds the nameorid to the delete access apikey nameorid params
func (o *DeleteAccessApikeyNameoridParams) SetNameorid(nameorid string) {
	o.Nameorid = nameorid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccessApikeyNameoridParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
