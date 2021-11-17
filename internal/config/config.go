package config

import (
	"os"
	"runtime"
	"log"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port	string `yaml:"port"`
	ConnStr	string `yaml:"conn_str"`
}

func Init() (*Config) {
	c := Config{}

	_, b, _, _ := runtime.Caller(0)
    basepath := filepath.Dir(b)
	file := filepath.Join(basepath, "../../config/dev.yml")

	// load from YAML config file, ToDo: read based on env variable
	bytes, err := ioutil.ReadFile(file)
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