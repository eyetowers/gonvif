package ptz

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var getConfigurations = &cobra.Command{
	Use:   "get-configurations",
	Short: "List PTZ device configurations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password)
		if err != nil {
			return nil
		}
		return runGetConfigurations(client)
	},
}

func runGetConfigurations(client wsdl.PTZ) error {
	resp, err := client.GetConfigurations(&wsdl.GetConfigurations{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
