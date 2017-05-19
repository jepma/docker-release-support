// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// releasePatchCmd represents the patch command
var releasePatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Bump your version with PATCH level.",
	Long: `By running the release patch command we will tag your code
and make sure it will be properly pushed upstream.

By running release patch we will bump it with PATCH level.

We will check if:
- there are no pending changes
- actually is a change since last commit`,
	Run: func(cmd *cobra.Command, args []string) {
		releasePatch()
	},
}

func init() {
	releaseCmd.AddCommand(releasePatchCmd)
}

func releasePatch() {

	tagLatest := releasePreCheck()

	// Run pre-bump code
	fmt.Print("@TODO: Implement pre-bump code\n")
	// application, arguments := createCommand(demoCommand)

	// Version Patch
	glblVersion.Patch()

	// Output
	fmt.Printf("Patched to version: %s\n", glblVersion.GetVersionString())

	tagStatus, tagErr := glblGitRepo.TagCreate(tagLatest, parseParameters(glblVersion.GetVersionTag()), "")
	if tagStatus == false {
		fmt.Printf("An error occured creating the tag: %s.\n", tagErr)
		os.Exit(1)
	}

	// Run post-bump code
	fmt.Print("@TODO: Implement post-bump code\n")

}
