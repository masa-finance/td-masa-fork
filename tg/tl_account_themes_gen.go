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

// AccountThemesNotModified represents TL type `account.themesNotModified#f41eb622`.
// No new themes were installed
//
// See https://core.telegram.org/constructor/account.themesNotModified for reference.
type AccountThemesNotModified struct {
}

// AccountThemesNotModifiedTypeID is TL type id of AccountThemesNotModified.
const AccountThemesNotModifiedTypeID = 0xf41eb622

func (t *AccountThemesNotModified) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountThemesNotModified) String() string {
	if t == nil {
		return "AccountThemesNotModified(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountThemesNotModified")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *AccountThemesNotModified) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themesNotModified#f41eb622 as nil")
	}
	b.PutID(AccountThemesNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountThemesNotModified) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themesNotModified#f41eb622 to nil")
	}
	if err := b.ConsumeID(AccountThemesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.themesNotModified#f41eb622: %w", err)
	}
	return nil
}

// construct implements constructor of AccountThemesClass.
func (t AccountThemesNotModified) construct() AccountThemesClass { return &t }

// Ensuring interfaces in compile-time for AccountThemesNotModified.
var (
	_ bin.Encoder = &AccountThemesNotModified{}
	_ bin.Decoder = &AccountThemesNotModified{}

	_ AccountThemesClass = &AccountThemesNotModified{}
)

// AccountThemes represents TL type `account.themes#7f676421`.
// Installed themes
//
// See https://core.telegram.org/constructor/account.themes for reference.
type AccountThemes struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Themes
	Themes []Theme
}

// AccountThemesTypeID is TL type id of AccountThemes.
const AccountThemesTypeID = 0x7f676421

func (t *AccountThemes) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Hash == 0) {
		return false
	}
	if !(t.Themes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountThemes) String() string {
	if t == nil {
		return "AccountThemes(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountThemes")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(t.Hash))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range t.Themes {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *AccountThemes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themes#7f676421 as nil")
	}
	b.PutID(AccountThemesTypeID)
	b.PutInt(t.Hash)
	b.PutVectorHeader(len(t.Themes))
	for idx, v := range t.Themes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.themes#7f676421: field themes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountThemes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themes#7f676421 to nil")
	}
	if err := b.ConsumeID(AccountThemesTypeID); err != nil {
		return fmt.Errorf("unable to decode account.themes#7f676421: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.themes#7f676421: field hash: %w", err)
		}
		t.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.themes#7f676421: field themes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Theme
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.themes#7f676421: field themes: %w", err)
			}
			t.Themes = append(t.Themes, value)
		}
	}
	return nil
}

// construct implements constructor of AccountThemesClass.
func (t AccountThemes) construct() AccountThemesClass { return &t }

// Ensuring interfaces in compile-time for AccountThemes.
var (
	_ bin.Encoder = &AccountThemes{}
	_ bin.Decoder = &AccountThemes{}

	_ AccountThemesClass = &AccountThemes{}
)

// AccountThemesClass represents account.Themes generic type.
//
// See https://core.telegram.org/type/account.Themes for reference.
//
// Example:
//  g, err := DecodeAccountThemes(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AccountThemesNotModified: // account.themesNotModified#f41eb622
//  case *AccountThemes: // account.themes#7f676421
//  default: panic(v)
//  }
type AccountThemesClass interface {
	bin.Encoder
	bin.Decoder
	construct() AccountThemesClass

	fmt.Stringer
	Zero() bool
}

// DecodeAccountThemes implements binary de-serialization for AccountThemesClass.
func DecodeAccountThemes(buf *bin.Buffer) (AccountThemesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountThemesNotModifiedTypeID:
		// Decoding account.themesNotModified#f41eb622.
		v := AccountThemesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", err)
		}
		return &v, nil
	case AccountThemesTypeID:
		// Decoding account.themes#7f676421.
		v := AccountThemes{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountThemes boxes the AccountThemesClass providing a helper.
type AccountThemesBox struct {
	Themes AccountThemesClass
}

// Decode implements bin.Decoder for AccountThemesBox.
func (b *AccountThemesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountThemesBox to nil")
	}
	v, err := DecodeAccountThemes(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Themes = v
	return nil
}

// Encode implements bin.Encode for AccountThemesBox.
func (b *AccountThemesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Themes == nil {
		return fmt.Errorf("unable to encode AccountThemesClass as nil")
	}
	return b.Themes.Encode(buf)
}
