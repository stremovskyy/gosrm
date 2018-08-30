package gosrm

import "errors"

// Errors which could be returned from OSRM Server
var RespCode = map[string]string{
	"NoRoute":        "No route found",
	"NoTable":        "No route found.",
	"NoMatch":        "No matchings found.",
	"NoTrips":        "No trips found because input coordinates are not connected.",
	"NotImplemented": "This request is not supported",
	"InvalidUrl":     "URL string is invalid",
	"Ok":             "All OK",
	"InvalidService": "Service name is invalid.",
	"InvalidVersion": "OSRM Version is not found.",
	"InvalidOptions": "Options are invalid.",
	"InvalidQuery":   "The query string is synctactically malformed.",
	"InvalidValue":   "The successfully parsed query parameters are invalid.",
	"NoSegment":      "One of the supplied input coordinates could not snap to street segment.",
	"TooBig":         "The request size violates one of the service specific request size restrictions",
}

const CodeOK = "Ok"

// Invalid request errors
var (
	ErrEmptyProfileName = errors.New("gosrm: the request should contain a profile name")
	ErrNoCoordinates    = errors.New("gosrm: the request should contain coordinates")
	ErrEmptyServiceName = errors.New("gosrm: the request should contain a service name")
)
