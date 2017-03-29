// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var workDir string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "release-support",
	Short: "Control your versioning",
	Long: `Control the versioning of your GIT repositories by using this
little tool. It gives you the ability to automatically build test and
tag your software. If everything succeeds, it will continue.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	defaultWorkdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// Global flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.release-support.yaml)")
	RootCmd.PersistentFlags().BoolP("dry-run", "d", false, "Enable dry-run mode")
	RootCmd.PersistentFlags().StringVar(&workDir, "w", defaultWorkdir, "Provide work-dir")

	// Set working-dir
	os.Chdir(workDir)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".release-support") // name of config file (without extension)
	viper.AddConfigPath("$HOME")            // adding home directory as first search path
	viper.AutomaticEnv()                    // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
