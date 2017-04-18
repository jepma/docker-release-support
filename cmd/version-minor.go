// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionMinorCmd represents the patch command
var versionMinorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Retrieve the version as it would be, if we minored it",
	Long: `And then we virtually bumped the version with a
tiny minor.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Temporary patch version
		glblVersion.Minor()
		// Print this string.
		fmt.Printf("%s\n", glblVersion.GetVersionString())
	},
}

func init() {
	versionCmd.AddCommand(versionMinorCmd)
}
