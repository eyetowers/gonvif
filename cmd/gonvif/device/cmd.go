package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
)

var cmd = &cobra.Command{
	Use:   "device",
	Short: "Manipulate Onvif device management features.",
}

func init() {
	root.RequireAuthFlags(cmd)
	root.Command.AddCommand(cmd)
	cmd.AddCommand(
		getServices,
	)
}

func ServiceClient(url, username, password string, verbose bool) (wsdl.Device, error) {
	serviceURL, err := root.ServiceURL(url, "onvif/device_service")
	if err != nil {
		return nil, err
	}
	return wsdl.NewDevice(root.AuthorizedSOAPClient(serviceURL, username, password, verbose)), nil
}
