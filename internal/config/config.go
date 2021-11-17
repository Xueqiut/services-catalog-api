package config

import (
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port	string `yaml:"port"`
	ConnStr	string `yaml:"conn_str"`
}

func Init() (*Config) {
	c := Config{}

	// load from YAML config file, ToDo: read based on env variable
	bytes, err := ioutil.ReadFile("./config/dev.yml")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	if err = yaml.Unmarshal(bytes, &c); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return &c
}