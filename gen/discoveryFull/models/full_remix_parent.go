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
)

// FullRemixParent full remix parent
//
// swagger:model full_remix_parent
type FullRemixParent struct {

	// tracks
	Tracks []*FullRemix `json:"tracks"`
}

// Validate validates this full remix parent
func (m *FullRemixParent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTracks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullRemixParent) validateTracks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tracks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tracks); i++ {
		if swag.IsZero(m.Tracks[i]) { // not required
			continue
		}

		if m.Tracks[i] != nil {
			if err := m.Tracks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tracks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tracks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this full remix parent based on the context it is used
func (m *FullRemixParent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTracks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullRemixParent) contextValidateTracks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tracks); i++ {

		if m.Tracks[i] != nil {

			if swag.IsZero(m.Tracks[i]) { // not required
				return nil
			}

			if err := m.Tracks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tracks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tracks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullRemixParent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullRemixParent) UnmarshalBinary(b []byte) error {
	var res FullRemixParent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
