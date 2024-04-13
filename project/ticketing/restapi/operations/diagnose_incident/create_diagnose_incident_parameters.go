// Code generated by go-swagger; DO NOT EDIT.

package diagnose_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"fmilheir/test_swagger/models"
)

// NewCreateDiagnoseIncidentParams creates a new CreateDiagnoseIncidentParams object
//
// There are no default values defined in the spec.
func NewCreateDiagnoseIncidentParams() CreateDiagnoseIncidentParams {

	return CreateDiagnoseIncidentParams{}
}

// CreateDiagnoseIncidentParams contains all the bound params for the create diagnose incident operation
// typically these are obtained from a http.Request
//
// swagger:parameters createDiagnoseIncident
type CreateDiagnoseIncidentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The DiagnoseIncident to be created
	  Required: true
	  In: body
	*/
	DiagnoseIncident *models.DiagnoseIncidentCreate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateDiagnoseIncidentParams() beforehand.
func (o *CreateDiagnoseIncidentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DiagnoseIncidentCreate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("diagnoseIncident", "body", ""))
			} else {
				res = append(res, errors.NewParseError("diagnoseIncident", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.DiagnoseIncident = &body
			}
		}
	} else {
		res = append(res, errors.Required("diagnoseIncident", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}