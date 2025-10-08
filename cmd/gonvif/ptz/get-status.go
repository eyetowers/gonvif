package ptz

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
	"github.com/eyetowers/gonvif/pkg/util"
)

var getStatus = &cobra.Command{
	Use:   "get-status",
	Short: "Get PTZ device status",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runGetStatus(client)
	},
}

func init() {
	getStatus.Flags().StringVarP(&profileToken, "profile_token", "t", "", "A reference to the MediaProfile where the operation should take place")
	root.MustMarkFlagRequired(getStatus, "profile_token")
}

func runGetStatus(client wsdl.PTZ) error {
	resp, err := client.GetStatus(&wsdl.GetStatus{
		ProfileToken: util.NewReferenceTokenPtr(profileToken),
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
