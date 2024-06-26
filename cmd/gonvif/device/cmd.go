package device

import (
	"context"

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
		getNetworkInterfaces,
		getServices,
		getSystemTime,
		setSystemFactoryDefault,
		systemReboot,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.Device, error) {
	onvif, err := gonvif.New(context.Background(), url, username, password, verbose)
	if err != nil {
		return nil, err
	}
	return onvif.Device()
}
