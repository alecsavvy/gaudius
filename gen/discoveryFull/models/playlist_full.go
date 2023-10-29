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

// PlaylistFull playlist full
//
// swagger:model playlist_full
type PlaylistFull struct {

	// added timestamps
	// Required: true
	AddedTimestamps []*PlaylistAddedTimestamp `json:"added_timestamps"`

	// artwork
	Artwork *PlaylistArtwork `json:"artwork,omitempty"`

	// blocknumber
	// Required: true
	Blocknumber *int64 `json:"blocknumber"`

	// cover art
	CoverArt string `json:"cover_art,omitempty"`

	// cover art cids
	CoverArtCids *PlaylistArtwork `json:"cover_art_cids,omitempty"`

	// cover art sizes
	CoverArtSizes string `json:"cover_art_sizes,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// favorite count
	// Required: true
	FavoriteCount *int64 `json:"favorite_count"`

	// followee favorites
	// Required: true
	FolloweeFavorites []*Favorite `json:"followee_favorites"`

	// followee reposts
	// Required: true
	FolloweeReposts []*Repost `json:"followee_reposts"`

	// has current user reposted
	// Required: true
	HasCurrentUserReposted *bool `json:"has_current_user_reposted"`

	// has current user saved
	// Required: true
	HasCurrentUserSaved *bool `json:"has_current_user_saved"`

	// id
	// Required: true
	ID *string `json:"id"`

	// is album
	// Required: true
	IsAlbum *bool `json:"is_album"`

	// is delete
	// Required: true
	IsDelete *bool `json:"is_delete"`

	// is image autogenerated
	// Required: true
	IsImageAutogenerated *bool `json:"is_image_autogenerated"`

	// is private
	// Required: true
	IsPrivate *bool `json:"is_private"`

	// permalink
	Permalink string `json:"permalink,omitempty"`

	// playlist contents
	// Required: true
	PlaylistContents []*PlaylistAddedTimestamp `json:"playlist_contents"`

	// playlist name
	// Required: true
	PlaylistName *string `json:"playlist_name"`

	// repost count
	// Required: true
	RepostCount *int64 `json:"repost_count"`

	// total play count
	// Required: true
	TotalPlayCount *int64 `json:"total_play_count"`

	// track count
	// Required: true
	TrackCount *int64 `json:"track_count"`

	// tracks
	// Required: true
	Tracks []*TrackFull `json:"tracks"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// user
	// Required: true
	User *UserFull `json:"user"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this playlist full
func (m *PlaylistFull) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddedTimestamps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocknumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoverArtCids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFavoriteCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolloweeFavorites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolloweeReposts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasCurrentUserReposted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasCurrentUserSaved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsAlbum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsImageAutogenerated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPrivate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlaylistContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlaylistName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepostCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalPlayCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTracks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlaylistFull) validateAddedTimestamps(formats strfmt.Registry) error {

	if err := validate.Required("added_timestamps", "body", m.AddedTimestamps); err != nil {
		return err
	}

	for i := 0; i < len(m.AddedTimestamps); i++ {
		if swag.IsZero(m.AddedTimestamps[i]) { // not required
			continue
		}

		if m.AddedTimestamps[i] != nil {
			if err := m.AddedTimestamps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("added_timestamps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("added_timestamps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) validateArtwork(formats strfmt.Registry) error {
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

func (m *PlaylistFull) validateBlocknumber(formats strfmt.Registry) error {

	if err := validate.Required("blocknumber", "body", m.Blocknumber); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateCoverArtCids(formats strfmt.Registry) error {
	if swag.IsZero(m.CoverArtCids) { // not required
		return nil
	}

	if m.CoverArtCids != nil {
		if err := m.CoverArtCids.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cover_art_cids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cover_art_cids")
			}
			return err
		}
	}

	return nil
}

func (m *PlaylistFull) validateFavoriteCount(formats strfmt.Registry) error {

	if err := validate.Required("favorite_count", "body", m.FavoriteCount); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateFolloweeFavorites(formats strfmt.Registry) error {

	if err := validate.Required("followee_favorites", "body", m.FolloweeFavorites); err != nil {
		return err
	}

	for i := 0; i < len(m.FolloweeFavorites); i++ {
		if swag.IsZero(m.FolloweeFavorites[i]) { // not required
			continue
		}

		if m.FolloweeFavorites[i] != nil {
			if err := m.FolloweeFavorites[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("followee_favorites" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("followee_favorites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) validateFolloweeReposts(formats strfmt.Registry) error {

	if err := validate.Required("followee_reposts", "body", m.FolloweeReposts); err != nil {
		return err
	}

	for i := 0; i < len(m.FolloweeReposts); i++ {
		if swag.IsZero(m.FolloweeReposts[i]) { // not required
			continue
		}

		if m.FolloweeReposts[i] != nil {
			if err := m.FolloweeReposts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("followee_reposts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("followee_reposts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) validateHasCurrentUserReposted(formats strfmt.Registry) error {

	if err := validate.Required("has_current_user_reposted", "body", m.HasCurrentUserReposted); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateHasCurrentUserSaved(formats strfmt.Registry) error {

	if err := validate.Required("has_current_user_saved", "body", m.HasCurrentUserSaved); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateIsAlbum(formats strfmt.Registry) error {

	if err := validate.Required("is_album", "body", m.IsAlbum); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateIsDelete(formats strfmt.Registry) error {

	if err := validate.Required("is_delete", "body", m.IsDelete); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateIsImageAutogenerated(formats strfmt.Registry) error {

	if err := validate.Required("is_image_autogenerated", "body", m.IsImageAutogenerated); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateIsPrivate(formats strfmt.Registry) error {

	if err := validate.Required("is_private", "body", m.IsPrivate); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validatePlaylistContents(formats strfmt.Registry) error {

	if err := validate.Required("playlist_contents", "body", m.PlaylistContents); err != nil {
		return err
	}

	for i := 0; i < len(m.PlaylistContents); i++ {
		if swag.IsZero(m.PlaylistContents[i]) { // not required
			continue
		}

		if m.PlaylistContents[i] != nil {
			if err := m.PlaylistContents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("playlist_contents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("playlist_contents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) validatePlaylistName(formats strfmt.Registry) error {

	if err := validate.Required("playlist_name", "body", m.PlaylistName); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateRepostCount(formats strfmt.Registry) error {

	if err := validate.Required("repost_count", "body", m.RepostCount); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateTotalPlayCount(formats strfmt.Registry) error {

	if err := validate.Required("total_play_count", "body", m.TotalPlayCount); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateTrackCount(formats strfmt.Registry) error {

	if err := validate.Required("track_count", "body", m.TrackCount); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistFull) validateTracks(formats strfmt.Registry) error {

	if err := validate.Required("tracks", "body", m.Tracks); err != nil {
		return err
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

func (m *PlaylistFull) validateUser(formats strfmt.Registry) error {

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

func (m *PlaylistFull) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this playlist full based on the context it is used
func (m *PlaylistFull) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddedTimestamps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArtwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoverArtCids(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFolloweeFavorites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFolloweeReposts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlaylistContents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTracks(ctx, formats); err != nil {
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

func (m *PlaylistFull) contextValidateAddedTimestamps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AddedTimestamps); i++ {

		if m.AddedTimestamps[i] != nil {

			if swag.IsZero(m.AddedTimestamps[i]) { // not required
				return nil
			}

			if err := m.AddedTimestamps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("added_timestamps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("added_timestamps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) contextValidateArtwork(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PlaylistFull) contextValidateCoverArtCids(ctx context.Context, formats strfmt.Registry) error {

	if m.CoverArtCids != nil {

		if swag.IsZero(m.CoverArtCids) { // not required
			return nil
		}

		if err := m.CoverArtCids.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cover_art_cids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cover_art_cids")
			}
			return err
		}
	}

	return nil
}

func (m *PlaylistFull) contextValidateFolloweeFavorites(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FolloweeFavorites); i++ {

		if m.FolloweeFavorites[i] != nil {

			if swag.IsZero(m.FolloweeFavorites[i]) { // not required
				return nil
			}

			if err := m.FolloweeFavorites[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("followee_favorites" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("followee_favorites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) contextValidateFolloweeReposts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FolloweeReposts); i++ {

		if m.FolloweeReposts[i] != nil {

			if swag.IsZero(m.FolloweeReposts[i]) { // not required
				return nil
			}

			if err := m.FolloweeReposts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("followee_reposts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("followee_reposts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) contextValidatePlaylistContents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlaylistContents); i++ {

		if m.PlaylistContents[i] != nil {

			if swag.IsZero(m.PlaylistContents[i]) { // not required
				return nil
			}

			if err := m.PlaylistContents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("playlist_contents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("playlist_contents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistFull) contextValidateTracks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PlaylistFull) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PlaylistFull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlaylistFull) UnmarshalBinary(b []byte) error {
	var res PlaylistFull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}