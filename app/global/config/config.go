package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"mini-news/app/global/structs"
)

var Config *structs.DBConfig

func Load() (err error) {
	envPathList := []string{
		"env/" + "local" + "/db.yaml",
	}

	for _, path := range envPathList {
		configFile, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err.Error())
		}

		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			panic(err.Error())
		}
	}

	return nil
}
