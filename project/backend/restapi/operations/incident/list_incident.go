// Code generated by go-swagger; DO NOT EDIT.

package incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListIncidentHandlerFunc turns a function with the right signature into a list incident handler
type ListIncidentHandlerFunc func(ListIncidentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListIncidentHandlerFunc) Handle(params ListIncidentParams) middleware.Responder {
	return fn(params)
}

// ListIncidentHandler interface for that can handle valid list incident params
type ListIncidentHandler interface {
	Handle(ListIncidentParams) middleware.Responder
}

// NewListIncident creates a new http.Handler for the list incident operation
func NewListIncident(ctx *middleware.Context, handler ListIncidentHandler) *ListIncident {
	return &ListIncident{Context: ctx, Handler: handler}
}

/*
	ListIncident swagger:route GET /incident incident listIncident

# List or find Incident objects

This operation list or find Incident entities
*/
type ListIncident struct {
	Context *middleware.Context
	Handler ListIncidentHandler
}

func (o *ListIncident) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListIncidentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
