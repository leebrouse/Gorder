package config

import (
	"github.com/spf13/viper"
	"strings"
)

/** For reading or setting config for the struct in the Project **/

// NewViperConfig New viper config
func NewViperConfig() error {
	//	set file name
	viper.SetConfigName("global")
	//	set file type
	viper.SetConfigType("yml")
	//	add file path
	viper.AddConfigPath("../common/config")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	//_ = viper.BindEnv("stripe-key", "STRIPE_KEY")
	//	read env arguments
	viper.AutomaticEnv()
	_ = viper.BindEnv("stripe-key", "STRIPE_KEY", "endpoint-stripe-secret", "ENDPOINT_STRIPE_SECRET")
	//	return read content in the config file
	return viper.ReadInConfig()
}
