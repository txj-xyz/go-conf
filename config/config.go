package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Load() {
  _, err := os.Stat("local.config.yaml")
  if err != nil {
    panic(fmt.Errorf("failed to locate local.config.yaml file", err))
  }
	viper.SetConfigName("local.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
  err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
  fmt.Println("loaded 'local.config.yaml' OK")

  if (viper.GetBool("debug.enabled")) {
    fmt.Println("debug mode: enabled")
    fmt.Println(viper.AllSettings())
  }
}

type Config struct {

}
