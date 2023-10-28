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
)

// StemsResponse stems response
//
// swagger:model stems_response
type StemsResponse struct {

	// data
	Data []*StemFull `json:"data"`

	// latest chain block
	// Required: true
	LatestChainBlock *int64 `json:"latest_chain_block"`

	// latest chain slot plays
	// Required: true
	LatestChainSlotPlays *int64 `json:"latest_chain_slot_plays"`

	// latest indexed block
	// Required: true
	LatestIndexedBlock *int64 `json:"latest_indexed_block"`

	// latest indexed slot plays
	// Required: true
	LatestIndexedSlotPlays *int64 `json:"latest_indexed_slot_plays"`

	// signature
	// Required: true
	Signature *string `json:"signature"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`

	// version
	// Required: true
	Version *VersionMetadata `json:"version"`
}

// Validate validates this stems response
func (m *StemsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestChainBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestChainSlotPlays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestIndexedBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestIndexedSlotPlays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StemsResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StemsResponse) validateLatestChainBlock(formats strfmt.Registry) error {

	if err := validate.Required("latest_chain_block", "body", m.LatestChainBlock); err != nil {
		return err
	}

	return nil
}

func (m *StemsResponse) validateLatestChainSlotPlays(formats strfmt.Registry) error {

	if err := validate.Required("latest_chain_slot_plays", "body", m.LatestChainSlotPlays); err != nil {
		return err
	}

	return nil
}

func (m *StemsResponse) validateLatestIndexedBlock(formats strfmt.Registry) error {

	if err := validate.Required("latest_indexed_block", "body", m.LatestIndexedBlock); err != nil {
		return err
	}

	return nil
}

func (m *StemsResponse) validateLatestIndexedSlotPlays(formats strfmt.Registry) error {

	if err := validate.Required("latest_indexed_slot_plays", "body", m.LatestIndexedSlotPlays); err != nil {
		return err
	}

	return nil
}

func (m *StemsResponse) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	return nil
}

func (m *StemsResponse) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *StemsResponse) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this stems response based on the context it is used
func (m *StemsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StemsResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {

			if swag.IsZero(m.Data[i]) { // not required
				return nil
			}

			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StemsResponse) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {

		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StemsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StemsResponse) UnmarshalBinary(b []byte) error {
	var res StemsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
