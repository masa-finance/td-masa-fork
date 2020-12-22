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

// HelpTermsOfServiceUpdateEmpty represents TL type `help.termsOfServiceUpdateEmpty#e3309f7f`.
// No changes were made to telegram's terms of service
//
// See https://core.telegram.org/constructor/help.termsOfServiceUpdateEmpty for reference.
type HelpTermsOfServiceUpdateEmpty struct {
	// New TOS updates will have to be queried using help.getTermsOfServiceUpdate in expires seconds
	Expires int
}

// HelpTermsOfServiceUpdateEmptyTypeID is TL type id of HelpTermsOfServiceUpdateEmpty.
const HelpTermsOfServiceUpdateEmptyTypeID = 0xe3309f7f

// String implements fmt.Stringer.
func (t *HelpTermsOfServiceUpdateEmpty) String() string {
	if t == nil {
		return "HelpTermsOfServiceUpdateEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpTermsOfServiceUpdateEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tExpires: ")
	sb.WriteString(fmt.Sprint(t.Expires))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *HelpTermsOfServiceUpdateEmpty) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfServiceUpdateEmpty#e3309f7f as nil")
	}
	b.PutID(HelpTermsOfServiceUpdateEmptyTypeID)
	b.PutInt(t.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (t *HelpTermsOfServiceUpdateEmpty) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfServiceUpdateEmpty#e3309f7f to nil")
	}
	if err := b.ConsumeID(HelpTermsOfServiceUpdateEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.termsOfServiceUpdateEmpty#e3309f7f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfServiceUpdateEmpty#e3309f7f: field expires: %w", err)
		}
		t.Expires = value
	}
	return nil
}

// construct implements constructor of HelpTermsOfServiceUpdateClass.
func (t HelpTermsOfServiceUpdateEmpty) construct() HelpTermsOfServiceUpdateClass { return &t }

// Ensuring interfaces in compile-time for HelpTermsOfServiceUpdateEmpty.
var (
	_ bin.Encoder = &HelpTermsOfServiceUpdateEmpty{}
	_ bin.Decoder = &HelpTermsOfServiceUpdateEmpty{}

	_ HelpTermsOfServiceUpdateClass = &HelpTermsOfServiceUpdateEmpty{}
)

// HelpTermsOfServiceUpdate represents TL type `help.termsOfServiceUpdate#28ecf961`.
// Info about an update of telegram's terms of service. If the terms of service are declined, then the account.deleteAccount method should be called with the reason "Decline ToS update"
//
// See https://core.telegram.org/constructor/help.termsOfServiceUpdate for reference.
type HelpTermsOfServiceUpdate struct {
	// New TOS updates will have to be queried using help.getTermsOfServiceUpdate in expires seconds
	Expires int
	// New terms of service
	TermsOfService HelpTermsOfService
}

// HelpTermsOfServiceUpdateTypeID is TL type id of HelpTermsOfServiceUpdate.
const HelpTermsOfServiceUpdateTypeID = 0x28ecf961

// String implements fmt.Stringer.
func (t *HelpTermsOfServiceUpdate) String() string {
	if t == nil {
		return "HelpTermsOfServiceUpdate(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpTermsOfServiceUpdate")
	sb.WriteString("{\n")
	sb.WriteString("\tExpires: ")
	sb.WriteString(fmt.Sprint(t.Expires))
	sb.WriteString(",\n")
	sb.WriteString("\tTermsOfService: ")
	sb.WriteString(t.TermsOfService.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *HelpTermsOfServiceUpdate) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfServiceUpdate#28ecf961 as nil")
	}
	b.PutID(HelpTermsOfServiceUpdateTypeID)
	b.PutInt(t.Expires)
	if err := t.TermsOfService.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.termsOfServiceUpdate#28ecf961: field terms_of_service: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *HelpTermsOfServiceUpdate) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfServiceUpdate#28ecf961 to nil")
	}
	if err := b.ConsumeID(HelpTermsOfServiceUpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode help.termsOfServiceUpdate#28ecf961: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfServiceUpdate#28ecf961: field expires: %w", err)
		}
		t.Expires = value
	}
	{
		if err := t.TermsOfService.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.termsOfServiceUpdate#28ecf961: field terms_of_service: %w", err)
		}
	}
	return nil
}

// construct implements constructor of HelpTermsOfServiceUpdateClass.
func (t HelpTermsOfServiceUpdate) construct() HelpTermsOfServiceUpdateClass { return &t }

// Ensuring interfaces in compile-time for HelpTermsOfServiceUpdate.
var (
	_ bin.Encoder = &HelpTermsOfServiceUpdate{}
	_ bin.Decoder = &HelpTermsOfServiceUpdate{}

	_ HelpTermsOfServiceUpdateClass = &HelpTermsOfServiceUpdate{}
)

// HelpTermsOfServiceUpdateClass represents help.TermsOfServiceUpdate generic type.
//
// See https://core.telegram.org/type/help.TermsOfServiceUpdate for reference.
//
// Example:
//  g, err := DecodeHelpTermsOfServiceUpdate(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpTermsOfServiceUpdateEmpty: // help.termsOfServiceUpdateEmpty#e3309f7f
//  case *HelpTermsOfServiceUpdate: // help.termsOfServiceUpdate#28ecf961
//  default: panic(v)
//  }
type HelpTermsOfServiceUpdateClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpTermsOfServiceUpdateClass
	fmt.Stringer
}

// DecodeHelpTermsOfServiceUpdate implements binary de-serialization for HelpTermsOfServiceUpdateClass.
func DecodeHelpTermsOfServiceUpdate(buf *bin.Buffer) (HelpTermsOfServiceUpdateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpTermsOfServiceUpdateEmptyTypeID:
		// Decoding help.termsOfServiceUpdateEmpty#e3309f7f.
		v := HelpTermsOfServiceUpdateEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpTermsOfServiceUpdateClass: %w", err)
		}
		return &v, nil
	case HelpTermsOfServiceUpdateTypeID:
		// Decoding help.termsOfServiceUpdate#28ecf961.
		v := HelpTermsOfServiceUpdate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpTermsOfServiceUpdateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpTermsOfServiceUpdateClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpTermsOfServiceUpdate boxes the HelpTermsOfServiceUpdateClass providing a helper.
type HelpTermsOfServiceUpdateBox struct {
	TermsOfServiceUpdate HelpTermsOfServiceUpdateClass
}

// Decode implements bin.Decoder for HelpTermsOfServiceUpdateBox.
func (b *HelpTermsOfServiceUpdateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpTermsOfServiceUpdateBox to nil")
	}
	v, err := DecodeHelpTermsOfServiceUpdate(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TermsOfServiceUpdate = v
	return nil
}

// Encode implements bin.Encode for HelpTermsOfServiceUpdateBox.
func (b *HelpTermsOfServiceUpdateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TermsOfServiceUpdate == nil {
		return fmt.Errorf("unable to encode HelpTermsOfServiceUpdateClass as nil")
	}
	return b.TermsOfServiceUpdate.Encode(buf)
}
