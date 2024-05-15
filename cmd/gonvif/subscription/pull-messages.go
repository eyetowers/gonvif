package subscription

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
)

var pullMessages = &cobra.Command{
	Use:   "pull-messages",
	Short: "Pulls Onvif events from a pull point subscription",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, subscription, root.Verbose)
		if err != nil {
			return err
		}
		return runPullMessages(client)
	},
}

func runPullMessages(client wsdl.PullPointSubscription) error {
	resp, err := client.PullMessages(&wsdl.PullMessages{
		MessageLimit: 100,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
