package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getVideoSourceConfigurations = &cobra.Command{
	Use:   "get-video-source-configurations",
	Short: "List Onvif device video source configurations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password)
		if err != nil {
			return nil
		}
		configurationToken, err := cmd.Flags().GetString("configuration_token")
		if err != nil {
			return nil
		}
		profileToken, err := cmd.Flags().GetString("profile_token")
		if err != nil {
			return nil
		}
		return runGetVideoSourceConfigurations(client, configurationToken, profileToken)
	},
}

func init() {
	getVideoSourceConfigurations.Flags().StringP("configuration_token", "c", "", "Token of the requested configuration")
	getVideoSourceConfigurations.Flags().StringP("profile_token", "t", "", "Contains the token of an existing media profile the configurations shall be compatible with")
}

func runGetVideoSourceConfigurations(
	client wsdl.Media2, configurationToken string, profileToken string,
) error {
	resp, err := client.GetVideoSourceConfigurations(&wsdl.GetConfiguration{
		ConfigurationToken: util.NewReferenceTokenPtr(configurationToken),
		ProfileToken:       util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
