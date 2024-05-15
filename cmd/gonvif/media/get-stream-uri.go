package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	protocol string
	stream   string
)

var getStreamURI = &cobra.Command{
	Use:   "get-stream-uri",
	Short: "Get Onvif device media stream URI",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetStreamURI(client, profileToken, protocol, stream)
	},
}

func init() {
	getStreamURI.Flags().StringVarP(&profileToken, "profile_token", "t", "", "The ProfileToken element indicates the media profile to use and will define the configuration of the content of the stream.")
	getStreamURI.Flags().StringVarP(&protocol, "protocol", "r", "RTSP", "Defines the network protocol for streaming (UDP, TCP, RTSP, HTTP).")
	getStreamURI.Flags().StringVarP(&stream, "stream", "s", "RTP-Unicast", "Defines the stream type for streaming (RTP-Unicast, RTP-Multicast).")
	root.MustMarkFlagRequired(getStreamURI, "profile_token")
}

func runGetStreamURI(client wsdl.Media, profileToken string, protocol string, stream string) error {
	p := schema.TransportProtocol(protocol)
	s := schema.StreamType(stream)
	resp, err := client.GetStreamUri(&wsdl.GetStreamUri{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
		StreamSetup: &schema.StreamSetup{
			Stream: &s,
			Transport: &schema.Transport{
				Protocol: &p,
			},
		},
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
