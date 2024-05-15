package events

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
)

var createPullPointSubscription = &cobra.Command{
	Use:   "create-pull-point-subscription",
	Short: "Create Onvif events pull point subscription",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runCreatePullPointSubscription(client)
	},
}

func runCreatePullPointSubscription(client wsdl.EventPortType) error {
	resp, err := client.CreatePullPointSubscription(&wsdl.CreatePullPointSubscription{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
