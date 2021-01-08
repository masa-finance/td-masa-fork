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

// UserFull represents TL type `userFull#edf17c12`.
// Extended user info
//
// See https://core.telegram.org/constructor/userFull for reference.
type UserFull struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether you have blocked this user
	Blocked bool
	// Whether this user can make VoIP calls
	PhoneCallsAvailable bool
	// Whether this user's privacy settings allow you to call him
	PhoneCallsPrivate bool
	// Whether you can pin messages in the chat with this user, you can do this only for a chat with yourself
	CanPinMessage bool
	// Whether scheduled messages¹ are available
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	HasScheduled bool
	// Whether the user can receive video calls
	VideoCallsAvailable bool
	// Remaining user info
	User UserClass
	// Bio of the user
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// Peer settings
	Settings PeerSettings
	// Profile photo
	//
	// Use SetProfilePhoto and GetProfilePhoto helpers.
	ProfilePhoto PhotoClass
	// Notification settings
	NotifySettings PeerNotifySettings
	// For bots, info about the bot (bot commands, etc)
	//
	// Use SetBotInfo and GetBotInfo helpers.
	BotInfo BotInfo
	// Message ID of the last pinned message¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pin
	//
	// Use SetPinnedMsgID and GetPinnedMsgID helpers.
	PinnedMsgID int
	// Chats in common with this user
	CommonChatsCount int
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
}

// UserFullTypeID is TL type id of UserFull.
const UserFullTypeID = 0xedf17c12

func (u *UserFull) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Blocked == false) {
		return false
	}
	if !(u.PhoneCallsAvailable == false) {
		return false
	}
	if !(u.PhoneCallsPrivate == false) {
		return false
	}
	if !(u.CanPinMessage == false) {
		return false
	}
	if !(u.HasScheduled == false) {
		return false
	}
	if !(u.VideoCallsAvailable == false) {
		return false
	}
	if !(u.User == nil) {
		return false
	}
	if !(u.About == "") {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}
	if !(u.ProfilePhoto == nil) {
		return false
	}
	if !(u.NotifySettings.Zero()) {
		return false
	}
	if !(u.BotInfo.Zero()) {
		return false
	}
	if !(u.PinnedMsgID == 0) {
		return false
	}
	if !(u.CommonChatsCount == 0) {
		return false
	}
	if !(u.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserFull) String() string {
	if u == nil {
		return "UserFull(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UserFull")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(u.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tUser: ")
	sb.WriteString(fmt.Sprint(u.User))
	sb.WriteString(",\n")
	if u.Flags.Has(1) {
		sb.WriteString("\tAbout: ")
		sb.WriteString(fmt.Sprint(u.About))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tSettings: ")
	sb.WriteString(fmt.Sprint(u.Settings))
	sb.WriteString(",\n")
	if u.Flags.Has(2) {
		sb.WriteString("\tProfilePhoto: ")
		sb.WriteString(fmt.Sprint(u.ProfilePhoto))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tNotifySettings: ")
	sb.WriteString(fmt.Sprint(u.NotifySettings))
	sb.WriteString(",\n")
	if u.Flags.Has(3) {
		sb.WriteString("\tBotInfo: ")
		sb.WriteString(fmt.Sprint(u.BotInfo))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(6) {
		sb.WriteString("\tPinnedMsgID: ")
		sb.WriteString(fmt.Sprint(u.PinnedMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tCommonChatsCount: ")
	sb.WriteString(fmt.Sprint(u.CommonChatsCount))
	sb.WriteString(",\n")
	if u.Flags.Has(11) {
		sb.WriteString("\tFolderID: ")
		sb.WriteString(fmt.Sprint(u.FolderID))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *UserFull) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFull#edf17c12 as nil")
	}
	b.PutID(UserFullTypeID)
	if !(u.Blocked == false) {
		u.Flags.Set(0)
	}
	if !(u.PhoneCallsAvailable == false) {
		u.Flags.Set(4)
	}
	if !(u.PhoneCallsPrivate == false) {
		u.Flags.Set(5)
	}
	if !(u.CanPinMessage == false) {
		u.Flags.Set(7)
	}
	if !(u.HasScheduled == false) {
		u.Flags.Set(12)
	}
	if !(u.VideoCallsAvailable == false) {
		u.Flags.Set(13)
	}
	if !(u.About == "") {
		u.Flags.Set(1)
	}
	if !(u.ProfilePhoto == nil) {
		u.Flags.Set(2)
	}
	if !(u.BotInfo.Zero()) {
		u.Flags.Set(3)
	}
	if !(u.PinnedMsgID == 0) {
		u.Flags.Set(6)
	}
	if !(u.FolderID == 0) {
		u.Flags.Set(11)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#edf17c12: field flags: %w", err)
	}
	if u.User == nil {
		return fmt.Errorf("unable to encode userFull#edf17c12: field user is nil")
	}
	if err := u.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#edf17c12: field user: %w", err)
	}
	if u.Flags.Has(1) {
		b.PutString(u.About)
	}
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#edf17c12: field settings: %w", err)
	}
	if u.Flags.Has(2) {
		if u.ProfilePhoto == nil {
			return fmt.Errorf("unable to encode userFull#edf17c12: field profile_photo is nil")
		}
		if err := u.ProfilePhoto.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#edf17c12: field profile_photo: %w", err)
		}
	}
	if err := u.NotifySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#edf17c12: field notify_settings: %w", err)
	}
	if u.Flags.Has(3) {
		if err := u.BotInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#edf17c12: field bot_info: %w", err)
		}
	}
	if u.Flags.Has(6) {
		b.PutInt(u.PinnedMsgID)
	}
	b.PutInt(u.CommonChatsCount)
	if u.Flags.Has(11) {
		b.PutInt(u.FolderID)
	}
	return nil
}

// SetBlocked sets value of Blocked conditional field.
func (u *UserFull) SetBlocked(value bool) {
	if value {
		u.Flags.Set(0)
		u.Blocked = true
	} else {
		u.Flags.Unset(0)
		u.Blocked = false
	}
}

// SetPhoneCallsAvailable sets value of PhoneCallsAvailable conditional field.
func (u *UserFull) SetPhoneCallsAvailable(value bool) {
	if value {
		u.Flags.Set(4)
		u.PhoneCallsAvailable = true
	} else {
		u.Flags.Unset(4)
		u.PhoneCallsAvailable = false
	}
}

// SetPhoneCallsPrivate sets value of PhoneCallsPrivate conditional field.
func (u *UserFull) SetPhoneCallsPrivate(value bool) {
	if value {
		u.Flags.Set(5)
		u.PhoneCallsPrivate = true
	} else {
		u.Flags.Unset(5)
		u.PhoneCallsPrivate = false
	}
}

// SetCanPinMessage sets value of CanPinMessage conditional field.
func (u *UserFull) SetCanPinMessage(value bool) {
	if value {
		u.Flags.Set(7)
		u.CanPinMessage = true
	} else {
		u.Flags.Unset(7)
		u.CanPinMessage = false
	}
}

// SetHasScheduled sets value of HasScheduled conditional field.
func (u *UserFull) SetHasScheduled(value bool) {
	if value {
		u.Flags.Set(12)
		u.HasScheduled = true
	} else {
		u.Flags.Unset(12)
		u.HasScheduled = false
	}
}

// SetVideoCallsAvailable sets value of VideoCallsAvailable conditional field.
func (u *UserFull) SetVideoCallsAvailable(value bool) {
	if value {
		u.Flags.Set(13)
		u.VideoCallsAvailable = true
	} else {
		u.Flags.Unset(13)
		u.VideoCallsAvailable = false
	}
}

// SetAbout sets value of About conditional field.
func (u *UserFull) SetAbout(value string) {
	u.Flags.Set(1)
	u.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetAbout() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.About, true
}

// SetProfilePhoto sets value of ProfilePhoto conditional field.
func (u *UserFull) SetProfilePhoto(value PhotoClass) {
	u.Flags.Set(2)
	u.ProfilePhoto = value
}

// GetProfilePhoto returns value of ProfilePhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetProfilePhoto() (value PhotoClass, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.ProfilePhoto, true
}

