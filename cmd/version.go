// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieve the current version of your repository",
	Long: `Retrieve the current version for your repository. It is also possible to fetch possible version number; For example:

"release-support version" : 0.0.1
"release-support version patch" : 0.0.2
"release-support version minor" : 0.1.1
"release-support version major" : 1.0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", glblVersion.GetVersionString())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
