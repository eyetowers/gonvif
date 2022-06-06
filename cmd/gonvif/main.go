package main

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"

	"github.com/eltrac-eu/gonvif/cmd/gonvif/completion"
	"github.com/eltrac-eu/gonvif/cmd/gonvif/ptz"
)

var RootCmd = &cobra.Command{
	Use:   "gonvif",
	Short: "Onvif CLI.",
	Long:  "Onvif CLI focused on PTZ cameras and implemented in pure Go.",
}

func init() {
	RootCmd.AddCommand(
		completion.Command,
		ptz.Command,
	)
}

func isTerminal(file *os.File) bool {
	o, _ := file.Stat()
	return (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func main() {
	if isTerminal(os.Stdout) && isTerminal(os.Stderr) {
		// Enable colored output if running in a terminal.
		cc.Init(&cc.Config{
			RootCmd:         RootCmd,
			Headings:        cc.HiCyan + cc.Bold + cc.Underline,
			Commands:        cc.HiYellow + cc.Bold,
			ExecName:        cc.Bold,
			Flags:           cc.Bold,
			NoBottomNewline: true,
		})
	}

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
