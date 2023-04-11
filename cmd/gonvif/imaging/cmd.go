package imaging

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var cmd = &cobra.Command{
	Use:   "imaging",
	Short: "Manipulate Onvif device imaging features.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		getImagingSettings,
		getMoveOptions,
		getOptions,
		getServiceCapabilities,
		getStatus,
		moveContinuous,
		setImagingSettingsAutoFocus,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.ImagingPort, error) {
	onvif, err := gonvif.New(url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Imaging()
}
