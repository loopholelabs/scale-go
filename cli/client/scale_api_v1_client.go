// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/cli/client/access"
	"github.com/loopholelabs/scale/cli/client/cloud"
	"github.com/loopholelabs/scale/cli/client/health"
	"github.com/loopholelabs/scale/cli/client/registry"
	"github.com/loopholelabs/scale/cli/client/userinfo"
)

// Default scale API v1 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.scale.sh"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new scale API v1 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ScaleAPIV1 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new scale API v1 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ScaleAPIV1 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new scale API v1 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ScaleAPIV1 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ScaleAPIV1)
	cli.Transport = transport
	cli.Access = access.New(transport, formats)
	cli.Cloud = cloud.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.Registry = registry.New(transport, formats)
	cli.Userinfo = userinfo.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// ScaleAPIV1 is a client for scale API v1
type ScaleAPIV1 struct {
	Access access.ClientService

	Cloud cloud.ClientService

	Health health.ClientService

	Registry registry.ClientService

	Userinfo userinfo.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ScaleAPIV1) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Access.SetTransport(transport)
	c.Cloud.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.Registry.SetTransport(transport)
	c.Userinfo.SetTransport(transport)
}
