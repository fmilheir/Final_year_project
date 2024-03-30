// Code generated by go-swagger; DO NOT EDIT.

package events_subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UnregisterListenerHandlerFunc turns a function with the right signature into a unregister listener handler
type UnregisterListenerHandlerFunc func(UnregisterListenerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UnregisterListenerHandlerFunc) Handle(params UnregisterListenerParams) middleware.Responder {
	return fn(params)
}

// UnregisterListenerHandler interface for that can handle valid unregister listener params
type UnregisterListenerHandler interface {
	Handle(UnregisterListenerParams) middleware.Responder
}

// NewUnregisterListener creates a new http.Handler for the unregister listener operation
func NewUnregisterListener(ctx *middleware.Context, handler UnregisterListenerHandler) *UnregisterListener {
	return &UnregisterListener{Context: ctx, Handler: handler}
}

/*
	UnregisterListener swagger:route DELETE /hub/{id} events subscription unregisterListener

# Unregister a listener

Resets the communication endpoint address the service instance must use to deliver information about its health state, execution state, failures and metrics.
*/
type UnregisterListener struct {
	Context *middleware.Context
	Handler UnregisterListenerHandler
}

func (o *UnregisterListener) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUnregisterListenerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
