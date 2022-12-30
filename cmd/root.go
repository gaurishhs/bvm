/*
Copyright © 2022 Gaurish Sethia <gaurishsethia@yahoo.com>
*/
package cmd

import (
	"bvm/lib"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "bvm",
	Long: `Bun Version Manager (bvm) is a tool to manage Bun releases.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(lib.InitalizeClient)
}
