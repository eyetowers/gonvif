package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	protocol string
)

var getStreamURI = &cobra.Command{
	Use:   "get-stream-uri",
	Short: "Get Onvif device media stream URI",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return nil
		}
		return runGetStreamURI(client, profileToken, protocol)
	},
}

func init() {
	getStreamURI.Flags().StringVarP(&profileToken, "profile_token", "t", "", "The ProfileToken element indicates the media profile to use and will define the configuration of the content of the stream.")
	getStreamURI.Flags().StringVarP(&protocol, "protocol", "r", "RTSP", "Defines the network protocol for streaming (RtspUnicast, RtspMulticast, RTSP, RtspOverHttp).")
	root.MustMarkFlagRequired(getStreamURI, "profile_token")
}

func runGetStreamURI(client wsdl.Media2, profileToken string, protocol string) error {
	resp, err := client.GetStreamUri(&wsdl.GetStreamUri{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
		Protocol:     protocol,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
