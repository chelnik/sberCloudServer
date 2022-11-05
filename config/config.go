package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Address struct {
	Port string `yaml:"port"`
}

func NewPointerAddress() *Address {
	return &Address{}
}

func (a *Address) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open error %s", err)
	}
	sliceBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("readAll error %s", err)
	}
	err = yaml.Unmarshal(sliceBytes, a)
	if err != nil {
		return fmt.Errorf("unmarshal error %s", err)
	}
	return nil
}
