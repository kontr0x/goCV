package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kontr0x/goCV/pkg/content"
	cli "github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "goCV",
		Usage: "CV generator using Go and LaTeX",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:  "content",
				Value: "./content.yaml",
			},
		},
		ArgsUsage: "content.yaml",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "output", Value: "./output", Usage: "Output directory for the rendered CV"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			data, err := content.ParseContentFromYaml(c.StringArg("content"))
			t := time.Now().Format("02-01-2006_15:04")
			if err != nil {
				panic(err)
			}
			for _, entry := range data {
				dir := fmt.Sprintf("%s/%s_%s", c.String("output"), entry.Content.Version, t)
				err := content.RenderTemplate("", "", dir, entry)
				if err != nil {
					panic(err)
				}
				err = content.BuildTemplate(dir)
				if err != nil {
					panic(err)
				}
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
