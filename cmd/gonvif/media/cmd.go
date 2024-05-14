package media

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var cmd = &cobra.Command{
	Use:   "media",
	Short: "Manipulate Onvif device media features.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		getProfiles,
		getSnapshotURI,
		getStreamURI,
		setSynchronizationPoint,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.Media, error) {
	onvif, err := gonvif.New(context.Background(), url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Media()
}
