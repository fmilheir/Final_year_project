// Code generated by go-swagger; DO NOT EDIT.

package diagnose_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ListDiagnoseIncidentURL generates an URL for the list diagnose incident operation
type ListDiagnoseIncidentURL struct {
	Fields *string
	Limit  *int64
	Offset *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListDiagnoseIncidentURL) WithBasePath(bp string) *ListDiagnoseIncidentURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListDiagnoseIncidentURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListDiagnoseIncidentURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/diagnoseIncident"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/tmf-api/Incident/v4/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var fieldsQ string
	if o.Fields != nil {
		fieldsQ = *o.Fields
	}
	if fieldsQ != "" {
		qs.Set("fields", fieldsQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt64(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListDiagnoseIncidentURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListDiagnoseIncidentURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListDiagnoseIncidentURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListDiagnoseIncidentURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListDiagnoseIncidentURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListDiagnoseIncidentURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
