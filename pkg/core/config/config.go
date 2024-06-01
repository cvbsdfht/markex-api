package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func NewConfig(defaultConfigFile string) *ConfigFile {
	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = defaultConfigFile
	}

	return &ConfigFile{
		Config: configFile,
	}
}

func (c *ConfigFile) LoadConfig() (*Configuration, error) {
	configuration := &Configuration{}

	fmt.Println("Load configuration file : " + c.Config)
	viper.SetConfigFile(c.Config)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	for _, key := range viper.AllKeys() {
		envKey := strings.ReplaceAll(key, ".", "_")
		err := viper.BindEnv(key, envKey)
		if err != nil {
			return nil, err
		}
	}

	err := viper.Unmarshal(configuration)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return configuration, nil
}
