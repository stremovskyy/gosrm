package gosrm

import "errors"

// Errors which could be returned from OSRM Server
const (
	ErrorCodeInvalidURL     = "InvalidUrl"
	ErrorCodeInvalidService = "InvalidService"
	ErrorCodeInvalidVersion = "InvalidVersion"
	ErrorCodeInvalidOptions = "InvalidOptions"
	ErrorCodeInvalidQuery   = "InvalidQuery"
	ErrorCodeInvalidValue   = "InvalidValue"
	ErrorCodeNoSegment      = "NoSegment"
	ErrorCodeTooBig         = "TooBig"
	ErrorCodeNoRoute        = "NoRoute"
	ErrorCodeNoTable        = "NoTable"
	ErrorCodeNoMatch        = "NoMatch"
)

// Invalid request errors
var (
	ErrEmptyProfileName = errors.New("gosrm: the request should contain a profile name")
	ErrNoCoordinates    = errors.New("gosrm: the request should contain coordinates")
	ErrEmptyServiceName = errors.New("gosrm: the request should contain a service name")
)
