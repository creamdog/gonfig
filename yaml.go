package gonfig

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

// FromYml reads the contents from the supplied reader.
// The content is parsed as json into a map[string]interface{}.
// It returns a JsonGonfig struct pointer and any error encountered
func FromYml(reader io.Reader) (Gonfig, error) {

	inputBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	if err := yaml.Unmarshal(inputBytes, &obj); err != nil {
		return nil, err
	}
	return &JsonGonfig{obj}, nil
}
