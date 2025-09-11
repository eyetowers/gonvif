package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	enable bool
)

var setMetadataAnalytics = &cobra.Command{
	Use:   "set-metadata-analytics",
	Short: "Enable/disable analytics in the Onvif device media metadata configuration",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runSetMetadataAnalytics(client)
	},
}

func init() {
	setMetadataAnalytics.Flags().StringVarP(&configurationToken, "configuration_token", "c", "", "Contains a reference to the MetadataConfiguration to add.")
	setMetadataAnalytics.Flags().BoolVarP(&enable, "enable", "", true, "Enable the analytics.")
	root.MustMarkFlagRequired(addMetadataConfiguration, "configuration_token")
}

func runSetMetadataAnalytics(client wsdl.Media) error {
	config, err := client.GetMetadataConfiguration(&wsdl.GetMetadataConfiguration{
		ConfigurationToken: util.NewReferenceTokenPtr(configurationToken),
	})
	if err != nil {
		return err
	}
	config.Configuration.Analytics = enable
	config.Configuration.Multicast = nil
	resp, err := client.SetMetadataConfiguration(&wsdl.SetMetadataConfiguration{
		Configuration: config.Configuration,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
