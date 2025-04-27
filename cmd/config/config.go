package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppFlags struct {
	ConfigPath string
}

func ParseFlags() AppFlags {
	configPath := flag.String("config", "", "Path to config")
	flag.Parse()
	return AppFlags{
		ConfigPath: *configPath,
	}
}

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:":8080"`
}

func MustLoad(cfgPath string, cfg any) {
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file is not exist: %s", cfgPath)
	}
	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		log.Fatalf("can not read config: %s", err)
	}
}
