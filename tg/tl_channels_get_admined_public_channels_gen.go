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
// Get channels/supergroups/geogroups¹ we're admin in. Usually called when the user exceeds the limit² for owned public channels/supergroups/geogroups³, and the user is given the choice to remove one of his channels/supergroups/geogroups.
//
// Links:
//  1) https://core.telegram.org/api/channel
//  2) https://core.telegram.org/constructor/config
//  3) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getAdminedPublicChannels for reference.
type ChannelsGetAdminedPublicChannelsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Get geogroups
	ByLocation bool
	// If set and the user has reached the limit of owned public channels/supergroups/geogroups¹, instead of returning the channel list one of the specified errors² will be returned.Useful to check if a new public channel can indeed be created, even before asking the user to enter a channel username to use in channels.checkUsername³/channels.updateUsername⁴.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	//  2) https://core.telegram.org#possible-errors
	//  3) https://core.telegram.org/method/channels.checkUsername
	//  4) https://core.telegram.org/method/channels.updateUsername
	CheckLimit bool
}

// ChannelsGetAdminedPublicChannelsRequestTypeID is TL type id of ChannelsGetAdminedPublicChannelsRequest.
const ChannelsGetAdminedPublicChannelsRequestTypeID = 0xf8b036af

func (g *ChannelsGetAdminedPublicChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ByLocation == false) {
		return false
	}
	if !(g.CheckLimit == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetAdminedPublicChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetAdminedPublicChannelsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsGetAdminedPublicChannelsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
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
	if !(g.ByLocation == false) {
		g.Flags.Set(0)
	}
	if !(g.CheckLimit == false) {
		g.Flags.Set(1)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getAdminedPublicChannels#f8b036af: field flags: %w", err)
	}
	return nil
}

// SetByLocation sets value of ByLocation conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetByLocation(value bool) {
	if value {
		g.Flags.Set(0)
		g.ByLocation = true
	} else {
		g.Flags.Unset(0)
		g.ByLocation = false
	}
}

// SetCheckLimit sets value of CheckLimit conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetCheckLimit(value bool) {
	if value {
		g.Flags.Set(1)
		g.CheckLimit = true
	} else {
		g.Flags.Unset(1)
		g.CheckLimit = false
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
// Get channels/supergroups/geogroups¹ we're admin in. Usually called when the user exceeds the limit² for owned public channels/supergroups/geogroups³, and the user is given the choice to remove one of his channels/supergroups/geogroups.
//
// Links:
//  1) https://core.telegram.org/api/channel
//  2) https://core.telegram.org/constructor/config
//  3) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNELS_ADMIN_LOCATED_TOO_MUCH: Returned if both the check_limit and the by_location flags are set and the user has reached the limit of public geogroups
//  400 CHANNELS_ADMIN_PUBLIC_TOO_MUCH: Returned if the check_limit flag is set and the user has reached the limit of public channels/supergroups
//
// See https://core.telegram.org/method/channels.getAdminedPublicChannels for reference.
func (c *Client) ChannelsGetAdminedPublicChannels(ctx context.Context, request *ChannelsGetAdminedPublicChannelsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
