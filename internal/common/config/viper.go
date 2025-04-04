package config

import "github.com/spf13/viper"

/** For reading or setting config for the struct in the Project **/

// NewViperConfig New viper config
func NewViperConfig() error {
	//	set file name
	viper.SetConfigName("global")
	//	set file type
	viper.SetConfigType("yml")
	//	add file path
	viper.AddConfigPath("../common/config")
	//	read env arguments
	viper.AutomaticEnv()
	//	return read content in the config file
	return viper.ReadInConfig()
}
