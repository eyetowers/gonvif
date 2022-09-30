package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
	"github.com/eyetowers/gonvif/pkg/gonvif"
)

var cmd = &cobra.Command{
	Use:   "device",
	Short: "Manipulate Onvif device management features.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		getDeviceInformation,
		getServices,
		systemReboot,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.Device, error) {
	onvif, err := gonvif.New(url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Device()
}
