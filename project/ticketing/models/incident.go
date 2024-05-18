// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jinzhu/gorm"
		 _ "github.com/jinzhu/gorm/dialects/postgres" 
	
)

// Incident An Incident is a record of an event that has altered the operational state of a entity (Resource, Service or Customers Product). An incident represents an issue that needs to be diagnosed and resolved.
//
// swagger:model Incident
type Incident struct {
	gorm.Model

	// When sub-classing, this defines the super-class
	AtBaseType string `json:"@baseType,omitempty" gorm:"size:255"` 

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation strfmt.URI `json:"@schemaLocation,omitempty" grom:"size:255"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType string `json:"@type,omitempty" gorm:"size:255"`

	// Provides the Acknowledgement State of the incident (unacknowledged | acknowledged).
	AckState IncidentAckStateType `json:"ackState,omitempty"`

	// The ackTime or acknowledgeTime of the incident. An acknowledged incident is being worked on, but is not yet resolved
	// Example: 2022-03-10T04:01:12Z
	// Format: date-time
	AckTime strfmt.DateTime `json:"ackTime,omitempty" gorm:"type:timestamp" `

	// List of affected entities
	// Example: e.g  a ref to a service or resource
	AffectedEntity []*EntityRef `json:"affectedEntity" gorm:"foreignKey:ID"`

	// The category of the incident  (category is the term used by ITU)
	// Example: BTS Software Fault
	Category string `json:"category,omitempty"`

	// The clear time of the incident
	// Example: 2022-03-10T04:01:12Z
	// Format: date-time
	ClearTime strfmt.DateTime `json:"clearTime,omitempty" gorm:"size:255"`

	// The domain of the incident, for example RAN, PON, OTN, Cross-Domain etc
	Domain string `json:"domain,omitempty" gorm:"size:255"`

	// The correlation event object such as alarm, externalAlarm, performance, etc.
	EventID []*ResourceEntity `json:"eventId" gorm:"foreignKey:ID"`

	// This is used for extend the incident with attributes
	ExtensionInfo []*Characteristic `json:"extensionInfo" gorm:"foreignKey:ID"`

	// An identification of an entity that is owned by or originates in a software system different from the current system, for example a ProductOrder handed off from a commerce platform into an order handling system. The structure identifies the system itself, the nature of the entity within the system (e.g. class name) and the unique ID of the entity within the system. It is anticipated that multiple external IDs can be held for a single entity, e.g. if the entity passed through multiple systems on the way to the current system. In this case the consumer is expected to sequence the IDs in the array in reverse order of provenance, i.e. most recent system first in the list.
	ExternalIdentifier []*ExternalIdentifier `json:"externalIdentifier" gorm:"foreignKey:ID"`

	// Hyperlink, a reference to the incident entity
	Href string `json:"href,omitempty" gorm:"size:255"`

	// unique identifier
	ID string `json:"id,omitempty" " gorm:"primaryKey"`

	// Impact which indicates the degree of impact on affected services or users. This field is optional. The options are extensive, significant, moderate, and minor
	Impact ImpactType `json:"impact,omitempty"`

	// A textual succinct description of the nature, symptoms, cause, or effect of the incident.
	IncidentDetail string `json:"incidentDetail,omitempty"`

	// Incident resolution suggestion or tip to resolve the incident
	IncidentResolutionSuggestion string `json:"incidentResolutionSuggestion,omitempty"`

	// The name of the incident. A short-form string that provides succinct, important information about the incident
	Name string `json:"name,omitempty"`

	// Indicates the time (as a date + time) at which the incident occurred at its source
	// Example: 2022-03-10T04:01:12Z
	// Format: date-time
	OccurTime strfmt.DateTime `json:"occurTime,omitempty"`

	// The priority  of the incident, priority  critical/high/medium/low
	Priority PriorityType `json:"priority,omitempty"`

	// A root cause is a fundamental or underlying reason behind why an incident occurred that identifies one or more failures. An incident many have multiple rootCauses
	RootCause []*RootCause `json:"rootCause" gorm:"foreignKey:ID"`

	// The root event object such as alarm, externalAlarm, performance, etc.
	RootEventID []*ResourceEntity `json:"rootEventId" gorm:"foreignKey:ID"`

	// The objects show the incident, it may be part of Network Equipment. Fault object, which may be an NE or a port.
	// Example: e.g. weak optical signals, the fault object is a PON port, the root cause is an optical splitter, and the affected object is an ONU
	SourceObject []*ResourceEntity `json:"sourceObject" gorm:"foreignKey:ID"`

	// Incident state. The options are raised | updated | cleared. Cleared means Resolved)
	State IncidentStateType `json:"state,omitempty"`

	// The last update time  of the incident
	// Example: 2022-03-10T04:01:12Z
	UpdateTime string `json:"updateTime,omitempty"`

	// Urgency is the speed required for resolving the service issues. A measure of how long it will be until an incident has a significant impact on the business. This field is optional. The options are critical, high, medium, and low
	Urgency UrgencyType `json:"urgency,omitempty"`
}

// Validate validates this incident
func (m *Incident) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAckState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAckTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAffectedEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClearTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOccurTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrgency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Incident) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateAckState(formats strfmt.Registry) error {
	if swag.IsZero(m.AckState) { // not required
		return nil
	}

	if err := m.AckState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ackState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ackState")
		}
		return err
	}

	return nil
}

