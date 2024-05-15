package media2

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getVideoEncoderConfigurations = &cobra.Command{
	Use:   "get-video-encoder-configurations",
	Short: "List Onvif device video encoder configurations",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetVideoEncoderConfigurations(client, configurationToken, profileToken)
	},
}

func init() {
	getVideoEncoderConfigurations.Flags().StringVarP(&configurationToken, "configuration_token", "c", "", "Token of the requested configuration")
	getVideoEncoderConfigurations.Flags().StringVarP(&profileToken, "profile_token", "t", "", "Contains the token of an existing media profile the configurations shall be compatible with")
}

func runGetVideoEncoderConfigurations(
	client wsdl.Media2, configurationToken string, profileToken string,
) error {
	resp, err := client.GetVideoEncoderConfigurations(&wsdl.GetConfiguration{
		ConfigurationToken: util.NewReferenceTokenPtr(configurationToken),
		ProfileToken:       util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
