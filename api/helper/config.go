package helper

/*
	General configuration with config.yaml
*/
import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

const configPath = "config.yaml"

type Config struct {
	DbUrl string `yaml:"DbURL"`
	//add parameter here when new parameter add in config.yaml
}

var AppConfig *Config

/**
Transform config.yaml to Config structure
*/
func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}
