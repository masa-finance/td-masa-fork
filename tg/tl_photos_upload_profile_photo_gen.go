// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// PhotosUploadProfilePhotoRequest represents TL type `photos.uploadProfilePhoto#89f30f69`.
// Updates current user profile photo.
//
// See https://core.telegram.org/method/photos.uploadProfilePhoto for reference.
type PhotosUploadProfilePhotoRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// File saved in parts by means of upload.saveFilePart method
	//
	// Use SetFile and GetFile helpers.
	File InputFileClass
	// Animated profile picture video
	//
	// Use SetVideo and GetVideo helpers.
	Video InputFileClass
	// Floating point UNIX timestamp in seconds, indicating the frame of the video that should be used as static preview.
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// PhotosUploadProfilePhotoRequestTypeID is TL type id of PhotosUploadProfilePhotoRequest.
const PhotosUploadProfilePhotoRequestTypeID = 0x89f30f69

// String implements fmt.Stringer.
func (u *PhotosUploadProfilePhotoRequest) String() string {
	if u == nil {
		return "PhotosUploadProfilePhotoRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhotosUploadProfilePhotoRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(u.Flags.String())
	sb.WriteString(",\n")
	if u.Flags.Has(0) {
		sb.WriteString("\tFile: ")
		sb.WriteString(u.File.String())
		sb.WriteString(",\n")
	}
	if u.Flags.Has(1) {
		sb.WriteString("\tVideo: ")
		sb.WriteString(u.Video.String())
		sb.WriteString(",\n")
	}
	if u.Flags.Has(2) {
		sb.WriteString("\tVideoStartTs: ")
		sb.WriteString(fmt.Sprint(u.VideoStartTs))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *PhotosUploadProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.uploadProfilePhoto#89f30f69 as nil")
	}
	b.PutID(PhotosUploadProfilePhotoRequestTypeID)
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.uploadProfilePhoto#89f30f69: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		if u.File == nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#89f30f69: field file is nil")
		}
		if err := u.File.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#89f30f69: field file: %w", err)
		}
	}
	if u.Flags.Has(1) {
		if u.Video == nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#89f30f69: field video is nil")
		}
		if err := u.Video.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#89f30f69: field video: %w", err)
		}
	}
	if u.Flags.Has(2) {
		b.PutDouble(u.VideoStartTs)
	}
	return nil
}

// SetFile sets value of File conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetFile(value InputFileClass) {
	u.Flags.Set(0)
	u.File = value
}

// GetFile returns value of File conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetFile() (value InputFileClass, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.File, true
}

// SetVideo sets value of Video conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetVideo(value InputFileClass) {
	u.Flags.Set(1)
	u.Video = value
}

// GetVideo returns value of Video conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetVideo() (value InputFileClass, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.Video, true
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetVideoStartTs(value float64) {
	u.Flags.Set(2)
	u.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetVideoStartTs() (value float64, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.VideoStartTs, true
}

// Decode implements bin.Decoder.
func (u *PhotosUploadProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.uploadProfilePhoto#89f30f69 to nil")
	}
	if err := b.ConsumeID(PhotosUploadProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.uploadProfilePhoto#89f30f69: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#89f30f69: field flags: %w", err)
		}
	}
	if u.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#89f30f69: field file: %w", err)
		}
		u.File = value
	}
	if u.Flags.Has(1) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#89f30f69: field video: %w", err)
		}
		u.Video = value
	}
	if u.Flags.Has(2) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#89f30f69: field video_start_ts: %w", err)
		}
		u.VideoStartTs = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhotosUploadProfilePhotoRequest.
var (
	_ bin.Encoder = &PhotosUploadProfilePhotoRequest{}
	_ bin.Decoder = &PhotosUploadProfilePhotoRequest{}
)

// PhotosUploadProfilePhoto invokes method photos.uploadProfilePhoto#89f30f69 returning error if any.
// Updates current user profile photo.
//
// See https://core.telegram.org/method/photos.uploadProfilePhoto for reference.
func (c *Client) PhotosUploadProfilePhoto(ctx context.Context, request *PhotosUploadProfilePhotoRequest) (*PhotosPhoto, error) {
	var result PhotosPhoto

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
