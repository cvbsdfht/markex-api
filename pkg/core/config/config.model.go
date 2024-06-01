package config

type ConfigFile struct {
	Config string
}

type Configuration struct {
	App   AppConfiguration
	Log   LogConfiguration
	Mongo MongoConfiguration
	Redis RedisConfiguration
}

type AppConfiguration struct {
	Name     string
	Port     int64
	Env      string
	ProdMode bool
}

type LogConfiguration struct {
	FilePath string
	Level    int
	Format   string
}

type MongoConfiguration struct {
	Uri      string
	Database string
}

type RedisConfiguration struct {
	Host     string
	Port     string
	Password string
	Prefix   string
	Database int
}
