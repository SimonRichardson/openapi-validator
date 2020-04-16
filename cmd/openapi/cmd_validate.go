package main

import (
	"flag"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"github.com/spoke-d/clui"
	"github.com/spoke-d/clui/commands"
	"github.com/spoke-d/clui/flagset"
	"github.com/spoke-d/task/group"
)

type validateCmd struct {
	ui      clui.UI
	flagSet *flagset.FlagSet

	file string
}

func validateCmdFn() func(clui.UI) clui.Command {
	return func(ui clui.UI) clui.Command {
		cmd := &validateCmd{
			ui:      ui,
			flagSet: flagset.New("validate", flag.ContinueOnError),
		}
		return cmd
	}
}

func (v *validateCmd) FlagSet() *flagset.FlagSet {
	return v.flagSet
}

func (v *validateCmd) Usages() []string {
	return make([]string, 0)
}

func (v *validateCmd) Help() string {
	return `
Validate OpenAPI yaml file to identify any issues.

Under the hood, this uses the kin-openapi validator,
as there isn't a equivalent validator from the OpenAPI
standards.
`
}

func (v *validateCmd) Synopsis() string {
	return "Validate OpenAPI yaml file to identify any issues."
}

func (v *validateCmd) Init(args []string, ctx commands.CommandContext) error {
	if len(args) != 1 {
		return errors.Errorf("expected path to yaml file: %q", args)
	}
	v.file = args[0]
	return nil
}

func (v *validateCmd) Run(g *group.Group) {
	g.Add(func() error {
		loader := openapi3.NewSwaggerLoader()
		swagger, err := loader.LoadSwaggerFromFile(v.file)
		if err != nil {
			return errors.WithStack(err)
		}

		return swagger.Validate(loader.Context)
	}, commands.Disguard)
}
