package events

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	wsnt "github.com/eyetowers/gonvif/pkg/generated/onvif/docs_oasisopen_org/wsn/b2"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

const (
	unsubscribeTimeout = 2 * time.Second
	pollTimeout        = "PT60S"
)

var (
	subscriptionTimeout wsnt.AbsoluteOrRelativeTimeType = "PT120S"
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
	resp, err := events.CreatePullPointSubscription(&wsdl.CreatePullPointSubscription{
		InitialTerminationTime: &subscriptionTimeout,
	})
	if err != nil {
		return err
	}
	headers := gonvif.ComposeHeaders(resp.SubscriptionReference)
	subscription, err := client.Subscription(string(*resp.SubscriptionReference.Address), headers...)
	if err != nil {
		return err
	}
	return processEvents(subscription)
}

func processEvents(subscription wsdl.PullPointSubscription) error {
	defer func() { _ = unsubscribe(subscription) }()
	for {
		resp, err := subscription.PullMessages(&wsdl.PullMessages{MessageLimit: 100, Timeout: pollTimeout})
		if err != nil {
			return err
		}
		err = root.OutputJSON(resp)
		if err != nil {
			return err
		}
		_, err = subscription.Renew(&wsnt.Renew{TerminationTime: &subscriptionTimeout})
		if err != nil {
			return err
		}
	}
}

func unsubscribe(subscription wsdl.PullPointSubscription) error {
	ctx, cancel := context.WithTimeout(context.Background(), unsubscribeTimeout)
	defer cancel()

	var empty wsdl.EmptyString
	_, err := subscription.UnsubscribeContext(ctx, &empty)
	return err
}
