package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfig = ".realworld.yml"
	envPrefix     = "realworld"
)

var (
	Version = "undefined"
	Date    = "undefined"
	Commit  = "undefined"
)

var (
	config string

	rootCmd = &cobra.Command{
		Use:   "realworld",
		Short: "the realworld example app",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&config, "config", "", fmt.Sprintf("config file (default \"%s\")", defaultConfig))
}

func initConfig() {
	cfgFile := config
	if cfgFile == "" {
		cfgFile = defaultConfig
	}
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		switch err.(type) {
		case *os.PathError:
			// ignore error when not using the config file
			if config != "" {
				fmt.Println(err)
			}
		default:
			fmt.Println(err)
		}
	}
}

func Execute() error {
	return rootCmd.Execute()
}
