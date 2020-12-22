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

// TopPeerCategoryBotsPM represents TL type `topPeerCategoryBotsPM#ab661b5b`.
// Most used bots
//
// See https://core.telegram.org/constructor/topPeerCategoryBotsPM for reference.
type TopPeerCategoryBotsPM struct {
}

// TopPeerCategoryBotsPMTypeID is TL type id of TopPeerCategoryBotsPM.
const TopPeerCategoryBotsPMTypeID = 0xab661b5b

// String implements fmt.Stringer.
func (t *TopPeerCategoryBotsPM) String() string {
	if t == nil {
		return "TopPeerCategoryBotsPM(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryBotsPM")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryBotsPM) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryBotsPM#ab661b5b as nil")
	}
	b.PutID(TopPeerCategoryBotsPMTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryBotsPM) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryBotsPM#ab661b5b to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryBotsPMTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryBotsPM#ab661b5b: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryBotsPM) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryBotsPM.
var (
	_ bin.Encoder = &TopPeerCategoryBotsPM{}
	_ bin.Decoder = &TopPeerCategoryBotsPM{}

	_ TopPeerCategoryClass = &TopPeerCategoryBotsPM{}
)

// TopPeerCategoryBotsInline represents TL type `topPeerCategoryBotsInline#148677e2`.
// Most used inline bots
//
// See https://core.telegram.org/constructor/topPeerCategoryBotsInline for reference.
type TopPeerCategoryBotsInline struct {
}

// TopPeerCategoryBotsInlineTypeID is TL type id of TopPeerCategoryBotsInline.
const TopPeerCategoryBotsInlineTypeID = 0x148677e2

// String implements fmt.Stringer.
func (t *TopPeerCategoryBotsInline) String() string {
	if t == nil {
		return "TopPeerCategoryBotsInline(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryBotsInline")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryBotsInline) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryBotsInline#148677e2 as nil")
	}
	b.PutID(TopPeerCategoryBotsInlineTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryBotsInline) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryBotsInline#148677e2 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryBotsInlineTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryBotsInline#148677e2: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryBotsInline) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryBotsInline.
var (
	_ bin.Encoder = &TopPeerCategoryBotsInline{}
	_ bin.Decoder = &TopPeerCategoryBotsInline{}

	_ TopPeerCategoryClass = &TopPeerCategoryBotsInline{}
)

// TopPeerCategoryCorrespondents represents TL type `topPeerCategoryCorrespondents#637b7ed`.
// Users we've chatted most frequently with
//
// See https://core.telegram.org/constructor/topPeerCategoryCorrespondents for reference.
type TopPeerCategoryCorrespondents struct {
}

// TopPeerCategoryCorrespondentsTypeID is TL type id of TopPeerCategoryCorrespondents.
const TopPeerCategoryCorrespondentsTypeID = 0x637b7ed

// String implements fmt.Stringer.
func (t *TopPeerCategoryCorrespondents) String() string {
	if t == nil {
		return "TopPeerCategoryCorrespondents(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryCorrespondents")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryCorrespondents) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryCorrespondents#637b7ed as nil")
	}
	b.PutID(TopPeerCategoryCorrespondentsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryCorrespondents) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryCorrespondents#637b7ed to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryCorrespondentsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryCorrespondents#637b7ed: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryCorrespondents) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryCorrespondents.
var (
	_ bin.Encoder = &TopPeerCategoryCorrespondents{}
	_ bin.Decoder = &TopPeerCategoryCorrespondents{}

	_ TopPeerCategoryClass = &TopPeerCategoryCorrespondents{}
)

// TopPeerCategoryGroups represents TL type `topPeerCategoryGroups#bd17a14a`.
// Often-opened groups and supergroups
//
// See https://core.telegram.org/constructor/topPeerCategoryGroups for reference.
type TopPeerCategoryGroups struct {
}

// TopPeerCategoryGroupsTypeID is TL type id of TopPeerCategoryGroups.
const TopPeerCategoryGroupsTypeID = 0xbd17a14a

// String implements fmt.Stringer.
func (t *TopPeerCategoryGroups) String() string {
	if t == nil {
		return "TopPeerCategoryGroups(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryGroups")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryGroups) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryGroups#bd17a14a as nil")
	}
	b.PutID(TopPeerCategoryGroupsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryGroups) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryGroups#bd17a14a to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryGroupsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryGroups#bd17a14a: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryGroups) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryGroups.
