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

// MessagesGetDialogUnreadMarksRequest represents TL type `messages.getDialogUnreadMarks#22e24e22`.
// Get dialogs manually marked as unread
//
// See https://core.telegram.org/method/messages.getDialogUnreadMarks for reference.
type MessagesGetDialogUnreadMarksRequest struct {
}

// MessagesGetDialogUnreadMarksRequestTypeID is TL type id of MessagesGetDialogUnreadMarksRequest.
const MessagesGetDialogUnreadMarksRequestTypeID = 0x22e24e22

// String implements fmt.Stringer.
func (g *MessagesGetDialogUnreadMarksRequest) String() string {
	if g == nil {
		return "MessagesGetDialogUnreadMarksRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetDialogUnreadMarksRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogUnreadMarksRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogUnreadMarks#22e24e22 as nil")
	}
	b.PutID(MessagesGetDialogUnreadMarksRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogUnreadMarksRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogUnreadMarks#22e24e22 to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogUnreadMarksRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogUnreadMarks#22e24e22: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDialogUnreadMarksRequest.
var (
	_ bin.Encoder = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.Decoder = &MessagesGetDialogUnreadMarksRequest{}
)

// MessagesGetDialogUnreadMarks invokes method messages.getDialogUnreadMarks#22e24e22 returning error if any.
// Get dialogs manually marked as unread
//
// See https://core.telegram.org/method/messages.getDialogUnreadMarks for reference.
func (c *Client) MessagesGetDialogUnreadMarks(ctx context.Context) ([]DialogPeerClass, error) {
	var result DialogPeerClassVector

	request := &MessagesGetDialogUnreadMarksRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
