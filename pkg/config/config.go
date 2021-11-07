package config

import "github.com/spf13/viper"

// Config stores configuration info
type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

// LoadConfig reads configuration from a file or env variables
func Load(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("api")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
