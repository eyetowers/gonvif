package subscription

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var (
	subscription string
)

var cmd = &cobra.Command{
	Use:   "subscription",
	Short: "Manipulate Onvif events pull point subscription.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		pullMessages,
	)
	cmd.PersistentFlags().StringVarP(&subscription, "subscription", "s", "", "URL of the Onvif pull point subscription.")
	root.MustMarkPersistentFlagRequired(cmd, "subscription")
}

func ServiceClient(url, username, password, subscription string, verbose bool) (wsdl.PullPointSubscription, error) {
	onvif, err := gonvif.New(url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Subscription(subscription)
}
