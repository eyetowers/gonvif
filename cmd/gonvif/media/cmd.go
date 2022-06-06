package media

import (
	"github.com/spf13/cobra"

	"github.com/eltrac-eu/gonvif/cmd/gonvif/root"
	"github.com/eltrac-eu/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
)

var Command = &cobra.Command{
	Use:   "media",
	Short: "Manipulate Onvif device media features.",
}

func init() {
	root.Command.AddCommand(Command)
	Command.AddCommand(
		getProfiles,
	)
}

func ServiceClient(url, username, password string) (wsdl.Media2, error) {
	serviceURL, err := root.ServiceURL(url, "onvif/Media")
	if err != nil {
		return nil, err
	}
	return wsdl.NewMedia2(root.AuthorizedSOAPClient(serviceURL, username, password)), nil
}
