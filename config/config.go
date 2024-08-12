package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	c    Config
	once sync.Once
)

type Config struct {
	App struct {
		Host               string `yaml:"host" json:"host" envconfig:"MDPAGES_HOST"`
		Port               int    `yaml:"port" json:"port" envconfig:"MDPAGES_PORT"`
		RootPassword       string `yaml:"root_password" json:"root_password" envconfig:"MDPAGES_ROOT_PASSWORD"`
		SessionKey         string `yaml:"session_key" json:"session_key" envconfig:"MDPAGES_SESSION_KEY"`
		Type               string `yaml:"type" json:"type" envconfig:"MDPAGES_TYPE"`
		EnableRegistration bool   `yaml:"enable_registration" json:"enable_registration" envconfig:"MDPAGES_ENABLE_REGISTRATION"`
	} `yaml:"app" json:"app"`
	Postgres struct {
		Username string `yaml:"username" json:"username" envconfig:"MDPAGES_POSTGRES_USERNAME"`
		Password string `yaml:"password" json:"password" envconfig:"MDPAGES_POSTGRES_PASSWORD"`
		Addr     string `yaml:"addr" json:"addr" envconfig:"MDPAGES_POSTGRES_ADDR"`
	} `yaml:"postgres" json:"postgres"`
	Redis struct {
		Addr     string `yaml:"addr" json:"addr" envconfig:"MDPAGES_REDIS_ADDR"`
		Db       int    `yaml:"db" json:"db" envconfig:"MDPAGES_REDIS_DB"`
		Password string `yaml:"password" json:"password" envconfig:"MDPAGES_REDIS_PASSWORD"`
	} `yaml:"redis" json:"redis"`
}

func Get() Config {
	once.Do(func() {
		if err := readYaml(&c, "config.yaml"); err != nil {
			panic(err)
		}
		if err := readEnv(&c); err != nil {
			panic(err)
		}

		s, _ := json.MarshalIndent(c, "# ", "    ")
		fmt.Println("# Config:")
		fmt.Println("# " + string(s))
	})
	return c
}

func readYaml(cfg *Config, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return yaml.NewDecoder(file).Decode(cfg)
}

func readEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}
