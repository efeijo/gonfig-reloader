package loader

import "gonfigreloader/pkg/yaml"

type LoaderType uint8

const (
	Yaml LoaderType = iota
)

type Loader interface {
	Load() error
}

func NewLoader(loaderType LoaderType, pathToFile string, configStruct any) Loader {

	switch loaderType {
	case Yaml:
		yaml.NewYamlLoader(pathToFile, configStruct)
		return yaml.NewYamlLoader(pathToFile, configStruct)
	}

	return nil
}
