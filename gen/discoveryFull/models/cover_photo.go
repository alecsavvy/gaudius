// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CoverPhoto cover photo
//
// swagger:model cover_photo
type CoverPhoto struct {

	// 2000x
	Nr2000x string `json:"2000x,omitempty"`

	// 640x
	Nr640x string `json:"640x,omitempty"`
}

// Validate validates this cover photo
func (m *CoverPhoto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cover photo based on context it is used
func (m *CoverPhoto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoverPhoto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoverPhoto) UnmarshalBinary(b []byte) error {
	var res CoverPhoto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}