package imaging

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	speed float32
)

var moveContinuous = &cobra.Command{
	Use:   "move-continuous",
	Short: "Start Onvif device continuous focusing",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runMoveContinuous(client, videoSourceToken, speed)
	},
}

func init() {
	moveContinuous.Flags().StringVarP(&videoSourceToken, "video_source_token", "t", "", "Video source token")
	root.MustMarkFlagRequired(moveContinuous, "video_source_token")
	moveContinuous.Flags().Float32VarP(&speed, "speed", "s", 0, "Continuous focus speed")
	root.MustMarkFlagRequired(moveContinuous, "speed")
}

func runMoveContinuous(client wsdl.ImagingPort, videoSourceToken string, speed float32) error {
	resp, err := client.Move(&wsdl.Move{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
		Focus: &schema.FocusMove{
			Continuous: &schema.ContinuousFocus{
				Speed: speed,
			},
		},
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
