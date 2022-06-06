package ptz

import (
	"errors"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "ptz",
	Short: "Manipulate a PTZ device",
	Long:  "Communicate with a PTZ device using the Onvif protocol",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("not yet implemented")
	},
}
