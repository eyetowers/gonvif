package media2

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
)

var (
	types []string
)

var getProfiles = &cobra.Command{
	Use:   "get-profiles",
	Short: "List Onvif device media profiles",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetProfiles(client, types)
	},
}

func init() {
	getProfiles.Flags().StringArrayVarP(&types, "types", "t", []string{"All"}, "Types of profile configurations to include")
}

func runGetProfiles(client wsdl.Media2, types []string) error {
	resp, err := client.GetProfiles(&wsdl.GetProfiles{
		Type: types,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
