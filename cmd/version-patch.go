// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var versionPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Retrieve the version as it would be, if we patched it",
	Long: `I got the feeling you wondered what the version would be if we bumped
it with just a tiny patch, right?`,
	Run: func(cmd *cobra.Command, args []string) {
		// Temporary patch version
		glblVersion.Patch()
		// Print this string.
		fmt.Printf("%s\n", glblVersion.GetVersionString())
	},
}

func init() {
	versionCmd.AddCommand(versionPatchCmd)
}
