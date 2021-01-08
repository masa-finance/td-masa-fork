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

// MessagesAffectedHistory represents TL type `messages.affectedHistory#b45c69d1`.
// Affected part of communication history with the user or in a chat.
//
// See https://core.telegram.org/constructor/messages.affectedHistory for reference.
type MessagesAffectedHistory struct {
	// Number of events occured in a text box
	Pts int
	// Number of affected events
	PtsCount int
	// If a parameter contains positive value, it is necessary to repeat the method call using the given value; during the proceeding of all the history the value itself shall gradually decrease
	Offset int
}

// MessagesAffectedHistoryTypeID is TL type id of MessagesAffectedHistory.
const MessagesAffectedHistoryTypeID = 0xb45c69d1

func (a *MessagesAffectedHistory) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Pts == 0) {
		return false
	}
	if !(a.PtsCount == 0) {
		return false
	}
	if !(a.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAffectedHistory) String() string {
	if a == nil {
		return "MessagesAffectedHistory(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesAffectedHistory")
	sb.WriteString("{\n")
	sb.WriteString("\tPts: ")
	sb.WriteString(fmt.Sprint(a.Pts))
	sb.WriteString(",\n")
	sb.WriteString("\tPtsCount: ")
	sb.WriteString(fmt.Sprint(a.PtsCount))
	sb.WriteString(",\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(a.Offset))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *MessagesAffectedHistory) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedHistory#b45c69d1 as nil")
	}
	b.PutID(MessagesAffectedHistoryTypeID)
	b.PutInt(a.Pts)
	b.PutInt(a.PtsCount)
	b.PutInt(a.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAffectedHistory) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedHistory#b45c69d1 to nil")
	}
	if err := b.ConsumeID(MessagesAffectedHistoryTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field pts: %w", err)
		}
		a.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field pts_count: %w", err)
		}
		a.PtsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field offset: %w", err)
		}
		a.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAffectedHistory.
var (
	_ bin.Encoder = &MessagesAffectedHistory{}
	_ bin.Decoder = &MessagesAffectedHistory{}
)
