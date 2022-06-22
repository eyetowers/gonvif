package ptz

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var Command = &cobra.Command{
	Use:   "ptz",
	Short: "Manipulate a PTZ device.",
}

func init() {
	root.Command.AddCommand(Command)
	Command.AddCommand(
		getConfigurations,
		getNodes,
	)
}

func ServiceClient(url, username, password string) (wsdl.PTZ, error) {
	u, err := root.ServiceURL(url, "onvif/PTZ")
	if err != nil {
		return nil, err
	}
	return wsdl.NewPTZ(root.AuthorizedSOAPClient(u, username, password)), nil
}
