// Package config provides ...
package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

//Decode ....
func Decode(file string, v interface{}) (toml.MetaData, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return toml.MetaData{}, err
	}

	meta, err := toml.Decode(string(data), v)
	return meta, err
}

//Server ......
type Server struct {
	IP   string
	Port string
}
