package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
)

var getDeviceInformation = &cobra.Command{
	Use:   "get-device-information",
	Short: "Show Onvif device information",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetDeviceInformation(client)
	},
}

func runGetDeviceInformation(client wsdl.Device) error {
	resp, err := client.GetDeviceInformation(&wsdl.GetDeviceInformation{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
