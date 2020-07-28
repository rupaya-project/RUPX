package rupexlending

import (
	"context"
	"errors"
	"sync"
	"time"
)


// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicRupeXLendingAPI provides the rupayaX RPC service that can be
// use publicly without security implications.
type PublicRupeXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicRupeXLendingAPI create a new RPC rupayaX service.
func NewPublicRupeXLendingAPI(t *Lending) *PublicRupeXLendingAPI {
	api := &PublicRupeXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicRupeXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
