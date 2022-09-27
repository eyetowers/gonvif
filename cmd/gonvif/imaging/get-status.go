package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getStatus = &cobra.Command{
	Use:   "get-status",
	Short: "Show Onvif device imaging status",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return nil
		}
		return runGetStatus(client, videoSourceToken)
	},
}

func init() {
	getStatus.Flags().StringVarP(&videoSourceToken, "video_source_token", "t", "", "Video source token")
	root.MustMarkFlagRequired(getStatus, "video_source_token")
}

func runGetStatus(client wsdl.ImagingPort, videoSourceToken string) error {
	resp, err := client.GetStatus(&wsdl.GetStatus{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
