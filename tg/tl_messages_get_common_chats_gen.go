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

// MessagesGetCommonChatsRequest represents TL type `messages.getCommonChats#d0a48c4`.
// Get chats in common with a user
//
// See https://core.telegram.org/method/messages.getCommonChats for reference.
type MessagesGetCommonChatsRequest struct {
	// User ID
	UserID InputUserClass
	// Maximum ID of chat to return (see pagination)
	MaxID int
	// Maximum number of results to return, see pagination
	Limit int
}

// MessagesGetCommonChatsRequestTypeID is TL type id of MessagesGetCommonChatsRequest.
const MessagesGetCommonChatsRequestTypeID = 0xd0a48c4

// String implements fmt.Stringer.
func (g *MessagesGetCommonChatsRequest) String() string {
	if g == nil {
		return "MessagesGetCommonChatsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetCommonChatsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(g.UserID.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMaxID: ")
	sb.WriteString(fmt.Sprint(g.MaxID))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetCommonChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getCommonChats#d0a48c4 as nil")
	}
	b.PutID(MessagesGetCommonChatsRequestTypeID)
	if g.UserID == nil {
		return fmt.Errorf("unable to encode messages.getCommonChats#d0a48c4: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getCommonChats#d0a48c4: field user_id: %w", err)
	}
	b.PutInt(g.MaxID)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetCommonChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getCommonChats#d0a48c4 to nil")
	}
	if err := b.ConsumeID(MessagesGetCommonChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetCommonChatsRequest.
var (
	_ bin.Encoder = &MessagesGetCommonChatsRequest{}
	_ bin.Decoder = &MessagesGetCommonChatsRequest{}
)

// MessagesGetCommonChats invokes method messages.getCommonChats#d0a48c4 returning error if any.
// Get chats in common with a user
//
// See https://core.telegram.org/method/messages.getCommonChats for reference.
func (c *Client) MessagesGetCommonChats(ctx context.Context, request *MessagesGetCommonChatsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
