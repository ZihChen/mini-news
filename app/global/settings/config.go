package settings

import (
	"embed"
	"gopkg.in/yaml.v2"
	"mini-news/app/global/structs"
)

var Config *structs.EnvConfig

func Load(f embed.FS) (err error) {
	envPathList := []string{
		"env/" + "local" + "/db.yaml",
		"env/" + "local" + "/other.yaml",
	}

	for k := range envPathList {
		configFile, err := f.ReadFile(envPathList[k])
		if err != nil {
			panic(err.Error())
		}

		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			panic(err.Error())
		}
	}

	return nil
}
