package rescue

import (
	"context"
	"runtime/debug"

	"github.com/hkloudou/xlib/xlog"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//
//	defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		xlog.ErrorStack(p)
	}
}

// RecoverCtx is used with defer to do cleanup on panics.
func RecoverCtx(ctx context.Context, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		xlog.WithContext(ctx).Errorf("%+v\n%s", p, debug.Stack())
	}
}
