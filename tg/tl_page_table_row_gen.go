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

// PageTableRow represents TL type `pageTableRow#e0c0c5e5`.
// Table row
//
// See https://core.telegram.org/constructor/pageTableRow for reference.
type PageTableRow struct {
	// Table cells
	Cells []PageTableCell
}

// PageTableRowTypeID is TL type id of PageTableRow.
const PageTableRowTypeID = 0xe0c0c5e5

// String implements fmt.Stringer.
func (p *PageTableRow) String() string {
	if p == nil {
		return "PageTableRow(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PageTableRow")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range p.Cells {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PageTableRow) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageTableRow#e0c0c5e5 as nil")
	}
	b.PutID(PageTableRowTypeID)
	b.PutVectorHeader(len(p.Cells))
	for idx, v := range p.Cells {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageTableRow#e0c0c5e5: field cells element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageTableRow) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageTableRow#e0c0c5e5 to nil")
	}
	if err := b.ConsumeID(PageTableRowTypeID); err != nil {
		return fmt.Errorf("unable to decode pageTableRow#e0c0c5e5: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pageTableRow#e0c0c5e5: field cells: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PageTableCell
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode pageTableRow#e0c0c5e5: field cells: %w", err)
			}
			p.Cells = append(p.Cells, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PageTableRow.
var (
	_ bin.Encoder = &PageTableRow{}
	_ bin.Decoder = &PageTableRow{}
)
