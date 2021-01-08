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

// MessagesGetFeaturedStickersRequest represents TL type `messages.getFeaturedStickers#2dacca4f`.
// Get featured stickers
//
// See https://core.telegram.org/method/messages.getFeaturedStickers for reference.
type MessagesGetFeaturedStickersRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetFeaturedStickersRequestTypeID is TL type id of MessagesGetFeaturedStickersRequest.
const MessagesGetFeaturedStickersRequestTypeID = 0x2dacca4f

func (g *MessagesGetFeaturedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetFeaturedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetFeaturedStickersRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetFeaturedStickersRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(g.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetFeaturedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFeaturedStickers#2dacca4f as nil")
	}
	b.PutID(MessagesGetFeaturedStickersRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetFeaturedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFeaturedStickers#2dacca4f to nil")
	}
	if err := b.ConsumeID(MessagesGetFeaturedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getFeaturedStickers#2dacca4f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getFeaturedStickers#2dacca4f: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetFeaturedStickersRequest.
var (
	_ bin.Encoder = &MessagesGetFeaturedStickersRequest{}
	_ bin.Decoder = &MessagesGetFeaturedStickersRequest{}
)

// MessagesGetFeaturedStickers invokes method messages.getFeaturedStickers#2dacca4f returning error if any.
// Get featured stickers
//
// See https://core.telegram.org/method/messages.getFeaturedStickers for reference.
func (c *Client) MessagesGetFeaturedStickers(ctx context.Context, hash int) (MessagesFeaturedStickersClass, error) {
	var result MessagesFeaturedStickersBox

	request := &MessagesGetFeaturedStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FeaturedStickers, nil
}
