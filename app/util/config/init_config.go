package config

import (
	"fmt"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"github.com/spf13/viper"
	"strings"
)

func (c *Config) InitConfig() (*domain.Config, error) {
	if err := c.setupViper(); err != nil {
		return nil, err
	}

	var cfg domain.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) setupViper() error {
	if err := c.loadEnv(); err != nil {
		return err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(constants.Dot, constants.Underscore))
	viper.SetConfigType("yaml")

	envConfigName := map[string]string{
		constants.Development: constants.ConfigNameDev,
		constants.Staging:     constants.ConfigNameStaging,
		constants.Production:  constants.ConfigNameProd,
	}

	env := viper.GetString(constants.Environment)
	configName, ok := envConfigName[env]
	if !ok {
		return fmt.Errorf("invalid ENV value: %s", env)
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(constants.ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	return nil
}

func (c *Config) loadEnv() error {
	v := viper.New()
	v.SetConfigType(constants.FileTypeEnv)
	v.SetConfigName(constants.ExtensionEnv)
	v.AddConfigPath(constants.Dot)

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read .env file: %w", err)
	}

	for _, key := range v.AllKeys() {
		viper.Set(key, v.Get(key))
	}

	return nil
}
