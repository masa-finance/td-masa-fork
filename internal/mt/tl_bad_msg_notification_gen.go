// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// BadMsgNotification represents TL type `bad_msg_notification#a7eff811`.
type BadMsgNotification struct {
	// BadMsgID field of BadMsgNotification.
	BadMsgID int64
	// BadMsgSeqno field of BadMsgNotification.
	BadMsgSeqno int
	// ErrorCode field of BadMsgNotification.
	ErrorCode int
}

// BadMsgNotificationTypeID is TL type id of BadMsgNotification.
const BadMsgNotificationTypeID = 0xa7eff811

// String implements fmt.Stringer.
func (b *BadMsgNotification) String() string {
	if b == nil {
		return "BadMsgNotification(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BadMsgNotification")
	sb.WriteString("{\n")
	sb.WriteString("\tBadMsgID: ")
	sb.WriteString(fmt.Sprint(b.BadMsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tBadMsgSeqno: ")
	sb.WriteString(fmt.Sprint(b.BadMsgSeqno))
	sb.WriteString(",\n")
	sb.WriteString("\tErrorCode: ")
	sb.WriteString(fmt.Sprint(b.ErrorCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BadMsgNotification) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bad_msg_notification#a7eff811 as nil")
	}
	buf.PutID(BadMsgNotificationTypeID)
	buf.PutLong(b.BadMsgID)
	buf.PutInt(b.BadMsgSeqno)
	buf.PutInt(b.ErrorCode)
	return nil
}

// Decode implements bin.Decoder.
func (b *BadMsgNotification) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bad_msg_notification#a7eff811 to nil")
	}
	if err := buf.ConsumeID(BadMsgNotificationTypeID); err != nil {
		return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: %w", err)
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: field bad_msg_id: %w", err)
		}
		b.BadMsgID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: field bad_msg_seqno: %w", err)
		}
		b.BadMsgSeqno = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: field error_code: %w", err)
		}
		b.ErrorCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for BadMsgNotification.
var (
	_ bin.Encoder = &BadMsgNotification{}
	_ bin.Decoder = &BadMsgNotification{}
)
