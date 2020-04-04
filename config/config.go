package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ApiConfiguration struct {
	AUTH0_API_IDENTIFIER string
	AUTH0_DOMAIN         string
}

func Get() ApiConfiguration {
	api := setAuth0Variables()
	return ApiConfiguration{
		AUTH0_API_IDENTIFIER: api.AUTH0_API_IDENTIFIER,
		AUTH0_DOMAIN:         api.AUTH0_DOMAIN,
	}
}
func setAuth0Variables() ApiConfiguration {
	var apiConfiguration ApiConfiguration

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	// Set undefined variables
	viper.SetDefault("AUTH0_API_IDENTIFIER", "AUTH0_DOMAIN")

	err := viper.Unmarshal(&apiConfiguration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	fmt.Println("AUTH0_API_IDENTIFIER is\t", apiConfiguration.AUTH0_API_IDENTIFIER)
	fmt.Println("AUTH0_DOMAIN is\t", apiConfiguration.AUTH0_DOMAIN)
	return apiConfiguration
}
