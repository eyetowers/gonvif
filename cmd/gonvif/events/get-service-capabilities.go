package events

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
)

var getServiceCapabilities = &cobra.Command{
	Use:   "get-service-capabilities",
	Short: "Show Onvif events capabilities",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetServiceCapabilities(client)
	},
}

func runGetServiceCapabilities(client wsdl.EventPortType) error {
	resp, err := client.GetServiceCapabilities(&wsdl.GetServiceCapabilities{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
