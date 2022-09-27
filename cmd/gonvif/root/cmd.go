package root

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "gonvif",
	Short: "Onvif CLI.",
	Long:  "Onvif CLI focused on PTZ cameras and implemented in pure Go.",
}

var (
	URL      string
	Username string
	Password string
	Verbose  bool
)

func init() {
	Command.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Print sent and received requests.")
}

func RequireAuthFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&URL, "url", "a", "", "Base URL of the Onvif device.")
	cmd.PersistentFlags().StringVarP(&Username, "username", "u", "", "Username for authentication with the Onvif device.")
	cmd.PersistentFlags().StringVarP(&Password, "password", "p", "", "Password for authentication with the Onvif device.")
	MustMarkPersistentFlagRequired(cmd, "url")
	MustMarkPersistentFlagRequired(cmd, "username")
	MustMarkPersistentFlagRequired(cmd, "password")
}

func MustMarkFlagRequired(cmd *cobra.Command, flag string) {
	err := cmd.MarkFlagRequired(flag)
	if err != nil {
		panic(err)
	}
}

func MustMarkPersistentFlagRequired(cmd *cobra.Command, flag string) {
	err := cmd.MarkPersistentFlagRequired(flag)
	if err != nil {
		panic(err)
	}
}

func OutputJSON(payload interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	return encoder.Encode(payload)
}
