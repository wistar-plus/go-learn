package main

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func init() {
	// config the viper
	viper.SetConfigName("config")
	//	viper.SetEnvPrefix("wweb")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.ReadInConfig()

	// viper.BindPFlag("addr", rootCmd.Flags().Lookup("addr"))
	// viper.BindPFlag("mode", rootCmd.Flags().Lookup("mode"))
	defaultConfig()

	// configure the cobra commands
	rootCmd = &cobra.Command{
		Use:   "main [command]",
		Short: "gin web site main section cli",
	}
	rootCmd.AddCommand(runCommand)
}

func main() {
	rootCmd.Execute()
}

func defaultConfig() {
	viper.SetDefault("addr", ":8080")
	viper.SetDefault("mode", "release")
	viper.SetDefault("log.filename", "logs/wweb.log")
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("DB.DRIVER", "mysql")
	viper.SetDefault("DB.USER", "root")
	viper.SetDefault("DB.PASSWORD", "root")
	viper.SetDefault("DB.NAME", "wweb")
	viper.SetDefault("DB.HOST", "127.0.0.1")
	viper.SetDefault("DB.PORT", "3306")
}
