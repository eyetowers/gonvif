package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
)

var systemReboot = &cobra.Command{
	Use:   "system-reboot",
	Short: "Reboot Onvif device",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runSystemReboot(client)
	},
}

func runSystemReboot(client wsdl.Device) error {
	resp, err := client.SystemReboot(&wsdl.SystemReboot{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
