// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResolveIncidentCreateEvent The notification data structure
//
// swagger:model ResolveIncidentCreateEvent
type ResolveIncidentCreateEvent struct {

	// The correlation id for this event.
	CorrelationID string `json:"correlationId,omitempty"`

	// An explnatory of the event.
	Description string `json:"description,omitempty"`

	// The domain of the event.
	Domain string `json:"domain,omitempty"`

	// The event payload linked to the involved resource object
	Event *ResolveIncidentCreateEventPayload `json:"event,omitempty"`

	// The identifier of the notification.
	EventID string `json:"eventId,omitempty"`

	// Time of the event occurrence.
	// Format: date-time
	EventTime strfmt.DateTime `json:"eventTime,omitempty"`

	// The type of the notification.
	EventType string `json:"eventType,omitempty"`

	// A priority.
	Priority string `json:"priority,omitempty"`

	// The time the event occured.
	// Format: date-time
	TimeOcurred strfmt.DateTime `json:"timeOcurred,omitempty"`

	// The title of the event.
	Title string `json:"title,omitempty"`
}

// Validate validates this resolve incident create event
func (m *ResolveIncidentCreateEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOcurred(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResolveIncidentCreateEvent) validateEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *ResolveIncidentCreateEvent) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResolveIncidentCreateEvent) validateTimeOcurred(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeOcurred) { // not required
		return nil
	}

	if err := validate.FormatOf("timeOcurred", "body", "date-time", m.TimeOcurred.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resolve incident create event based on the context it is used
func (m *ResolveIncidentCreateEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResolveIncidentCreateEvent) contextValidateEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.Event != nil {

		if swag.IsZero(m.Event) { // not required
			return nil
		}

		if err := m.Event.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResolveIncidentCreateEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolveIncidentCreateEvent) UnmarshalBinary(b []byte) error {
	var res ResolveIncidentCreateEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
