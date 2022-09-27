package gonvif

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/eyetowers/gowsdl/soap"
	"github.com/motemen/go-loghttp"

	device "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
	analytics "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/analytics/wsdl"
	imaging "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/imaging/wsdl"
	media2 "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/media/wsdl"
	ptz "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver20/ptz/wsdl"
)

var (
	ErrServiceNotSupported = errors.New("onvif service not supported")

	verboseHTTPClient = &http.Client{
		Transport: &loghttp.Transport{
			LogResponse: logResponse,
			LogRequest:  logRequest,
		},
	}
)

type Client interface {
	Analytics() (analytics.AnalyticsEnginePort, error)
	Device() (device.Device, error)
	Imaging() (imaging.ImagingPort, error)
	Media2() (media2.Media2, error)
	PTZ() (ptz.PTZ, error)
}

type impl struct {
	analytics analytics.AnalyticsEnginePort
	device    device.Device
	imaging   imaging.ImagingPort
	media2    media2.Media2
	ptz       ptz.PTZ
}

func New(baseURL, username, password string, verbose bool) (Client, error) {
	soapClient, err := serviceSOAPClient(baseURL, "onvif/device_service", username, password, verbose)
	if err != nil {
		return nil, err
	}
	d := device.NewDevice(soapClient)
	resp, err := d.GetServices(&device.GetServices{})
	if err != nil {
		return nil, fmt.Errorf("listing available Onvif services: %w", err)
	}

	var result impl
	for _, svc := range resp.Service {
		svcClient, err := serviceSOAPClient(baseURL, svc.XAddr, username, password, verbose)
		if err != nil {
			return nil, err
		}
		if svc.Namespace == "http://www.onvif.org/ver20/analytics/wsdl" {
			result.analytics = analytics.NewAnalyticsEnginePort(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver10/device/wsdl" {
			result.device = device.NewDevice(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver20/imaging/wsdl" {
			result.imaging = imaging.NewImagingPort(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver20/media/wsdl" {
			result.media2 = media2.NewMedia2(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver20/ptz/wsdl" {
			result.ptz = ptz.NewPTZ(svcClient)
		}
	}

	return &result, nil
}

func (c *impl) Analytics() (analytics.AnalyticsEnginePort, error) {
	if c.analytics == nil {
		return nil, ErrServiceNotSupported
	}
	return c.analytics, nil
}

func (c *impl) Device() (device.Device, error) {
	if c.device == nil {
		return nil, ErrServiceNotSupported
	}
	return c.device, nil
}

func (c *impl) Imaging() (imaging.ImagingPort, error) {
	if c.imaging == nil {
		return nil, ErrServiceNotSupported
	}
	return c.imaging, nil
}

func (c *impl) Media2() (media2.Media2, error) {
	if c.media2 == nil {
		return nil, ErrServiceNotSupported
	}
	return c.media2, nil
}

func (c *impl) PTZ() (ptz.PTZ, error) {
	if c.ptz == nil {
		return nil, ErrServiceNotSupported
	}
	return c.ptz, nil
}

func serviceURL(baseURL, suffix string) (string, error) {
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

func sanitizeServiceURL(baseURL, advertisedURL string) (string, error) {
	u, err := url.Parse(advertisedURL)
	if err != nil {
		return "", fmt.Errorf("malformed service advertised URL: %w", err)
	}
	return serviceURL(baseURL, u.Path)
}

func serviceSOAPClient(baseURL, advertisedURL, username, password string, verbose bool) (*soap.Client, error) {
	u, err := sanitizeServiceURL(baseURL, advertisedURL)
	if err != nil {
		return nil, err
	}
	return AuthorizedSOAPClient(u, username, password, verbose), nil
}

func AuthorizedSOAPClient(serviceURL, username, password string, verbose bool) *soap.Client {
	httpClient := http.DefaultClient
	if verbose {
		httpClient = verboseHTTPClient
	}
	client := soap.NewClient(serviceURL, soap.WithHTTPClient(httpClient))
	client.SetHeaders(soap.NewSecurity(username, password))
	return client
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
