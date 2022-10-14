package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getSnapshotURI = &cobra.Command{
	Use:   "get-snapshot-uri",
	Short: "Get Onvif device media snapshot URI",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetSnapshotURI(client, profileToken)
	},
}

func init() {
	getSnapshotURI.Flags().StringVarP(&profileToken, "profile_token", "t", "", "The ProfileToken element indicates the media profile to use and will define the source and dimensions of the snapshot.")
	root.MustMarkFlagRequired(getSnapshotURI, "profile_token")
}

func runGetSnapshotURI(client wsdl.Media2, profileToken string) error {
	resp, err := client.GetSnapshotUri(&wsdl.GetSnapshotUri{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
