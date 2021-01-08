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

// UploadGetCdnFileRequest represents TL type `upload.getCdnFile#2000bcc3`.
// Download a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.getCdnFile for reference.
type UploadGetCdnFileRequest struct {
	// File token
	FileToken []byte
	// Offset of chunk to download
	Offset int
	// Length of chunk to download
	Limit int
}

// UploadGetCdnFileRequestTypeID is TL type id of UploadGetCdnFileRequest.
const UploadGetCdnFileRequestTypeID = 0x2000bcc3

func (g *UploadGetCdnFileRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FileToken == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UploadGetCdnFileRequest) String() string {
	if g == nil {
		return "UploadGetCdnFileRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UploadGetCdnFileRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFileToken: ")
	sb.WriteString(fmt.Sprint(g.FileToken))
	sb.WriteString(",\n")
	sb.WriteString("\tOffset: ")
	sb.WriteString(fmt.Sprint(g.Offset))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *UploadGetCdnFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getCdnFile#2000bcc3 as nil")
	}
	b.PutID(UploadGetCdnFileRequestTypeID)
	b.PutBytes(g.FileToken)
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetCdnFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getCdnFile#2000bcc3 to nil")
	}
	if err := b.ConsumeID(UploadGetCdnFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: field file_token: %w", err)
		}
		g.FileToken = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadGetCdnFileRequest.
var (
	_ bin.Encoder = &UploadGetCdnFileRequest{}
	_ bin.Decoder = &UploadGetCdnFileRequest{}
)

// UploadGetCdnFile invokes method upload.getCdnFile#2000bcc3 returning error if any.
// Download a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.getCdnFile for reference.
func (c *Client) UploadGetCdnFile(ctx context.Context, request *UploadGetCdnFileRequest) (UploadCdnFileClass, error) {
	var result UploadCdnFileBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.CdnFile, nil
}
