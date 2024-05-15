package ptz

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var getNodes = &cobra.Command{
	Use:   "get-nodes",
	Short: "List PTZ device nodes",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
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
