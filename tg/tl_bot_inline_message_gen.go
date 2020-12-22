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

// BotInlineMessageMediaAuto represents TL type `botInlineMessageMediaAuto#764cf810`.
// Send whatever media is attached to the botInlineMediaResult
//
// See https://core.telegram.org/constructor/botInlineMessageMediaAuto for reference.
type BotInlineMessageMediaAuto struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Caption
	Message string
	// Message entities for styled text
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Inline keyboard
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaAutoTypeID is TL type id of BotInlineMessageMediaAuto.
const BotInlineMessageMediaAutoTypeID = 0x764cf810

// String implements fmt.Stringer.
func (b *BotInlineMessageMediaAuto) String() string {
	if b == nil {
		return "BotInlineMessageMediaAuto(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotInlineMessageMediaAuto")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(b.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(b.Message))
	sb.WriteString(",\n")
	if b.Flags.Has(1) {
		sb.WriteByte('[')
		for _, v := range b.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if b.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(b.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaAuto) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaAuto#764cf810 as nil")
	}
	buf.PutID(BotInlineMessageMediaAutoTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field flags: %w", err)
	}
	buf.PutString(b.Message)
	if b.Flags.Has(1) {
		buf.PutVectorHeader(len(b.Entities))
		for idx, v := range b.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(buf); err != nil {
				return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaAuto#764cf810: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetEntities sets value of Entities conditional field.
func (b *BotInlineMessageMediaAuto) SetEntities(value []MessageEntityClass) {
	b.Flags.Set(1)
	b.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaAuto) GetEntities() (value []MessageEntityClass, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Entities, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaAuto) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaAuto) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaAuto) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaAuto#764cf810 to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaAutoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field flags: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field message: %w", err)
		}
		b.Message = value
	}
	if b.Flags.Has(1) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field entities: %w", err)
			}
			b.Entities = append(b.Entities, value)
		}
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaAuto#764cf810: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaAuto) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaAuto.
var (
	_ bin.Encoder = &BotInlineMessageMediaAuto{}
	_ bin.Decoder = &BotInlineMessageMediaAuto{}

	_ BotInlineMessageClass = &BotInlineMessageMediaAuto{}
)

// BotInlineMessageText represents TL type `botInlineMessageText#8c7f65e2`.
// Send a simple text message
//
// See https://core.telegram.org/constructor/botInlineMessageText for reference.
type BotInlineMessageText struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Disable webpage preview
	NoWebpage bool
	// The message
	Message string
	// Message entities for styled text
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Inline keyboard
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageTextTypeID is TL type id of BotInlineMessageText.
const BotInlineMessageTextTypeID = 0x8c7f65e2

// String implements fmt.Stringer.
func (b *BotInlineMessageText) String() string {
	if b == nil {
		return "BotInlineMessageText(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotInlineMessageText")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(b.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(b.Message))
	sb.WriteString(",\n")
	if b.Flags.Has(1) {
		sb.WriteByte('[')
		for _, v := range b.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if b.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(b.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BotInlineMessageText) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageText#8c7f65e2 as nil")
	}
	buf.PutID(BotInlineMessageTextTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field flags: %w", err)
	}
	buf.PutString(b.Message)
	if b.Flags.Has(1) {
		buf.PutVectorHeader(len(b.Entities))
		for idx, v := range b.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(buf); err != nil {
				return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageText#8c7f65e2: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (b *BotInlineMessageText) SetNoWebpage(value bool) {
	if value {
		b.Flags.Set(0)
	} else {
		b.Flags.Unset(0)
	}
}

// SetEntities sets value of Entities conditional field.
func (b *BotInlineMessageText) SetEntities(value []MessageEntityClass) {
	b.Flags.Set(1)
	b.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageText) GetEntities() (value []MessageEntityClass, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Entities, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageText) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageText) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageText) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageText#8c7f65e2 to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageTextTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field flags: %w", err)
		}
	}
	b.NoWebpage = b.Flags.Has(0)
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field message: %w", err)
		}
		b.Message = value
	}
	if b.Flags.Has(1) {
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(buf)
			if err != nil {
				return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field entities: %w", err)
			}
			b.Entities = append(b.Entities, value)
		}
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageText#8c7f65e2: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageText) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageText.
var (
	_ bin.Encoder = &BotInlineMessageText{}
	_ bin.Decoder = &BotInlineMessageText{}

	_ BotInlineMessageClass = &BotInlineMessageText{}
)

