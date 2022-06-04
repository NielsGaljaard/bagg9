package configreader

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Secrets Secrets `mapstructure:"Secrets"`
}

type Secrets struct {
	BitbucketPassword string `mapstructure:"bbAppPassword"`
	GithubPassword    string `mapstructure:"ghAppPassword"`
	GitlabPassword    string `mapstructure:"glAppPassword"`
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
