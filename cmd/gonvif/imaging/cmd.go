package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
)

var Command = &cobra.Command{
	Use:   "imaging",
	Short: "Manipulate Onvif device imaging features.",
}

func init() {
	root.Command.AddCommand(Command)
	Command.AddCommand(
		getImagingSettings,
	)
}

func ServiceClient(url, username, password string, vebose bool) (wsdl.ImagingPort, error) {
	serviceURL, err := root.ServiceURL(url, "onvif/Imaging")
	if err != nil {
		return nil, err
	}
	return wsdl.NewImagingPort(root.AuthorizedSOAPClient(serviceURL, username, password, vebose)), nil
}
