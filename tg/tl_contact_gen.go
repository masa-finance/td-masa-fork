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

// Contact represents TL type `contact#f911c994`.
// A contact of the current user that is registered in the system.
//
// See https://core.telegram.org/constructor/contact for reference.
type Contact struct {
	// User identifier
	UserID int
	// Current user is in the user's contact list
	Mutual bool
}

// ContactTypeID is TL type id of Contact.
const ContactTypeID = 0xf911c994

func (c *Contact) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Mutual == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Contact) String() string {
	if c == nil {
		return "Contact(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Contact")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(c.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tMutual: ")
	sb.WriteString(fmt.Sprint(c.Mutual))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *Contact) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contact#f911c994 as nil")
	}
	b.PutID(ContactTypeID)
	b.PutInt(c.UserID)
	b.PutBool(c.Mutual)
	return nil
}

// Decode implements bin.Decoder.
func (c *Contact) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contact#f911c994 to nil")
	}
	if err := b.ConsumeID(ContactTypeID); err != nil {
		return fmt.Errorf("unable to decode contact#f911c994: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contact#f911c994: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode contact#f911c994: field mutual: %w", err)
		}
		c.Mutual = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Contact.
var (
	_ bin.Encoder = &Contact{}
	_ bin.Decoder = &Contact{}
)