// SetBotInfo sets value of BotInfo conditional field.
func (u *UserFull) SetBotInfo(value BotInfo) {
	u.Flags.Set(3)
	u.BotInfo = value
}

// GetBotInfo returns value of BotInfo conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetBotInfo() (value BotInfo, ok bool) {
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.BotInfo, true
}

// SetPinnedMsgID sets value of PinnedMsgID conditional field.
func (u *UserFull) SetPinnedMsgID(value int) {
	u.Flags.Set(6)
	u.PinnedMsgID = value
}

// GetPinnedMsgID returns value of PinnedMsgID conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPinnedMsgID() (value int, ok bool) {
	if !u.Flags.Has(6) {
		return value, false
	}
	return u.PinnedMsgID, true
}

// SetFolderID sets value of FolderID conditional field.
func (u *UserFull) SetFolderID(value int) {
	u.Flags.Set(11)
	u.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetFolderID() (value int, ok bool) {
	if !u.Flags.Has(11) {
		return value, false
	}
	return u.FolderID, true
}

// Decode implements bin.Decoder.
func (u *UserFull) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFull#edf17c12 to nil")
	}
	if err := b.ConsumeID(UserFullTypeID); err != nil {
		return fmt.Errorf("unable to decode userFull#edf17c12: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field flags: %w", err)
		}
	}
	u.Blocked = u.Flags.Has(0)
	u.PhoneCallsAvailable = u.Flags.Has(4)
	u.PhoneCallsPrivate = u.Flags.Has(5)
	u.CanPinMessage = u.Flags.Has(7)
	u.HasScheduled = u.Flags.Has(12)
	u.VideoCallsAvailable = u.Flags.Has(13)
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field user: %w", err)
		}
		u.User = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field about: %w", err)
		}
		u.About = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field settings: %w", err)
		}
	}
	if u.Flags.Has(2) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field profile_photo: %w", err)
		}
		u.ProfilePhoto = value
	}
	{
		if err := u.NotifySettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field notify_settings: %w", err)
		}
	}
	if u.Flags.Has(3) {
		if err := u.BotInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field bot_info: %w", err)
		}
	}
	if u.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field pinned_msg_id: %w", err)
		}
		u.PinnedMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field common_chats_count: %w", err)
		}
		u.CommonChatsCount = value
	}
	if u.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#edf17c12: field folder_id: %w", err)
		}
		u.FolderID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UserFull.
var (
	_ bin.Encoder = &UserFull{}
	_ bin.Decoder = &UserFull{}
)
