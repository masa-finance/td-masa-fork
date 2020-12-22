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

// AccountConfirmPasswordEmailRequest represents TL type `account.confirmPasswordEmail#8fdf1920`.
// Verify an email to use as 2FA recovery method.
//
// See https://core.telegram.org/method/account.confirmPasswordEmail for reference.
type AccountConfirmPasswordEmailRequest struct {
	// The phone code that was received after setting a recovery email
	Code string
}

// AccountConfirmPasswordEmailRequestTypeID is TL type id of AccountConfirmPasswordEmailRequest.
const AccountConfirmPasswordEmailRequestTypeID = 0x8fdf1920

// String implements fmt.Stringer.
func (c *AccountConfirmPasswordEmailRequest) String() string {
	if c == nil {
		return "AccountConfirmPasswordEmailRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountConfirmPasswordEmailRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tCode: ")
	sb.WriteString(fmt.Sprint(c.Code))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *AccountConfirmPasswordEmailRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.confirmPasswordEmail#8fdf1920 as nil")
	}
	b.PutID(AccountConfirmPasswordEmailRequestTypeID)
	b.PutString(c.Code)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountConfirmPasswordEmailRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.confirmPasswordEmail#8fdf1920 to nil")
	}
	if err := b.ConsumeID(AccountConfirmPasswordEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.confirmPasswordEmail#8fdf1920: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPasswordEmail#8fdf1920: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountConfirmPasswordEmailRequest.
var (
	_ bin.Encoder = &AccountConfirmPasswordEmailRequest{}
	_ bin.Decoder = &AccountConfirmPasswordEmailRequest{}
)

// AccountConfirmPasswordEmail invokes method account.confirmPasswordEmail#8fdf1920 returning error if any.
// Verify an email to use as 2FA recovery method.
//
// See https://core.telegram.org/method/account.confirmPasswordEmail for reference.
func (c *Client) AccountConfirmPasswordEmail(ctx context.Context, code string) (bool, error) {
	var result BoolBox

	request := &AccountConfirmPasswordEmailRequest{
		Code: code,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
