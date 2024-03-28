/*
Copyright © 2024 Juha Ruotsalainen <juha.ruotsalainen@iki.fi>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var groupId string
var artifactId string
var version string

const GID = "group-id"
const AID = "artifact-id"
const VER = "parent-version"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pommod [flags] PATH-TO-POM",
	Short:   "A simple program to change parent in a pom.xml",
	Version: "v1.0.0",
	Args:    cobra.MinimumNArgs(1),
	Run:     rootRunner,
}

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
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMilli})
		return nil
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pommod.yaml)")
	rootCmd.PersistentFlags().StringVarP(&groupId, GID, "g", "", "New parent's group ID")
	rootCmd.PersistentFlags().StringVarP(&artifactId, AID, "a", "", "New parent's artifact ID")
	rootCmd.PersistentFlags().StringVarP(&version, VER, "V", "", "New parent's version")
	rootCmd.MarkPersistentFlagRequired(GID)
	rootCmd.MarkPersistentFlagRequired(AID)
	rootCmd.MarkPersistentFlagRequired(VER)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pommod")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
