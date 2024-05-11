package xlog

import (
	"sync/atomic"
	"time"

	"github.com/hkloudou/xlib/xsync"
	"github.com/hkloudou/xlib/xtime"
)

type limitedExecutor struct {
	threshold time.Duration
	lastTime  *xsync.AtomicDuration
	discarded uint32
}

func newLimitedExecutor(milliseconds int) *limitedExecutor {
	return &limitedExecutor{
		threshold: time.Duration(milliseconds) * time.Millisecond,
		lastTime:  xsync.NewAtomicDuration(),
	}
}

func (le *limitedExecutor) logOrDiscard(execute func()) {
	if le == nil || le.threshold <= 0 {
		execute()
		return
	}

	now := xtime.Now()
	if now-le.lastTime.Load() <= le.threshold {
		atomic.AddUint32(&le.discarded, 1)
	} else {
		le.lastTime.Set(now)
		discarded := atomic.SwapUint32(&le.discarded, 0)
		if discarded > 0 {
			Errorf("Discarded %d error messages", discarded)
		}

		execute()
	}
}
