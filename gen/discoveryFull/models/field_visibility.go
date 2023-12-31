// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FieldVisibility field visibility
//
// swagger:model field_visibility
type FieldVisibility struct {

	// genre
	Genre bool `json:"genre,omitempty"`

	// mood
	Mood bool `json:"mood,omitempty"`

	// play count
	PlayCount bool `json:"play_count,omitempty"`

	// remixes
	Remixes bool `json:"remixes,omitempty"`

	// share
	Share bool `json:"share,omitempty"`

	// tags
	Tags bool `json:"tags,omitempty"`
}

// Validate validates this field visibility
func (m *FieldVisibility) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this field visibility based on context it is used
func (m *FieldVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FieldVisibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldVisibility) UnmarshalBinary(b []byte) error {
	var res FieldVisibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
