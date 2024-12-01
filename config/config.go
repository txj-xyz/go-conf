package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

func Load() (*ConfigBase, error) {
	_, err := os.Stat("local.config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to load local.config.yaml %w", err)
	}
	viper.SetConfigName("local.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	cfg, err := unmarshal()
	if err != nil {
		return nil, fmt.Errorf("error with marshal thingy: %w", err)
	}

	// if viper.GetBool("debug.enabled") {
	// 	fmt.Println("debug mode: enabled")
	// 	cfg.PrettyPrint("")
	// }

	return cfg, nil

}

func (v *ConfigBase) PrettyPrint(indent string) {
	val := reflect.ValueOf(v)
	typ := val.Type()

	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct!")
		return
	}

	fmt.Printf("%s%s {\n", indent, typ.Name())
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("%s  %s: %v\n", indent, field.Name, value.Interface())
	}
	fmt.Printf("%s}\n", indent)
}

func unmarshal() (*ConfigBase, error) {
	var c ConfigBase
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal file contents into struct")
	}
	return &c, nil
}

// TODO: Something like this to to check config types and maybe a regex
// func ModeCheck() ConfigSettings {
//   _, err := os.Stat("prod.config.yaml")
//   if err != nil {
//     return ConfigSettings{ Prod: false }
//   }
//
//   return ConfigSettings{}
// }
//
// type ConfigSettings struct {
//   Prod bool `json:"prod"`
// }
