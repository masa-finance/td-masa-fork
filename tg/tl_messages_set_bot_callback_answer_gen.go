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

// MessagesSetBotCallbackAnswerRequest represents TL type `messages.setBotCallbackAnswer#d58f130a`.
// Set the callback answer to a user button press (bots only)
//
// See https://core.telegram.org/method/messages.setBotCallbackAnswer for reference.
type MessagesSetBotCallbackAnswerRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether to show the message as a popup instead of a toast notification
	Alert bool
	// Query ID
	QueryID int64
	// Popup to show
	//
	// Use SetMessage and GetMessage helpers.
	Message string
	// URL to open
	//
	// Use SetURL and GetURL helpers.
	URL string
	// Cache validity
	CacheTime int
}

// MessagesSetBotCallbackAnswerRequestTypeID is TL type id of MessagesSetBotCallbackAnswerRequest.
const MessagesSetBotCallbackAnswerRequestTypeID = 0xd58f130a

// String implements fmt.Stringer.
func (s *MessagesSetBotCallbackAnswerRequest) String() string {
	if s == nil {
		return "MessagesSetBotCallbackAnswerRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSetBotCallbackAnswerRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(s.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tQueryID: ")
	sb.WriteString(fmt.Sprint(s.QueryID))
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tMessage: ")
		sb.WriteString(fmt.Sprint(s.Message))
		sb.WriteString(",\n")
	}
	if s.Flags.Has(2) {
		sb.WriteString("\tURL: ")
		sb.WriteString(fmt.Sprint(s.URL))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tCacheTime: ")
	sb.WriteString(fmt.Sprint(s.CacheTime))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSetBotCallbackAnswerRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotCallbackAnswer#d58f130a as nil")
	}
	b.PutID(MessagesSetBotCallbackAnswerRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setBotCallbackAnswer#d58f130a: field flags: %w", err)
	}
	b.PutLong(s.QueryID)
	if s.Flags.Has(0) {
		b.PutString(s.Message)
	}
	if s.Flags.Has(2) {
		b.PutString(s.URL)
	}
	b.PutInt(s.CacheTime)
	return nil
}

// SetAlert sets value of Alert conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetAlert(value bool) {
	if value {
		s.Flags.Set(1)
	} else {
		s.Flags.Unset(1)
	}
}

// SetMessage sets value of Message conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetMessage(value string) {
	s.Flags.Set(0)
	s.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotCallbackAnswerRequest) GetMessage() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Message, true
}

// SetURL sets value of URL conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetURL(value string) {
	s.Flags.Set(2)
	s.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotCallbackAnswerRequest) GetURL() (value string, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.URL, true
}

// Decode implements bin.Decoder.
func (s *MessagesSetBotCallbackAnswerRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotCallbackAnswer#d58f130a to nil")
	}
	if err := b.ConsumeID(MessagesSetBotCallbackAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field flags: %w", err)
		}
	}
	s.Alert = s.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field query_id: %w", err)
		}
		s.QueryID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field url: %w", err)
		}
		s.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field cache_time: %w", err)
		}
		s.CacheTime = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetBotCallbackAnswerRequest.
var (
	_ bin.Encoder = &MessagesSetBotCallbackAnswerRequest{}
	_ bin.Decoder = &MessagesSetBotCallbackAnswerRequest{}
)

// MessagesSetBotCallbackAnswer invokes method messages.setBotCallbackAnswer#d58f130a returning error if any.
// Set the callback answer to a user button press (bots only)
//
// See https://core.telegram.org/method/messages.setBotCallbackAnswer for reference.
func (c *Client) MessagesSetBotCallbackAnswer(ctx context.Context, request *MessagesSetBotCallbackAnswerRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
