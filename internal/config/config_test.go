package config

import (
	"bytes"
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestType(t *testing.T) {
	config := &SvangType{}
	content, err := yaml.Marshal(config)
	if err != nil {
		_ = fmt.Errorf("error while marshalling data %v", err)
	}
	fmt.Println(bytes.NewBuffer(content).String())
}
