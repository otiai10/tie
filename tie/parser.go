package tie

import (
	"os"

	"github.com/go-yaml/yaml"
)

type (
	ConfigV1 struct {
		Version  string `json:"version" yaml:"version"`
		Commands []struct {
			Name string            `json:"name" yaml:"name"`
			Run  string            `json:"run" yaml:"run"`
			Env  map[string]string `json:"env" yaml:"env"`
			// DependsOn []string          `json:"depends_on" yaml:"depends_on"`
		} `json:"commands" yaml:"commands"`
	}
)

func Parse(fpath string) (c ConfigV1, err error) {
	f, err := os.Open(fpath)
	if err != nil {
		return
	}
	defer f.Close()

	c = ConfigV1{}
	err = yaml.NewDecoder(f).Decode(&c)
	return
}