func (m *Incident) validateAckTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AckTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ackTime", "body", "date-time", m.AckTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateAffectedEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.AffectedEntity) { // not required
		return nil
	}

	for i := 0; i < len(m.AffectedEntity); i++ {
		if swag.IsZero(m.AffectedEntity[i]) { // not required
			continue
		}

		if m.AffectedEntity[i] != nil {
			if err := m.AffectedEntity[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affectedEntity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("affectedEntity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateClearTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ClearTime) { // not required
		return nil
	}

	if err := validate.FormatOf("clearTime", "body", "date-time", m.ClearTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validateEventID(formats strfmt.Registry) error {
	if swag.IsZero(m.EventID) { // not required
		return nil
	}

	for i := 0; i < len(m.EventID); i++ {
		if swag.IsZero(m.EventID[i]) { // not required
			continue
		}

		if m.EventID[i] != nil {
			if err := m.EventID[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventId" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("eventId" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateExtensionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtensionInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtensionInfo); i++ {
		if swag.IsZero(m.ExtensionInfo[i]) { // not required
			continue
		}

		if m.ExtensionInfo[i] != nil {
			if err := m.ExtensionInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extensionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateExternalIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalIdentifier) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalIdentifier); i++ {
		if swag.IsZero(m.ExternalIdentifier[i]) { // not required
			continue
		}

		if m.ExternalIdentifier[i] != nil {
			if err := m.ExternalIdentifier[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateImpact(formats strfmt.Registry) error {
	if swag.IsZero(m.Impact) { // not required
		return nil
	}

	if err := m.Impact.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("impact")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("impact")
		}
		return err
	}

	return nil
}

func (m *Incident) validateOccurTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OccurTime) { // not required
		return nil
	}

	if err := validate.FormatOf("occurTime", "body", "date-time", m.OccurTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Incident) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := m.Priority.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("priority")
		}
		return err
	}

	return nil
}

func (m *Incident) validateRootCause(formats strfmt.Registry) error {
	if swag.IsZero(m.RootCause) { // not required
		return nil
	}

	for i := 0; i < len(m.RootCause); i++ {
		if swag.IsZero(m.RootCause[i]) { // not required
			continue
		}

		if m.RootCause[i] != nil {
			if err := m.RootCause[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootCause" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rootCause" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateRootEventID(formats strfmt.Registry) error {
	if swag.IsZero(m.RootEventID) { // not required
		return nil
	}

	for i := 0; i < len(m.RootEventID); i++ {
		if swag.IsZero(m.RootEventID[i]) { // not required
			continue
		}

		if m.RootEventID[i] != nil {
			if err := m.RootEventID[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootEventId" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rootEventId" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateSourceObject(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceObject) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceObject); i++ {
		if swag.IsZero(m.SourceObject[i]) { // not required
			continue
		}

		if m.SourceObject[i] != nil {
			if err := m.SourceObject[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceObject" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceObject" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Incident) validateUrgency(formats strfmt.Registry) error {
	if swag.IsZero(m.Urgency) { // not required
		return nil
	}

	if err := m.Urgency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("urgency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("urgency")
		}
		return err
	}

	return nil
}

// ContextValidate validate this incident based on the context it is used
func (m *Incident) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAckState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAffectedEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImpact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootEventID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUrgency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Incident) contextValidateAckState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AckState) { // not required
		return nil
	}

	if err := m.AckState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ackState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ackState")
		}
		return err
	}

	return nil
}

func (m *Incident) contextValidateAffectedEntity(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AffectedEntity); i++ {

		if m.AffectedEntity[i] != nil {

			if swag.IsZero(m.AffectedEntity[i]) { // not required
				return nil
			}

			if err := m.AffectedEntity[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affectedEntity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("affectedEntity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateEventID(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EventID); i++ {

		if m.EventID[i] != nil {

			if swag.IsZero(m.EventID[i]) { // not required
				return nil
			}

			if err := m.EventID[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventId" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("eventId" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateExtensionInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtensionInfo); i++ {

		if m.ExtensionInfo[i] != nil {

			if swag.IsZero(m.ExtensionInfo[i]) { // not required
				return nil
			}

			if err := m.ExtensionInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extensionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateExternalIdentifier(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalIdentifier); i++ {

		if m.ExternalIdentifier[i] != nil {

			if swag.IsZero(m.ExternalIdentifier[i]) { // not required
				return nil
			}

			if err := m.ExternalIdentifier[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateImpact(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Impact) { // not required
		return nil
	}

	if err := m.Impact.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("impact")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("impact")
		}
		return err
	}

	return nil
}

func (m *Incident) contextValidatePriority(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := m.Priority.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("priority")
		}
		return err
	}

	return nil
}

func (m *Incident) contextValidateRootCause(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RootCause); i++ {

		if m.RootCause[i] != nil {

			if swag.IsZero(m.RootCause[i]) { // not required
				return nil
			}

			if err := m.RootCause[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootCause" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rootCause" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateRootEventID(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RootEventID); i++ {

		if m.RootEventID[i] != nil {

			if swag.IsZero(m.RootEventID[i]) { // not required
				return nil
			}

			if err := m.RootEventID[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootEventId" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rootEventId" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateSourceObject(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SourceObject); i++ {

		if m.SourceObject[i] != nil {

			if swag.IsZero(m.SourceObject[i]) { // not required
				return nil
			}

			if err := m.SourceObject[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceObject" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceObject" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Incident) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Incident) contextValidateUrgency(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Urgency) { // not required
		return nil
	}

	if err := m.Urgency.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("urgency")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("urgency")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Incident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Incident) UnmarshalBinary(b []byte) error {
	var res Incident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
