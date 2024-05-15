package ptz

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	profileToken string
)

var getPresets = &cobra.Command{
	Use:   "get-presets",
	Short: "List PTZ device presets",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetPresets(client, profileToken)
	},
}

func init() {
	getPresets.Flags().StringVarP(&profileToken, "profile_token", "t", "", "A reference to the MediaProfile where the operation should take place")
	root.MustMarkFlagRequired(getPresets, "profile_token")
}

func runGetPresets(client wsdl.PTZ, profileToken string) error {
	resp, err := client.GetPresets(&wsdl.GetPresets{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
