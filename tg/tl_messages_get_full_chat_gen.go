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

// MessagesGetFullChatRequest represents TL type `messages.getFullChat#3b831c66`.
// Returns full chat info according to its ID.
//
// See https://core.telegram.org/method/messages.getFullChat for reference.
type MessagesGetFullChatRequest struct {
	// Chat ID
	ChatID int
}

// MessagesGetFullChatRequestTypeID is TL type id of MessagesGetFullChatRequest.
const MessagesGetFullChatRequestTypeID = 0x3b831c66

// String implements fmt.Stringer.
func (g *MessagesGetFullChatRequest) String() string {
	if g == nil {
		return "MessagesGetFullChatRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetFullChatRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(g.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetFullChatRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFullChat#3b831c66 as nil")
	}
	b.PutID(MessagesGetFullChatRequestTypeID)
	b.PutInt(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetFullChatRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFullChat#3b831c66 to nil")
	}
	if err := b.ConsumeID(MessagesGetFullChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getFullChat#3b831c66: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getFullChat#3b831c66: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetFullChatRequest.
var (
	_ bin.Encoder = &MessagesGetFullChatRequest{}
	_ bin.Decoder = &MessagesGetFullChatRequest{}
)

// MessagesGetFullChat invokes method messages.getFullChat#3b831c66 returning error if any.
// Returns full chat info according to its ID.
//
// See https://core.telegram.org/method/messages.getFullChat for reference.
func (c *Client) MessagesGetFullChat(ctx context.Context, chatid int) (*MessagesChatFull, error) {
	var result MessagesChatFull

	request := &MessagesGetFullChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
