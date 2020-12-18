// Package config provides ...
package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

func Decode(file string) (Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if _, err := toml.Decode(string(data), &config); err != nil {
		log.Fatalln(err)
	}
	return config, nil
}

type Config struct {
	Server Server
	DB     Database
}

type Server struct {
	IP   string
	Port string
}

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}
