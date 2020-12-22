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

// ChannelsGetParticipantRequest represents TL type `channels.getParticipant#546dd7a6`.
// Get info about a channel/supergroup participant
//
// See https://core.telegram.org/method/channels.getParticipant for reference.
type ChannelsGetParticipantRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// ID of participant to get info about
	UserID InputUserClass
}

// ChannelsGetParticipantRequestTypeID is TL type id of ChannelsGetParticipantRequest.
const ChannelsGetParticipantRequestTypeID = 0x546dd7a6

// String implements fmt.Stringer.
func (g *ChannelsGetParticipantRequest) String() string {
	if g == nil {
		return "ChannelsGetParticipantRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetParticipantRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(g.Channel.String())
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(g.UserID.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ChannelsGetParticipantRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getParticipant#546dd7a6 as nil")
	}
	b.PutID(ChannelsGetParticipantRequestTypeID)
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getParticipant#546dd7a6: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipant#546dd7a6: field channel: %w", err)
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode channels.getParticipant#546dd7a6: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipant#546dd7a6: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetParticipantRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getParticipant#546dd7a6 to nil")
	}
	if err := b.ConsumeID(ChannelsGetParticipantRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getParticipant#546dd7a6: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipant#546dd7a6: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipant#546dd7a6: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetParticipantRequest.
var (
	_ bin.Encoder = &ChannelsGetParticipantRequest{}
	_ bin.Decoder = &ChannelsGetParticipantRequest{}
)

// ChannelsGetParticipant invokes method channels.getParticipant#546dd7a6 returning error if any.
// Get info about a channel/supergroup participant
//
// See https://core.telegram.org/method/channels.getParticipant for reference.
func (c *Client) ChannelsGetParticipant(ctx context.Context, request *ChannelsGetParticipantRequest) (*ChannelsChannelParticipant, error) {
	var result ChannelsChannelParticipant

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
