package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
)

var getProfiles = &cobra.Command{
	Use:   "get-profiles",
	Short: "List Onvif device media profiles",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetProfiles(client)
	},
}

func runGetProfiles(client wsdl.Media) error {
	resp, err := client.GetProfiles(&wsdl.GetProfiles{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
