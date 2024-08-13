package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
}

func LoadConfig(path, configName, configType string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal the config into a struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
