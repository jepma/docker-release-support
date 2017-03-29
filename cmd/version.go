// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"

	"github.com/jepma/go-versioning"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieve the current version of your repository",
	Long: `Retrieve the current version for your repository. It is also possible to fetch possible version number; For example:

"release-support version" : 0.0.1
"release-support version --patch" : 0.0.2
"release-support version --minor" : 0.1.1
"release-support version --major" : 1.0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("version called")

		// Run version.get
		version := versioning.CreateVersion("1.0.0")
		fmt.Printf("%s\n", version.GetVersionString())
		fmt.Printf("workdir: %s\n", workDir)

	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	versionCmd.Flags().BoolP("patch", "p", false, "Retrieve next patched version")

}
