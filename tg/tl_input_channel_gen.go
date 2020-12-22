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

// InputChannelEmpty represents TL type `inputChannelEmpty#ee8c1e86`.
// Represents the absence of a channel
//
// See https://core.telegram.org/constructor/inputChannelEmpty for reference.
type InputChannelEmpty struct {
}

// InputChannelEmptyTypeID is TL type id of InputChannelEmpty.
const InputChannelEmptyTypeID = 0xee8c1e86

// String implements fmt.Stringer.
func (i *InputChannelEmpty) String() string {
	if i == nil {
		return "InputChannelEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputChannelEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputChannelEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelEmpty#ee8c1e86 as nil")
	}
	b.PutID(InputChannelEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelEmpty#ee8c1e86 to nil")
	}
	if err := b.ConsumeID(InputChannelEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelEmpty#ee8c1e86: %w", err)
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannelEmpty) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelEmpty.
var (
	_ bin.Encoder = &InputChannelEmpty{}
	_ bin.Decoder = &InputChannelEmpty{}

	_ InputChannelClass = &InputChannelEmpty{}
)

// InputChannel represents TL type `inputChannel#afeb712e`.
// Represents a channel
//
// See https://core.telegram.org/constructor/inputChannel for reference.
type InputChannel struct {
	// Channel ID
	ChannelID int
	// Access hash taken from the channel constructor
	AccessHash int64
}

// InputChannelTypeID is TL type id of InputChannel.
const InputChannelTypeID = 0xafeb712e

// String implements fmt.Stringer.
func (i *InputChannel) String() string {
	if i == nil {
		return "InputChannel(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputChannel")
	sb.WriteString("{\n")
	sb.WriteString("\tChannelID: ")
	sb.WriteString(fmt.Sprint(i.ChannelID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputChannel) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannel#afeb712e as nil")
	}
	b.PutID(InputChannelTypeID)
	b.PutInt(i.ChannelID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannel) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannel#afeb712e to nil")
	}
	if err := b.ConsumeID(InputChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannel#afeb712e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#afeb712e: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#afeb712e: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannel) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannel.
var (
	_ bin.Encoder = &InputChannel{}
	_ bin.Decoder = &InputChannel{}

	_ InputChannelClass = &InputChannel{}
)

// InputChannelFromMessage represents TL type `inputChannelFromMessage#2a286531`.
// Defines a min channel that was seen in a certain message of a certain chat.
//
// See https://core.telegram.org/constructor/inputChannelFromMessage for reference.
type InputChannelFromMessage struct {
	// The chat where the channel was seen
	Peer InputPeerClass
	// The message ID in the chat where the channel was seen
	MsgID int
	// The channel ID
	ChannelID int
}

// InputChannelFromMessageTypeID is TL type id of InputChannelFromMessage.
const InputChannelFromMessageTypeID = 0x2a286531

// String implements fmt.Stringer.
func (i *InputChannelFromMessage) String() string {
	if i == nil {
		return "InputChannelFromMessage(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputChannelFromMessage")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(i.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(i.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tChannelID: ")
	sb.WriteString(fmt.Sprint(i.ChannelID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputChannelFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelFromMessage#2a286531 as nil")
	}
	b.PutID(InputChannelFromMessageTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#2a286531: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#2a286531: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutInt(i.ChannelID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelFromMessage#2a286531 to nil")
	}
	if err := b.ConsumeID(InputChannelFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannelFromMessage) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelFromMessage.
var (
	_ bin.Encoder = &InputChannelFromMessage{}
	_ bin.Decoder = &InputChannelFromMessage{}

	_ InputChannelClass = &InputChannelFromMessage{}
)

// InputChannelClass represents InputChannel generic type.
//
// See https://core.telegram.org/type/InputChannel for reference.
//
// Example:
//  g, err := DecodeInputChannel(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputChannelEmpty: // inputChannelEmpty#ee8c1e86
//  case *InputChannel: // inputChannel#afeb712e
//  case *InputChannelFromMessage: // inputChannelFromMessage#2a286531
//  default: panic(v)
//  }
type InputChannelClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputChannelClass
	fmt.Stringer
}

// DecodeInputChannel implements binary de-serialization for InputChannelClass.
func DecodeInputChannel(buf *bin.Buffer) (InputChannelClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChannelEmptyTypeID:
		// Decoding inputChannelEmpty#ee8c1e86.
		v := InputChannelEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelTypeID:
		// Decoding inputChannel#afeb712e.
		v := InputChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelFromMessageTypeID:
		// Decoding inputChannelFromMessage#2a286531.
		v := InputChannelFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChannelClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputChannel boxes the InputChannelClass providing a helper.
type InputChannelBox struct {
	InputChannel InputChannelClass
}

// Decode implements bin.Decoder for InputChannelBox.
func (b *InputChannelBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChannelBox to nil")
	}
	v, err := DecodeInputChannel(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChannel = v
	return nil
}

// Encode implements bin.Encode for InputChannelBox.
func (b *InputChannelBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputChannel == nil {
		return fmt.Errorf("unable to encode InputChannelClass as nil")
	}
	return b.InputChannel.Encode(buf)
}
