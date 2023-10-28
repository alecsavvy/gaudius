// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActivityFull activity full
//
// swagger:model activity_full
type ActivityFull struct {

	// item
	Item interface{} `json:"item,omitempty"`

	// item type
	ItemType interface{} `json:"item_type,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this activity full
func (m *ActivityFull) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this activity full based on context it is used
func (m *ActivityFull) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActivityFull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityFull) UnmarshalBinary(b []byte) error {
	var res ActivityFull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
