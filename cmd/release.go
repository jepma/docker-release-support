// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func releasePreCheck() (tagLatest string, err error) {

	// Check if repository is dirty
	if glblVersion.GetDirty() == true {
		return "", ErrChangesPending
	}

	// Check if diff from latest tag
	tagLatest, tagLatestErr := glblGitRepo.TagLatest(fmt.Sprintf("%s", viper.Get("tag-prefix")))
	if tagLatestErr != nil {
		tagLatest = versionFormat
	}

	// Check if tag exists, if not: go-go-go!
	if glblGitRepo.TagExists(tagLatest) != true {
		return tagLatest, nil
	}

	vDiff := glblGitRepo.DiffSinceRelease(tagLatest)

	if vDiff == false {
		fmt.Printf("Repository contains no changes since TAG %s.\n", tagLatest)
		return tagLatest, ErrNoChanges
	}

	return tagLatest, nil

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
func parseTagParameters(value string) string {
	value = strings.Replace(value, "@TAG@", glblVersion.GetVersionTag(), -1)
	value = strings.Replace(value, "@VERSION@", glblVersion.GetVersionString(), -1)

	return value
}

func parseParametersList(valueList []string) []string {

	returnList := make([]string, len(valueList))

	for index, element := range valueList {
		returnList[index] = parseTagParameters(element)
	}

	return returnList
}
