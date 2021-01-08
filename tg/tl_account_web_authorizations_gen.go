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

// AccountWebAuthorizations represents TL type `account.webAuthorizations#ed56c9fc`.
// Web authorizations
//
// See https://core.telegram.org/constructor/account.webAuthorizations for reference.
type AccountWebAuthorizations struct {
	// Web authorization list
	Authorizations []WebAuthorization
	// Users
	Users []UserClass
}

// AccountWebAuthorizationsTypeID is TL type id of AccountWebAuthorizations.
const AccountWebAuthorizationsTypeID = 0xed56c9fc

func (w *AccountWebAuthorizations) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Authorizations == nil) {
		return false
	}
	if !(w.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *AccountWebAuthorizations) String() string {
	if w == nil {
		return "AccountWebAuthorizations(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountWebAuthorizations")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range w.Authorizations {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range w.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (w *AccountWebAuthorizations) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.webAuthorizations#ed56c9fc as nil")
	}
	b.PutID(AccountWebAuthorizationsTypeID)
	b.PutVectorHeader(len(w.Authorizations))
	for idx, v := range w.Authorizations {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.webAuthorizations#ed56c9fc: field authorizations element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(w.Users))
	for idx, v := range w.Users {
		if v == nil {
			return fmt.Errorf("unable to encode account.webAuthorizations#ed56c9fc: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.webAuthorizations#ed56c9fc: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *AccountWebAuthorizations) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.webAuthorizations#ed56c9fc to nil")
	}
	if err := b.ConsumeID(AccountWebAuthorizationsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field authorizations: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value WebAuthorization
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field authorizations: %w", err)
			}
			w.Authorizations = append(w.Authorizations, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field users: %w", err)
			}
			w.Users = append(w.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountWebAuthorizations.
var (
	_ bin.Encoder = &AccountWebAuthorizations{}
	_ bin.Decoder = &AccountWebAuthorizations{}
)
