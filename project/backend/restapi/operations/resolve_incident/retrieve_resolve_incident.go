// Code generated by go-swagger; DO NOT EDIT.

package resolve_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrieveResolveIncidentHandlerFunc turns a function with the right signature into a retrieve resolve incident handler
type RetrieveResolveIncidentHandlerFunc func(RetrieveResolveIncidentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveResolveIncidentHandlerFunc) Handle(params RetrieveResolveIncidentParams) middleware.Responder {
	return fn(params)
}

// RetrieveResolveIncidentHandler interface for that can handle valid retrieve resolve incident params
type RetrieveResolveIncidentHandler interface {
	Handle(RetrieveResolveIncidentParams) middleware.Responder
}

// NewRetrieveResolveIncident creates a new http.Handler for the retrieve resolve incident operation
func NewRetrieveResolveIncident(ctx *middleware.Context, handler RetrieveResolveIncidentHandler) *RetrieveResolveIncident {
	return &RetrieveResolveIncident{Context: ctx, Handler: handler}
}

/*
	RetrieveResolveIncident swagger:route GET /resolveIncident/{id} resolveIncident retrieveResolveIncident

# Retrieves a ResolveIncident by ID

This operation retrieves a ResolveIncident entity. Attribute selection is enabled for all first level attributes.
*/
type RetrieveResolveIncident struct {
	Context *middleware.Context
	Handler RetrieveResolveIncidentHandler
}

func (o *RetrieveResolveIncident) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRetrieveResolveIncidentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
