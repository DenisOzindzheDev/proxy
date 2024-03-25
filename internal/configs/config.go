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
	Mode   string
	Server struct {
		Host        string
		Listen_port string
	}
	Resources []resource
}

var Config *configuration

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