// BotInlineMessageMediaGeo represents TL type `botInlineMessageMediaGeo#51846fd`.
// Send a geolocation
//
// See https://core.telegram.org/constructor/botInlineMessageMediaGeo for reference.
type BotInlineMessageMediaGeo struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Geolocation
	Geo GeoPointClass
	// For live locations, a direction in which the location moves, in degrees; 1-360.
	//
	// Use SetHeading and GetHeading helpers.
	Heading int
	// Validity period
	//
	// Use SetPeriod and GetPeriod helpers.
	Period int
	// For live locations, a maximum distance to another chat member for proximity alerts, in meters (0-100000).
	//
	// Use SetProximityNotificationRadius and GetProximityNotificationRadius helpers.
	ProximityNotificationRadius int
	// Inline keyboard
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaGeoTypeID is TL type id of BotInlineMessageMediaGeo.
const BotInlineMessageMediaGeoTypeID = 0x51846fd

// String implements fmt.Stringer.
func (b *BotInlineMessageMediaGeo) String() string {
	if b == nil {
		return "BotInlineMessageMediaGeo(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotInlineMessageMediaGeo")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(b.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tGeo: ")
	sb.WriteString(b.Geo.String())
	sb.WriteString(",\n")
	if b.Flags.Has(0) {
		sb.WriteString("\tHeading: ")
		sb.WriteString(fmt.Sprint(b.Heading))
		sb.WriteString(",\n")
	}
	if b.Flags.Has(1) {
		sb.WriteString("\tPeriod: ")
		sb.WriteString(fmt.Sprint(b.Period))
		sb.WriteString(",\n")
	}
	if b.Flags.Has(3) {
		sb.WriteString("\tProximityNotificationRadius: ")
		sb.WriteString(fmt.Sprint(b.ProximityNotificationRadius))
		sb.WriteString(",\n")
	}
	if b.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(b.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaGeo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaGeo#51846fd as nil")
	}
	buf.PutID(BotInlineMessageMediaGeoTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field flags: %w", err)
	}
	if b.Geo == nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field geo is nil")
	}
	if err := b.Geo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field geo: %w", err)
	}
	if b.Flags.Has(0) {
		buf.PutInt(b.Heading)
	}
	if b.Flags.Has(1) {
		buf.PutInt(b.Period)
	}
	if b.Flags.Has(3) {
		buf.PutInt(b.ProximityNotificationRadius)
	}
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaGeo#51846fd: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetHeading sets value of Heading conditional field.
func (b *BotInlineMessageMediaGeo) SetHeading(value int) {
	b.Flags.Set(0)
	b.Heading = value
}

// GetHeading returns value of Heading conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetHeading() (value int, ok bool) {
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.Heading, true
}

// SetPeriod sets value of Period conditional field.
func (b *BotInlineMessageMediaGeo) SetPeriod(value int) {
	b.Flags.Set(1)
	b.Period = value
}

// GetPeriod returns value of Period conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetPeriod() (value int, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Period, true
}

// SetProximityNotificationRadius sets value of ProximityNotificationRadius conditional field.
func (b *BotInlineMessageMediaGeo) SetProximityNotificationRadius(value int) {
	b.Flags.Set(3)
	b.ProximityNotificationRadius = value
}

// GetProximityNotificationRadius returns value of ProximityNotificationRadius conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetProximityNotificationRadius() (value int, ok bool) {
	if !b.Flags.Has(3) {
		return value, false
	}
	return b.ProximityNotificationRadius, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaGeo) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaGeo) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaGeo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaGeo#51846fd to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaGeoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field flags: %w", err)
		}
	}
	{
		value, err := DecodeGeoPoint(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field geo: %w", err)
		}
		b.Geo = value
	}
	if b.Flags.Has(0) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field heading: %w", err)
		}
		b.Heading = value
	}
	if b.Flags.Has(1) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field period: %w", err)
		}
		b.Period = value
	}
	if b.Flags.Has(3) {
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field proximity_notification_radius: %w", err)
		}
		b.ProximityNotificationRadius = value
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaGeo#51846fd: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaGeo) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaGeo.
var (
	_ bin.Encoder = &BotInlineMessageMediaGeo{}
	_ bin.Decoder = &BotInlineMessageMediaGeo{}

	_ BotInlineMessageClass = &BotInlineMessageMediaGeo{}
)

