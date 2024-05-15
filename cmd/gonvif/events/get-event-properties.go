package events

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
)

var getEventProperties = &cobra.Command{
	Use:   "get-event-properties",
	Short: "Show Onvif detailed event propeties",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetEventProperties(client)
	},
}

func runGetEventProperties(client wsdl.EventPortType) error {
	resp, err := client.GetEventProperties(&wsdl.GetEventProperties{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
