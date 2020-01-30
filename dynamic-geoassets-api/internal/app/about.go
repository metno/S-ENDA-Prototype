package app

import (
	"net/url"

	"github.com/metno/S-ENDA-Prototype/dynamic-geoassets-api/pkg/metaservice"
)

func about() *metaservice.About {
	return &metaservice.About{
		Name:           "Dynamic geo assets API",
		Description:    "The purpose of this service is to list, create and handle datasets and services for geodynamic assets.",
		Responsible:    "S-ENDA sprint team <s-enda-sprint@met.no>",
		Documentation:  &url.URL{Path: "/"},
		TermsOfService: &url.URL{Path: "/docs/termsofservice"},
	}
}
