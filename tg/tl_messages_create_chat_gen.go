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

// MessagesCreateChatRequest represents TL type `messages.createChat#9cb126e`.
// Creates a new chat.
//
// See https://core.telegram.org/method/messages.createChat for reference.
type MessagesCreateChatRequest struct {
	// List of user IDs to be invited
	Users []InputUserClass
	// Chat name
	Title string
}

// MessagesCreateChatRequestTypeID is TL type id of MessagesCreateChatRequest.
const MessagesCreateChatRequestTypeID = 0x9cb126e

// String implements fmt.Stringer.
func (c *MessagesCreateChatRequest) String() string {
	if c == nil {
		return "MessagesCreateChatRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesCreateChatRequest")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range c.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(c.Title))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *MessagesCreateChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.createChat#9cb126e as nil")
	}
	b.PutID(MessagesCreateChatRequestTypeID)
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.createChat#9cb126e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.createChat#9cb126e: field users element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.Title)
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesCreateChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.createChat#9cb126e to nil")
	}
	if err := b.ConsumeID(MessagesCreateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.createChat#9cb126e: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#9cb126e: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.createChat#9cb126e: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.createChat#9cb126e: field title: %w", err)
		}
		c.Title = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesCreateChatRequest.
var (
	_ bin.Encoder = &MessagesCreateChatRequest{}
	_ bin.Decoder = &MessagesCreateChatRequest{}
)

// MessagesCreateChat invokes method messages.createChat#9cb126e returning error if any.
// Creates a new chat.
//
// See https://core.telegram.org/method/messages.createChat for reference.
func (c *Client) MessagesCreateChat(ctx context.Context, request *MessagesCreateChatRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
