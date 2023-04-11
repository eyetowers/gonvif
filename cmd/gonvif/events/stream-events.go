package events

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var streamEvents = &cobra.Command{
	Use:   "stream-events",
	Short: "Stream live Onvif events from the device.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		onvif, err := gonvif.New(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runStreamEvents(onvif)
	},
}

func runStreamEvents(client gonvif.Client) error {
	events, err := client.Events()
	if err != nil {
		return err
	}
	resp, err := events.CreatePullPointSubscription(&wsdl.CreatePullPointSubscription{})
	if err != nil {
		return err
	}
	subscription, err := client.Subscription(string(*resp.SubscriptionReference.Address))
	if err != nil {
		return err
	}
	return processEvents(subscription)
}

func processEvents(subscription wsdl.PullPointSubscription) error {
	for {
		resp, err := subscription.PullMessages(&wsdl.PullMessages{MessageLimit: 100})
		if err != nil {
			return err
		}
		err = root.OutputJSON(resp)
		if err != nil {
			return err
		}
	}
}
