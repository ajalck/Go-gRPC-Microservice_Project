package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSrvUrl    string `mapstructure:"AUTH_SRV_URL"`
	ProductSrvUrl string `mapstructure:"PRODUCT_SRV_URL"`
	OrderSrvUrl   string `mapstructure:"ORDER_SRV_URL"`
}

/*
 mapstructure is a tag used in go structs to mension the struct field value should be of the input map key .
	eg:	struct filed value : "port"
		map key : "PORT"
		 >> The value to "port" should be the value of "PORT"
*/

func LoadConfig() (c Config, err error) {

	// AddConfigPath adds the directory where the configuration file is located
	viper.AddConfigPath("./pkg/config/env")
	// SetConfigName sets the name of the configuration file to be read
	viper.SetConfigName("dev")
	// SetConfigType sets the type of the configuration file
	viper.SetConfigType("env")
	// AutomaticEnv enables automatic binding of environment variables to configuration values.
	viper.AutomaticEnv()
	// ReadInConfig reads the configuration file with the specified name and type.
	err = viper.ReadInConfig()

	// Check if there was an error reading the configuration file.
	if err != nil {
		return
	}

	// Unmarshal reads the configuration settings into the Config struct.
	err = viper.Unmarshal(&c)
	return
}
