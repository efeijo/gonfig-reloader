package main

import (
	"fmt"
	"gonfigreloader/cmd/config"
	"gonfigreloader/pkg/yaml"
	"log/slog"
	"os"
)

func main() {
	f, err := os.Getwd()
	loader := yaml.NewYamlLoader(f+"/cmd/config/config.yaml", &config.Config{})
	err = loader.Load()

	if err != nil {
		slog.Error("error:", err)
		os.Exit(1)
	}
	fmt.Println(loader.ConfigFile)

}
