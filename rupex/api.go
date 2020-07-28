package rupex

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicRupeXAPI provides the rupayaX RPC service that can be
// use publicly without security implications.
type PublicRupeXAPI struct {
	t        *RupeX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicRupeXAPI create a new RPC rupayaX service.
func NewPublicRupeXAPI(t *RupeX) *PublicRupeXAPI {
	api := &PublicRupeXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the RupeX sub-protocol version.
func (api *PublicRupeXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
