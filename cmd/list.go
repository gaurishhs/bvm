/*
Copyright © 2022 Gaurish Sethia <gaurishsethia@yahoo.com>
*/
package cmd

import (
	"bvm/lib"
	"fmt"
	"os"

	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "A brief description of your command",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		releases, err := lib.GetAllReleases()
		if err != nil {
			fmt.Println(err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "VERSION\tDATE\t")
		for _, release := range releases {
			fmt.Fprintf(w, "%s\t%s\t \n", *release.TagName, *release.PublishedAt)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
