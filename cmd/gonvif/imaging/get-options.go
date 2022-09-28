package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getOptions = &cobra.Command{
	Use:   "get-options",
	Short: "Show Onvif device imaging options",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetOptions(client, videoSourceToken)
	},
}

func init() {
	getOptions.Flags().StringVarP(&videoSourceToken, "video_source_token", "t", "", "Video source token")
	root.MustMarkFlagRequired(getOptions, "video_source_token")
}

func runGetOptions(client wsdl.ImagingPort, videoSourceToken string) error {
	resp, err := client.GetOptions(&wsdl.GetOptions{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
