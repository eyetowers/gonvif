package media

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	autoFocus string
)

var setImagingSettingsAutoFocus = &cobra.Command{
	Use:   "set-imaging-settings-auto-focus",
	Short: "Set Onvif device imaging settings autofocus",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runSetImagingSettingsAutoFocus(client, videoSourceToken, autoFocus)
	},
}

func init() {
	setImagingSettingsAutoFocus.Flags().StringVarP(&videoSourceToken, "video_source_token", "t", "", "Video source token")
	root.MustMarkFlagRequired(setImagingSettingsAutoFocus, "video_source_token")
	setImagingSettingsAutoFocus.Flags().StringVarP(&autoFocus, "auto_focus", "f", "AUTO", "Auto focus mode")
}

func runSetImagingSettingsAutoFocus(client wsdl.ImagingPort, videoSourceToken string, autoFocus string) error {
	// First, read the current settings.
	original, err := client.GetImagingSettings(&wsdl.GetImagingSettings{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
	})
	if err != nil {
		return err
	}

	req := &wsdl.SetImagingSettings{
		VideoSourceToken: util.NewReferenceTokenPtr(videoSourceToken),
		ImagingSettings:  original.ImagingSettings,
	}
	af := schema.AutoFocusMode(autoFocus)
	req.ImagingSettings.Focus.AutoFocusMode = &af

	resp, err := client.SetImagingSettings(req)
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
