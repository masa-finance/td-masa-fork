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

// ChannelsGetAdminedPublicChannelsRequest represents TL type `channels.getAdminedPublicChannels#f8b036af`.
// Get channels/supergroups/geogroups we're admin in. Usually called when the user exceeds the limit for owned public channels/supergroups/geogroups, and the user is given the choice to remove one of his channels/supergroups/geogroups.
//
// See https://core.telegram.org/method/channels.getAdminedPublicChannels for reference.
type ChannelsGetAdminedPublicChannelsRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Get geogroups
	ByLocation bool
	// If set and the user has reached the limit of owned public channels/supergroups/geogroups, instead of returning the channel list one of the specified errors will be returned.Useful to check if a new public channel can indeed be created, even before asking the user to enter a channel username to use in channels.checkUsername/channels.updateUsername.
	CheckLimit bool
}

// ChannelsGetAdminedPublicChannelsRequestTypeID is TL type id of ChannelsGetAdminedPublicChannelsRequest.
const ChannelsGetAdminedPublicChannelsRequestTypeID = 0xf8b036af

// String implements fmt.Stringer.
func (g *ChannelsGetAdminedPublicChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetAdminedPublicChannelsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetAdminedPublicChannelsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(g.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ChannelsGetAdminedPublicChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getAdminedPublicChannels#f8b036af as nil")
	}
	b.PutID(ChannelsGetAdminedPublicChannelsRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getAdminedPublicChannels#f8b036af: field flags: %w", err)
	}
	return nil
}

// SetByLocation sets value of ByLocation conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetByLocation(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// SetCheckLimit sets value of CheckLimit conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetCheckLimit(value bool) {
	if value {
		g.Flags.Set(1)
	} else {
		g.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (g *ChannelsGetAdminedPublicChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getAdminedPublicChannels#f8b036af to nil")
	}
	if err := b.ConsumeID(ChannelsGetAdminedPublicChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getAdminedPublicChannels#f8b036af: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.getAdminedPublicChannels#f8b036af: field flags: %w", err)
		}
	}
	g.ByLocation = g.Flags.Has(0)
	g.CheckLimit = g.Flags.Has(1)
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetAdminedPublicChannelsRequest.
var (
	_ bin.Encoder = &ChannelsGetAdminedPublicChannelsRequest{}
	_ bin.Decoder = &ChannelsGetAdminedPublicChannelsRequest{}
)

// ChannelsGetAdminedPublicChannels invokes method channels.getAdminedPublicChannels#f8b036af returning error if any.
// Get channels/supergroups/geogroups we're admin in. Usually called when the user exceeds the limit for owned public channels/supergroups/geogroups, and the user is given the choice to remove one of his channels/supergroups/geogroups.
//
// See https://core.telegram.org/method/channels.getAdminedPublicChannels for reference.
func (c *Client) ChannelsGetAdminedPublicChannels(ctx context.Context, request *ChannelsGetAdminedPublicChannelsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
