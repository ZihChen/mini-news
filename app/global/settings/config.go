package settings

import (
	"embed"
	"gopkg.in/yaml.v2"
	"log"
	"mini-news/app/global/errorcode"
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
			log.Fatalf(errorcode.ReadFileError, err.Error())
		}

		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			log.Fatalf(errorcode.YamlUnmarshalError, err.Error())
		}
	}

	return nil
}
