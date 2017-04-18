// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionMajorCmd represents the patch command
var versionMajorCmd = &cobra.Command{
	Use:   "major",
	Short: "Retrieve the version as it would be, if we patched it",
	Long: `And then we virtually bumped the version with a
tiny major, major!`,
	Run: func(cmd *cobra.Command, args []string) {
		// Temporary patch version
		glblVersion.Major()
		// Print this string.
		fmt.Printf("%s\n", glblVersion.GetVersionString())
	},
}

func init() {
	versionCmd.AddCommand(versionMajorCmd)
}
