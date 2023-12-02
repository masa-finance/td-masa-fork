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

// TextQuote represents TL type `textQuote#8675b8b2`.
type TextQuote struct {
	// Text of the quote. Only Bold, Italic, Underline, Strikethrough, Spoiler, and
	// CustomEmoji entities can be present in the text
	Text FormattedText
	// Approximate quote position in the original message in UTF-16 code units
	Position int32
	// True, if the quote was manually chosen by the message sender
	IsManual bool
}

// TextQuoteTypeID is TL type id of TextQuote.
const TextQuoteTypeID = 0x8675b8b2

// Ensuring interfaces in compile-time for TextQuote.
var (
	_ bin.Encoder     = &TextQuote{}
	_ bin.Decoder     = &TextQuote{}
	_ bin.BareEncoder = &TextQuote{}
	_ bin.BareDecoder = &TextQuote{}
)

func (t *TextQuote) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Text.Zero()) {
		return false
	}
	if !(t.Position == 0) {
		return false
	}
	if !(t.IsManual == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TextQuote) String() string {
	if t == nil {
		return "TextQuote(nil)"
	}
	type Alias TextQuote
	return fmt.Sprintf("TextQuote%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TextQuote) TypeID() uint32 {
	return TextQuoteTypeID
}

// TypeName returns name of type in TL schema.
func (*TextQuote) TypeName() string {
	return "textQuote"
}

// TypeInfo returns info about TL type.
func (t *TextQuote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "textQuote",
		ID:   TextQuoteTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Position",
			SchemaName: "position",
		},
		{
			Name:       "IsManual",
			SchemaName: "is_manual",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TextQuote) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textQuote#8675b8b2 as nil")
	}
	b.PutID(TextQuoteTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TextQuote) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textQuote#8675b8b2 as nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textQuote#8675b8b2: field text: %w", err)
	}
	b.PutInt32(t.Position)
	b.PutBool(t.IsManual)
	return nil
}

// Decode implements bin.Decoder.
func (t *TextQuote) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textQuote#8675b8b2 to nil")
	}
	if err := b.ConsumeID(TextQuoteTypeID); err != nil {
		return fmt.Errorf("unable to decode textQuote#8675b8b2: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TextQuote) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textQuote#8675b8b2 to nil")
	}
	{
		if err := t.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode textQuote#8675b8b2: field text: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode textQuote#8675b8b2: field position: %w", err)
		}
		t.Position = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode textQuote#8675b8b2: field is_manual: %w", err)
		}
		t.IsManual = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TextQuote) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode textQuote#8675b8b2 as nil")
	}
	b.ObjStart()
	b.PutID("textQuote")
	b.Comma()
	b.FieldStart("text")
	if err := t.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode textQuote#8675b8b2: field text: %w", err)
	}
	b.Comma()
	b.FieldStart("position")
	b.PutInt32(t.Position)
	b.Comma()
	b.FieldStart("is_manual")
	b.PutBool(t.IsManual)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TextQuote) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode textQuote#8675b8b2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("textQuote"); err != nil {
				return fmt.Errorf("unable to decode textQuote#8675b8b2: %w", err)
			}
		case "text":
			if err := t.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode textQuote#8675b8b2: field text: %w", err)
			}
		case "position":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode textQuote#8675b8b2: field position: %w", err)
			}
			t.Position = value
		case "is_manual":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode textQuote#8675b8b2: field is_manual: %w", err)
			}
			t.IsManual = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (t *TextQuote) GetText() (value FormattedText) {
	if t == nil {
		return
	}
	return t.Text
}

// GetPosition returns value of Position field.
func (t *TextQuote) GetPosition() (value int32) {
	if t == nil {
		return
	}
	return t.Position
}

// GetIsManual returns value of IsManual field.
func (t *TextQuote) GetIsManual() (value bool) {
	if t == nil {
		return
	}
	return t.IsManual
}
