package gosrm

import "net/url"

// OSRM client options structure
type Options struct {
	// Url of osrm server with / on the end
	Url url.URL `json:"url"`

	// One of the following values:  route ,  nearest ,  table ,  match ,  trip ,  tile
	Service string `json:"service"`

	//	Version of the protocol implemented by the service
	Version string `json:"version"`

	// Mode of transportation, is determined statically by the Lua profile that is used to prepare the data using  osrm-extract
	Profile string `json:"profile"`

	// Timeout for request in seconds
	RequestTimeout int `json:"request_timeout"`
}
