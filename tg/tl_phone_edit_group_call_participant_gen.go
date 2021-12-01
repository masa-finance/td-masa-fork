// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// PhoneEditGroupCallParticipantRequest represents TL type `phone.editGroupCallParticipant#a5273abf`.
// Edit information about a given group call participant
// Note: flags¹.N?Bool² parameters can have three possible values:
//
// Links:
//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
//  2) https://core.telegram.org/type/Bool
//
// See https://core.telegram.org/method/phone.editGroupCallParticipant for reference.
type PhoneEditGroupCallParticipantRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The group call
	Call InputGroupCall
	// The group call participant (can also be the user itself)
	Participant InputPeerClass
	// Whether to mute or unmute the specified participant
	//
	// Use SetMuted and GetMuted helpers.
	Muted bool
	// New volume
	//
	// Use SetVolume and GetVolume helpers.
	Volume int
	// Raise or lower hand
	//
	// Use SetRaiseHand and GetRaiseHand helpers.
	RaiseHand bool
	// Start or stop the video stream
	//
	// Use SetVideoStopped and GetVideoStopped helpers.
	VideoStopped bool
	// Pause or resume the video stream
	//
	// Use SetVideoPaused and GetVideoPaused helpers.
	VideoPaused bool
	// Pause or resume the screen sharing stream
	//
	// Use SetPresentationPaused and GetPresentationPaused helpers.
	PresentationPaused bool
}

// PhoneEditGroupCallParticipantRequestTypeID is TL type id of PhoneEditGroupCallParticipantRequest.
const PhoneEditGroupCallParticipantRequestTypeID = 0xa5273abf

// Ensuring interfaces in compile-time for PhoneEditGroupCallParticipantRequest.
var (
	_ bin.Encoder     = &PhoneEditGroupCallParticipantRequest{}
	_ bin.Decoder     = &PhoneEditGroupCallParticipantRequest{}
	_ bin.BareEncoder = &PhoneEditGroupCallParticipantRequest{}
	_ bin.BareDecoder = &PhoneEditGroupCallParticipantRequest{}
)

func (e *PhoneEditGroupCallParticipantRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Call.Zero()) {
		return false
	}
	if !(e.Participant == nil) {
		return false
	}
	if !(e.Muted == false) {
		return false
	}
	if !(e.Volume == 0) {
		return false
	}
	if !(e.RaiseHand == false) {
		return false
	}
	if !(e.VideoStopped == false) {
		return false
	}
	if !(e.VideoPaused == false) {
		return false
	}
	if !(e.PresentationPaused == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PhoneEditGroupCallParticipantRequest) String() string {
	if e == nil {
		return "PhoneEditGroupCallParticipantRequest(nil)"
	}
	type Alias PhoneEditGroupCallParticipantRequest
	return fmt.Sprintf("PhoneEditGroupCallParticipantRequest%+v", Alias(*e))
}

// FillFrom fills PhoneEditGroupCallParticipantRequest from given interface.
func (e *PhoneEditGroupCallParticipantRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetParticipant() (value InputPeerClass)
	GetMuted() (value bool, ok bool)
	GetVolume() (value int, ok bool)
	GetRaiseHand() (value bool, ok bool)
	GetVideoStopped() (value bool, ok bool)
	GetVideoPaused() (value bool, ok bool)
	GetPresentationPaused() (value bool, ok bool)
}) {
	e.Call = from.GetCall()
	e.Participant = from.GetParticipant()
	if val, ok := from.GetMuted(); ok {
		e.Muted = val
	}

	if val, ok := from.GetVolume(); ok {
		e.Volume = val
	}

	if val, ok := from.GetRaiseHand(); ok {
		e.RaiseHand = val
	}

	if val, ok := from.GetVideoStopped(); ok {
		e.VideoStopped = val
	}

	if val, ok := from.GetVideoPaused(); ok {
		e.VideoPaused = val
	}

	if val, ok := from.GetPresentationPaused(); ok {
		e.PresentationPaused = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneEditGroupCallParticipantRequest) TypeID() uint32 {
	return PhoneEditGroupCallParticipantRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneEditGroupCallParticipantRequest) TypeName() string {
	return "phone.editGroupCallParticipant"
}

// TypeInfo returns info about TL type.
func (e *PhoneEditGroupCallParticipantRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.editGroupCallParticipant",
		ID:   PhoneEditGroupCallParticipantRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Participant",
			SchemaName: "participant",
		},
		{
			Name:       "Muted",
			SchemaName: "muted",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "Volume",
			SchemaName: "volume",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "RaiseHand",
			SchemaName: "raise_hand",
			Null:       !e.Flags.Has(2),
		},
		{
			Name:       "VideoStopped",
			SchemaName: "video_stopped",
			Null:       !e.Flags.Has(3),
		},
		{
			Name:       "VideoPaused",
			SchemaName: "video_paused",
			Null:       !e.Flags.Has(4),
		},
		{
			Name:       "PresentationPaused",
			SchemaName: "presentation_paused",
			Null:       !e.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *PhoneEditGroupCallParticipantRequest) SetFlags() {
	if !(e.Muted == false) {
		e.Flags.Set(0)
	}
	if !(e.Volume == 0) {
		e.Flags.Set(1)
	}
	if !(e.RaiseHand == false) {
		e.Flags.Set(2)
	}
	if !(e.VideoStopped == false) {
		e.Flags.Set(3)
	}
	if !(e.VideoPaused == false) {
		e.Flags.Set(4)
	}
	if !(e.PresentationPaused == false) {
		e.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (e *PhoneEditGroupCallParticipantRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.editGroupCallParticipant#a5273abf as nil")
	}
	b.PutID(PhoneEditGroupCallParticipantRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PhoneEditGroupCallParticipantRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.editGroupCallParticipant#a5273abf as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallParticipant#a5273abf: field flags: %w", err)
	}
	if err := e.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallParticipant#a5273abf: field call: %w", err)
	}
	if e.Participant == nil {
		return fmt.Errorf("unable to encode phone.editGroupCallParticipant#a5273abf: field participant is nil")
	}
	if err := e.Participant.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallParticipant#a5273abf: field participant: %w", err)
	}
	if e.Flags.Has(0) {
		b.PutBool(e.Muted)
	}
	if e.Flags.Has(1) {
		b.PutInt(e.Volume)
	}
	if e.Flags.Has(2) {
		b.PutBool(e.RaiseHand)
	}
	if e.Flags.Has(3) {
		b.PutBool(e.VideoStopped)
	}
	if e.Flags.Has(4) {
		b.PutBool(e.VideoPaused)
	}
	if e.Flags.Has(5) {
		b.PutBool(e.PresentationPaused)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *PhoneEditGroupCallParticipantRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.editGroupCallParticipant#a5273abf to nil")
	}
	if err := b.ConsumeID(PhoneEditGroupCallParticipantRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PhoneEditGroupCallParticipantRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.editGroupCallParticipant#a5273abf to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field flags: %w", err)
		}
	}
	{
		if err := e.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field call: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field participant: %w", err)
		}
		e.Participant = value
	}
	if e.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field muted: %w", err)
		}
		e.Muted = value
	}
	if e.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field volume: %w", err)
		}
		e.Volume = value
	}
	if e.Flags.Has(2) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field raise_hand: %w", err)
		}
		e.RaiseHand = value
	}
	if e.Flags.Has(3) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field video_stopped: %w", err)
		}
		e.VideoStopped = value
	}
	if e.Flags.Has(4) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field video_paused: %w", err)
		}
		e.VideoPaused = value
	}
	if e.Flags.Has(5) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallParticipant#a5273abf: field presentation_paused: %w", err)
		}
		e.PresentationPaused = value
	}
	return nil
}

