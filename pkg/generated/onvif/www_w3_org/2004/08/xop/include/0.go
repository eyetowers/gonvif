// Code generated by gowsdl DO NOT EDIT.

package include

import (
	"context"
	"encoding/xml"
	"github.com/eyetowers/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name
var _ context.Context
var _ soap.SOAPEnvelope

type Include struct {
	Items []string `xml:",any" json:"items,omitempty"`

	Href string `xml:"href,attr,omitempty" json:"href,omitempty"`
}
