package events

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
)

var (
	address string
)

var getEventBrokers = &cobra.Command{
	Use:   "get-event-brokers",
	Short: "List Onvif event brokers",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetEventBrokers(client, address)
	},
}

func init() {
	getEventBrokers.Flags().StringVarP(&address, "broker_address", "b", "", "Optional address to query for.")
}

func runGetEventBrokers(client wsdl.EventPortType, address string) error {
	resp, err := client.GetEventBrokers(&wsdl.GetEventBrokers{Address: address})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
