package watcher_test

import (
	"context"
	"gonfigreloader/pkg/watcher"
	"io/fs"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileWatcher(t *testing.T) {

	fileName := "test.txt"
	//Create file
	_, err := os.Create(fileName)
	defer os.Remove(fileName)
	assert.NoError(t, err)

	fw := watcher.NewFileWatcher(fileName)
	ctx := context.Background()
	eventChannel := fw.Start(ctx)

	err = os.WriteFile(fileName, []byte("Cenas para file"), fs.ModeAppend)

	time.Sleep(5 * time.Second)

	<-eventChannel

}

func TestFileWatcherWithCustomWatchInterval(t *testing.T) {
	fileName := "test.txt"
	//Create file
	_, err := os.Create(fileName)
	defer os.Remove(fileName)
	assert.NoError(t, err)

	fw :=
		watcher.
			NewFileWatcher(fileName).
			WithTimeInterval(time.Second)

	ctx := context.Background()
	eventChannel := fw.Start(ctx)

	err = os.WriteFile(fileName, []byte("Cenas para file"), fs.ModeAppend)

	time.Sleep(1 * time.Second)

	<-eventChannel

}
