// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Invoker can invoke raw MTProto rpc calls.
type Invoker interface {
	InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client implement methods for calling functions from TL schema via Invoker.
type Client struct {
	rpc Invoker
}

func NewClient(invoker Invoker) *Client {
	return &Client{
		rpc: invoker,
	}
}
