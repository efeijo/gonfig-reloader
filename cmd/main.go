package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"gonfigreloader/cmd/config"
	"gonfigreloader/internal/reloader"
	"gonfigreloader/pkg/loader"
	"gonfigreloader/pkg/watcher"
)

func main() {
	f, err := os.Getwd()
	if err != nil {
		return
	}
	configFilePath := f + "/cmd/config/config.yaml"

	c := &config.Config{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	loader := loader.NewLoader(loader.Yaml, configFilePath, c)
	fileWatcher := watcher.NewFileWatcher(configFilePath)
	reloader.NewReloader(ctx, loader, fileWatcher)

	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		fmt.Println(c)
	}

}
