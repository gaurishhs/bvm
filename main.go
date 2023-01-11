package main

import (
	"bvm/bun"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/bowerswilkins/humantime"
	"github.com/urfave/cli/v2"
)

func main() {
	bvm := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "node",
				Aliases: []string{"n"},
				Usage:   "Manage node versions",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list node versions",
						Action: func(c *cli.Context) error {
							return nil
						},
					},
				},
			},
			{
				Name:    "bun",
				Aliases: []string{"b"},
				Usage:   "Manage Bun versions",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list bun versions",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "latest",
								Value:   false,
								Aliases: []string{"l"},
								Usage:   "list latest version",
							},
						},
						Action: func(c *cli.Context) error {
							releases, err := bun.ListVersions()
							if err != nil {
								return err
							}
							latest := c.Bool("latest")
							if latest {
								parsedTime, err := time.Parse(time.RFC3339, releases[0].PublishedAt)
								if err != nil {
									return err
								}
								fmt.Println("Latest version:", releases[0].Name+" ("+releases[0].TagName+")")
								fmt.Println("Release date (relative):", humantime.Time(parsedTime))
							} else {
								w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
								fmt.Fprintln(w, "VERSION\tRELATIVE TIME\t")
								for _, release := range releases {
									parsedTime, err := time.Parse(time.RFC3339, release.PublishedAt)
									if err != nil {
										return err
									}
									fmt.Fprintf(w, "%s\t%s\t \n", release.TagName, humantime.Time(parsedTime))
								}
								w.Flush()
							}
							return nil
						},
					},
				},
			},
		},
	}

	if err := bvm.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
