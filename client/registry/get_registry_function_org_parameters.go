// Code generated by go-swagger; DO NOT EDIT.

package registry

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

// NewGetRegistryFunctionOrgParams creates a new GetRegistryFunctionOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistryFunctionOrgParams() *GetRegistryFunctionOrgParams {
	return &GetRegistryFunctionOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistryFunctionOrgParamsWithTimeout creates a new GetRegistryFunctionOrgParams object
// with the ability to set a timeout on a request.
func NewGetRegistryFunctionOrgParamsWithTimeout(timeout time.Duration) *GetRegistryFunctionOrgParams {
	return &GetRegistryFunctionOrgParams{
		timeout: timeout,
	}
}

// NewGetRegistryFunctionOrgParamsWithContext creates a new GetRegistryFunctionOrgParams object
// with the ability to set a context for a request.
func NewGetRegistryFunctionOrgParamsWithContext(ctx context.Context) *GetRegistryFunctionOrgParams {
	return &GetRegistryFunctionOrgParams{
		Context: ctx,
	}
}

// NewGetRegistryFunctionOrgParamsWithHTTPClient creates a new GetRegistryFunctionOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistryFunctionOrgParamsWithHTTPClient(client *http.Client) *GetRegistryFunctionOrgParams {
	return &GetRegistryFunctionOrgParams{
		HTTPClient: client,
	}
}

/*
GetRegistryFunctionOrgParams contains all the parameters to send to the API endpoint

	for the get registry function org operation.

	Typically these are written to a http.Request.
*/
type GetRegistryFunctionOrgParams struct {

	/* Org.

	   org
	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registry function org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryFunctionOrgParams) WithDefaults() *GetRegistryFunctionOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registry function org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryFunctionOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registry function org params
func (o *GetRegistryFunctionOrgParams) WithTimeout(timeout time.Duration) *GetRegistryFunctionOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registry function org params
func (o *GetRegistryFunctionOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registry function org params
func (o *GetRegistryFunctionOrgParams) WithContext(ctx context.Context) *GetRegistryFunctionOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registry function org params
func (o *GetRegistryFunctionOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registry function org params
func (o *GetRegistryFunctionOrgParams) WithHTTPClient(client *http.Client) *GetRegistryFunctionOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registry function org params
func (o *GetRegistryFunctionOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the get registry function org params
func (o *GetRegistryFunctionOrgParams) WithOrg(org string) *GetRegistryFunctionOrgParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the get registry function org params
func (o *GetRegistryFunctionOrgParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistryFunctionOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
