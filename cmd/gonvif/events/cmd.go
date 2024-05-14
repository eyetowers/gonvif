package events

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var cmd = &cobra.Command{
	Use:   "events",
	Short: "Manipulate Onvif events streams.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		createPullPointSubscription,
		getEventBrokers,
		getEventProperties,
		getServiceCapabilities,
		streamEvents,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.EventPortType, error) {
	onvif, err := gonvif.New(context.Background(), url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Events()
}
