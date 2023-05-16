package gonvif

import (
	"encoding/xml"

	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_w3_org/2005/08/addressing"
	"github.com/eyetowers/gonvif/pkg/gonvif/axis"
)

func ComposeHeaders(ref *addressing.EndpointReferenceType) []any {
	// NOTE: Axis returns a SubscriptionId that has to be passed into the PullMessages call as SOAP
	// header as-is. We try to mimic this behavior here.
	var headers []any
	if axisHeader, ok := axis.ToHeader(ref); ok {
		headers = append(headers, axisHeader)
	}
	headers = append(headers,
		ReplyToAddressingHeader{Address: "http://www.w3.org/2005/08/addressing/anonymous"},
		ToAddressingHeader{Value: string(*ref.Address)},
	)
	return headers
}

type ToAddressingHeader struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing To"`

	Value string `xml:",chardata"`
}

type ReplyToAddressingHeader struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ReplyTo"`

	Address string `xml:"http://www.w3.org/2005/08/addressing Address"`
}
