package util

import "github.com/eyetowers/gonvif/pkg/generated/onvif/www_onvif_org/ver10/schema"

func NewReferenceTokenPtr(token string) *schema.ReferenceToken {
	if token == "" {
		return nil
	}
	result := schema.ReferenceToken(token)
	return &result
}
