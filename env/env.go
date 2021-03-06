package env

import (
	"github.com/spf13/viper"
)

//Load loads ENV Vars data from ${path}/${name}.env
func Load(path string, name string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")
	return viper.ReadInConfig()
}

//Get return the ENV var for the given key.
func Get(key string) string {
	return viper.GetString(key)
}
