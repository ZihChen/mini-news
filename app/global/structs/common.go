package structs

type EnvConfig struct {
	DBConfig    DBConfig    `yaml:"db_config"`
	JwtConfig   JwtConfig   `yaml:"jwt_config"`
	RedisConfig RedisConfig `yaml:"redis_config"`
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

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
