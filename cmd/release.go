// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Tag and bag your code",
	Long: `By running the release command we will tag your code
and make sure it will be properly pushed upstream.

We will check if:
- there are no pending changes
- actually is a change since last commit

By default we will release patch it for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Please provide your release type [patch, minor, major]")
	},
}

func init() {
	RootCmd.AddCommand(releaseCmd)
}

func releasePreCheck() (tagLatest string) {

	// Check if repository is dirty
	if glblVersion.GetDirty() == true {
		fmt.Printf("Repository contains changes, cannot release patch with a dirty tree.\n")
		os.Exit(1)
	}

	// Check if diff from latest tag
	tagLatest, tagLatestErr := glblGitRepo.TagLatest()
	if tagLatestErr != nil {
		tagLatest = versionFormat
	}
	vDiff := glblGitRepo.DiffSinceRelease(tagLatest)

	if vDiff == false {
		fmt.Printf("Repository contains no changes since TAG %s.\n", tagLatest)
		os.Exit(1)
	}

	return tagLatest

}

func createCommand(command string) (application string, arguments []string) {

	// split command string
	parts := strings.Fields(command)
	application = parts[0]
	arguments = parts[1:len(parts)]

	// check for parameters
	arguments = parseParametersList(arguments)

	return application, arguments
}

// Replace parameters within given value string
func parseParameters(value string) string {
	value = strings.Replace(value, "$TAG", glblVersion.GetVersionTag(), -1)
	value = strings.Replace(value, "$VERSION", glblVersion.GetVersionString(), -1)
	value = strings.Replace(value, "$HASH", glblGitRepo.GetRevParse(), -1)
	return value
}

func parseParametersList(valueList []string) []string {

	returnList := make([]string, len(valueList))

	for index, element := range valueList {
		returnList[index] = parseParameters(element)
	}

	return returnList
}
