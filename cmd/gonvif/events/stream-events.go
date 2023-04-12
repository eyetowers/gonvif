package events

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_w3_org/2005/08/addressing"
	"github.com/eyetowers/gonvif/pkg/gonvif"
	"github.com/eyetowers/gonvif/pkg/gonvif/axis"
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
	headers := extractHeaders(resp.SubscriptionReference)
	subscription, err := client.Subscription(string(*resp.SubscriptionReference.Address), headers...)
	if err != nil {
		return err
	}
	return processEvents(subscription)
}

func processEvents(subscription wsdl.PullPointSubscription) error {
	for {
		resp, err := subscription.PullMessages(&wsdl.PullMessages{MessageLimit: 100, Timeout: "PT60S"})
		if err != nil {
			return err
		}
		err = root.OutputJSON(resp)
		if err != nil {
			return err
		}
	}
}

func extractHeaders(ref *addressing.EndpointReferenceType) []any {
	// NOTE: Axis returns a SubscriptionId that has to be passed into the PullMessages call as SOAP
	// header as-is. We try to mimic this behavior here.
	var headers []any
	if axisHeader, ok := axis.ToHeader(ref); ok {
		headers = append(headers, axisHeader)
	}
	return headers
}
