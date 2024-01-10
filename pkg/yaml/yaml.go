package yaml

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

type YamlLoader struct {
	PathToFile string
	ConfigFile any
}

func NewYamlLoader(pathToFile string, configFile any) *YamlLoader {
	return &YamlLoader{
		PathToFile: pathToFile,
		ConfigFile: configFile,
	}
}

func (y *YamlLoader) Load() error {
	fs := afero.NewOsFs()
	fileContent, err := afero.ReadFile(fs, y.PathToFile)
	if err != nil {
		log.Println(err)
		return err
	}

	m := make(map[string]any)
	err = yaml.Unmarshal(fileContent, m)
	if err != nil {
		return err
	}
	err = mapstructure.Decode(m, y.ConfigFile)

	return err
}
