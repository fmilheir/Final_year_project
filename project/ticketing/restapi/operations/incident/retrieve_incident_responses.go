// Code generated by go-swagger; DO NOT EDIT.

package incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"fmilheir/test_swagger/models"
)

// RetrieveIncidentOKCode is the HTTP code returned for type RetrieveIncidentOK
const RetrieveIncidentOKCode int = 200

/*
RetrieveIncidentOK Success

swagger:response retrieveIncidentOK
*/
type RetrieveIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.Incident `json:"body,omitempty"`
}

// NewRetrieveIncidentOK creates RetrieveIncidentOK with default headers values
func NewRetrieveIncidentOK() *RetrieveIncidentOK {

	return &RetrieveIncidentOK{}
}

// WithPayload adds the payload to the retrieve incident o k response
func (o *RetrieveIncidentOK) WithPayload(payload *models.Incident) *RetrieveIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident o k response
func (o *RetrieveIncidentOK) SetPayload(payload *models.Incident) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentBadRequestCode is the HTTP code returned for type RetrieveIncidentBadRequest
const RetrieveIncidentBadRequestCode int = 400

/*
RetrieveIncidentBadRequest Bad Request

swagger:response retrieveIncidentBadRequest
*/
type RetrieveIncidentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentBadRequest creates RetrieveIncidentBadRequest with default headers values
func NewRetrieveIncidentBadRequest() *RetrieveIncidentBadRequest {

	return &RetrieveIncidentBadRequest{}
}

// WithPayload adds the payload to the retrieve incident bad request response
func (o *RetrieveIncidentBadRequest) WithPayload(payload *models.Error) *RetrieveIncidentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident bad request response
func (o *RetrieveIncidentBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentUnauthorizedCode is the HTTP code returned for type RetrieveIncidentUnauthorized
const RetrieveIncidentUnauthorizedCode int = 401

/*
RetrieveIncidentUnauthorized Unauthorized

swagger:response retrieveIncidentUnauthorized
*/
type RetrieveIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentUnauthorized creates RetrieveIncidentUnauthorized with default headers values
func NewRetrieveIncidentUnauthorized() *RetrieveIncidentUnauthorized {

	return &RetrieveIncidentUnauthorized{}
}

// WithPayload adds the payload to the retrieve incident unauthorized response
func (o *RetrieveIncidentUnauthorized) WithPayload(payload *models.Error) *RetrieveIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident unauthorized response
func (o *RetrieveIncidentUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentForbiddenCode is the HTTP code returned for type RetrieveIncidentForbidden
const RetrieveIncidentForbiddenCode int = 403

/*
RetrieveIncidentForbidden Forbidden

swagger:response retrieveIncidentForbidden
*/
type RetrieveIncidentForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentForbidden creates RetrieveIncidentForbidden with default headers values
func NewRetrieveIncidentForbidden() *RetrieveIncidentForbidden {

	return &RetrieveIncidentForbidden{}
}

// WithPayload adds the payload to the retrieve incident forbidden response
func (o *RetrieveIncidentForbidden) WithPayload(payload *models.Error) *RetrieveIncidentForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident forbidden response
func (o *RetrieveIncidentForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentNotFoundCode is the HTTP code returned for type RetrieveIncidentNotFound
const RetrieveIncidentNotFoundCode int = 404

/*
RetrieveIncidentNotFound Not Found

swagger:response retrieveIncidentNotFound
*/
type RetrieveIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentNotFound creates RetrieveIncidentNotFound with default headers values
func NewRetrieveIncidentNotFound() *RetrieveIncidentNotFound {

	return &RetrieveIncidentNotFound{}
}

// WithPayload adds the payload to the retrieve incident not found response
func (o *RetrieveIncidentNotFound) WithPayload(payload *models.Error) *RetrieveIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident not found response
func (o *RetrieveIncidentNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentMethodNotAllowedCode is the HTTP code returned for type RetrieveIncidentMethodNotAllowed
const RetrieveIncidentMethodNotAllowedCode int = 405

/*
RetrieveIncidentMethodNotAllowed Method Not allowed

swagger:response retrieveIncidentMethodNotAllowed
*/
type RetrieveIncidentMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentMethodNotAllowed creates RetrieveIncidentMethodNotAllowed with default headers values
func NewRetrieveIncidentMethodNotAllowed() *RetrieveIncidentMethodNotAllowed {

	return &RetrieveIncidentMethodNotAllowed{}
}

// WithPayload adds the payload to the retrieve incident method not allowed response
func (o *RetrieveIncidentMethodNotAllowed) WithPayload(payload *models.Error) *RetrieveIncidentMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident method not allowed response
func (o *RetrieveIncidentMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentConflictCode is the HTTP code returned for type RetrieveIncidentConflict
const RetrieveIncidentConflictCode int = 409

/*
RetrieveIncidentConflict Conflict

swagger:response retrieveIncidentConflict
*/
type RetrieveIncidentConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentConflict creates RetrieveIncidentConflict with default headers values
func NewRetrieveIncidentConflict() *RetrieveIncidentConflict {

	return &RetrieveIncidentConflict{}
}

// WithPayload adds the payload to the retrieve incident conflict response
func (o *RetrieveIncidentConflict) WithPayload(payload *models.Error) *RetrieveIncidentConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident conflict response
func (o *RetrieveIncidentConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveIncidentInternalServerErrorCode is the HTTP code returned for type RetrieveIncidentInternalServerError
const RetrieveIncidentInternalServerErrorCode int = 500

/*
RetrieveIncidentInternalServerError Internal Server Error

swagger:response retrieveIncidentInternalServerError
*/
type RetrieveIncidentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveIncidentInternalServerError creates RetrieveIncidentInternalServerError with default headers values
func NewRetrieveIncidentInternalServerError() *RetrieveIncidentInternalServerError {

	return &RetrieveIncidentInternalServerError{}
}

// WithPayload adds the payload to the retrieve incident internal server error response
func (o *RetrieveIncidentInternalServerError) WithPayload(payload *models.Error) *RetrieveIncidentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve incident internal server error response
func (o *RetrieveIncidentInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveIncidentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
