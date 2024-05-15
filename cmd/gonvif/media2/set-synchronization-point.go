package media2

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var setSynchronizationPoint = &cobra.Command{
	Use:   "set-synchronization-point",
	Short: "Request Onvif device to insert a key frame into all streams for the given profile",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runSetSynchronizationPoint(client, profileToken)
	},
}

func init() {
	setSynchronizationPoint.Flags().StringVarP(&profileToken, "profile_token", "t", "", "The ProfileToken element indicates the media profile to use and will define the source and dimensions of the snapshot.")
	root.MustMarkFlagRequired(setSynchronizationPoint, "profile_token")
}

func runSetSynchronizationPoint(client wsdl.Media2, profileToken string) error {
	resp, err := client.SetSynchronizationPoint(&wsdl.SetSynchronizationPoint{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
