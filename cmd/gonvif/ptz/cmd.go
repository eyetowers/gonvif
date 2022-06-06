package ptz

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "ptz",
	Short: "Manipulate a PTZ device",
}

func init() {
	Command.AddCommand(
		getConfigurations,
		getNodes,
	)
}
