package yaml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	Emanuel []string `mapstructure:"emanuel"`
	Coding  string   `mapstructure:"coding"`
}

func TestNewYamlLoader(t *testing.T) {
	configFile := &Config{}
	loader := NewYamlLoader("./yaml_test.yaml", configFile)

	err := loader.Load()
	assert.NoError(t, err)

	assert.Equal(t, configFile.Emanuel, []string{"great", "coder"})
	assert.Equal(t, configFile.Coding, "a lot")

}
