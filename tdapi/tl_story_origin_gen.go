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

// StoryOriginPublicStory represents TL type `storyOriginPublicStory#2c379fbe`.
type StoryOriginPublicStory struct {
	// Identifier of the chat that posted original story
	ChatID int64
	// Story identifier of the original story
	StoryID int32
}

// StoryOriginPublicStoryTypeID is TL type id of StoryOriginPublicStory.
const StoryOriginPublicStoryTypeID = 0x2c379fbe

// construct implements constructor of StoryOriginClass.
func (s StoryOriginPublicStory) construct() StoryOriginClass { return &s }

// Ensuring interfaces in compile-time for StoryOriginPublicStory.
var (
	_ bin.Encoder     = &StoryOriginPublicStory{}
	_ bin.Decoder     = &StoryOriginPublicStory{}
	_ bin.BareEncoder = &StoryOriginPublicStory{}
	_ bin.BareDecoder = &StoryOriginPublicStory{}

	_ StoryOriginClass = &StoryOriginPublicStory{}
)

func (s *StoryOriginPublicStory) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryOriginPublicStory) String() string {
	if s == nil {
		return "StoryOriginPublicStory(nil)"
	}
	type Alias StoryOriginPublicStory
	return fmt.Sprintf("StoryOriginPublicStory%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryOriginPublicStory) TypeID() uint32 {
	return StoryOriginPublicStoryTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryOriginPublicStory) TypeName() string {
	return "storyOriginPublicStory"
}

// TypeInfo returns info about TL type.
func (s *StoryOriginPublicStory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyOriginPublicStory",
		ID:   StoryOriginPublicStoryTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryOriginPublicStory) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyOriginPublicStory#2c379fbe as nil")
	}
	b.PutID(StoryOriginPublicStoryTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryOriginPublicStory) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyOriginPublicStory#2c379fbe as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt32(s.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryOriginPublicStory) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyOriginPublicStory#2c379fbe to nil")
	}
	if err := b.ConsumeID(StoryOriginPublicStoryTypeID); err != nil {
		return fmt.Errorf("unable to decode storyOriginPublicStory#2c379fbe: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryOriginPublicStory) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyOriginPublicStory#2c379fbe to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storyOriginPublicStory#2c379fbe: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyOriginPublicStory#2c379fbe: field story_id: %w", err)
		}
		s.StoryID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryOriginPublicStory) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyOriginPublicStory#2c379fbe as nil")
	}
	b.ObjStart()
	b.PutID("storyOriginPublicStory")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(s.StoryID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryOriginPublicStory) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyOriginPublicStory#2c379fbe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyOriginPublicStory"); err != nil {
				return fmt.Errorf("unable to decode storyOriginPublicStory#2c379fbe: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storyOriginPublicStory#2c379fbe: field chat_id: %w", err)
			}
			s.ChatID = value
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyOriginPublicStory#2c379fbe: field story_id: %w", err)
			}
			s.StoryID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *StoryOriginPublicStory) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetStoryID returns value of StoryID field.
func (s *StoryOriginPublicStory) GetStoryID() (value int32) {
	if s == nil {
		return
	}
	return s.StoryID
}

// StoryOriginHiddenUser represents TL type `storyOriginHiddenUser#5a1f89ec`.
type StoryOriginHiddenUser struct {
	// Name of the story sender
	SenderName string
}

// StoryOriginHiddenUserTypeID is TL type id of StoryOriginHiddenUser.
const StoryOriginHiddenUserTypeID = 0x5a1f89ec

// construct implements constructor of StoryOriginClass.
func (s StoryOriginHiddenUser) construct() StoryOriginClass { return &s }

// Ensuring interfaces in compile-time for StoryOriginHiddenUser.
var (
	_ bin.Encoder     = &StoryOriginHiddenUser{}
	_ bin.Decoder     = &StoryOriginHiddenUser{}
	_ bin.BareEncoder = &StoryOriginHiddenUser{}
	_ bin.BareDecoder = &StoryOriginHiddenUser{}

	_ StoryOriginClass = &StoryOriginHiddenUser{}
)

