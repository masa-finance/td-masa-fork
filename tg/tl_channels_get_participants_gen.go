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

// ChannelsGetParticipantsRequest represents TL type `channels.getParticipants#123e05e9`.
// Get the participants of a supergroup/channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getParticipants for reference.
type ChannelsGetParticipantsRequest struct {
	// Channel
	Channel InputChannelClass
	// Which participant types to fetch
	Filter ChannelParticipantsFilterClass
	// Offset¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
	// Limit¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int
}

// ChannelsGetParticipantsRequestTypeID is TL type id of ChannelsGetParticipantsRequest.
const ChannelsGetParticipantsRequestTypeID = 0x123e05e9

func (g *ChannelsGetParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetParticipantsRequest) String() string {
	if g == nil {
		return "ChannelsGetParticipantsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetParticipantsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(g.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tFilter: ")
	sb.WriteString(fmt.Sprint(g.Filter))
	sb.WriteString(",\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(g.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ChannelsGetParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getParticipants#123e05e9 as nil")
	}
	b.PutID(ChannelsGetParticipantsRequestTypeID)
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field channel: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field filter: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getParticipants#123e05e9 to nil")
	}
	if err := b.ConsumeID(ChannelsGetParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeChannelParticipantsFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetParticipantsRequest.
var (
	_ bin.Encoder = &ChannelsGetParticipantsRequest{}
	_ bin.Decoder = &ChannelsGetParticipantsRequest{}
)

// ChannelsGetParticipants invokes method channels.getParticipants#123e05e9 returning error if any.
// Get the participants of a supergroup/channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 INPUT_CONSTRUCTOR_INVALID: The provided constructor is invalid
//
// See https://core.telegram.org/method/channels.getParticipants for reference.
// Can be used by bots.
func (c *Client) ChannelsGetParticipants(ctx context.Context, request *ChannelsGetParticipantsRequest) (ChannelsChannelParticipantsClass, error) {
	var result ChannelsChannelParticipantsBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChannelParticipants, nil
}
