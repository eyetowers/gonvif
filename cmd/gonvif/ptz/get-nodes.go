package ptz

import (
	"github.com/spf13/cobra"

	"github.com/eltrac-eu/gonvif/cmd/gonvif/root"
	"github.com/eltrac-eu/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var getNodes = &cobra.Command{
	Use:   "get-nodes",
	Short: "List PTZ device nodes",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password)
		if err != nil {
			return nil
		}
		return runGetNodes(client)
	},
}

func runGetNodes(client wsdl.PTZ) error {
	resp, err := client.GetNodes(&wsdl.GetNodes{})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
