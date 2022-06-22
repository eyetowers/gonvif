package main

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"

	_ "github.com/eyetowers/gonvif/cmd/gonvif/completion"
	_ "github.com/eyetowers/gonvif/cmd/gonvif/media"
	_ "github.com/eyetowers/gonvif/cmd/gonvif/ptz"
	"github.com/eyetowers/gonvif/cmd/gonvif/root"
)

func isTerminal(file *os.File) bool {
	o, _ := file.Stat()
	return (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func main() {
	if isTerminal(os.Stdout) && isTerminal(os.Stderr) {
		// Enable colored output if running in a terminal.
		cc.Init(&cc.Config{
			RootCmd:         root.Command,
			Headings:        cc.HiCyan + cc.Bold + cc.Underline,
			Commands:        cc.HiYellow + cc.Bold,
			ExecName:        cc.Bold,
			Flags:           cc.Bold,
			NoBottomNewline: true,
		})
	}

	if err := root.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