func (s *StoryOriginHiddenUser) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SenderName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryOriginHiddenUser) String() string {
	if s == nil {
		return "StoryOriginHiddenUser(nil)"
	}
	type Alias StoryOriginHiddenUser
	return fmt.Sprintf("StoryOriginHiddenUser%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryOriginHiddenUser) TypeID() uint32 {
	return StoryOriginHiddenUserTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryOriginHiddenUser) TypeName() string {
	return "storyOriginHiddenUser"
}

// TypeInfo returns info about TL type.
func (s *StoryOriginHiddenUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyOriginHiddenUser",
		ID:   StoryOriginHiddenUserTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderName",
			SchemaName: "sender_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryOriginHiddenUser) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyOriginHiddenUser#5a1f89ec as nil")
	}
	b.PutID(StoryOriginHiddenUserTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryOriginHiddenUser) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyOriginHiddenUser#5a1f89ec as nil")
	}
	b.PutString(s.SenderName)
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryOriginHiddenUser) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyOriginHiddenUser#5a1f89ec to nil")
	}
	if err := b.ConsumeID(StoryOriginHiddenUserTypeID); err != nil {
		return fmt.Errorf("unable to decode storyOriginHiddenUser#5a1f89ec: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryOriginHiddenUser) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyOriginHiddenUser#5a1f89ec to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode storyOriginHiddenUser#5a1f89ec: field sender_name: %w", err)
		}
		s.SenderName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryOriginHiddenUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyOriginHiddenUser#5a1f89ec as nil")
	}
	b.ObjStart()
	b.PutID("storyOriginHiddenUser")
	b.Comma()
	b.FieldStart("sender_name")
	b.PutString(s.SenderName)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryOriginHiddenUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyOriginHiddenUser#5a1f89ec to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyOriginHiddenUser"); err != nil {
				return fmt.Errorf("unable to decode storyOriginHiddenUser#5a1f89ec: %w", err)
			}
		case "sender_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode storyOriginHiddenUser#5a1f89ec: field sender_name: %w", err)
			}
			s.SenderName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderName returns value of SenderName field.
func (s *StoryOriginHiddenUser) GetSenderName() (value string) {
	if s == nil {
		return
	}
	return s.SenderName
}

// StoryOriginClassName is schema name of StoryOriginClass.
const StoryOriginClassName = "StoryOrigin"

// StoryOriginClass represents StoryOrigin generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStoryOrigin(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StoryOriginPublicStory: // storyOriginPublicStory#2c379fbe
//	case *tdapi.StoryOriginHiddenUser: // storyOriginHiddenUser#5a1f89ec
//	default: panic(v)
//	}
type StoryOriginClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryOriginClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeStoryOrigin implements binary de-serialization for StoryOriginClass.
func DecodeStoryOrigin(buf *bin.Buffer) (StoryOriginClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryOriginPublicStoryTypeID:
		// Decoding storyOriginPublicStory#2c379fbe.
		v := StoryOriginPublicStory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryOriginClass: %w", err)
		}
		return &v, nil
	case StoryOriginHiddenUserTypeID:
		// Decoding storyOriginHiddenUser#5a1f89ec.
		v := StoryOriginHiddenUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryOriginClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryOriginClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStoryOrigin implements binary de-serialization for StoryOriginClass.
func DecodeTDLibJSONStoryOrigin(buf tdjson.Decoder) (StoryOriginClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "storyOriginPublicStory":
		// Decoding storyOriginPublicStory#2c379fbe.
		v := StoryOriginPublicStory{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryOriginClass: %w", err)
		}
		return &v, nil
	case "storyOriginHiddenUser":
		// Decoding storyOriginHiddenUser#5a1f89ec.
		v := StoryOriginHiddenUser{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryOriginClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryOriginClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StoryOrigin boxes the StoryOriginClass providing a helper.
type StoryOriginBox struct {
	StoryOrigin StoryOriginClass
}

// Decode implements bin.Decoder for StoryOriginBox.
func (b *StoryOriginBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryOriginBox to nil")
	}
	v, err := DecodeStoryOrigin(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryOrigin = v
	return nil
}

// Encode implements bin.Encode for StoryOriginBox.
func (b *StoryOriginBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryOrigin == nil {
		return fmt.Errorf("unable to encode StoryOriginClass as nil")
	}
	return b.StoryOrigin.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StoryOriginBox.
func (b *StoryOriginBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryOriginBox to nil")
	}
	v, err := DecodeTDLibJSONStoryOrigin(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryOrigin = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StoryOriginBox.
func (b *StoryOriginBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StoryOrigin == nil {
		return fmt.Errorf("unable to encode StoryOriginClass as nil")
	}
	return b.StoryOrigin.EncodeTDLibJSON(buf)
}
