// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateSegmentHandlerFunc turns a function with the right signature into a create segment handler
type CreateSegmentHandlerFunc func(CreateSegmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateSegmentHandlerFunc) Handle(params CreateSegmentParams) middleware.Responder {
	return fn(params)
}

// CreateSegmentHandler interface for that can handle valid create segment params
type CreateSegmentHandler interface {
	Handle(CreateSegmentParams) middleware.Responder
}

// NewCreateSegment creates a new http.Handler for the create segment operation
func NewCreateSegment(ctx *middleware.Context, handler CreateSegmentHandler) *CreateSegment {
	return &CreateSegment{Context: ctx, Handler: handler}
}

/*
	CreateSegment swagger:route POST /flags/{flagID}/segments segment createSegment

CreateSegment create segment API
*/
type CreateSegment struct {
	Context *middleware.Context
	Handler CreateSegmentHandler
}

func (o *CreateSegment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateSegmentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
