/*
Copyright © 2022 Gaurish Sethia <gaurishsethia@yahoo.com>
*/
package cmd

import (
	"bvm/lib"
	"fmt"

	"github.com/spf13/cobra"
)

var latestCmd = &cobra.Command{
	Use:     "latest",
	Short:   "Get the latest release of Bun",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		latest, err := lib.GetLatestRelease()
		if err != nil {
			fmt.Println("Failed to get latest release: ", err)
			return
		}
		fmt.Printf("Latest release: %s (%s)\nTo install: bvm install --release %[1]s\n", *latest.TagName, latest.PublishedAt)
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
}
