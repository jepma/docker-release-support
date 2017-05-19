// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"
	"os"
	"strings"

	git "github.com/jepma/go-git"
	versioning "github.com/jepma/go-versioning"
	vipersaveconfig "github.com/jepma/go-viper-saveconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var noCfgFile bool
var workDir string
var actionn bool
var dryRun bool
var debug bool
var versionFormat string

// Type of release
var patchRelease bool
var minorRelease bool
var majorRelease bool

// Init
var initConfigFile string

var glblVersion versioning.Version
var glblGitRepo git.Repo

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "release-support",
	Short: "Control your versioning",
	Long: `Control the versioning of your GIT repositories by using this
little tool. It gives you the ability to automatically build test and
tag your software. If everything succeeds, it will continue.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Check if we need to update our release.yaml
		if initConfigFile != "" {

			initReleaseFile(initConfigFile)

		}

		// Check if we got actual tag command, if not display current version
		if patchRelease == false && minorRelease == false && majorRelease == false {
			fmt.Printf("%s\n", glblVersion.GetVersionString())
		} else {

			if _, err := releasePreCheck(); err != nil {

				// NoChanges should not trigger an error and thus no os.Exit(1)
				if err != ErrNoChanges {
					fmt.Printf("ERROR: %s\n", err)
					os.Exit(1)
				}

			} else {

				if patchRelease == true {
					glblVersion.Patch()
				} else if minorRelease == true {
					glblVersion.Minor()
				} else if majorRelease == true {
					glblVersion.Major()
				}

				tagStatus, tagErr := glblGitRepo.TagCreate(parseTagParameters(glblVersion.GetVersionTag()), "")
				if tagStatus == false {
					fmt.Printf("An error occured creating the tag: %s.\n", tagErr)
					os.Exit(1)
				}

				fmt.Printf("Bumped to version %s with tag %s\n", glblVersion.GetVersionString(), glblVersion.GetVersionTag())
			}
		}

	},
}

func initReleaseFile(tagPrefix string) {
	// Check if there is a config-file present
	if _, err := os.Stat(fmt.Sprintf("%s/release.yaml", workDir)); err == nil {
		fmt.Printf("ERROR: Release file already exists!\n")
		os.Exit(1)
	}

	// Set tag-prefix
	viper.Set("tag-prefix", tagPrefix)

	nwViper := viper.New()
	nwViper.SetConfigFile(fmt.Sprintf("%s/release.yaml", workDir))
	nwViper.Set("tag-prefix", viper.Get("tag-prefix"))
	nwViper.Set("pre-tag-command", viper.Get("pre-tag-command"))

	// Save config
	if err := vipersaveconfig.SaveConfig(*nwViper); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("INFO: Written to configfile: %s\n", fmt.Sprintf("%s/release.yaml", workDir))
	os.Exit(0)
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
		os.Exit(1)
	}

	RootCmd.Flags().BoolVar(&patchRelease, "patch-release", false, "patch bump")
	RootCmd.Flags().BoolVar(&minorRelease, "minor-release", false, "minor bump")
	RootCmd.Flags().BoolVar(&majorRelease, "major-release", false, "major bump")
	RootCmd.Flags().StringVar(&initConfigFile, "init", "", "initialize release file")

	// Global flags
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $(pwd)/release.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug information")
	RootCmd.PersistentFlags().StringVar(&workDir, "workdir", defaultWorkdir, "Provide work-dir")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	var cfgTagPrefix = "v"

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
	viper.SetDefault("tag-prefix", "v")

	// Set working-dir
	os.Chdir(workDir)

	// Create Global GIT Repository
	glblGitRepo, gitRepoErr := git.CreateRepoObject(workDir)
	if gitRepoErr != nil {
		fmt.Printf("There was an error reading GIT Repository %s", workDir)
		os.Exit(1)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && debug == true {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		cfgTagPrefix = fmt.Sprintf("%s", viper.Get("tag-prefix"))
		versionFormat = fmt.Sprintf("%s0.0.0", cfgTagPrefix)
		tagLatest, tagLatestErr := glblGitRepo.TagLatest(cfgTagPrefix)
		if tagLatestErr != nil {
			tagLatest = versionFormat
		}

		viper.SetDefault("version", strings.Split(tagLatest, cfgTagPrefix)[1])

	}

	// Set current version struct
	glblVersion = versioning.CreateVersion(fmt.Sprintf("%s", viper.Get("version")))
	glblVersion.SetDirty(glblGitRepo.IsDirty())
	glblVersion.SetPrefix(fmt.Sprintf("%s", viper.Get("tag-prefix")))

	// Check if there is a diff
	tagLatest, tagLatestErr := glblGitRepo.TagLatest(cfgTagPrefix)
	if tagLatestErr != nil {
		tagLatest = "v0.0.0"
	}

	vDiff := glblGitRepo.DiffSinceRelease(tagLatest)
	if vDiff == true {
		glblVersion.SetHash(glblGitRepo.GetRevParse())
	}

}
