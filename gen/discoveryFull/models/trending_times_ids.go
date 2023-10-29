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

// TrendingTimesIds trending times ids
//
// swagger:model trending_times_ids
type TrendingTimesIds struct {

	// month
	Month []*TrackID `json:"month"`

	// week
	Week []*TrackID `json:"week"`

	// year
	Year []*TrackID `json:"year"`
}

// Validate validates this trending times ids
func (m *TrendingTimesIds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYear(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrendingTimesIds) validateMonth(formats strfmt.Registry) error {
	if swag.IsZero(m.Month) { // not required
		return nil
	}

	for i := 0; i < len(m.Month); i++ {
		if swag.IsZero(m.Month[i]) { // not required
			continue
		}

		if m.Month[i] != nil {
			if err := m.Month[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("month" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("month" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrendingTimesIds) validateWeek(formats strfmt.Registry) error {
	if swag.IsZero(m.Week) { // not required
		return nil
	}

	for i := 0; i < len(m.Week); i++ {
		if swag.IsZero(m.Week[i]) { // not required
			continue
		}

		if m.Week[i] != nil {
			if err := m.Week[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("week" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("week" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrendingTimesIds) validateYear(formats strfmt.Registry) error {
	if swag.IsZero(m.Year) { // not required
		return nil
	}

	for i := 0; i < len(m.Year); i++ {
		if swag.IsZero(m.Year[i]) { // not required
			continue
		}

		if m.Year[i] != nil {
			if err := m.Year[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("year" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("year" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this trending times ids based on the context it is used
func (m *TrendingTimesIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMonth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeek(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateYear(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrendingTimesIds) contextValidateMonth(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Month); i++ {

		if m.Month[i] != nil {

			if swag.IsZero(m.Month[i]) { // not required
				return nil
			}

			if err := m.Month[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("month" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("month" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrendingTimesIds) contextValidateWeek(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Week); i++ {

		if m.Week[i] != nil {

			if swag.IsZero(m.Week[i]) { // not required
				return nil
			}

			if err := m.Week[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("week" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("week" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TrendingTimesIds) contextValidateYear(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Year); i++ {

		if m.Year[i] != nil {

			if swag.IsZero(m.Year[i]) { // not required
				return nil
			}

			if err := m.Year[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("year" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("year" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrendingTimesIds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrendingTimesIds) UnmarshalBinary(b []byte) error {
	var res TrendingTimesIds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}