// GetCall returns value of Call field.
func (e *PhoneEditGroupCallParticipantRequest) GetCall() (value InputGroupCall) {
	return e.Call
}

// GetParticipant returns value of Participant field.
func (e *PhoneEditGroupCallParticipantRequest) GetParticipant() (value InputPeerClass) {
	return e.Participant
}

// SetMuted sets value of Muted conditional field.
func (e *PhoneEditGroupCallParticipantRequest) SetMuted(value bool) {
	e.Flags.Set(0)
	e.Muted = value
}

// GetMuted returns value of Muted conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallParticipantRequest) GetMuted() (value bool, ok bool) {
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.Muted, true
}

// SetVolume sets value of Volume conditional field.
func (e *PhoneEditGroupCallParticipantRequest) SetVolume(value int) {
	e.Flags.Set(1)
	e.Volume = value
}

// GetVolume returns value of Volume conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallParticipantRequest) GetVolume() (value int, ok bool) {
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.Volume, true
}

// SetRaiseHand sets value of RaiseHand conditional field.
func (e *PhoneEditGroupCallParticipantRequest) SetRaiseHand(value bool) {
	e.Flags.Set(2)
	e.RaiseHand = value
}

// GetRaiseHand returns value of RaiseHand conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallParticipantRequest) GetRaiseHand() (value bool, ok bool) {
	if !e.Flags.Has(2) {
		return value, false
	}
	return e.RaiseHand, true
}

// SetVideoStopped sets value of VideoStopped conditional field.
func (e *PhoneEditGroupCallParticipantRequest) SetVideoStopped(value bool) {
	e.Flags.Set(3)
	e.VideoStopped = value
}

// GetVideoStopped returns value of VideoStopped conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallParticipantRequest) GetVideoStopped() (value bool, ok bool) {
	if !e.Flags.Has(3) {
		return value, false
	}
	return e.VideoStopped, true
}

// SetVideoPaused sets value of VideoPaused conditional field.
func (e *PhoneEditGroupCallParticipantRequest) SetVideoPaused(value bool) {
	e.Flags.Set(4)
	e.VideoPaused = value
}

// GetVideoPaused returns value of VideoPaused conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallParticipantRequest) GetVideoPaused() (value bool, ok bool) {
	if !e.Flags.Has(4) {
		return value, false
	}
	return e.VideoPaused, true
}

// SetPresentationPaused sets value of PresentationPaused conditional field.
func (e *PhoneEditGroupCallParticipantRequest) SetPresentationPaused(value bool) {
	e.Flags.Set(5)
	e.PresentationPaused = value
}

// GetPresentationPaused returns value of PresentationPaused conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallParticipantRequest) GetPresentationPaused() (value bool, ok bool) {
	if !e.Flags.Has(5) {
		return value, false
	}
	return e.PresentationPaused, true
}

// PhoneEditGroupCallParticipant invokes method phone.editGroupCallParticipant#a5273abf returning error if any.
// Edit information about a given group call participant
// Note: flags¹.N?Bool² parameters can have three possible values:
//
// Links:
//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
//  2) https://core.telegram.org/type/Bool
//
// Possible errors:
//  400 USER_VOLUME_INVALID: The specified user volume is invalid.
//
// See https://core.telegram.org/method/phone.editGroupCallParticipant for reference.
func (c *Client) PhoneEditGroupCallParticipant(ctx context.Context, request *PhoneEditGroupCallParticipantRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
