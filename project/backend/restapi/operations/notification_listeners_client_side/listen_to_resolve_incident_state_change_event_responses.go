// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/fmilheir/final_year_project/backend/models"
)

// ListenToResolveIncidentStateChangeEventCreatedCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventCreated
const ListenToResolveIncidentStateChangeEventCreatedCode int = 201

/*
ListenToResolveIncidentStateChangeEventCreated Notified

swagger:response listenToResolveIncidentStateChangeEventCreated
*/
type ListenToResolveIncidentStateChangeEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventCreated creates ListenToResolveIncidentStateChangeEventCreated with default headers values
func NewListenToResolveIncidentStateChangeEventCreated() *ListenToResolveIncidentStateChangeEventCreated {

	return &ListenToResolveIncidentStateChangeEventCreated{}
}

// WithPayload adds the payload to the listen to resolve incident state change event created response
func (o *ListenToResolveIncidentStateChangeEventCreated) WithPayload(payload *models.EventSubscription) *ListenToResolveIncidentStateChangeEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event created response
func (o *ListenToResolveIncidentStateChangeEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventBadRequestCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventBadRequest
const ListenToResolveIncidentStateChangeEventBadRequestCode int = 400

/*
ListenToResolveIncidentStateChangeEventBadRequest Bad Request

swagger:response listenToResolveIncidentStateChangeEventBadRequest
*/
type ListenToResolveIncidentStateChangeEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventBadRequest creates ListenToResolveIncidentStateChangeEventBadRequest with default headers values
func NewListenToResolveIncidentStateChangeEventBadRequest() *ListenToResolveIncidentStateChangeEventBadRequest {

	return &ListenToResolveIncidentStateChangeEventBadRequest{}
}

// WithPayload adds the payload to the listen to resolve incident state change event bad request response
func (o *ListenToResolveIncidentStateChangeEventBadRequest) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event bad request response
func (o *ListenToResolveIncidentStateChangeEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventUnauthorizedCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventUnauthorized
const ListenToResolveIncidentStateChangeEventUnauthorizedCode int = 401

/*
ListenToResolveIncidentStateChangeEventUnauthorized Unauthorized

swagger:response listenToResolveIncidentStateChangeEventUnauthorized
*/
type ListenToResolveIncidentStateChangeEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventUnauthorized creates ListenToResolveIncidentStateChangeEventUnauthorized with default headers values
func NewListenToResolveIncidentStateChangeEventUnauthorized() *ListenToResolveIncidentStateChangeEventUnauthorized {

	return &ListenToResolveIncidentStateChangeEventUnauthorized{}
}

// WithPayload adds the payload to the listen to resolve incident state change event unauthorized response
func (o *ListenToResolveIncidentStateChangeEventUnauthorized) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event unauthorized response
func (o *ListenToResolveIncidentStateChangeEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventForbiddenCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventForbidden
const ListenToResolveIncidentStateChangeEventForbiddenCode int = 403

/*
ListenToResolveIncidentStateChangeEventForbidden Forbidden

swagger:response listenToResolveIncidentStateChangeEventForbidden
*/
type ListenToResolveIncidentStateChangeEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventForbidden creates ListenToResolveIncidentStateChangeEventForbidden with default headers values
func NewListenToResolveIncidentStateChangeEventForbidden() *ListenToResolveIncidentStateChangeEventForbidden {

	return &ListenToResolveIncidentStateChangeEventForbidden{}
}

// WithPayload adds the payload to the listen to resolve incident state change event forbidden response
func (o *ListenToResolveIncidentStateChangeEventForbidden) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event forbidden response
func (o *ListenToResolveIncidentStateChangeEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventNotFoundCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventNotFound
const ListenToResolveIncidentStateChangeEventNotFoundCode int = 404

/*
ListenToResolveIncidentStateChangeEventNotFound Not Found

swagger:response listenToResolveIncidentStateChangeEventNotFound
*/
type ListenToResolveIncidentStateChangeEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventNotFound creates ListenToResolveIncidentStateChangeEventNotFound with default headers values
func NewListenToResolveIncidentStateChangeEventNotFound() *ListenToResolveIncidentStateChangeEventNotFound {

	return &ListenToResolveIncidentStateChangeEventNotFound{}
}

// WithPayload adds the payload to the listen to resolve incident state change event not found response
func (o *ListenToResolveIncidentStateChangeEventNotFound) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event not found response
func (o *ListenToResolveIncidentStateChangeEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventMethodNotAllowedCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventMethodNotAllowed
const ListenToResolveIncidentStateChangeEventMethodNotAllowedCode int = 405

/*
ListenToResolveIncidentStateChangeEventMethodNotAllowed Method Not allowed

swagger:response listenToResolveIncidentStateChangeEventMethodNotAllowed
*/
type ListenToResolveIncidentStateChangeEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventMethodNotAllowed creates ListenToResolveIncidentStateChangeEventMethodNotAllowed with default headers values
func NewListenToResolveIncidentStateChangeEventMethodNotAllowed() *ListenToResolveIncidentStateChangeEventMethodNotAllowed {

	return &ListenToResolveIncidentStateChangeEventMethodNotAllowed{}
}

// WithPayload adds the payload to the listen to resolve incident state change event method not allowed response
func (o *ListenToResolveIncidentStateChangeEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event method not allowed response
func (o *ListenToResolveIncidentStateChangeEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventConflictCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventConflict
const ListenToResolveIncidentStateChangeEventConflictCode int = 409

/*
ListenToResolveIncidentStateChangeEventConflict Conflict

swagger:response listenToResolveIncidentStateChangeEventConflict
*/
type ListenToResolveIncidentStateChangeEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventConflict creates ListenToResolveIncidentStateChangeEventConflict with default headers values
func NewListenToResolveIncidentStateChangeEventConflict() *ListenToResolveIncidentStateChangeEventConflict {

	return &ListenToResolveIncidentStateChangeEventConflict{}
}

// WithPayload adds the payload to the listen to resolve incident state change event conflict response
func (o *ListenToResolveIncidentStateChangeEventConflict) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event conflict response
func (o *ListenToResolveIncidentStateChangeEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListenToResolveIncidentStateChangeEventInternalServerErrorCode is the HTTP code returned for type ListenToResolveIncidentStateChangeEventInternalServerError
const ListenToResolveIncidentStateChangeEventInternalServerErrorCode int = 500

/*
ListenToResolveIncidentStateChangeEventInternalServerError Internal Server Error

swagger:response listenToResolveIncidentStateChangeEventInternalServerError
*/
type ListenToResolveIncidentStateChangeEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListenToResolveIncidentStateChangeEventInternalServerError creates ListenToResolveIncidentStateChangeEventInternalServerError with default headers values
func NewListenToResolveIncidentStateChangeEventInternalServerError() *ListenToResolveIncidentStateChangeEventInternalServerError {

	return &ListenToResolveIncidentStateChangeEventInternalServerError{}
}

// WithPayload adds the payload to the listen to resolve incident state change event internal server error response
func (o *ListenToResolveIncidentStateChangeEventInternalServerError) WithPayload(payload *models.Error) *ListenToResolveIncidentStateChangeEventInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to resolve incident state change event internal server error response
func (o *ListenToResolveIncidentStateChangeEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToResolveIncidentStateChangeEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}