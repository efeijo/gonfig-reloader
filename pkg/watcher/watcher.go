package watcher

import (
	"context"
	"os"
	"time"
)

const (
	defaultInterval = 5 * time.Second
)

type FileWatcher struct {
	filePath         string
	lastModifiedDate time.Time
	watchInterval    time.Duration
}

func NewFileWatcher(filePath string) *FileWatcher {
	fileWatcher := &FileWatcher{
		filePath:      filePath,
		watchInterval: defaultInterval,
	}

	return fileWatcher
}

func (f *FileWatcher) WithTimeInterval(interval time.Duration) *FileWatcher {
	f.watchInterval = interval
	return f
}

func (f *FileWatcher) Start(ctx context.Context) chan struct{} {
	ch := make(chan struct{})
	go func(ctx context.Context, eventChannel chan struct{}) {
		ticker := time.NewTicker(f.watchInterval)

		for {
			select {
			case <-ticker.C:
				changed, err := f.modTimeChanged()
				if err != nil {
					//log error
					ticker.Stop()
					return
				}
				if changed {
					eventChannel <- struct{}{}
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}

	}(ctx, ch)

	return ch
}

func (f *FileWatcher) modTimeChanged() (bool, error) {
	fileInfo, err := os.Stat(f.filePath)
	if err != nil {
		// log error
		return false, err
	}

	modTime := fileInfo.ModTime()

	if f.lastModifiedDate.Equal(modTime) {
		return false, nil
	}

	f.lastModifiedDate = modTime

	return true, nil
}
