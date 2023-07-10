// Code generated by go-swagger; DO NOT EDIT.

package constraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// PutConstraintURL generates an URL for the put constraint operation
type PutConstraintURL struct {
	ConstraintID int64
	FlagID       int64
	SegmentID    int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutConstraintURL) WithBasePath(bp string) *PutConstraintURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutConstraintURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PutConstraintURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/flags/{flagID}/segments/{segmentID}/constraints/{constraintID}"

	constraintID := swag.FormatInt64(o.ConstraintID)
	if constraintID != "" {
		_path = strings.Replace(_path, "{constraintID}", constraintID, -1)
	} else {
		return nil, errors.New("constraintId is required on PutConstraintURL")
	}

	flagID := swag.FormatInt64(o.FlagID)
	if flagID != "" {
		_path = strings.Replace(_path, "{flagID}", flagID, -1)
	} else {
		return nil, errors.New("flagId is required on PutConstraintURL")
	}

	segmentID := swag.FormatInt64(o.SegmentID)
	if segmentID != "" {
		_path = strings.Replace(_path, "{segmentID}", segmentID, -1)
	} else {
		return nil, errors.New("segmentId is required on PutConstraintURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PutConstraintURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PutConstraintURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PutConstraintURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PutConstraintURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PutConstraintURL")
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
func (o *PutConstraintURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}