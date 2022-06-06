package root

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/hooklift/gowsdl/soap"
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
)

func init() {
	Command.PersistentFlags().StringVarP(&URL, "url", "a", "", "Base URL of the Onvif device.")
	Command.PersistentFlags().StringVarP(&Username, "username", "u", "", "Username for authentication with the Onvif device.")
	Command.PersistentFlags().StringVarP(&Password, "password", "p", "", "Password for authentication with the Onvif device.")
	if err := Command.MarkPersistentFlagRequired("url"); err != nil {
		panic(err)
	}
	if err := Command.MarkPersistentFlagRequired("username"); err != nil {
		panic(err)
	}
	if err := Command.MarkPersistentFlagRequired("password"); err != nil {
		panic(err)
	}
}

func ServiceURL(baseURL, suffix string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("malformed base URL: %w", err)
	}
	u, err := url.Parse(suffix)
	if err != nil {
		return "", fmt.Errorf("malformed service suffix URL: %w", err)
	}
	return base.ResolveReference(u).String(), nil
}

func AuthorizedSOAPClient(serviceURL, username, password string) *soap.Client {
	client := soap.NewClient(serviceURL)
	client.SetHeaders(soap.NewSecurity(username, password))
	return client
}

func OutputJSON(payload interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	return encoder.Encode(payload)
}
