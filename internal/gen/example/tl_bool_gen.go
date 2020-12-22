// Code generated by gotdgen, DO NOT EDIT.

package td

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

// False represents TL type `false#bc799737`.
//
// See https://localhost:80/doc/constructor/false for reference.
type False struct {
}

// FalseTypeID is TL type id of False.
const FalseTypeID = 0xbc799737

// String implements fmt.Stringer.
func (f *False) String() string {
	if f == nil {
		return "False(nil)"
	}
	var sb strings.Builder
	sb.WriteString("False")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (f *False) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode false#bc799737 as nil")
	}
	b.PutID(FalseTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *False) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode false#bc799737 to nil")
	}
	if err := b.ConsumeID(FalseTypeID); err != nil {
		return fmt.Errorf("unable to decode false#bc799737: %w", err)
	}
	return nil
}

// construct implements constructor of BoolClass.
func (f False) construct() BoolClass { return &f }

// Ensuring interfaces in compile-time for False.
var (
	_ bin.Encoder = &False{}
	_ bin.Decoder = &False{}

	_ BoolClass = &False{}
)

// True represents TL type `true#997275b5`.
//
// See https://localhost:80/doc/constructor/true for reference.
type True struct {
}

// TrueTypeID is TL type id of True.
const TrueTypeID = 0x997275b5

// String implements fmt.Stringer.
func (t *True) String() string {
	if t == nil {
		return "True(nil)"
	}
	var sb strings.Builder
	sb.WriteString("True")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *True) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode true#997275b5 as nil")
	}
	b.PutID(TrueTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *True) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode true#997275b5 to nil")
	}
	if err := b.ConsumeID(TrueTypeID); err != nil {
		return fmt.Errorf("unable to decode true#997275b5: %w", err)
	}
	return nil
}

// construct implements constructor of BoolClass.
func (t True) construct() BoolClass { return &t }

// Ensuring interfaces in compile-time for True.
var (
	_ bin.Encoder = &True{}
	_ bin.Decoder = &True{}

	_ BoolClass = &True{}
)

// BoolClass represents Bool generic type.
//
// See https://localhost:80/doc/type/Bool for reference.
//
// Example:
//  g, err := DecodeBool(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *False: // false#bc799737
//  case *True: // true#997275b5
//  default: panic(v)
//  }
type BoolClass interface {
	bin.Encoder
	bin.Decoder
	construct() BoolClass
	fmt.Stringer
}

// DecodeBool implements binary de-serialization for BoolClass.
func DecodeBool(buf *bin.Buffer) (BoolClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case FalseTypeID:
		// Decoding false#bc799737.
		v := False{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	case TrueTypeID:
		// Decoding true#997275b5.
		v := True{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BoolClass: %w", bin.NewUnexpectedID(id))
	}
}

// Bool boxes the BoolClass providing a helper.
type BoolBox struct {
	Bool BoolClass
}

// Decode implements bin.Decoder for BoolBox.
func (b *BoolBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BoolBox to nil")
	}
	v, err := DecodeBool(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Bool = v
	return nil
}

// Encode implements bin.Encode for BoolBox.
func (b *BoolBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Bool == nil {
		return fmt.Errorf("unable to encode BoolClass as nil")
	}
	return b.Bool.Encode(buf)
}
