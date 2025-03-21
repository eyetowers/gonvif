package gonvif

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/eyetowers/gowsdl/soap"
	"github.com/motemen/go-loghttp"

	device "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/device/wsdl"
	events "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/events/wsdl"
	media "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/media/wsdl"
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
	Events() (events.EventPortType, error)
	Subscription(url string, headers ...any) (events.PullPointSubscription, error)
	Imaging() (imaging.ImagingPort, error)
	Media() (media.Media, error)
	Media2() (media2.Media2, error)
	PTZ() (ptz.PTZ, error)
}

type impl struct {
	baseURL  string
	username string
	password string
	verbose  bool
	diff     time.Duration

	analytics analytics.AnalyticsEnginePort
	device    device.Device
	events    events.EventPortType
	imaging   imaging.ImagingPort
	media     media.Media
	media2    media2.Media2
	ptz       ptz.PTZ
}

func New(ctx context.Context, baseURL, username, password string, verbose bool) (Client, error) {
	soapClient, err := serviceSOAPClient(baseURL, "onvif/device_service", username, password, verbose, 0)
	if err != nil {
		return nil, err
	}
	d := device.NewDevice(soapClient)
	diff, err := getTimeDiff(ctx, d)
	if err != nil {
		return nil, err
	}
	soapClient.SetTimeDiff(diff)
	resp, err := d.GetServicesContext(ctx, &device.GetServices{})
	if err != nil {
		return nil, fmt.Errorf("listing available Onvif services: %w", err)
	}

	result := impl{
		baseURL:  baseURL,
		username: username,
		password: password,
		verbose:  verbose,
		diff:     diff,
	}

	for _, svc := range resp.Service {
		svcClient, err := serviceSOAPClient(baseURL, svc.XAddr, username, password, verbose, diff)
		if err != nil {
			return nil, err
		}
		if svc.Namespace == "http://www.onvif.org/ver20/analytics/wsdl" {
			result.analytics = analytics.NewAnalyticsEnginePort(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver10/device/wsdl" {
			result.device = device.NewDevice(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver10/events/wsdl" {
			result.events = events.NewEventPortType(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver20/imaging/wsdl" {
			result.imaging = imaging.NewImagingPort(svcClient)
		}
		if svc.Namespace == "http://www.onvif.org/ver10/media/wsdl" {
			result.media = media.NewMedia(svcClient)
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

func getTimeDiff(ctx context.Context, d device.Device) (time.Duration, error) {
	resp, err := d.GetSystemDateAndTimeContext(ctx, &device.GetSystemDateAndTime{})
	if err != nil {
		return 0, fmt.Errorf("getting Onvif device time: %w", err)
	}
	if resp.SystemDateAndTime == nil ||
		resp.SystemDateAndTime.UTCDateTime == nil ||
		resp.SystemDateAndTime.UTCDateTime.Date == nil ||
		resp.SystemDateAndTime.UTCDateTime.Time == nil {
		return 0, fmt.Errorf("no time returned by Onvif device")
	}
	utc := resp.SystemDateAndTime.UTCDateTime
	deviceTime := time.Date(
		int(utc.Date.Year),
		time.Month(utc.Date.Month),
		int(utc.Date.Day),
		int(utc.Time.Hour),
		int(utc.Time.Minute),
		int(utc.Time.Second),
		0,
		time.UTC,
	)
	return time.Until(deviceTime), nil
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

func (c *impl) Events() (events.EventPortType, error) {
	if c.events == nil {
		return nil, ErrServiceNotSupported
	}
	return c.events, nil
}

func (c *impl) Subscription(url string, headers ...any) (events.PullPointSubscription, error) {
	if c.events == nil {
		return nil, ErrServiceNotSupported
	}
	client, err := serviceSOAPClient(c.baseURL, url, c.username, c.password, c.verbose, c.diff)
	if err != nil {
		return nil, err
	}
	client.SetHeaders(headers)

	return events.NewPullPointSubscription(client), nil
}

func (c *impl) Imaging() (imaging.ImagingPort, error) {
	if c.imaging == nil {
		return nil, ErrServiceNotSupported
	}
	return c.imaging, nil
}

func (c *impl) Media() (media.Media, error) {
	if c.media == nil {
		return nil, ErrServiceNotSupported
	}
	return c.media, nil
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
	return serviceURL(baseURL, u.RequestURI())
}

func serviceSOAPClient(
	baseURL, advertisedURL, username, password string, verbose bool, diff time.Duration,
) (*soap.Client, error) {
	u, err := sanitizeServiceURL(baseURL, advertisedURL)
	if err != nil {
		return nil, err
	}
	return AuthorizedSOAPClient(u, username, password, verbose, diff), nil
}

func AuthorizedSOAPClient(
	serviceURL, username, password string, verbose bool, diff time.Duration,
) *soap.Client {
	httpClient := http.DefaultClient
	if verbose {
		httpClient = verboseHTTPClient
	}
	client := soap.NewClient(serviceURL,
		soap.WithHTTPClient(httpClient),
		soap.WithSOAPAuth(username, password),
		soap.WithTimeDiff(diff),
	)
	return client
}

func logResponse(resp *http.Response) {
	log.Printf("<-- %d %s", resp.StatusCode, resp.Request.URL)
	resp.Body = logBodyAndHeaders(resp.Body, resp.Header)
}

func logRequest(req *http.Request) {
	log.Printf("--> %s %s", req.Method, req.URL)
	req.Body = logBodyAndHeaders(req.Body, req.Header)
}

func logBodyAndHeaders(body io.ReadCloser, headers http.Header) io.ReadCloser {
	defer body.Close()
	bytes, dupe := readAndDuplicate(body)
	log.Printf("HEADERS:\n")
	for k, vals := range headers {
		for _, val := range vals {
			log.Printf("%s: %s\n", k, val)
		}
	}
	log.Printf("BODY:\n%s", string(bytes))
	return dupe
}

func readAndDuplicate(body io.ReadCloser) ([]byte, io.ReadCloser) {
	defer body.Close()
	buf, _ := io.ReadAll(body)

	return buf, io.NopCloser(bytes.NewReader(buf))
}
