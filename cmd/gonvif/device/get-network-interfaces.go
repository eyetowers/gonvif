package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
)

var getNetworkInterfaces = &cobra.Command{
	Use:   "get-network-interfaces",
	Short: "List Onvif device network interfaces",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetNetworkInterfaces(client)
	},
}

func runGetNetworkInterfaces(client wsdl.Device) error {
	resp, err := client.GetNetworkInterfaces(&wsdl.GetNetworkInterfaces{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
