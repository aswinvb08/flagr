// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteFlagHandlerFunc turns a function with the right signature into a delete flag handler
type DeleteFlagHandlerFunc func(DeleteFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFlagHandlerFunc) Handle(params DeleteFlagParams) middleware.Responder {
	return fn(params)
}

// DeleteFlagHandler interface for that can handle valid delete flag params
type DeleteFlagHandler interface {
	Handle(DeleteFlagParams) middleware.Responder
}

// NewDeleteFlag creates a new http.Handler for the delete flag operation
func NewDeleteFlag(ctx *middleware.Context, handler DeleteFlagHandler) *DeleteFlag {
	return &DeleteFlag{Context: ctx, Handler: handler}
}

/*
	DeleteFlag swagger:route DELETE /flags/{flagID} flag deleteFlag

DeleteFlag delete flag API
*/
type DeleteFlag struct {
	Context *middleware.Context
	Handler DeleteFlagHandler
}

func (o *DeleteFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteFlagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
