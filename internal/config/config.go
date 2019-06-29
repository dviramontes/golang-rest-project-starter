package config

import (
	"log"

	"github.com/spf13/viper"
)

// Read attempts to read a config file from path.
// The defaults provided will be used if no config file doesn't exist
func Read(path string, defaults map[string]interface{}) *viper.Viper {
	v := viper.New()

	v.SetConfigFile(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Println(err.Error(), "proceeding with defaults")
		for key, val := range defaults {
			v.SetDefault(key, val)
		}
	}

	return v
}
