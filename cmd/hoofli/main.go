package main

import (
	"os"
	"strings"

	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"

	"github.com/dnnrly/hoofli"
)

type args struct {
	Help  bool      `cli:"!h,help" usage:"show help"`
	Input clix.File `cli:"*i,input" usage:"the location of a HAR file to parse"`
}

// AutoHelp implements cli.AutoHelper interface
func (a *args) AutoHelp() bool {
	return a.Help
}

func main() {
	os.Exit(cli.Run(new(args), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*args)
		ir := strings.NewReader(argv.Input.String())
		har, err := hoofli.NewHar(ir)
		if err != nil {
			return nil
		}

		_ = har.Draw(os.Stdout)
		return nil
	}))
}
