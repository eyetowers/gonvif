package ptz

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	tt "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var (
	pan          float32
	tilt         float32
	zoom         float32
	panTiltSpeed float32
	zoomSpeed    float32

	hasPan          bool
	hasTilt         bool
	hasZoom         bool
	hasPanTiltSpeed bool
	hasZoomSpeed    bool

	panTiltSpace string
	zoomSpace    string
)

var absoluteMove = &cobra.Command{
	Use:     "absolute-move",
	Short:   "Move PTZ device to an absolute position",
	Args:    cobra.NoArgs,
	PreRunE: validateAbsoluteMoveFlags,
	RunE: func(cmd *cobra.Command, _ []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runAbsoluteMove(client, profileToken, newAbsoluteMovePosition(), newAbsoluteMoveSpeed())
	},
}

func init() {
	absoluteMove.Flags().StringVarP(&profileToken, "profile_token", "t", "", "A reference to the MediaProfile where the operation should take place")
	root.MustMarkFlagRequired(absoluteMove, "profile_token")

	absoluteMove.Flags().Float32Var(&pan, "pan", 0, "Absolute pan coordinate")
	absoluteMove.Flags().Float32Var(&tilt, "tilt", 0, "Absolute tilt coordinate")
	absoluteMove.Flags().Float32Var(&zoom, "zoom", 0, "Absolute zoom coordinate")

	absoluteMove.Flags().Float32Var(&panTiltSpeed, "pan_tilt_speed", 0, "Pan and tilt move speed")
	absoluteMove.Flags().Float32Var(&zoomSpeed, "zoom_speed", 0, "Zoom move speed")

	absoluteMove.Flags().StringVar(&panTiltSpace, "pan_tilt_space", "", "Pan/tilt coordinate space URI")
	absoluteMove.Flags().StringVar(&zoomSpace, "zoom_space", "", "Zoom coordinate space URI")
}

func validateAbsoluteMoveFlags(cmd *cobra.Command, _ []string) error {
	hasPan = cmd.Flags().Changed("pan")
	hasTilt = cmd.Flags().Changed("tilt")
	hasZoom = cmd.Flags().Changed("zoom")
	hasPanTiltSpeed = cmd.Flags().Changed("pan_tilt_speed")
	hasZoomSpeed = cmd.Flags().Changed("zoom_speed")

	if hasPan != hasTilt {
		return errors.New("pan and tilt must be provided together")
	}
	if !hasPan && !hasZoom {
		return errors.New("at least pan and tilt, or zoom, must be provided")
	}
	return nil
}

func runAbsoluteMove(client wsdl.PTZ, profileToken string, position *tt.PTZVector, speed *tt.PTZSpeed) error {
	resp, err := client.AbsoluteMove(&wsdl.AbsoluteMove{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
		Position:     position,
		Speed:        speed,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}

func newAbsoluteMovePosition() *tt.PTZVector {
	position := &tt.PTZVector{}
	if hasPan {
		position.PanTilt = &tt.Vector2D{
			X:     pan,
			Y:     tilt,
			Space: panTiltSpace,
		}
	}
	if hasZoom {
		position.Zoom = &tt.Vector1D{
			X:     zoom,
			Space: zoomSpace,
		}
	}
	return position
}

func newAbsoluteMoveSpeed() *tt.PTZSpeed {
	if !hasPanTiltSpeed && !hasZoomSpeed {
		return nil
	}
	speed := &tt.PTZSpeed{}
	if hasPanTiltSpeed {
		speed.PanTilt = &tt.Vector2D{
			X: panTiltSpeed,
			Y: panTiltSpeed,
		}
	}
	if hasZoomSpeed {
		speed.Zoom = &tt.Vector1D{
			X: zoomSpeed,
		}
	}
	return speed
}
