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

// IncidentRef Incident reference.
//
// swagger:model IncidentRef
type IncidentRef struct {

	// When sub-classing, this defines the super-class
	AtBaseType string `json:"@baseType,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	AtReferredType string `json:"@referredType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation strfmt.URI `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType string `json:"@type,omitempty"`

	// Hyperlink reference
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// unique identifier
	// Required: true
	ID *string `json:"id"`

	// Name of the related incident
	Name string `json:"name,omitempty"`

	// resolution (edited by me in order to be able to get the resolution from the incident)
	IncidentResolutionSuggestion string `json:"resolution,omitempty"`
}

// Validate validates this incident ref
func (m *IncidentRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentRef) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentRef) validateHref(formats strfmt.Registry) error {
	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this incident ref based on context it is used
func (m *IncidentRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IncidentRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentRef) UnmarshalBinary(b []byte) error {
	var res IncidentRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
