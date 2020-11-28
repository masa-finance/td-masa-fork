// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// AccountGetThemesRequest represents TL type `account.getThemes#285946f8`.
type AccountGetThemesRequest struct {
	// Format field of AccountGetThemesRequest.
	Format string
	// Hash field of AccountGetThemesRequest.
	Hash int
}

// AccountGetThemesRequestTypeID is TL type id of AccountGetThemesRequest.
const AccountGetThemesRequestTypeID = 0x285946f8

// Encode implements bin.Encoder.
func (g *AccountGetThemesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getThemes#285946f8 as nil")
	}
	b.PutID(AccountGetThemesRequestTypeID)
	b.PutString(g.Format)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetThemesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getThemes#285946f8 to nil")
	}
	if err := b.ConsumeID(AccountGetThemesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getThemes#285946f8: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getThemes#285946f8: field format: %w", err)
		}
		g.Format = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.getThemes#285946f8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetThemesRequest.
var (
	_ bin.Encoder = &AccountGetThemesRequest{}
	_ bin.Decoder = &AccountGetThemesRequest{}
)