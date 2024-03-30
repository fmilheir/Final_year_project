// Code generated by go-swagger; DO NOT EDIT.

package events_subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/fmilheir/final_year_project/backend/models"
)

// UnregisterListenerNoContentCode is the HTTP code returned for type UnregisterListenerNoContent
const UnregisterListenerNoContentCode int = 204

/*
UnregisterListenerNoContent Deleted

swagger:response unregisterListenerNoContent
*/
type UnregisterListenerNoContent struct {
}

// NewUnregisterListenerNoContent creates UnregisterListenerNoContent with default headers values
func NewUnregisterListenerNoContent() *UnregisterListenerNoContent {

	return &UnregisterListenerNoContent{}
}

// WriteResponse to the client
func (o *UnregisterListenerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// UnregisterListenerBadRequestCode is the HTTP code returned for type UnregisterListenerBadRequest
const UnregisterListenerBadRequestCode int = 400

/*
UnregisterListenerBadRequest Bad request

swagger:response unregisterListenerBadRequest
*/
type UnregisterListenerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnregisterListenerBadRequest creates UnregisterListenerBadRequest with default headers values
func NewUnregisterListenerBadRequest() *UnregisterListenerBadRequest {

	return &UnregisterListenerBadRequest{}
}

// WithPayload adds the payload to the unregister listener bad request response
func (o *UnregisterListenerBadRequest) WithPayload(payload *models.Error) *UnregisterListenerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unregister listener bad request response
func (o *UnregisterListenerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnregisterListenerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnregisterListenerUnauthorizedCode is the HTTP code returned for type UnregisterListenerUnauthorized
const UnregisterListenerUnauthorizedCode int = 401

/*
UnregisterListenerUnauthorized Unauthorized

swagger:response unregisterListenerUnauthorized
*/
type UnregisterListenerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnregisterListenerUnauthorized creates UnregisterListenerUnauthorized with default headers values
func NewUnregisterListenerUnauthorized() *UnregisterListenerUnauthorized {

	return &UnregisterListenerUnauthorized{}
}

// WithPayload adds the payload to the unregister listener unauthorized response
func (o *UnregisterListenerUnauthorized) WithPayload(payload *models.Error) *UnregisterListenerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unregister listener unauthorized response
func (o *UnregisterListenerUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnregisterListenerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnregisterListenerForbiddenCode is the HTTP code returned for type UnregisterListenerForbidden
const UnregisterListenerForbiddenCode int = 403

/*
UnregisterListenerForbidden Forbidden

swagger:response unregisterListenerForbidden
*/
type UnregisterListenerForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnregisterListenerForbidden creates UnregisterListenerForbidden with default headers values
func NewUnregisterListenerForbidden() *UnregisterListenerForbidden {

	return &UnregisterListenerForbidden{}
}

// WithPayload adds the payload to the unregister listener forbidden response
func (o *UnregisterListenerForbidden) WithPayload(payload *models.Error) *UnregisterListenerForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unregister listener forbidden response
func (o *UnregisterListenerForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnregisterListenerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnregisterListenerNotFoundCode is the HTTP code returned for type UnregisterListenerNotFound
const UnregisterListenerNotFoundCode int = 404

/*
UnregisterListenerNotFound Not Found

swagger:response unregisterListenerNotFound
*/
type UnregisterListenerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnregisterListenerNotFound creates UnregisterListenerNotFound with default headers values
func NewUnregisterListenerNotFound() *UnregisterListenerNotFound {

	return &UnregisterListenerNotFound{}
}

// WithPayload adds the payload to the unregister listener not found response
func (o *UnregisterListenerNotFound) WithPayload(payload *models.Error) *UnregisterListenerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unregister listener not found response
func (o *UnregisterListenerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnregisterListenerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnregisterListenerMethodNotAllowedCode is the HTTP code returned for type UnregisterListenerMethodNotAllowed
const UnregisterListenerMethodNotAllowedCode int = 405

/*
UnregisterListenerMethodNotAllowed Method not allowed

swagger:response unregisterListenerMethodNotAllowed
*/
type UnregisterListenerMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnregisterListenerMethodNotAllowed creates UnregisterListenerMethodNotAllowed with default headers values
func NewUnregisterListenerMethodNotAllowed() *UnregisterListenerMethodNotAllowed {

	return &UnregisterListenerMethodNotAllowed{}
}

// WithPayload adds the payload to the unregister listener method not allowed response
func (o *UnregisterListenerMethodNotAllowed) WithPayload(payload *models.Error) *UnregisterListenerMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unregister listener method not allowed response
func (o *UnregisterListenerMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnregisterListenerMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UnregisterListenerInternalServerErrorCode is the HTTP code returned for type UnregisterListenerInternalServerError
const UnregisterListenerInternalServerErrorCode int = 500

/*
UnregisterListenerInternalServerError Internal Server Error

swagger:response unregisterListenerInternalServerError
*/
type UnregisterListenerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnregisterListenerInternalServerError creates UnregisterListenerInternalServerError with default headers values
func NewUnregisterListenerInternalServerError() *UnregisterListenerInternalServerError {

	return &UnregisterListenerInternalServerError{}
}

// WithPayload adds the payload to the unregister listener internal server error response
func (o *UnregisterListenerInternalServerError) WithPayload(payload *models.Error) *UnregisterListenerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unregister listener internal server error response
func (o *UnregisterListenerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnregisterListenerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}