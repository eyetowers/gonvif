package device

import (
	"github.com/spf13/cobra"

	"github.com/eyetowers/gonvif/cmd/gonvif/root"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"
)

var (
	hard bool
)

var setSystemFactoryDefault = &cobra.Command{
	Use:   "set-system-factory-default",
	Short: "Set Onvif device system factory default",
	Args:  cobra.NoArgs,
	RunE: func(*cobra.Command, []string) error {
		client, err := ServiceClient(root.URL, root.Username, root.Password, root.Verbose)
		if err != nil {
			return err
		}
		return runSetSystemFactoryDefault(client, hard)
	},
}

func init() {
	setSystemFactoryDefault.Flags().BoolVarP(&hard, "hard", "", false,
		"Require a hard factory reset which will also reset camera network settings and thus its reachability.")
}

func runSetSystemFactoryDefault(client wsdl.Device, hard bool) error {
	reset := schema.FactoryDefaultTypeSoft
	if hard {
		reset = schema.FactoryDefaultTypeHard
	}

	resp, err := client.SetSystemFactoryDefault(&wsdl.SetSystemFactoryDefault{
		FactoryDefault: &reset,
	})
	if err != nil {
		return err
	}
	return root.OutputJSON(resp)
}
