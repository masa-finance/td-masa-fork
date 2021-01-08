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

// StatsMessageStats represents TL type `stats.messageStats#8999f295`.
// Message statistics
//
// See https://core.telegram.org/constructor/stats.messageStats for reference.
type StatsMessageStats struct {
	// Message view graph
	ViewsGraph StatsGraphClass
}

// StatsMessageStatsTypeID is TL type id of StatsMessageStats.
const StatsMessageStatsTypeID = 0x8999f295

func (m *StatsMessageStats) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ViewsGraph == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *StatsMessageStats) String() string {
	if m == nil {
		return "StatsMessageStats(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsMessageStats")
	sb.WriteString("{\n")
	sb.WriteString("\tViewsGraph: ")
	sb.WriteString(fmt.Sprint(m.ViewsGraph))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *StatsMessageStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.messageStats#8999f295 as nil")
	}
	b.PutID(StatsMessageStatsTypeID)
	if m.ViewsGraph == nil {
		return fmt.Errorf("unable to encode stats.messageStats#8999f295: field views_graph is nil")
	}
	if err := m.ViewsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.messageStats#8999f295: field views_graph: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *StatsMessageStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.messageStats#8999f295 to nil")
	}
	if err := b.ConsumeID(StatsMessageStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.messageStats#8999f295: %w", err)
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.messageStats#8999f295: field views_graph: %w", err)
		}
		m.ViewsGraph = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsMessageStats.
var (
	_ bin.Encoder = &StatsMessageStats{}
	_ bin.Decoder = &StatsMessageStats{}
)