// BotInlineMessageMediaVenue represents TL type `botInlineMessageMediaVenue#8a86659c`.
// Send a venue
//
// See https://core.telegram.org/constructor/botInlineMessageMediaVenue for reference.
type BotInlineMessageMediaVenue struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Geolocation of venue
	Geo GeoPointClass
	// Venue name
	Title string
	// Address
	Address string
	// Venue provider: currently only "foursquare" needs to be supported
	Provider string
	// Venue ID in the provider's database
	VenueID string
	// Venue type in the provider's database
	VenueType string
	// Inline keyboard
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaVenueTypeID is TL type id of BotInlineMessageMediaVenue.
const BotInlineMessageMediaVenueTypeID = 0x8a86659c

// String implements fmt.Stringer.
func (b *BotInlineMessageMediaVenue) String() string {
	if b == nil {
		return "BotInlineMessageMediaVenue(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotInlineMessageMediaVenue")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(b.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tGeo: ")
	sb.WriteString(b.Geo.String())
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(b.Title))
	sb.WriteString(",\n")
	sb.WriteString("\tAddress: ")
	sb.WriteString(fmt.Sprint(b.Address))
	sb.WriteString(",\n")
	sb.WriteString("\tProvider: ")
	sb.WriteString(fmt.Sprint(b.Provider))
	sb.WriteString(",\n")
	sb.WriteString("\tVenueID: ")
	sb.WriteString(fmt.Sprint(b.VenueID))
	sb.WriteString(",\n")
	sb.WriteString("\tVenueType: ")
	sb.WriteString(fmt.Sprint(b.VenueType))
	sb.WriteString(",\n")
	if b.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(b.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaVenue) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaVenue#8a86659c as nil")
	}
	buf.PutID(BotInlineMessageMediaVenueTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field flags: %w", err)
	}
	if b.Geo == nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field geo is nil")
	}
	if err := b.Geo.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field geo: %w", err)
	}
	buf.PutString(b.Title)
	buf.PutString(b.Address)
	buf.PutString(b.Provider)
	buf.PutString(b.VenueID)
	buf.PutString(b.VenueType)
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaVenue#8a86659c: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaVenue) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaVenue) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaVenue) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaVenue#8a86659c to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field flags: %w", err)
		}
	}
	{
		value, err := DecodeGeoPoint(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field geo: %w", err)
		}
		b.Geo = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field title: %w", err)
		}
		b.Title = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field address: %w", err)
		}
		b.Address = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field provider: %w", err)
		}
		b.Provider = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field venue_id: %w", err)
		}
		b.VenueID = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field venue_type: %w", err)
		}
		b.VenueType = value
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaVenue#8a86659c: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaVenue) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaVenue.
var (
	_ bin.Encoder = &BotInlineMessageMediaVenue{}
	_ bin.Decoder = &BotInlineMessageMediaVenue{}

	_ BotInlineMessageClass = &BotInlineMessageMediaVenue{}
)

// BotInlineMessageMediaContact represents TL type `botInlineMessageMediaContact#18d1cdc2`.
// Send a contact
//
// See https://core.telegram.org/constructor/botInlineMessageMediaContact for reference.
type BotInlineMessageMediaContact struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Phone number
	PhoneNumber string
	// First name
	FirstName string
	// Last name
	LastName string
	// VCard info
	Vcard string
	// Inline keyboard
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
}

// BotInlineMessageMediaContactTypeID is TL type id of BotInlineMessageMediaContact.
const BotInlineMessageMediaContactTypeID = 0x18d1cdc2