var (
	_ bin.Encoder = &TopPeerCategoryGroups{}
	_ bin.Decoder = &TopPeerCategoryGroups{}

	_ TopPeerCategoryClass = &TopPeerCategoryGroups{}
)

// TopPeerCategoryChannels represents TL type `topPeerCategoryChannels#161d9628`.
// Most frequently visited channels
//
// See https://core.telegram.org/constructor/topPeerCategoryChannels for reference.
type TopPeerCategoryChannels struct {
}

// TopPeerCategoryChannelsTypeID is TL type id of TopPeerCategoryChannels.
const TopPeerCategoryChannelsTypeID = 0x161d9628

// String implements fmt.Stringer.
func (t *TopPeerCategoryChannels) String() string {
	if t == nil {
		return "TopPeerCategoryChannels(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryChannels")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryChannels) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryChannels#161d9628 as nil")
	}
	b.PutID(TopPeerCategoryChannelsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryChannels) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryChannels#161d9628 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryChannelsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryChannels#161d9628: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryChannels) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryChannels.
var (
	_ bin.Encoder = &TopPeerCategoryChannels{}
	_ bin.Decoder = &TopPeerCategoryChannels{}

	_ TopPeerCategoryClass = &TopPeerCategoryChannels{}
)

// TopPeerCategoryPhoneCalls represents TL type `topPeerCategoryPhoneCalls#1e76a78c`.
// Most frequently called users
//
// See https://core.telegram.org/constructor/topPeerCategoryPhoneCalls for reference.
type TopPeerCategoryPhoneCalls struct {
}

// TopPeerCategoryPhoneCallsTypeID is TL type id of TopPeerCategoryPhoneCalls.
const TopPeerCategoryPhoneCallsTypeID = 0x1e76a78c

// String implements fmt.Stringer.
func (t *TopPeerCategoryPhoneCalls) String() string {
	if t == nil {
		return "TopPeerCategoryPhoneCalls(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryPhoneCalls")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryPhoneCalls) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryPhoneCalls#1e76a78c as nil")
	}
	b.PutID(TopPeerCategoryPhoneCallsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryPhoneCalls) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryPhoneCalls#1e76a78c to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryPhoneCallsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryPhoneCalls#1e76a78c: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryPhoneCalls) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryPhoneCalls.
var (
	_ bin.Encoder = &TopPeerCategoryPhoneCalls{}
	_ bin.Decoder = &TopPeerCategoryPhoneCalls{}

	_ TopPeerCategoryClass = &TopPeerCategoryPhoneCalls{}
)

// TopPeerCategoryForwardUsers represents TL type `topPeerCategoryForwardUsers#a8406ca9`.
// Users to which the users often forwards messages to
//
// See https://core.telegram.org/constructor/topPeerCategoryForwardUsers for reference.
type TopPeerCategoryForwardUsers struct {
}

// TopPeerCategoryForwardUsersTypeID is TL type id of TopPeerCategoryForwardUsers.
const TopPeerCategoryForwardUsersTypeID = 0xa8406ca9

// String implements fmt.Stringer.
func (t *TopPeerCategoryForwardUsers) String() string {
	if t == nil {
		return "TopPeerCategoryForwardUsers(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryForwardUsers")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryForwardUsers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryForwardUsers#a8406ca9 as nil")
	}
	b.PutID(TopPeerCategoryForwardUsersTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryForwardUsers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryForwardUsers#a8406ca9 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryForwardUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryForwardUsers#a8406ca9: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryForwardUsers) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryForwardUsers.
var (
	_ bin.Encoder = &TopPeerCategoryForwardUsers{}
	_ bin.Decoder = &TopPeerCategoryForwardUsers{}

	_ TopPeerCategoryClass = &TopPeerCategoryForwardUsers{}
)

// TopPeerCategoryForwardChats represents TL type `topPeerCategoryForwardChats#fbeec0f0`.
// Chats to which the users often forwards messages to
//
// See https://core.telegram.org/constructor/topPeerCategoryForwardChats for reference.
type TopPeerCategoryForwardChats struct {
}

// TopPeerCategoryForwardChatsTypeID is TL type id of TopPeerCategoryForwardChats.
const TopPeerCategoryForwardChatsTypeID = 0xfbeec0f0

// String implements fmt.Stringer.
func (t *TopPeerCategoryForwardChats) String() string {
	if t == nil {
		return "TopPeerCategoryForwardChats(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TopPeerCategoryForwardChats")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryForwardChats) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryForwardChats#fbeec0f0 as nil")
	}
	b.PutID(TopPeerCategoryForwardChatsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryForwardChats) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryForwardChats#fbeec0f0 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryForwardChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryForwardChats#fbeec0f0: %w", err)
	}
	return nil
}

