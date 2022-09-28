package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
)

var getServices = &cobra.Command{
	Use:   "get-services",
	Short: "List Onvif device services",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetServices(client)
	},
}

func runGetServices(client wsdl.Device) error {
	resp, err := client.GetServices(&wsdl.GetServices{
		IncludeCapability: true,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
