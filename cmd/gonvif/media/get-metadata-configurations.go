package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
)

var getMetadataConfigurations = &cobra.Command{
	Use:   "get-metadata-configurations",
	Short: "List Onvif device media metadata configurations",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetMetadataConfigurations(client)
	},
}

func runGetMetadataConfigurations(client wsdl.Media) error {
	resp, err := client.GetMetadataConfigurations(&wsdl.GetMetadataConfigurations{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
