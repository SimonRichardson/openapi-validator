package main

import (
	"fmt"
	"os"

	"github.com/spoke-d/clui"
	"github.com/spoke-d/clui/autocomplete/fsys"
)

const clientVersion = "0.0.1-alpha"

func main() {
	fsys := fsys.NewLocalFileSystem()

	cli := clui.New("openapi", clientVersion, header, clui.OptionFileSystem(fsys))
	cli.Add("version", versionCmdFn(clientVersion))
	cli.Add("validate", validateCmdFn())

	code, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(code.Code())
}

const header = `
_██████╗ ██████╗ ███████╗███╗   ██╗ █████╗ ██████╗ ██╗
██╔═══██╗██╔══██╗██╔════╝████╗  ██║██╔══██╗██╔══██╗██║
██║   ██║██████╔╝█████╗  ██╔██╗ ██║███████║██████╔╝██║
██║   ██║██╔═══╝ ██╔══╝  ██║╚██╗██║██╔══██║██╔═══╝ ██║
╚██████╔╝██║     ███████╗██║ ╚████║██║  ██║██║     ██║
 ╚═════╝ ╚═╝     ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝
`
