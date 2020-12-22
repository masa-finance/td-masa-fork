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

// AuthSentCodeTypeApp represents TL type `auth.sentCodeTypeApp#3dbb5986`.
// The code was sent through the telegram app
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeApp for reference.
type AuthSentCodeTypeApp struct {
	// Length of the code in bytes
	Length int
}

// AuthSentCodeTypeAppTypeID is TL type id of AuthSentCodeTypeApp.
const AuthSentCodeTypeAppTypeID = 0x3dbb5986

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeApp) String() string {
	if s == nil {
		return "AuthSentCodeTypeApp(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthSentCodeTypeApp")
	sb.WriteString("{\n")
	sb.WriteString("\tLength: ")
	sb.WriteString(fmt.Sprint(s.Length))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeApp) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeApp#3dbb5986 as nil")
	}
	b.PutID(AuthSentCodeTypeAppTypeID)
	b.PutInt(s.Length)
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeApp) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeApp#3dbb5986 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeAppTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeApp#3dbb5986: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeApp#3dbb5986: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeApp) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeApp.
var (
	_ bin.Encoder = &AuthSentCodeTypeApp{}
	_ bin.Decoder = &AuthSentCodeTypeApp{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeApp{}
)

// AuthSentCodeTypeSms represents TL type `auth.sentCodeTypeSms#c000bba2`.
// The code was sent via SMS
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeSms for reference.
type AuthSentCodeTypeSms struct {
	// Length of the code in bytes
	Length int
}

// AuthSentCodeTypeSmsTypeID is TL type id of AuthSentCodeTypeSms.
const AuthSentCodeTypeSmsTypeID = 0xc000bba2

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeSms) String() string {
	if s == nil {
		return "AuthSentCodeTypeSms(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthSentCodeTypeSms")
	sb.WriteString("{\n")
	sb.WriteString("\tLength: ")
	sb.WriteString(fmt.Sprint(s.Length))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeSms) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeSms#c000bba2 as nil")
	}
	b.PutID(AuthSentCodeTypeSmsTypeID)
	b.PutInt(s.Length)
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeSms) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeSms#c000bba2 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeSmsTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeSms#c000bba2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeSms#c000bba2: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeSms) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeSms.
var (
	_ bin.Encoder = &AuthSentCodeTypeSms{}
	_ bin.Decoder = &AuthSentCodeTypeSms{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeSms{}
)

// AuthSentCodeTypeCall represents TL type `auth.sentCodeTypeCall#5353e5a7`.
// The code will be sent via a phone call: a synthesized voice will tell the user which verification code to input.
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeCall for reference.
type AuthSentCodeTypeCall struct {
	// Length of the verification code
	Length int
}

// AuthSentCodeTypeCallTypeID is TL type id of AuthSentCodeTypeCall.
const AuthSentCodeTypeCallTypeID = 0x5353e5a7

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeCall) String() string {
	if s == nil {
		return "AuthSentCodeTypeCall(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthSentCodeTypeCall")
	sb.WriteString("{\n")
	sb.WriteString("\tLength: ")
	sb.WriteString(fmt.Sprint(s.Length))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeCall) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeCall#5353e5a7 as nil")
	}
	b.PutID(AuthSentCodeTypeCallTypeID)
	b.PutInt(s.Length)
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeCall) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeCall#5353e5a7 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeCall#5353e5a7: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeCall#5353e5a7: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeCall) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeCall.
var (
	_ bin.Encoder = &AuthSentCodeTypeCall{}
	_ bin.Decoder = &AuthSentCodeTypeCall{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeCall{}
)

// AuthSentCodeTypeFlashCall represents TL type `auth.sentCodeTypeFlashCall#ab03c6d9`.
// The code will be sent via a flash phone call, that will be closed immediately. The phone code will then be the phone number itself, just make sure that the phone number matches the specified pattern.
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeFlashCall for reference.
type AuthSentCodeTypeFlashCall struct {
	// pattern to match
	Pattern string
}

// AuthSentCodeTypeFlashCallTypeID is TL type id of AuthSentCodeTypeFlashCall.
const AuthSentCodeTypeFlashCallTypeID = 0xab03c6d9

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeFlashCall) String() string {
	if s == nil {
		return "AuthSentCodeTypeFlashCall(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AuthSentCodeTypeFlashCall")
	sb.WriteString("{\n")
	sb.WriteString("\tPattern: ")
	sb.WriteString(fmt.Sprint(s.Pattern))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeFlashCall) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeFlashCall#ab03c6d9 as nil")
	}
	b.PutID(AuthSentCodeTypeFlashCallTypeID)
	b.PutString(s.Pattern)
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeFlashCall) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeFlashCall#ab03c6d9 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeFlashCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeFlashCall#ab03c6d9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeFlashCall#ab03c6d9: field pattern: %w", err)
		}
		s.Pattern = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeFlashCall) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeFlashCall.
var (
	_ bin.Encoder = &AuthSentCodeTypeFlashCall{}
	_ bin.Decoder = &AuthSentCodeTypeFlashCall{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeFlashCall{}
)

// AuthSentCodeTypeClass represents auth.SentCodeType generic type.
//
// See https://core.telegram.org/type/auth.SentCodeType for reference.
//
// Example:
//  g, err := DecodeAuthSentCodeType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthSentCodeTypeApp: // auth.sentCodeTypeApp#3dbb5986
//  case *AuthSentCodeTypeSms: // auth.sentCodeTypeSms#c000bba2
//  case *AuthSentCodeTypeCall: // auth.sentCodeTypeCall#5353e5a7
//  case *AuthSentCodeTypeFlashCall: // auth.sentCodeTypeFlashCall#ab03c6d9
//  default: panic(v)
//  }
type AuthSentCodeTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthSentCodeTypeClass
	fmt.Stringer
}

// DecodeAuthSentCodeType implements binary de-serialization for AuthSentCodeTypeClass.
func DecodeAuthSentCodeType(buf *bin.Buffer) (AuthSentCodeTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthSentCodeTypeAppTypeID:
		// Decoding auth.sentCodeTypeApp#3dbb5986.
		v := AuthSentCodeTypeApp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeSmsTypeID:
		// Decoding auth.sentCodeTypeSms#c000bba2.
		v := AuthSentCodeTypeSms{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeCallTypeID:
		// Decoding auth.sentCodeTypeCall#5353e5a7.
		v := AuthSentCodeTypeCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeFlashCallTypeID:
		// Decoding auth.sentCodeTypeFlashCall#ab03c6d9.
		v := AuthSentCodeTypeFlashCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthSentCodeType boxes the AuthSentCodeTypeClass providing a helper.
type AuthSentCodeTypeBox struct {
	SentCodeType AuthSentCodeTypeClass
}

// Decode implements bin.Decoder for AuthSentCodeTypeBox.
func (b *AuthSentCodeTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthSentCodeTypeBox to nil")
	}
	v, err := DecodeAuthSentCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SentCodeType = v
	return nil
}

// Encode implements bin.Encode for AuthSentCodeTypeBox.
func (b *AuthSentCodeTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SentCodeType == nil {
		return fmt.Errorf("unable to encode AuthSentCodeTypeClass as nil")
	}
	return b.SentCodeType.Encode(buf)
}
