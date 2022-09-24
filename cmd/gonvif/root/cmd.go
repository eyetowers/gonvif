package root

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/hooklift/gowsdl/soap"
	"github.com/motemen/go-loghttp"
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
	cmd.MarkPersistentFlagRequired("url")
	cmd.MarkPersistentFlagRequired("username")
	cmd.MarkPersistentFlagRequired("password")
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

func AuthorizedSOAPClient(serviceURL, username, password string, verbose bool) *soap.Client {
	httpClient := http.DefaultClient
	if verbose {
		httpClient = &http.Client{
			Transport: &loghttp.Transport{
				LogResponse: logResponse,
				LogRequest:  logRequest,
			},
		}
	}
	client := soap.NewClient(serviceURL, soap.WithHTTPClient(httpClient))
	client.SetHeaders(soap.NewSecurity(username, password))
	return client
}

func OutputJSON(payload interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	return encoder.Encode(payload)
}

func logResponse(resp *http.Response) {
	log.Printf("<-- %d %s", resp.StatusCode, resp.Request.URL)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("BODY:\n%s", string(body))
	resp.Body = io.NopCloser(bytes.NewReader(body))
}

func logRequest(req *http.Request) {
	log.Printf("--> %s %s", req.Method, req.URL)
	defer req.Body.Close()
	body, _ := io.ReadAll(req.Body)
	log.Printf("BODY:\n%s", string(body))
	req.Body = io.NopCloser(bytes.NewReader(body))
}
