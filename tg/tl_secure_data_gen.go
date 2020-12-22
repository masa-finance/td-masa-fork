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

// SecureData represents TL type `secureData#8aeabec3`.
// Secure passport data, for more info see the passport docs »
//
// See https://core.telegram.org/constructor/secureData for reference.
type SecureData struct {
	// Data
	Data []byte
	// Data hash
	DataHash []byte
	// Secret
	Secret []byte
}

// SecureDataTypeID is TL type id of SecureData.
const SecureDataTypeID = 0x8aeabec3

// String implements fmt.Stringer.
func (s *SecureData) String() string {
	if s == nil {
		return "SecureData(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SecureData")
	sb.WriteString("{\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(s.Data))
	sb.WriteString(",\n")
	sb.WriteString("\tDataHash: ")
	sb.WriteString(fmt.Sprint(s.DataHash))
	sb.WriteString(",\n")
	sb.WriteString("\tSecret: ")
	sb.WriteString(fmt.Sprint(s.Secret))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *SecureData) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureData#8aeabec3 as nil")
	}
	b.PutID(SecureDataTypeID)
	b.PutBytes(s.Data)
	b.PutBytes(s.DataHash)
	b.PutBytes(s.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureData) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureData#8aeabec3 to nil")
	}
	if err := b.ConsumeID(SecureDataTypeID); err != nil {
		return fmt.Errorf("unable to decode secureData#8aeabec3: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureData#8aeabec3: field data: %w", err)
		}
		s.Data = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureData#8aeabec3: field data_hash: %w", err)
		}
		s.DataHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureData#8aeabec3: field secret: %w", err)
		}
		s.Secret = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureData.
var (
	_ bin.Encoder = &SecureData{}
	_ bin.Decoder = &SecureData{}
)
