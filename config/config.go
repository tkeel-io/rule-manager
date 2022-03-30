package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var DSN = "root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
var RuleTopic = "TOPIC_FOR_RULE"

type Config struct {
	Server struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Env       string `yaml:"env"`
		Signature bool   `yaml:"signature"`
		NodeName  string `yaml:"node_name"`
	}
	Api struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		NodeName string `yaml:"node_name"`
	}

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	Log struct {
		File struct {
			Path       string `yaml:"path"`
			MaxHour    int64  `yaml:"max_hour"`
			RotateHour int64  `yaml:"rotate_hour"`
		}
		LogPath string `yaml:"log_path"`
		Level   string `yaml:"level"`
	}
	Etcd struct {
		Address []string `yaml:"address,flow"`
	}
	Redis struct {
		Addr     string `yaml:"addr"`     // ip:port
		Password string `yaml:"password"` //
		DB       int    `yaml:"db"`       //
	}
	Websocket struct {
		Address string `yaml:"address,flow"`
	}
	Accesses []struct {
		Name      string `yaml:"name"`
		Endpoint  string `yaml:"endpoint"`
		Enabled   bool   `yaml:"enabled"`
		ProtoType string `yaml:"type"`
	}
}

var (
	conf       *Config
	once       sync.Once
	configPath string
)

func InitConfig(path string) {
	configPath = path
}

func GetConfig() *Config {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, &conf)
		if err != nil {
			panic(err)
		}
	})
	return conf
}
