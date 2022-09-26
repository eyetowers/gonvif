// Code generated by gowsdl DO NOT EDIT.

package bf2

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"

	wsa "github.com/eyetowers/gonvif/pkg/generated/onvif/www_w3_org/2005/08/addressing"
)

// against "unused imports"
var _ time.Time
var _ xml.Name
var _ context.Context
var _ soap.SOAPEnvelope

type BaseFault BaseFaultType

type BaseFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsrf/bf-2 BaseFault" json:"-"`

	Timestamp soap.XSDDateTime `xml:"Timestamp,omitempty" json:"Timestamp,omitempty"`

	Originator *wsa.EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty"`

	ErrorCode struct {
		string

		Dialect string `xml:"dialect,attr,omitempty" json:"dialect,omitempty"`
	} `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`

	Description []struct {
		Value string `xml:",chardata" json:",omitempty"`

		EmptyString string `xml:",attr,omitempty" json:",omitempty"`
	} `xml:"Description,omitempty" json:"Description,omitempty"`

	FaultCause struct {
	} `xml:"FaultCause,omitempty" json:"FaultCause,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}
