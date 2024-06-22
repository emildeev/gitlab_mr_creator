/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/emil110778/gitlab_mr_creator/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	LogLevelInfo    = "info"
	LogLevelWarning = "warning"
	LogLevelError   = "error"
	LogLevelDebug   = "debug"
	DefaultLogLevel = LogLevelInfo
	LogLevelLen     = 3
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "tools",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		TraverseChildren: true,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	logLevelUsage := fmt.Sprintf(
		"set app log level (%s, %s, %s, %s) default is %s",
		LogLevelError, LogLevelWarning, LogLevelInfo, LogLevelDebug, DefaultLogLevel,
	)

	rootCmd.PersistentFlags().StringP(
		"log_level", "l", DefaultLogLevel, logLevelUsage,
	)
}

func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".fullstack")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	InitLogger()
}

func InitLogger() {
	logLevel, _ := rootCmd.PersistentFlags().GetString("log_level")
	slog.SetLogLoggerLevel(getSlogLogLevel(logLevel))
}

func getSlogLogLevel(strLevel string) slog.Level {
	strLevel = makeLogLevelStr(strLevel)

	switch strLevel {
	case makeLogLevelStr(LogLevelInfo):
		return slog.LevelInfo
	case makeLogLevelStr(LogLevelWarning):
		return slog.LevelWarn
	case makeLogLevelStr(LogLevelError):
		return slog.LevelError
	case makeLogLevelStr(LogLevelDebug):
		return slog.LevelDebug
	default:
		return getSlogLogLevel(DefaultLogLevel)
	}
}

func makeLogLevelStr(s string) string {
	return strings.ToLower(helper.StringTruncate(s, LogLevelLen))
}
