// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindSegmentsHandlerFunc turns a function with the right signature into a find segments handler
type FindSegmentsHandlerFunc func(FindSegmentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindSegmentsHandlerFunc) Handle(params FindSegmentsParams) middleware.Responder {
	return fn(params)
}

// FindSegmentsHandler interface for that can handle valid find segments params
type FindSegmentsHandler interface {
	Handle(FindSegmentsParams) middleware.Responder
}

// NewFindSegments creates a new http.Handler for the find segments operation
func NewFindSegments(ctx *middleware.Context, handler FindSegmentsHandler) *FindSegments {
	return &FindSegments{Context: ctx, Handler: handler}
}

/*
	FindSegments swagger:route GET /flags/{flagID}/segments segment findSegments

FindSegments find segments API
*/
type FindSegments struct {
	Context *middleware.Context
	Handler FindSegmentsHandler
}

func (o *FindSegments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindSegmentsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
