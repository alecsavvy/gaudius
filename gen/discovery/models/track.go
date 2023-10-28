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

// Track track
//
// swagger:model Track
type Track struct {

	// artwork
	Artwork *TrackArtwork `json:"artwork,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// downloadable
	Downloadable bool `json:"downloadable,omitempty"`

	// duration
	// Required: true
	Duration *int64 `json:"duration"`

	// favorite count
	// Required: true
	FavoriteCount *int64 `json:"favorite_count"`

	// genre
	Genre string `json:"genre,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// is streamable
	IsStreamable bool `json:"is_streamable,omitempty"`

	// mood
	Mood string `json:"mood,omitempty"`

	// permalink
	Permalink string `json:"permalink,omitempty"`

	// play count
	// Required: true
	PlayCount *int64 `json:"play_count"`

	// preview cid
	PreviewCid string `json:"preview_cid,omitempty"`

	// release date
	ReleaseDate string `json:"release_date,omitempty"`

	// remix of
	RemixOf *RemixParent `json:"remix_of,omitempty"`

	// repost count
	// Required: true
	RepostCount *int64 `json:"repost_count"`

	// tags
	Tags string `json:"tags,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// track cid
	TrackCid string `json:"track_cid,omitempty"`

	// user
	// Required: true
	User *User `json:"user"`
}

// Validate validates this track
func (m *Track) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFavoriteCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemixOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepostCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Track) validateArtwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Artwork) { // not required
		return nil
	}

	if m.Artwork != nil {
		if err := m.Artwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artwork")
			}
			return err
		}
	}

	return nil
}

func (m *Track) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *Track) validateFavoriteCount(formats strfmt.Registry) error {

	if err := validate.Required("favorite_count", "body", m.FavoriteCount); err != nil {
		return err
	}

	return nil
}

func (m *Track) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Track) validatePlayCount(formats strfmt.Registry) error {

	if err := validate.Required("play_count", "body", m.PlayCount); err != nil {
		return err
	}

	return nil
}

func (m *Track) validateRemixOf(formats strfmt.Registry) error {
	if swag.IsZero(m.RemixOf) { // not required
		return nil
	}

	if m.RemixOf != nil {
		if err := m.RemixOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remix_of")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remix_of")
			}
			return err
		}
	}

	return nil
}

func (m *Track) validateRepostCount(formats strfmt.Registry) error {

	if err := validate.Required("repost_count", "body", m.RepostCount); err != nil {
		return err
	}

	return nil
}

func (m *Track) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Track) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this track based on the context it is used
func (m *Track) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArtwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemixOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Track) contextValidateArtwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Artwork != nil {

		if swag.IsZero(m.Artwork) { // not required
			return nil
		}

		if err := m.Artwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artwork")
			}
			return err
		}
	}

	return nil
}

func (m *Track) contextValidateRemixOf(ctx context.Context, formats strfmt.Registry) error {

	if m.RemixOf != nil {

		if swag.IsZero(m.RemixOf) { // not required
			return nil
		}

		if err := m.RemixOf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remix_of")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remix_of")
			}
			return err
		}
	}

	return nil
}

func (m *Track) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Track) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Track) UnmarshalBinary(b []byte) error {
	var res Track
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
