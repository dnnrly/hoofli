package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"

	"github.com/dnnrly/hoofli"
)

type args struct {
	Help       bool      `cli:"!h,help" usage:"show help"`
	Input      clix.File `cli:"*i,input" usage:"the location of a HAR file to parse"`
	ExcludeURL []string  `cli:"exclude-url" usage:"regular expression of URLs to exclude from the diagram"`
}

// AutoHelp implements cli.AutoHelper interface
func (a *args) AutoHelp() bool {
	return a.Help
}

func main() {
	os.Exit(cli.Run(new(args), rootCmd))
}

func rootCmd(ctx *cli.Context) error {
	argv := ctx.Argv().(*args)
	ir := strings.NewReader(argv.Input.String())
	har, err := hoofli.NewHar(ir)
	if err != nil {
		return fmt.Errorf("unable to parse HAR file: %w", err)
	}

	for _, pattern := range argv.ExcludeURL {
		har.Log.Entries = har.Log.Entries.ExcludeByURL(pattern)
	}

	_ = har.Draw(os.Stdout)
	return nil
}
