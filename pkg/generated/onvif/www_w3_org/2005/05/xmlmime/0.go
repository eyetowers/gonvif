// Code generated by gowsdl DO NOT EDIT.

package xmlmime

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name
var _ context.Context
var _ soap.SOAPEnvelope

type Base64Binary struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/05/xmlmime Data" json:"-"`

	Value []byte `xml:",chardata" json:",omitempty"`

	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
}

type HexBinary struct {
	Value []byte `xml:",chardata" json:",omitempty"`

	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
}
