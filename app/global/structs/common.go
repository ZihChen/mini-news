package structs

type EnvConfig struct {
	DBConfig  DBConfig  `yaml:"db_config"`
	JwtConfig JwtConfig `yaml:"jwt_config"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type JwtConfig struct {
	Key string `yaml:"key"`
}