// construct implements constructor of TopPeerCategoryClass.
func (t TopPeerCategoryForwardChats) construct() TopPeerCategoryClass { return &t }

// Ensuring interfaces in compile-time for TopPeerCategoryForwardChats.
var (
	_ bin.Encoder = &TopPeerCategoryForwardChats{}
	_ bin.Decoder = &TopPeerCategoryForwardChats{}

	_ TopPeerCategoryClass = &TopPeerCategoryForwardChats{}
)

// TopPeerCategoryClass represents TopPeerCategory generic type.
//
// See https://core.telegram.org/type/TopPeerCategory for reference.
//
// Example:
//  g, err := DecodeTopPeerCategory(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *TopPeerCategoryBotsPM: // topPeerCategoryBotsPM#ab661b5b
//  case *TopPeerCategoryBotsInline: // topPeerCategoryBotsInline#148677e2
//  case *TopPeerCategoryCorrespondents: // topPeerCategoryCorrespondents#637b7ed
//  case *TopPeerCategoryGroups: // topPeerCategoryGroups#bd17a14a
//  case *TopPeerCategoryChannels: // topPeerCategoryChannels#161d9628
//  case *TopPeerCategoryPhoneCalls: // topPeerCategoryPhoneCalls#1e76a78c
//  case *TopPeerCategoryForwardUsers: // topPeerCategoryForwardUsers#a8406ca9
//  case *TopPeerCategoryForwardChats: // topPeerCategoryForwardChats#fbeec0f0
//  default: panic(v)
//  }
type TopPeerCategoryClass interface {
	bin.Encoder
	bin.Decoder
	construct() TopPeerCategoryClass
	fmt.Stringer
}

// DecodeTopPeerCategory implements binary de-serialization for TopPeerCategoryClass.
func DecodeTopPeerCategory(buf *bin.Buffer) (TopPeerCategoryClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case TopPeerCategoryBotsPMTypeID:
		// Decoding topPeerCategoryBotsPM#ab661b5b.
		v := TopPeerCategoryBotsPM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryBotsInlineTypeID:
		// Decoding topPeerCategoryBotsInline#148677e2.
		v := TopPeerCategoryBotsInline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryCorrespondentsTypeID:
		// Decoding topPeerCategoryCorrespondents#637b7ed.
		v := TopPeerCategoryCorrespondents{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryGroupsTypeID:
		// Decoding topPeerCategoryGroups#bd17a14a.
		v := TopPeerCategoryGroups{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryChannelsTypeID:
		// Decoding topPeerCategoryChannels#161d9628.
		v := TopPeerCategoryChannels{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryPhoneCallsTypeID:
		// Decoding topPeerCategoryPhoneCalls#1e76a78c.
		v := TopPeerCategoryPhoneCalls{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryForwardUsersTypeID:
		// Decoding topPeerCategoryForwardUsers#a8406ca9.
		v := TopPeerCategoryForwardUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	case TopPeerCategoryForwardChatsTypeID:
		// Decoding topPeerCategoryForwardChats#fbeec0f0.
		v := TopPeerCategoryForwardChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode TopPeerCategoryClass: %w", bin.NewUnexpectedID(id))
	}
}

// TopPeerCategory boxes the TopPeerCategoryClass providing a helper.
type TopPeerCategoryBox struct {
	TopPeerCategory TopPeerCategoryClass
}

// Decode implements bin.Decoder for TopPeerCategoryBox.
func (b *TopPeerCategoryBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode TopPeerCategoryBox to nil")
	}
	v, err := DecodeTopPeerCategory(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TopPeerCategory = v
	return nil
}

// Encode implements bin.Encode for TopPeerCategoryBox.
func (b *TopPeerCategoryBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TopPeerCategory == nil {
		return fmt.Errorf("unable to encode TopPeerCategoryClass as nil")
	}
	return b.TopPeerCategory.Encode(buf)
}
