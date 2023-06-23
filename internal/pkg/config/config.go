package config

import (
	"fmt"
	"os"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// NewConfig new a config
func NewConfig(path string) (*viper.Viper, error) {
	var (
		v = viper.New()
	)

	if path != "" {
		v.SetConfigFile(path)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, errors.Wrap(err, "get user home dir failed")
		}

		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName(".project-goapi")
	}

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "read config file failed")
	}

	fmt.Println("v.ConfigFileUsed(): ", v.ConfigFileUsed())

	return v, nil
}

// ProviderSet is config providers
var ProviderSet = wire.NewSet(NewConfig)
