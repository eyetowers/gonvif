package ptz

import (
	"encoding/json"
	"os"

	"github.com/hooklift/gowsdl/soap"
	"github.com/spf13/cobra"

	"github.com/eltrac-eu/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var getConfigurations = &cobra.Command{
	Use:   "get-configurations",
	Short: "List PTZ device configurations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGetConfigurations("TODO", "TODO", "TODO")
	},
}

func runGetConfigurations(url, username, password string) error {
	cl := soap.NewClient(url)
	cl.SetHeaders(soap.NewSecurity(username, password))
	p := wsdl.NewPTZ(cl)

	resp, err := p.GetConfigurations(&wsdl.GetConfigurations{})
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
