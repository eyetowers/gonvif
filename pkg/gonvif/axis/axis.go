package axis

import (
	"encoding/xml"

	"github.com/eyetowers/gonvif/pkg/generated/onvif/www_w3_org/2005/08/addressing"
)

type SubscriptionIDHeader struct {
	XMLName xml.Name `xml:"http://www.axis.com/2009/event SubscriptionId"`

	Value string `xml:",chardata"`
}

func ToHeader(ref *addressing.EndpointReferenceType) (SubscriptionIDHeader, bool) {
	if ref == nil || ref.ReferenceParameters == nil || len(ref.ReferenceParameters.Items) != 1 {
		return SubscriptionIDHeader{}, false
	}
	return SubscriptionIDHeader{
		Value: ref.ReferenceParameters.Items[0],
	}, true
}
