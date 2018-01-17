package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var globalCfg *viper.Viper

func init() {
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}

	pflag.StringP("env", "e", "dev", "environment defalut to development")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	env := viper.GetString("env")

	globalCfg = viper.Sub(env)
	if globalCfg == nil {
		log.Fatalln("invalid env value:", env)

	}
	globalCfg.SetEnvPrefix("msg")
	globalCfg.BindEnv("az_name")
	globalCfg.BindEnv("az_key")

}

func main() {
	// os.Setenv("MSG_AZ_NAME", "dfd")
	// os.Setenv("MSG_AZ_KEY", "key")
	fmt.Println(
		globalCfg.GetString("env"),
		globalCfg.GetString("az_name"),
		globalCfg.GetString("az_key"),
		globalCfg.GetString("host.port"),
		globalCfg.GetString("d"),
	)
}
