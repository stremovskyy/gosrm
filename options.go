package gosrm

import (
	"net/url"
	"time"
)

// Options GOSRM client options structure
type Options struct {
	// Url of osrm server with / on the end
	Url url.URL

	// One of the following values:  route ,  nearest ,  table ,  match ,  trip ,  tile
	Service string

	//	Version of the protocol implemented by the service
	Version string

	// Mode of transportation, is determined statically by the Lua profile that is used to prepare the data using  osrm-extract
	Profile string

	// Timeout for request in seconds
	RequestTimeout int

	// if need debug you should turn this on
	Debug bool

	// if true http client uses AcceptEncoding gzip
	UseGzip bool

	ClientMaxIdleConnections  *int
	ClientTLSHandshakeTimeout *time.Duration
}
