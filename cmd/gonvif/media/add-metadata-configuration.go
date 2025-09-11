package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	configurationToken string
)

var addMetadataConfiguration = &cobra.Command{
	Use:   "add-metadata-configuration",
	Short: "Associate Onvif device media metadata configuration to the profile",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runAddMetadataConfiguration(client)
	},
}

func init() {
	addMetadataConfiguration.Flags().StringVarP(&profileToken, "profile_token", "t", "", "Reference to the profile where the configuration should be added.")
	addMetadataConfiguration.Flags().StringVarP(&configurationToken, "configuration_token", "c", "", "Contains a reference to the MetadataConfiguration to add.")
	root.MustMarkFlagRequired(addMetadataConfiguration, "profile_token")
	root.MustMarkFlagRequired(addMetadataConfiguration, "configuration_token")
}

func runAddMetadataConfiguration(client wsdl.Media) error {
	resp, err := client.AddMetadataConfiguration(&wsdl.AddMetadataConfiguration{
		ProfileToken:       util.NewReferenceTokenPtr(profileToken),
		ConfigurationToken: util.NewReferenceTokenPtr(configurationToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
