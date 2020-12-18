// Package config provides ...
package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

//Decode ....
func Decode(file string) (Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if _, err := toml.Decode(string(data), &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

//Config ....
type Config struct {
	Server Server
	DB     Database
}

//Server ......
type Server struct {
	IP   string
	Port string
}

//Database .....
type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}
