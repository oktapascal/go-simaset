package welcome

import (
	"github.com/google/wire"
	"sync"
)

var (
	hdl     *Handler
	hdlOnce sync.Once

	ProviderSet = wire.NewSet(ProvideHandler)
)

// ProvideHandler returns a singleton instance of the Handler struct.
//
// It uses the hdlOnce variable to ensure that the Handler is only created once.
// The function returns a pointer to the Handler struct.
func ProvideHandler() *Handler {
	hdlOnce.Do(func() {
		hdl = new(Handler)
	})

	return hdl
}
