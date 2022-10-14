package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getAnalyticsConfigurations = &cobra.Command{
	Use:   "get-analytics-configurations",
	Short: "Show Onvif device video analytics configurations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetAnalyticsConfigurations(client, configurationToken, profileToken)
	},
}

func init() {
	getAnalyticsConfigurations.Flags().StringVarP(&configurationToken, "configuration_token", "c", "", "Token of the requested configuration")
	getAnalyticsConfigurations.Flags().StringVarP(&profileToken, "profile_token", "t", "", "Contains the token of an existing media profile the configurations shall be compatible with")
}

func runGetAnalyticsConfigurations(
	client wsdl.Media2, configurationToken string, profileToken string,
) error {
	resp, err := client.GetAnalyticsConfigurations(&wsdl.GetConfiguration{
		ConfigurationToken: util.NewReferenceTokenPtr(configurationToken),
		ProfileToken:       util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
