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

// PaymentsPaymentReceipt represents TL type `payments.paymentReceipt#500911e1`.
// Receipt
//
// See https://core.telegram.org/constructor/payments.paymentReceipt for reference.
type PaymentsPaymentReceipt struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Date of generation
	Date int
	// Bot ID
	BotID int
	// Invoice
	Invoice Invoice
	// Provider ID
	ProviderID int
	// Info
	//
	// Use SetInfo and GetInfo helpers.
	Info PaymentRequestedInfo
	// Selected shipping option
	//
	// Use SetShipping and GetShipping helpers.
	Shipping ShippingOption
	// Three-letter ISO 4217 currency code
	Currency string
	// Total amount in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int64
	// Payment credential name
	CredentialsTitle string
	// Users
	Users []UserClass
}

// PaymentsPaymentReceiptTypeID is TL type id of PaymentsPaymentReceipt.
const PaymentsPaymentReceiptTypeID = 0x500911e1

// String implements fmt.Stringer.
func (p *PaymentsPaymentReceipt) String() string {
	if p == nil {
		return "PaymentsPaymentReceipt(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PaymentsPaymentReceipt")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(p.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(p.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tBotID: ")
	sb.WriteString(fmt.Sprint(p.BotID))
	sb.WriteString(",\n")
	sb.WriteString("\tInvoice: ")
	sb.WriteString(p.Invoice.String())
	sb.WriteString(",\n")
	sb.WriteString("\tProviderID: ")
	sb.WriteString(fmt.Sprint(p.ProviderID))
	sb.WriteString(",\n")
	if p.Flags.Has(0) {
		sb.WriteString("\tInfo: ")
		sb.WriteString(p.Info.String())
		sb.WriteString(",\n")
	}
	if p.Flags.Has(1) {
		sb.WriteString("\tShipping: ")
		sb.WriteString(p.Shipping.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("\tCurrency: ")
	sb.WriteString(fmt.Sprint(p.Currency))
	sb.WriteString(",\n")
	sb.WriteString("\tTotalAmount: ")
	sb.WriteString(fmt.Sprint(p.TotalAmount))
	sb.WriteString(",\n")
	sb.WriteString("\tCredentialsTitle: ")
	sb.WriteString(fmt.Sprint(p.CredentialsTitle))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range p.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PaymentsPaymentReceipt) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentReceipt#500911e1 as nil")
	}
	b.PutID(PaymentsPaymentReceiptTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field flags: %w", err)
	}
	b.PutInt(p.Date)
	b.PutInt(p.BotID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field invoice: %w", err)
	}
	b.PutInt(p.ProviderID)
	if p.Flags.Has(0) {
		if err := p.Info.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.Shipping.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field shipping: %w", err)
		}
	}
	b.PutString(p.Currency)
	b.PutLong(p.TotalAmount)
	b.PutString(p.CredentialsTitle)
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetInfo sets value of Info conditional field.
func (p *PaymentsPaymentReceipt) SetInfo(value PaymentRequestedInfo) {
	p.Flags.Set(0)
	p.Info = value
}

// GetInfo returns value of Info conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentReceipt) GetInfo() (value PaymentRequestedInfo, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Info, true
}

// SetShipping sets value of Shipping conditional field.
func (p *PaymentsPaymentReceipt) SetShipping(value ShippingOption) {
	p.Flags.Set(1)
	p.Shipping = value
}

// GetShipping returns value of Shipping conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentReceipt) GetShipping() (value ShippingOption, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Shipping, true
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentReceipt) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentReceipt#500911e1 to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentReceiptTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field bot_id: %w", err)
		}
		p.BotID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field invoice: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field provider_id: %w", err)
		}
		p.ProviderID = value
	}
	if p.Flags.Has(0) {
		if err := p.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.Shipping.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field shipping: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field currency: %w", err)
		}
		p.Currency = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field total_amount: %w", err)
		}
		p.TotalAmount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field credentials_title: %w", err)
		}
		p.CredentialsTitle = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsPaymentReceipt.
var (
	_ bin.Encoder = &PaymentsPaymentReceipt{}
	_ bin.Decoder = &PaymentsPaymentReceipt{}
)
