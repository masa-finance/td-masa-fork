// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ToggleChatViewAsTopicsRequest represents TL type `toggleChatViewAsTopics#2b2783dc`.
type ToggleChatViewAsTopicsRequest struct {
	// Chat identifier
	ChatID int64
	// New value of view_as_topics
	ViewAsTopics bool
}

// ToggleChatViewAsTopicsRequestTypeID is TL type id of ToggleChatViewAsTopicsRequest.
const ToggleChatViewAsTopicsRequestTypeID = 0x2b2783dc

// Ensuring interfaces in compile-time for ToggleChatViewAsTopicsRequest.
var (
	_ bin.Encoder     = &ToggleChatViewAsTopicsRequest{}
	_ bin.Decoder     = &ToggleChatViewAsTopicsRequest{}
	_ bin.BareEncoder = &ToggleChatViewAsTopicsRequest{}
	_ bin.BareDecoder = &ToggleChatViewAsTopicsRequest{}
)

func (t *ToggleChatViewAsTopicsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.ViewAsTopics == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleChatViewAsTopicsRequest) String() string {
	if t == nil {
		return "ToggleChatViewAsTopicsRequest(nil)"
	}
	type Alias ToggleChatViewAsTopicsRequest
	return fmt.Sprintf("ToggleChatViewAsTopicsRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleChatViewAsTopicsRequest) TypeID() uint32 {
	return ToggleChatViewAsTopicsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleChatViewAsTopicsRequest) TypeName() string {
	return "toggleChatViewAsTopics"
}

// TypeInfo returns info about TL type.
func (t *ToggleChatViewAsTopicsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleChatViewAsTopics",
		ID:   ToggleChatViewAsTopicsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "ViewAsTopics",
			SchemaName: "view_as_topics",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleChatViewAsTopicsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatViewAsTopics#2b2783dc as nil")
	}
	b.PutID(ToggleChatViewAsTopicsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleChatViewAsTopicsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatViewAsTopics#2b2783dc as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutBool(t.ViewAsTopics)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleChatViewAsTopicsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatViewAsTopics#2b2783dc to nil")
	}
	if err := b.ConsumeID(ToggleChatViewAsTopicsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleChatViewAsTopics#2b2783dc: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleChatViewAsTopicsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatViewAsTopics#2b2783dc to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatViewAsTopics#2b2783dc: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatViewAsTopics#2b2783dc: field view_as_topics: %w", err)
		}
		t.ViewAsTopics = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleChatViewAsTopicsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatViewAsTopics#2b2783dc as nil")
	}
	b.ObjStart()
	b.PutID("toggleChatViewAsTopics")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("view_as_topics")
	b.PutBool(t.ViewAsTopics)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleChatViewAsTopicsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatViewAsTopics#2b2783dc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleChatViewAsTopics"); err != nil {
				return fmt.Errorf("unable to decode toggleChatViewAsTopics#2b2783dc: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatViewAsTopics#2b2783dc: field chat_id: %w", err)
			}
			t.ChatID = value
		case "view_as_topics":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatViewAsTopics#2b2783dc: field view_as_topics: %w", err)
			}
			t.ViewAsTopics = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleChatViewAsTopicsRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetViewAsTopics returns value of ViewAsTopics field.
func (t *ToggleChatViewAsTopicsRequest) GetViewAsTopics() (value bool) {
	if t == nil {
		return
	}
	return t.ViewAsTopics
}

// ToggleChatViewAsTopics invokes method toggleChatViewAsTopics#2b2783dc returning error if any.
func (c *Client) ToggleChatViewAsTopics(ctx context.Context, request *ToggleChatViewAsTopicsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
