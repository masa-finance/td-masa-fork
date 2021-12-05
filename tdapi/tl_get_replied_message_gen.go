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

// GetRepliedMessageRequest represents TL type `getRepliedMessage#d9bd19bd`.
type GetRepliedMessageRequest struct {
	// Identifier of the chat the message belongs to
	ChatID int64
	// Identifier of the reply message
	MessageID int64
}

// GetRepliedMessageRequestTypeID is TL type id of GetRepliedMessageRequest.
const GetRepliedMessageRequestTypeID = 0xd9bd19bd

// Ensuring interfaces in compile-time for GetRepliedMessageRequest.
var (
	_ bin.Encoder     = &GetRepliedMessageRequest{}
	_ bin.Decoder     = &GetRepliedMessageRequest{}
	_ bin.BareEncoder = &GetRepliedMessageRequest{}
	_ bin.BareDecoder = &GetRepliedMessageRequest{}
)

func (g *GetRepliedMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetRepliedMessageRequest) String() string {
	if g == nil {
		return "GetRepliedMessageRequest(nil)"
	}
	type Alias GetRepliedMessageRequest
	return fmt.Sprintf("GetRepliedMessageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetRepliedMessageRequest) TypeID() uint32 {
	return GetRepliedMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetRepliedMessageRequest) TypeName() string {
	return "getRepliedMessage"
}

// TypeInfo returns info about TL type.
func (g *GetRepliedMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getRepliedMessage",
		ID:   GetRepliedMessageRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetRepliedMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRepliedMessage#d9bd19bd as nil")
	}
	b.PutID(GetRepliedMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetRepliedMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRepliedMessage#d9bd19bd as nil")
	}
	b.PutLong(g.ChatID)
	b.PutLong(g.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetRepliedMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRepliedMessage#d9bd19bd to nil")
	}
	if err := b.ConsumeID(GetRepliedMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getRepliedMessage#d9bd19bd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetRepliedMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRepliedMessage#d9bd19bd to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getRepliedMessage#d9bd19bd: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getRepliedMessage#d9bd19bd: field message_id: %w", err)
		}
		g.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetRepliedMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getRepliedMessage#d9bd19bd as nil")
	}
	b.ObjStart()
	b.PutID("getRepliedMessage")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_id")
	b.PutLong(g.MessageID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetRepliedMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getRepliedMessage#d9bd19bd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getRepliedMessage"); err != nil {
				return fmt.Errorf("unable to decode getRepliedMessage#d9bd19bd: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getRepliedMessage#d9bd19bd: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getRepliedMessage#d9bd19bd: field message_id: %w", err)
			}
			g.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetRepliedMessageRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetRepliedMessageRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetRepliedMessage invokes method getRepliedMessage#d9bd19bd returning error if any.
func (c *Client) GetRepliedMessage(ctx context.Context, request *GetRepliedMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}