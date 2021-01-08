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

// ContactsGetBlockedRequest represents TL type `contacts.getBlocked#f57c350f`.
// Returns the list of blocked users.
//
// See https://core.telegram.org/method/contacts.getBlocked for reference.
type ContactsGetBlockedRequest struct {
	// The number of list elements to be skipped
	Offset int
	// The number of list elements to be returned
	Limit int
}

// ContactsGetBlockedRequestTypeID is TL type id of ContactsGetBlockedRequest.
const ContactsGetBlockedRequestTypeID = 0xf57c350f

func (g *ContactsGetBlockedRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetBlockedRequest) String() string {
	if g == nil {
		return "ContactsGetBlockedRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsGetBlockedRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ContactsGetBlockedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getBlocked#f57c350f as nil")
	}
	b.PutID(ContactsGetBlockedRequestTypeID)
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetBlockedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getBlocked#f57c350f to nil")
	}
	if err := b.ConsumeID(ContactsGetBlockedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getBlocked#f57c350f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getBlocked#f57c350f: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getBlocked#f57c350f: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetBlockedRequest.
var (
	_ bin.Encoder = &ContactsGetBlockedRequest{}
	_ bin.Decoder = &ContactsGetBlockedRequest{}
)

// ContactsGetBlocked invokes method contacts.getBlocked#f57c350f returning error if any.
// Returns the list of blocked users.
//
// See https://core.telegram.org/method/contacts.getBlocked for reference.
func (c *Client) ContactsGetBlocked(ctx context.Context, request *ContactsGetBlockedRequest) (ContactsBlockedClass, error) {
	var result ContactsBlockedBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Blocked, nil
}
