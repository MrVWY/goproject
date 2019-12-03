package pkger

import (
	"context"

	"github.com/micro/go-micro/config/source"
)

type pkgerPathKey struct{}

// WithPath sets the path to pkger
func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, pkgerPathKey{}, p)
	}
}
