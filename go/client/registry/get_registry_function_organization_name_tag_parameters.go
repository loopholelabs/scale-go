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

// NewGetRegistryFunctionOrganizationNameTagParams creates a new GetRegistryFunctionOrganizationNameTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistryFunctionOrganizationNameTagParams() *GetRegistryFunctionOrganizationNameTagParams {
	return &GetRegistryFunctionOrganizationNameTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistryFunctionOrganizationNameTagParamsWithTimeout creates a new GetRegistryFunctionOrganizationNameTagParams object
// with the ability to set a timeout on a request.
func NewGetRegistryFunctionOrganizationNameTagParamsWithTimeout(timeout time.Duration) *GetRegistryFunctionOrganizationNameTagParams {
	return &GetRegistryFunctionOrganizationNameTagParams{
		timeout: timeout,
	}
}

// NewGetRegistryFunctionOrganizationNameTagParamsWithContext creates a new GetRegistryFunctionOrganizationNameTagParams object
// with the ability to set a context for a request.
func NewGetRegistryFunctionOrganizationNameTagParamsWithContext(ctx context.Context) *GetRegistryFunctionOrganizationNameTagParams {
	return &GetRegistryFunctionOrganizationNameTagParams{
		Context: ctx,
	}
}

// NewGetRegistryFunctionOrganizationNameTagParamsWithHTTPClient creates a new GetRegistryFunctionOrganizationNameTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistryFunctionOrganizationNameTagParamsWithHTTPClient(client *http.Client) *GetRegistryFunctionOrganizationNameTagParams {
	return &GetRegistryFunctionOrganizationNameTagParams{
		HTTPClient: client,
	}
}

/*
GetRegistryFunctionOrganizationNameTagParams contains all the parameters to send to the API endpoint

	for the get registry function organization name tag operation.

	Typically these are written to a http.Request.
*/
type GetRegistryFunctionOrganizationNameTagParams struct {

	/* Name.

	   name
	*/
	Name string

	/* Organization.

	   organization
	*/
	Organization string

	/* Tag.

	   tag
	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registry function organization name tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryFunctionOrganizationNameTagParams) WithDefaults() *GetRegistryFunctionOrganizationNameTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registry function organization name tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistryFunctionOrganizationNameTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) WithTimeout(timeout time.Duration) *GetRegistryFunctionOrganizationNameTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) WithContext(ctx context.Context) *GetRegistryFunctionOrganizationNameTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) WithHTTPClient(client *http.Client) *GetRegistryFunctionOrganizationNameTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) WithName(name string) *GetRegistryFunctionOrganizationNameTagParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) WithOrganization(organization string) *GetRegistryFunctionOrganizationNameTagParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithTag adds the tag to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) WithTag(tag string) *GetRegistryFunctionOrganizationNameTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get registry function organization name tag params
func (o *GetRegistryFunctionOrganizationNameTagParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistryFunctionOrganizationNameTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
