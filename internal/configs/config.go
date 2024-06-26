package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type resource struct {
	Name            string
	Endpoint        string
	Destination_URL string
}
type configuration struct {
	Server struct {
		Host        string
		Listen_port string
	}
	Resources []resource
}

var Config *configuration

// Returns a configuration object that represents the configuration of the application
func NewConfig() (*configuration, error) {
	viper.AddConfigPath("data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}
	return Config, nil
}
