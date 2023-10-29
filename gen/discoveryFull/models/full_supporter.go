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

// FullSupporter full supporter
//
// swagger:model full_supporter
type FullSupporter struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// rank
	// Required: true
	Rank *int64 `json:"rank"`

	// sender
	// Required: true
	Sender *UserFull `json:"sender"`
}

// Validate validates this full supporter
func (m *FullSupporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullSupporter) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *FullSupporter) validateRank(formats strfmt.Registry) error {

	if err := validate.Required("rank", "body", m.Rank); err != nil {
		return err
	}

	return nil
}

func (m *FullSupporter) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this full supporter based on the context it is used
func (m *FullSupporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FullSupporter) contextValidateSender(ctx context.Context, formats strfmt.Registry) error {

	if m.Sender != nil {

		if err := m.Sender.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FullSupporter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FullSupporter) UnmarshalBinary(b []byte) error {
	var res FullSupporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}