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

// MessagesSendMultiMediaRequest represents TL type `messages.sendMultiMedia#cc0110cb`.
// Send an album or grouped media¹
//
// Links:
//  1) https://core.telegram.org/api/files#albums-grouped-media
//
// See https://core.telegram.org/method/messages.sendMultiMedia for reference.
type MessagesSendMultiMediaRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send the album silently (no notification triggered)
	Silent bool
	// Send in background?
	Background bool
	// Whether to clear drafts¹
	//
	// Links:
	//  1) https://core.telegram.org/api/drafts
	ClearDraft bool
	// The destination chat
	Peer InputPeerClass
	// The message to reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// The medias to send
	MultiMedia []InputSingleMedia
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// MessagesSendMultiMediaRequestTypeID is TL type id of MessagesSendMultiMediaRequest.
const MessagesSendMultiMediaRequestTypeID = 0xcc0110cb

func (s *MessagesSendMultiMediaRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.MultiMedia == nil) {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendMultiMediaRequest) String() string {
	if s == nil {
		return "MessagesSendMultiMediaRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSendMultiMediaRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(s.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tReplyToMsgID: ")
		sb.WriteString(fmt.Sprint(s.ReplyToMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteByte('[')
	for _, v := range s.MultiMedia {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	if s.Flags.Has(10) {
		sb.WriteString("\tScheduleDate: ")
		sb.WriteString(fmt.Sprint(s.ScheduleDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSendMultiMediaRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMultiMedia#cc0110cb as nil")
	}
	b.PutID(MessagesSendMultiMediaRequestTypeID)
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.ReplyToMsgID == 0) {
		s.Flags.Set(0)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#cc0110cb: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#cc0110cb: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#cc0110cb: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	b.PutVectorHeader(len(s.MultiMedia))
	for idx, v := range s.MultiMedia {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#cc0110cb: field multi_media element with index %d: %w", idx, err)
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMultiMediaRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMultiMediaRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMultiMediaRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendMultiMediaRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMultiMediaRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetScheduleDate() (value int, ok bool) {
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (s *MessagesSendMultiMediaRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMultiMedia#cc0110cb to nil")
	}
	if err := b.ConsumeID(MessagesSendMultiMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: field multi_media: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputSingleMedia
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: field multi_media: %w", err)
			}
			s.MultiMedia = append(s.MultiMedia, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#cc0110cb: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendMultiMediaRequest.
var (
	_ bin.Encoder = &MessagesSendMultiMediaRequest{}
	_ bin.Decoder = &MessagesSendMultiMediaRequest{}
)

// MessagesSendMultiMedia invokes method messages.sendMultiMedia#cc0110cb returning error if any.
// Send an album or grouped media¹
//
// Links:
//  1) https://core.telegram.org/api/files#albums-grouped-media
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 MEDIA_EMPTY: The provided media object is invalid
//  400 MEDIA_INVALID: Media invalid
//  400 MULTI_MEDIA_TOO_LONG: Too many media files for album
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 RANDOM_ID_EMPTY: Random ID empty
//
// See https://core.telegram.org/method/messages.sendMultiMedia for reference.
// Can be used by bots.
func (c *Client) MessagesSendMultiMedia(ctx context.Context, request *MessagesSendMultiMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
