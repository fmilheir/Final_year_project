// Code generated by go-swagger; DO NOT EDIT.

package diagnose_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrieveDiagnoseIncidentHandlerFunc turns a function with the right signature into a retrieve diagnose incident handler
type RetrieveDiagnoseIncidentHandlerFunc func(RetrieveDiagnoseIncidentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveDiagnoseIncidentHandlerFunc) Handle(params RetrieveDiagnoseIncidentParams) middleware.Responder {
	return fn(params)
}

// RetrieveDiagnoseIncidentHandler interface for that can handle valid retrieve diagnose incident params
type RetrieveDiagnoseIncidentHandler interface {
	Handle(RetrieveDiagnoseIncidentParams) middleware.Responder
}

// NewRetrieveDiagnoseIncident creates a new http.Handler for the retrieve diagnose incident operation
func NewRetrieveDiagnoseIncident(ctx *middleware.Context, handler RetrieveDiagnoseIncidentHandler) *RetrieveDiagnoseIncident {
	return &RetrieveDiagnoseIncident{Context: ctx, Handler: handler}
}

/*
	RetrieveDiagnoseIncident swagger:route GET /diagnoseIncident/{id} diagnoseIncident retrieveDiagnoseIncident

# Retrieves a DiagnoseIncident by ID

This operation retrieves a DiagnoseIncident entity. Attribute selection is enabled for all first level attributes.
*/
type RetrieveDiagnoseIncident struct {
	Context *middleware.Context
	Handler RetrieveDiagnoseIncidentHandler
}

func (o *RetrieveDiagnoseIncident) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRetrieveDiagnoseIncidentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
