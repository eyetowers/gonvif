package media

import (
	"github.com/spf13/cobra"

	"github.com/eltrac-eu/gonvif/cmd/gonvif/root"
	"github.com/eltrac-eu/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
)

var getProfiles = &cobra.Command{
	Use:   "get-profiles",
	Short: "List Onvif device media profiles",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password)
		if err != nil {
			return nil
		}
		return runGetProfiles(client)
	},
}

func runGetProfiles(client wsdl.Media2) error {
	resp, err := client.GetProfiles(&wsdl.GetProfiles{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
