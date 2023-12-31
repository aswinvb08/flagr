// Code generated by go-swagger; DO NOT EDIT.

package variant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteVariantHandlerFunc turns a function with the right signature into a delete variant handler
type DeleteVariantHandlerFunc func(DeleteVariantParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteVariantHandlerFunc) Handle(params DeleteVariantParams) middleware.Responder {
	return fn(params)
}

// DeleteVariantHandler interface for that can handle valid delete variant params
type DeleteVariantHandler interface {
	Handle(DeleteVariantParams) middleware.Responder
}

// NewDeleteVariant creates a new http.Handler for the delete variant operation
func NewDeleteVariant(ctx *middleware.Context, handler DeleteVariantHandler) *DeleteVariant {
	return &DeleteVariant{Context: ctx, Handler: handler}
}

/*
	DeleteVariant swagger:route DELETE /flags/{flagID}/variants/{variantID} variant deleteVariant

DeleteVariant delete variant API
*/
type DeleteVariant struct {
	Context *middleware.Context
	Handler DeleteVariantHandler
}

func (o *DeleteVariant) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteVariantParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
