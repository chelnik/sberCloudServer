package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Addresss struct {
	Port string `yaml:"port"`
}

func NewPointerAddress() *Addresss {
	return &Addresss{}
}

func (a *Addresss) LoadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("error with LoadConfig", err)
		return
	}
	slyceBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("error with LoadConfig", err)
		return
	}

	err = yaml.Unmarshal(slyceBytes, a)
	if err != nil {
		log.Println("error with LoadConfig", err)
		return
	}
}
