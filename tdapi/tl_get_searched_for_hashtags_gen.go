// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// GetSearchedForHashtagsRequest represents TL type `getSearchedForHashtags#ea843efd`.
type GetSearchedForHashtagsRequest struct {
	// Prefix of hashtags or cashtags to return
	Prefix string
	// The maximum number of items to be returned
	Limit int32
}

// GetSearchedForHashtagsRequestTypeID is TL type id of GetSearchedForHashtagsRequest.
const GetSearchedForHashtagsRequestTypeID = 0xea843efd

// Ensuring interfaces in compile-time for GetSearchedForHashtagsRequest.
var (
	_ bin.Encoder     = &GetSearchedForHashtagsRequest{}
	_ bin.Decoder     = &GetSearchedForHashtagsRequest{}
	_ bin.BareEncoder = &GetSearchedForHashtagsRequest{}
	_ bin.BareDecoder = &GetSearchedForHashtagsRequest{}
)

func (g *GetSearchedForHashtagsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Prefix == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSearchedForHashtagsRequest) String() string {
	if g == nil {
		return "GetSearchedForHashtagsRequest(nil)"
	}
	type Alias GetSearchedForHashtagsRequest
	return fmt.Sprintf("GetSearchedForHashtagsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSearchedForHashtagsRequest) TypeID() uint32 {
	return GetSearchedForHashtagsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSearchedForHashtagsRequest) TypeName() string {
	return "getSearchedForHashtags"
}

// TypeInfo returns info about TL type.
func (g *GetSearchedForHashtagsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSearchedForHashtags",
		ID:   GetSearchedForHashtagsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Prefix",
			SchemaName: "prefix",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSearchedForHashtagsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSearchedForHashtags#ea843efd as nil")
	}
	b.PutID(GetSearchedForHashtagsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSearchedForHashtagsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSearchedForHashtags#ea843efd as nil")
	}
	b.PutString(g.Prefix)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSearchedForHashtagsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSearchedForHashtags#ea843efd to nil")
	}
	if err := b.ConsumeID(GetSearchedForHashtagsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSearchedForHashtags#ea843efd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSearchedForHashtagsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSearchedForHashtags#ea843efd to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getSearchedForHashtags#ea843efd: field prefix: %w", err)
		}
		g.Prefix = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getSearchedForHashtags#ea843efd: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSearchedForHashtagsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSearchedForHashtags#ea843efd as nil")
	}
	b.ObjStart()
	b.PutID("getSearchedForHashtags")
	b.Comma()
	b.FieldStart("prefix")
	b.PutString(g.Prefix)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSearchedForHashtagsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSearchedForHashtags#ea843efd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSearchedForHashtags"); err != nil {
				return fmt.Errorf("unable to decode getSearchedForHashtags#ea843efd: %w", err)
			}
		case "prefix":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getSearchedForHashtags#ea843efd: field prefix: %w", err)
			}
			g.Prefix = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getSearchedForHashtags#ea843efd: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPrefix returns value of Prefix field.
func (g *GetSearchedForHashtagsRequest) GetPrefix() (value string) {
	if g == nil {
		return
	}
	return g.Prefix
}

// GetLimit returns value of Limit field.
func (g *GetSearchedForHashtagsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetSearchedForHashtags invokes method getSearchedForHashtags#ea843efd returning error if any.
func (c *Client) GetSearchedForHashtags(ctx context.Context, request *GetSearchedForHashtagsRequest) (*Hashtags, error) {
	var result Hashtags

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}