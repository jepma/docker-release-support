// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"
	"os"
	"strings"

	git "github.com/jepma/go-git"
	versioning "github.com/jepma/go-versioning"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var noCfgFile bool
var workDir string
var dryRun bool
var debug bool
var versionFormat string

// var glblVersionCurrent versioning.Version // DEPRECATED
// var glblVersionRelease versioning.Version // DEPRECATED
var glblVersion versioning.Version
var glblGitRepo git.Repo

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
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Determine working-dir
	defaultWorkdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// Global flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $(pwd)/release.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "D", false, "Enable dry-run mode")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable dry-run mode")
	RootCmd.PersistentFlags().BoolVarP(&noCfgFile, "no-config", "n", false, "Do not use a config-file but rely on GIT")
	RootCmd.PersistentFlags().StringVar(&workDir, "workdir", defaultWorkdir, "Provide work-dir")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Some debug shizzle
	if debug == true {
		fmt.Println("DEBUG ENABLED!")
		fmt.Println("Using workdir:", workDir)
	}

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("release") // name of config file (without extension)
	viper.AddConfigPath(workDir)   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	// Set working-dir
	os.Chdir(workDir)

	// Create Global GIT Repository
	glblGitRepo, gitRepoErr := git.CreateRepoObject(workDir)
	if gitRepoErr != nil {
		fmt.Printf("There was an error reading GIT Repository %s", workDir)
		os.Exit(1)
		// glog.Fatal("Error with GIT Repo: ", gitRepoErr)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && debug == true {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		versionFormat = "v0.0.1"

		tagLatest, tagLatestErr := glblGitRepo.TagLatest()
		if tagLatestErr != nil {
			tagLatest = versionFormat
		}

		viper.SetDefault("version", strings.Split(tagLatest, "v")[1])
	}

	// Set current version struct
	glblVersion = versioning.CreateVersion(fmt.Sprintf("%s", viper.Get("version")))
	glblVersion.SetDirty(glblGitRepo.IsDirty())

	// Check if there is a diff
	tagLatest, tagLatestErr := glblGitRepo.TagLatest()
	if tagLatestErr != nil {
		tagLatest = "v0.0.1"
	}
	vDiff := glblGitRepo.DiffSinceRelease(tagLatest)
	if vDiff == true {
		glblVersion.SetHash(glblGitRepo.GetRevParse())
	}

}
