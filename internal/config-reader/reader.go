package configreader

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Secrets []Secret `mapstructure:"Secrets"`
}
type Secret struct {
	AppPassword string `mapstructure:"appPassword"`
	id          string `mapstructure:"id"`
}

type Cluster struct {
	Name      string      `mapstructure:"name"`
	NameSpace []NameSpace `mapstructure:"namespaces"`
}

type NameSpace struct {
	Name     string `mapstructure:"name"`
	Kind     string `mapstructure:"kind"`
	URL      string `mapstructure:"URL"`
	SecretID string `mapstructure:"secretID"`
}

func ReadConfig() (*Config, error) {
	viper.AddConfigPath("/etc/bagg9/")
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("can't find home dir")
	}
	viper.AddConfigPath(home + "/.config/bagg9/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("can't read config file")
		os.Exit(1)
	}
	viper.SetConfigName("secrets")
	err = viper.MergeInConfig()
	if err != nil {
		fmt.Println("can't read secrets file")
		os.Exit(1)
	}
	var conf = &Config{}
	viper.Unmarshal(conf)
	return conf, nil
}
