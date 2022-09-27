package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getMoveOptions = &cobra.Command{
	Use:   "get-move-options",
	Short: "Show Onvif device focus move options",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return nil
		}
		return runGetMoveOptions(client, videoSourceToken)
	},
}

func init() {
	getMoveOptions.Flags().StringVarP(&videoSourceToken, "video_source_token", "t", "", "Video source token")
	root.MustMarkFlagRequired(getMoveOptions, "video_source_token")
}

func runGetMoveOptions(client wsdl.ImagingPort, videoSourceToken string) error {
	resp, err := client.GetMoveOptions(&wsdl.GetMoveOptions{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
