package imaging

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	videoSourceToken string
)

var getImagingSettings = &cobra.Command{
	Use:   "get-imaging-settings",
	Short: "Show Onvif device imaging settings",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetImagingSettings(client, videoSourceToken)
	},
}

func init() {
	getImagingSettings.Flags().StringVarP(&videoSourceToken, "video_source_token", "t", "", "Video source token")
	root.MustMarkFlagRequired(getImagingSettings, "video_source_token")
}

func runGetImagingSettings(client wsdl.ImagingPort, videoSourceToken string) error {
	resp, err := client.GetImagingSettings(&wsdl.GetImagingSettings{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
