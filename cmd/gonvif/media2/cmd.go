package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var cmd = &cobra.Command{
	Use:   "media2",
	Short: "Manipulate Onvif device media features.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		getAnalyticsConfigurations,
		getProfiles,
		getSnapshotURI,
		getStreamURI,
		getVideoEncoderConfigurations,
		getVideoSourceConfigurations,
		setSynchronizationPoint,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.Media2, error) {
	onvif, err := gonvif.New(url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Media2()
}
