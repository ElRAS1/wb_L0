package store

type Config struct {
	DatabaseUrl string `toml: "dburl"`
}

func NewConfig() *Config {
	return &Config{}
}
