package reloader

import (
	"context"

	"gonfigreloader/pkg/loader"
	"gonfigreloader/pkg/watcher"
)

func NewReloader(ctx context.Context, l loader.Loader, fw *watcher.FileWatcher) {
	go func(ctx context.Context) {
		l.Load()
		reloadChan := fw.Start(ctx)
		for {
			select {
			case _ = <-reloadChan:
				l.Load()
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

}
