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

// MessagesGetWebPagePreviewRequest represents TL type `messages.getWebPagePreview#8b68b0cc`.
// Get preview of webpage
//
// See https://core.telegram.org/method/messages.getWebPagePreview for reference.
type MessagesGetWebPagePreviewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Message from which to extract the preview
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// MessagesGetWebPagePreviewRequestTypeID is TL type id of MessagesGetWebPagePreviewRequest.
const MessagesGetWebPagePreviewRequestTypeID = 0x8b68b0cc

// Ensuring interfaces in compile-time for MessagesGetWebPagePreviewRequest.
var (
	_ bin.Encoder     = &MessagesGetWebPagePreviewRequest{}
	_ bin.Decoder     = &MessagesGetWebPagePreviewRequest{}
	_ bin.BareEncoder = &MessagesGetWebPagePreviewRequest{}
	_ bin.BareDecoder = &MessagesGetWebPagePreviewRequest{}
)

func (g *MessagesGetWebPagePreviewRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Message == "") {
		return false
	}
	if !(g.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetWebPagePreviewRequest) String() string {
	if g == nil {
		return "MessagesGetWebPagePreviewRequest(nil)"
	}
	type Alias MessagesGetWebPagePreviewRequest
	return fmt.Sprintf("MessagesGetWebPagePreviewRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetWebPagePreviewRequest from given interface.
func (g *MessagesGetWebPagePreviewRequest) FillFrom(from interface {
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
}) {
	g.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		g.Entities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetWebPagePreviewRequest) TypeID() uint32 {
	return MessagesGetWebPagePreviewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetWebPagePreviewRequest) TypeName() string {
	return "messages.getWebPagePreview"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetWebPagePreviewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getWebPagePreview",
		ID:   MessagesGetWebPagePreviewRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !g.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetWebPagePreviewRequest) SetFlags() {
	if !(g.Entities == nil) {
		g.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetWebPagePreviewRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPagePreview#8b68b0cc as nil")
	}
	b.PutID(MessagesGetWebPagePreviewRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetWebPagePreviewRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPagePreview#8b68b0cc as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field flags: %w", err)
	}
	b.PutString(g.Message)
	if g.Flags.Has(3) {
		b.PutVectorHeader(len(g.Entities))
		for idx, v := range g.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetWebPagePreviewRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPagePreview#8b68b0cc to nil")
	}
	if err := b.ConsumeID(MessagesGetWebPagePreviewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetWebPagePreviewRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPagePreview#8b68b0cc to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field message: %w", err)
		}
		g.Message = value
	}
	if g.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field entities: %w", err)
		}

		if headerLen > 0 {
			g.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field entities: %w", err)
			}
			g.Entities = append(g.Entities, value)
		}
	}
	return nil
}

// GetMessage returns value of Message field.
func (g *MessagesGetWebPagePreviewRequest) GetMessage() (value string) {
	return g.Message
}

// SetEntities sets value of Entities conditional field.
func (g *MessagesGetWebPagePreviewRequest) SetEntities(value []MessageEntityClass) {
	g.Flags.Set(3)
	g.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (g *MessagesGetWebPagePreviewRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.Entities, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (g *MessagesGetWebPagePreviewRequest) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(g.Entities), true
}

// MessagesGetWebPagePreview invokes method messages.getWebPagePreview#8b68b0cc returning error if any.
// Get preview of webpage
//
// Possible errors:
//  400 MESSAGE_EMPTY: The provided message is empty.
//
// See https://core.telegram.org/method/messages.getWebPagePreview for reference.
func (c *Client) MessagesGetWebPagePreview(ctx context.Context, request *MessagesGetWebPagePreviewRequest) (MessageMediaClass, error) {
	var result MessageMediaBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.MessageMedia, nil
}
