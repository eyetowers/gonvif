package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getCompatibleMetadataConfigurations = &cobra.Command{
	Use:   "get-compatible-metadata-configurations",
	Short: "List Onvif device media metadata configurations compatible with the profile",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetCompatibleMetadataConfigurations(client)
	},
}

func init() {
	getCompatibleMetadataConfigurations.Flags().StringVarP(&profileToken, "profile_token", "t", "", "Contains the token of an existing media profile the configurations shall be compatible with.")
	root.MustMarkFlagRequired(getCompatibleMetadataConfigurations, "profile_token")
}

func runGetCompatibleMetadataConfigurations(client wsdl.Media) error {
	resp, err := client.GetCompatibleMetadataConfigurations(&wsdl.GetCompatibleMetadataConfigurations{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
