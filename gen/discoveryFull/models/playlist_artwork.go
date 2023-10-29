// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PlaylistArtwork playlist artwork
//
// swagger:model playlist_artwork
type PlaylistArtwork struct {

	// 1000x1000
	Nr1000x1000 string `json:"1000x1000,omitempty"`

	// 150x150
	Nr150x150 string `json:"150x150,omitempty"`

	// 480x480
	Nr480x480 string `json:"480x480,omitempty"`
}

// Validate validates this playlist artwork
func (m *PlaylistArtwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this playlist artwork based on context it is used
func (m *PlaylistArtwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlaylistArtwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlaylistArtwork) UnmarshalBinary(b []byte) error {
	var res PlaylistArtwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}