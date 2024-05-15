package device

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
)

var getSystemTime = &cobra.Command{
	Use:   "get-system-time",
	Short: "Show Onvif device system date and time",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetSystemTime(client)
	},
}

func runGetSystemTime(client wsdl.Device) error {
	resp, err := client.GetSystemDateAndTime(&wsdl.GetSystemDateAndTime{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
