// Code generated by go-swagger; DO NOT EDIT.

package incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"fmilheir/test_swagger/models"
)

// ListIncidentOKCode is the HTTP code returned for type ListIncidentOK
const ListIncidentOKCode int = 200

/*
ListIncidentOK Success

swagger:response listIncidentOK
*/
type ListIncidentOK struct {
	/*Actual number of items returned in the response body

	 */
	XResultCount int64 `json:"X-Result-Count"`
	/*Total number of items matching criteria

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.Incident `json:"body,omitempty"`
}

// NewListIncidentOK creates ListIncidentOK with default headers values
func NewListIncidentOK() *ListIncidentOK {

	return &ListIncidentOK{}
}

// WithXResultCount adds the xResultCount to the list incident o k response
func (o *ListIncidentOK) WithXResultCount(xResultCount int64) *ListIncidentOK {
	o.XResultCount = xResultCount
	return o
}

// SetXResultCount sets the xResultCount to the list incident o k response
func (o *ListIncidentOK) SetXResultCount(xResultCount int64) {
	o.XResultCount = xResultCount
}

// WithXTotalCount adds the xTotalCount to the list incident o k response
func (o *ListIncidentOK) WithXTotalCount(xTotalCount int64) *ListIncidentOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the list incident o k response
func (o *ListIncidentOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the list incident o k response
func (o *ListIncidentOK) WithPayload(payload []*models.Incident) *ListIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident o k response
func (o *ListIncidentOK) SetPayload(payload []*models.Incident) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Result-Count

	xResultCount := swag.FormatInt64(o.XResultCount)
	if xResultCount != "" {
		rw.Header().Set("X-Result-Count", xResultCount)
	}

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Incident, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListIncidentBadRequestCode is the HTTP code returned for type ListIncidentBadRequest
const ListIncidentBadRequestCode int = 400

/*
ListIncidentBadRequest Bad Request

swagger:response listIncidentBadRequest
*/
type ListIncidentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentBadRequest creates ListIncidentBadRequest with default headers values
func NewListIncidentBadRequest() *ListIncidentBadRequest {

	return &ListIncidentBadRequest{}
}

// WithPayload adds the payload to the list incident bad request response
func (o *ListIncidentBadRequest) WithPayload(payload *models.Error) *ListIncidentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident bad request response
func (o *ListIncidentBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIncidentUnauthorizedCode is the HTTP code returned for type ListIncidentUnauthorized
const ListIncidentUnauthorizedCode int = 401

/*
ListIncidentUnauthorized Unauthorized

swagger:response listIncidentUnauthorized
*/
type ListIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentUnauthorized creates ListIncidentUnauthorized with default headers values
func NewListIncidentUnauthorized() *ListIncidentUnauthorized {

	return &ListIncidentUnauthorized{}
}

// WithPayload adds the payload to the list incident unauthorized response
func (o *ListIncidentUnauthorized) WithPayload(payload *models.Error) *ListIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident unauthorized response
func (o *ListIncidentUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIncidentForbiddenCode is the HTTP code returned for type ListIncidentForbidden
const ListIncidentForbiddenCode int = 403

/*
ListIncidentForbidden Forbidden

swagger:response listIncidentForbidden
*/
type ListIncidentForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentForbidden creates ListIncidentForbidden with default headers values
func NewListIncidentForbidden() *ListIncidentForbidden {

	return &ListIncidentForbidden{}
}

// WithPayload adds the payload to the list incident forbidden response
func (o *ListIncidentForbidden) WithPayload(payload *models.Error) *ListIncidentForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident forbidden response
func (o *ListIncidentForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIncidentNotFoundCode is the HTTP code returned for type ListIncidentNotFound
const ListIncidentNotFoundCode int = 404

/*
ListIncidentNotFound Not Found

swagger:response listIncidentNotFound
*/
type ListIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentNotFound creates ListIncidentNotFound with default headers values
func NewListIncidentNotFound() *ListIncidentNotFound {

	return &ListIncidentNotFound{}
}

// WithPayload adds the payload to the list incident not found response
func (o *ListIncidentNotFound) WithPayload(payload *models.Error) *ListIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident not found response
func (o *ListIncidentNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIncidentMethodNotAllowedCode is the HTTP code returned for type ListIncidentMethodNotAllowed
const ListIncidentMethodNotAllowedCode int = 405

/*
ListIncidentMethodNotAllowed Method Not allowed

swagger:response listIncidentMethodNotAllowed
*/
type ListIncidentMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentMethodNotAllowed creates ListIncidentMethodNotAllowed with default headers values
func NewListIncidentMethodNotAllowed() *ListIncidentMethodNotAllowed {

	return &ListIncidentMethodNotAllowed{}
}

// WithPayload adds the payload to the list incident method not allowed response
func (o *ListIncidentMethodNotAllowed) WithPayload(payload *models.Error) *ListIncidentMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident method not allowed response
func (o *ListIncidentMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIncidentConflictCode is the HTTP code returned for type ListIncidentConflict
const ListIncidentConflictCode int = 409

/*
ListIncidentConflict Conflict

swagger:response listIncidentConflict
*/
type ListIncidentConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentConflict creates ListIncidentConflict with default headers values
func NewListIncidentConflict() *ListIncidentConflict {

	return &ListIncidentConflict{}
}

// WithPayload adds the payload to the list incident conflict response
func (o *ListIncidentConflict) WithPayload(payload *models.Error) *ListIncidentConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident conflict response
func (o *ListIncidentConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIncidentInternalServerErrorCode is the HTTP code returned for type ListIncidentInternalServerError
const ListIncidentInternalServerErrorCode int = 500

/*
ListIncidentInternalServerError Internal Server Error

swagger:response listIncidentInternalServerError
*/
type ListIncidentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListIncidentInternalServerError creates ListIncidentInternalServerError with default headers values
func NewListIncidentInternalServerError() *ListIncidentInternalServerError {

	return &ListIncidentInternalServerError{}
}

// WithPayload adds the payload to the list incident internal server error response
func (o *ListIncidentInternalServerError) WithPayload(payload *models.Error) *ListIncidentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list incident internal server error response
func (o *ListIncidentInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIncidentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
