// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Retrieve debug information for the current project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Workdir:\t\t%s\n", workDir)
		fmt.Printf("Debug:\t\t\t%t\n", debug)
		fmt.Printf("Dry-run:\t\t%t\n", dryRun)
		fmt.Printf("Version:\t\t%s\n", glblVersion.GetVersionString())
		fmt.Printf("Dirty:\t\t\t%t\n", glblGitRepo.IsDirty())
		fmt.Printf("Hash:\t\t\t%s\n", glblGitRepo.GetRevParse())
		fmt.Printf("Config:\t\t\t%s\n", viper.ConfigFileUsed())

		for _, element := range viper.AllKeys() {
			// element is the element from someSlice for where we are
			fmt.Printf(" - %s\t\t%s\n", element, viper.Get(element))
		}

		// DEBUG
		demoCommand := "docker ps -a -s -f $VERSION $HASH"
		application, arguments := createCommand(demoCommand)

		fmt.Printf("%s\n", application)
		fmt.Printf("%s\n", arguments)

		// Executing command
		// out, err := exec.Command(application, arguments...).Output()
		// if err != nil {
		// 	fmt.Printf("%s", err)
		// }
		// fmt.Printf("%s", out)
		// wg.Done() // Need to signal to waitgroup that this goroutine is done

		/////////

	},
}

func init() {
	RootCmd.AddCommand(debugCmd)
}
