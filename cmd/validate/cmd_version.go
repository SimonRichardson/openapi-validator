package main

import (
	"flag"

	"github.com/spoke-d/clui"
	"github.com/spoke-d/clui/commands"
	"github.com/spoke-d/clui/flagset"
	"github.com/spoke-d/clui/ui"
	"github.com/spoke-d/task/group"
)

type versionCmd struct {
	ui      clui.UI
	flagSet *flagset.FlagSet

	clientVersion string
	template      string
}

func versionCmdFn(version string) func(clui.UI) clui.Command {
	return func(ui clui.UI) clui.Command {
		cmd := &versionCmd{
			ui:      ui,
			flagSet: flagset.New("version", flag.ContinueOnError),

			clientVersion: version,
		}
		cmd.init()
		return cmd
	}
}

func (v *versionCmd) init() {
	v.flagSet.StringVar(&v.template, "template", "{{.ClientVersion}}", "Template for the version template")
}

func (v *versionCmd) FlagSet() *flagset.FlagSet {
	return v.flagSet
}

func (v *versionCmd) Usages() []string {
	return make([]string, 0)
}

func (v *versionCmd) Help() string {
	return `
Show the current client version.
`
}

func (v *versionCmd) Synopsis() string {
	return "Show client version."
}

func (v *versionCmd) Init([]string, commands.CommandContext) error {
	return nil
}

func (v *versionCmd) Run(g *group.Group) {
	type version struct {
		ClientVersion string
	}

	template := ui.NewTemplate(versionTemplate, ui.OptionFormat(v.template))
	g.Add(func() error {
		return v.ui.Output(template, version{
			ClientVersion: v.clientVersion,
		})
	}, commands.Disguard)
}

const versionTemplate = `
Client version: %s
`
