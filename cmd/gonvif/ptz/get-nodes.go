package ptz

import (
	"encoding/json"
	"os"

	"github.com/hooklift/gowsdl/soap"
	"github.com/spf13/cobra"

	"github.com/eltrac-eu/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var getNodes = &cobra.Command{
	Use:   "get-nodes",
	Short: "List PTZ device nodes",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGetConfigurations("TODO", "TODO", "TODO")
	},
}

func runGetNodes(url, username, password string) error {
	cl := soap.NewClient(url)
	cl.SetHeaders(soap.NewSecurity(username, password))
	p := wsdl.NewPTZ(cl)

	resp, err := p.GetNodes(&wsdl.GetNodes{})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(resp)
	if err != nil {
		return err
	}
	return nil
}
