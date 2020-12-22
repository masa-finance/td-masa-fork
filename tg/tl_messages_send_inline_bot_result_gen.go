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

// MessagesSendInlineBotResultRequest represents TL type `messages.sendInlineBotResult#220815b0`.
// Send a result obtained using messages.getInlineBotResults.
//
// See https://core.telegram.org/method/messages.sendInlineBotResult for reference.
type MessagesSendInlineBotResultRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether to send the message silently (no notification will be triggered on the other client)
	Silent bool
	// Whether to send the message in background
	Background bool
	// Whether to clear the draft
	ClearDraft bool
	// Whether to hide the via @botname in the resulting message (only for bot usernames encountered in the config)
	HideVia bool
	// Destination
	Peer InputPeerClass
	// ID of the message this message should reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// Random ID to avoid resending the same query
	RandomID int64
	// Query ID from messages.getInlineBotResults
	QueryID int64
	// Result ID from messages.getInlineBotResults
	ID string
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// MessagesSendInlineBotResultRequestTypeID is TL type id of MessagesSendInlineBotResultRequest.
const MessagesSendInlineBotResultRequestTypeID = 0x220815b0

// String implements fmt.Stringer.
func (s *MessagesSendInlineBotResultRequest) String() string {
	if s == nil {
		return "MessagesSendInlineBotResultRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSendInlineBotResultRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(s.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(s.Peer.String())
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tReplyToMsgID: ")
		sb.WriteString(fmt.Sprint(s.ReplyToMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(s.RandomID))
	sb.WriteString(",\n")
	sb.WriteString("\tQueryID: ")
	sb.WriteString(fmt.Sprint(s.QueryID))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(s.ID))
	sb.WriteString(",\n")
	if s.Flags.Has(10) {
		sb.WriteString("\tScheduleDate: ")
		sb.WriteString(fmt.Sprint(s.ScheduleDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSendInlineBotResultRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendInlineBotResult#220815b0 as nil")
	}
	b.PutID(MessagesSendInlineBotResultRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendInlineBotResult#220815b0: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	b.PutLong(s.RandomID)
	b.PutLong(s.QueryID)
	b.PutString(s.ID)
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendInlineBotResultRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
	} else {
		s.Flags.Unset(5)
	}
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendInlineBotResultRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
	} else {
		s.Flags.Unset(6)
	}
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendInlineBotResultRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
	} else {
		s.Flags.Unset(7)
	}
}

// SetHideVia sets value of HideVia conditional field.
func (s *MessagesSendInlineBotResultRequest) SetHideVia(value bool) {
	if value {
		s.Flags.Set(11)
	} else {
		s.Flags.Unset(11)
	}
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendInlineBotResultRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendInlineBotResultRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendInlineBotResultRequest) GetScheduleDate() (value int, ok bool) {
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (s *MessagesSendInlineBotResultRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendInlineBotResult#220815b0 to nil")
	}
	if err := b.ConsumeID(MessagesSendInlineBotResultRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.HideVia = s.Flags.Has(11)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field id: %w", err)
		}
		s.ID = value
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendInlineBotResult#220815b0: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendInlineBotResultRequest.
var (
	_ bin.Encoder = &MessagesSendInlineBotResultRequest{}
	_ bin.Decoder = &MessagesSendInlineBotResultRequest{}
)

// MessagesSendInlineBotResult invokes method messages.sendInlineBotResult#220815b0 returning error if any.
// Send a result obtained using messages.getInlineBotResults.
//
// See https://core.telegram.org/method/messages.sendInlineBotResult for reference.
func (c *Client) MessagesSendInlineBotResult(ctx context.Context, request *MessagesSendInlineBotResultRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