// String implements fmt.Stringer.
func (b *BotInlineMessageMediaContact) String() string {
	if b == nil {
		return "BotInlineMessageMediaContact(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotInlineMessageMediaContact")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(b.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPhoneNumber: ")
	sb.WriteString(fmt.Sprint(b.PhoneNumber))
	sb.WriteString(",\n")
	sb.WriteString("\tFirstName: ")
	sb.WriteString(fmt.Sprint(b.FirstName))
	sb.WriteString(",\n")
	sb.WriteString("\tLastName: ")
	sb.WriteString(fmt.Sprint(b.LastName))
	sb.WriteString(",\n")
	sb.WriteString("\tVcard: ")
	sb.WriteString(fmt.Sprint(b.Vcard))
	sb.WriteString(",\n")
	if b.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(b.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BotInlineMessageMediaContact) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMessageMediaContact#18d1cdc2 as nil")
	}
	buf.PutID(BotInlineMessageMediaContactTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMessageMediaContact#18d1cdc2: field flags: %w", err)
	}
	buf.PutString(b.PhoneNumber)
	buf.PutString(b.FirstName)
	buf.PutString(b.LastName)
	buf.PutString(b.Vcard)
	if b.Flags.Has(2) {
		if b.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaContact#18d1cdc2: field reply_markup is nil")
		}
		if err := b.ReplyMarkup.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMessageMediaContact#18d1cdc2: field reply_markup: %w", err)
		}
	}
	return nil
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (b *BotInlineMessageMediaContact) SetReplyMarkup(value ReplyMarkupClass) {
	b.Flags.Set(2)
	b.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (b *BotInlineMessageMediaContact) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.ReplyMarkup, true
}

// Decode implements bin.Decoder.
func (b *BotInlineMessageMediaContact) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMessageMediaContact#18d1cdc2 to nil")
	}
	if err := buf.ConsumeID(BotInlineMessageMediaContactTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field flags: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field phone_number: %w", err)
		}
		b.PhoneNumber = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field first_name: %w", err)
		}
		b.FirstName = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field last_name: %w", err)
		}
		b.LastName = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field vcard: %w", err)
		}
		b.Vcard = value
	}
	if b.Flags.Has(2) {
		value, err := DecodeReplyMarkup(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMessageMediaContact#18d1cdc2: field reply_markup: %w", err)
		}
		b.ReplyMarkup = value
	}
	return nil
}

// construct implements constructor of BotInlineMessageClass.
func (b BotInlineMessageMediaContact) construct() BotInlineMessageClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMessageMediaContact.
var (
	_ bin.Encoder = &BotInlineMessageMediaContact{}
	_ bin.Decoder = &BotInlineMessageMediaContact{}

	_ BotInlineMessageClass = &BotInlineMessageMediaContact{}
)

// BotInlineMessageClass represents BotInlineMessage generic type.
//
// See https://core.telegram.org/type/BotInlineMessage for reference.
//
// Example:
//  g, err := DecodeBotInlineMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *BotInlineMessageMediaAuto: // botInlineMessageMediaAuto#764cf810
//  case *BotInlineMessageText: // botInlineMessageText#8c7f65e2
//  case *BotInlineMessageMediaGeo: // botInlineMessageMediaGeo#51846fd
//  case *BotInlineMessageMediaVenue: // botInlineMessageMediaVenue#8a86659c
//  case *BotInlineMessageMediaContact: // botInlineMessageMediaContact#18d1cdc2
//  default: panic(v)
//  }
type BotInlineMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() BotInlineMessageClass
	fmt.Stringer
}

// DecodeBotInlineMessage implements binary de-serialization for BotInlineMessageClass.
func DecodeBotInlineMessage(buf *bin.Buffer) (BotInlineMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BotInlineMessageMediaAutoTypeID:
		// Decoding botInlineMessageMediaAuto#764cf810.
		v := BotInlineMessageMediaAuto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageTextTypeID:
		// Decoding botInlineMessageText#8c7f65e2.
		v := BotInlineMessageText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageMediaGeoTypeID:
		// Decoding botInlineMessageMediaGeo#51846fd.
		v := BotInlineMessageMediaGeo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageMediaVenueTypeID:
		// Decoding botInlineMessageMediaVenue#8a86659c.
		v := BotInlineMessageMediaVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	case BotInlineMessageMediaContactTypeID:
		// Decoding botInlineMessageMediaContact#18d1cdc2.
		v := BotInlineMessageMediaContact{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotInlineMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// BotInlineMessage boxes the BotInlineMessageClass providing a helper.
type BotInlineMessageBox struct {
	BotInlineMessage BotInlineMessageClass
}

// Decode implements bin.Decoder for BotInlineMessageBox.
func (b *BotInlineMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotInlineMessageBox to nil")
	}
	v, err := DecodeBotInlineMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotInlineMessage = v
	return nil
}

// Encode implements bin.Encode for BotInlineMessageBox.
func (b *BotInlineMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BotInlineMessage == nil {
		return fmt.Errorf("unable to encode BotInlineMessageClass as nil")
	}
	return b.BotInlineMessage.Encode(buf)
}
