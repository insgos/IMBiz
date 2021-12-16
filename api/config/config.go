package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Mysql struct {
		Host string
		Port int
		Db   string
		Usr  string
		Pwd  string
	}
	Mongo struct {
		Host  string
		Port  int
		Db    string
		LogDb string
	}

	Redis struct {
		Host    string
		Port    int
		Db      int
		StatsDb int
		BadgeDb int
	}

	Logger struct {
		Path string
	}

	System struct {
		Port int
	}
}

func LoadConfiguration(file string) (config Config, err error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return
}
