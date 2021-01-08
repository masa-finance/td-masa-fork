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

// EmojiKeyword represents TL type `emojiKeyword#d5b3b9f9`.
// Emoji keyword
//
// See https://core.telegram.org/constructor/emojiKeyword for reference.
type EmojiKeyword struct {
	// Keyword
	Keyword string
	// Emojis associated to keyword
	Emoticons []string
}

// EmojiKeywordTypeID is TL type id of EmojiKeyword.
const EmojiKeywordTypeID = 0xd5b3b9f9

func (e *EmojiKeyword) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Keyword == "") {
		return false
	}
	if !(e.Emoticons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiKeyword) String() string {
	if e == nil {
		return "EmojiKeyword(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EmojiKeyword")
	sb.WriteString("{\n")
	sb.WriteString("\tKeyword: ")
	sb.WriteString(fmt.Sprint(e.Keyword))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range e.Emoticons {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *EmojiKeyword) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeyword#d5b3b9f9 as nil")
	}
	b.PutID(EmojiKeywordTypeID)
	b.PutString(e.Keyword)
	b.PutVectorHeader(len(e.Emoticons))
	for _, v := range e.Emoticons {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiKeyword) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeyword#d5b3b9f9 to nil")
	}
	if err := b.ConsumeID(EmojiKeywordTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: field keyword: %w", err)
		}
		e.Keyword = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: field emoticons: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: field emoticons: %w", err)
			}
			e.Emoticons = append(e.Emoticons, value)
		}
	}
	return nil
}

// construct implements constructor of EmojiKeywordClass.
func (e EmojiKeyword) construct() EmojiKeywordClass { return &e }

// Ensuring interfaces in compile-time for EmojiKeyword.
var (
	_ bin.Encoder = &EmojiKeyword{}
	_ bin.Decoder = &EmojiKeyword{}

	_ EmojiKeywordClass = &EmojiKeyword{}
)

// EmojiKeywordDeleted represents TL type `emojiKeywordDeleted#236df622`.
// Deleted emoji keyword
//
// See https://core.telegram.org/constructor/emojiKeywordDeleted for reference.
type EmojiKeywordDeleted struct {
	// Keyword
	Keyword string
	// Emojis that were associated to keyword
	Emoticons []string
}

// EmojiKeywordDeletedTypeID is TL type id of EmojiKeywordDeleted.
const EmojiKeywordDeletedTypeID = 0x236df622

func (e *EmojiKeywordDeleted) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Keyword == "") {
		return false
	}
	if !(e.Emoticons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiKeywordDeleted) String() string {
	if e == nil {
		return "EmojiKeywordDeleted(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EmojiKeywordDeleted")
	sb.WriteString("{\n")
	sb.WriteString("\tKeyword: ")
	sb.WriteString(fmt.Sprint(e.Keyword))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range e.Emoticons {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *EmojiKeywordDeleted) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywordDeleted#236df622 as nil")
	}
	b.PutID(EmojiKeywordDeletedTypeID)
	b.PutString(e.Keyword)
	b.PutVectorHeader(len(e.Emoticons))
	for _, v := range e.Emoticons {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiKeywordDeleted) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywordDeleted#236df622 to nil")
	}
	if err := b.ConsumeID(EmojiKeywordDeletedTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: field keyword: %w", err)
		}
		e.Keyword = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: field emoticons: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: field emoticons: %w", err)
			}
			e.Emoticons = append(e.Emoticons, value)
		}
	}
	return nil
}

// construct implements constructor of EmojiKeywordClass.
func (e EmojiKeywordDeleted) construct() EmojiKeywordClass { return &e }

// Ensuring interfaces in compile-time for EmojiKeywordDeleted.
var (
	_ bin.Encoder = &EmojiKeywordDeleted{}
	_ bin.Decoder = &EmojiKeywordDeleted{}

	_ EmojiKeywordClass = &EmojiKeywordDeleted{}
)

// EmojiKeywordClass represents EmojiKeyword generic type.
//
// See https://core.telegram.org/type/EmojiKeyword for reference.
//
// Example:
//  g, err := DecodeEmojiKeyword(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *EmojiKeyword: // emojiKeyword#d5b3b9f9
//  case *EmojiKeywordDeleted: // emojiKeywordDeleted#236df622
//  default: panic(v)
//  }
type EmojiKeywordClass interface {
	bin.Encoder
	bin.Decoder
	construct() EmojiKeywordClass

	fmt.Stringer
	Zero() bool
}

// DecodeEmojiKeyword implements binary de-serialization for EmojiKeywordClass.
func DecodeEmojiKeyword(buf *bin.Buffer) (EmojiKeywordClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EmojiKeywordTypeID:
		// Decoding emojiKeyword#d5b3b9f9.
		v := EmojiKeyword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiKeywordClass: %w", err)
		}
		return &v, nil
	case EmojiKeywordDeletedTypeID:
		// Decoding emojiKeywordDeleted#236df622.
		v := EmojiKeywordDeleted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiKeywordClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmojiKeywordClass: %w", bin.NewUnexpectedID(id))
	}
}

// EmojiKeyword boxes the EmojiKeywordClass providing a helper.
type EmojiKeywordBox struct {
	EmojiKeyword EmojiKeywordClass
}

// Decode implements bin.Decoder for EmojiKeywordBox.
func (b *EmojiKeywordBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmojiKeywordBox to nil")
	}
	v, err := DecodeEmojiKeyword(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiKeyword = v
	return nil
}

// Encode implements bin.Encode for EmojiKeywordBox.
func (b *EmojiKeywordBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmojiKeyword == nil {
		return fmt.Errorf("unable to encode EmojiKeywordClass as nil")
	}
	return b.EmojiKeyword.Encode(buf)
}
