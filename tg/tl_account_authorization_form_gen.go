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

// AccountAuthorizationForm represents TL type `account.authorizationForm#ad2e1cd8`.
// Telegram Passport¹ authorization form
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/account.authorizationForm for reference.
type AccountAuthorizationForm struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Required Telegram Passport¹ documents
	//
	// Links:
	//  1) https://core.telegram.org/passport
	RequiredTypes []SecureRequiredTypeClass
	// Already submitted Telegram Passport¹ documents
	//
	// Links:
	//  1) https://core.telegram.org/passport
	Values []SecureValue
	// Telegram Passport¹ errors
	//
	// Links:
	//  1) https://core.telegram.org/passport
	Errors []SecureValueErrorClass
	// Info about the bot to which the form will be submitted
	Users []UserClass
	// URL of the service's privacy policy
	//
	// Use SetPrivacyPolicyURL and GetPrivacyPolicyURL helpers.
	PrivacyPolicyURL string
}

// AccountAuthorizationFormTypeID is TL type id of AccountAuthorizationForm.
const AccountAuthorizationFormTypeID = 0xad2e1cd8

func (a *AccountAuthorizationForm) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.RequiredTypes == nil) {
		return false
	}
	if !(a.Values == nil) {
		return false
	}
	if !(a.Errors == nil) {
		return false
	}
	if !(a.Users == nil) {
		return false
	}
	if !(a.PrivacyPolicyURL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccountAuthorizationForm) String() string {
	if a == nil {
		return "AccountAuthorizationForm(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountAuthorizationForm")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(a.Flags))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range a.RequiredTypes {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range a.Values {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range a.Errors {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range a.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	if a.Flags.Has(0) {
		sb.WriteString("\tPrivacyPolicyURL: ")
		sb.WriteString(fmt.Sprint(a.PrivacyPolicyURL))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *AccountAuthorizationForm) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.authorizationForm#ad2e1cd8 as nil")
	}
	b.PutID(AccountAuthorizationFormTypeID)
	if !(a.PrivacyPolicyURL == "") {
		a.Flags.Set(0)
	}
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field flags: %w", err)
	}
	b.PutVectorHeader(len(a.RequiredTypes))
	for idx, v := range a.RequiredTypes {
		if v == nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field required_types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field required_types element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Values))
	for idx, v := range a.Values {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field values element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Errors))
	for idx, v := range a.Errors {
		if v == nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field errors element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field errors element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Users))
	for idx, v := range a.Users {
		if v == nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.authorizationForm#ad2e1cd8: field users element with index %d: %w", idx, err)
		}
	}
	if a.Flags.Has(0) {
		b.PutString(a.PrivacyPolicyURL)
	}
	return nil
}

// SetPrivacyPolicyURL sets value of PrivacyPolicyURL conditional field.
func (a *AccountAuthorizationForm) SetPrivacyPolicyURL(value string) {
	a.Flags.Set(0)
	a.PrivacyPolicyURL = value
}

// GetPrivacyPolicyURL returns value of PrivacyPolicyURL conditional field and
// boolean which is true if field was set.
func (a *AccountAuthorizationForm) GetPrivacyPolicyURL() (value string, ok bool) {
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.PrivacyPolicyURL, true
}

// Decode implements bin.Decoder.
func (a *AccountAuthorizationForm) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.authorizationForm#ad2e1cd8 to nil")
	}
	if err := b.ConsumeID(AccountAuthorizationFormTypeID); err != nil {
		return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field flags: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field required_types: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureRequiredType(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field required_types: %w", err)
			}
			a.RequiredTypes = append(a.RequiredTypes, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field values: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SecureValue
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field values: %w", err)
			}
			a.Values = append(a.Values, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field errors: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureValueError(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field errors: %w", err)
			}
			a.Errors = append(a.Errors, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field users: %w", err)
			}
			a.Users = append(a.Users, value)
		}
	}
	if a.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.authorizationForm#ad2e1cd8: field privacy_policy_url: %w", err)
		}
		a.PrivacyPolicyURL = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountAuthorizationForm.
var (
	_ bin.Encoder = &AccountAuthorizationForm{}
	_ bin.Decoder = &AccountAuthorizationForm{}
)
