/*
Copyright © 2022 Gaurish Sethia <gaurishsethia@yahoo.com>
*/
package cmd

import (
	"bvm/lib"
	"fmt"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a specific release of Bun",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		grabClient := grab.NewClient()
		platform := lib.GetPlatform()
		arch := lib.GetArch()
		if platform == "other" || arch == "other" {
			fmt.Println("Unsupported platform")
			return
		}
		release, _ := cmd.Flags().GetString("release")
		if release == "" {
			lRelease, err := lib.GetLatestRelease()
			if err != nil {
				fmt.Println("Failed to get latest release: ", err)
				return
			}
			release = *lRelease.TagName
		}
		req, _ := grab.NewRequest(".", fmt.Sprintf("https://github.com/oven-sh/bun/releases/download/%s/bun-%s-%s.zip", release, platform, arch))

		fmt.Printf("Downloading %v...\n", req.URL())

		resp := grabClient.Do(req)
		fmt.Printf("  %v\n", resp.HTTPResponse.Status)

		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()

	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("  Downloading [%v / %v bytes] (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size(),
					100*resp.Progress())

			case <-resp.Done:
				break Loop
			}
		}
		if err := resp.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Download saved to ./%v \n", resp.Filename)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("release", "r", "", "Release to install")
}
