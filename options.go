package gosrm

import (
	"log"
	"net/url"
	"time"
)

// Options defines the configuration for the GOSRM client.
type Options struct {
	// Url represents the URL of the OSRM server. It should end with a slash (/).
	BaseURL url.URL

	// Version represents the version of the protocol implemented by the service.
	Version string

	// RequestTimeout represents the timeout for the request in seconds.
	RequestTimeout int

	// Debug, when set to true, enables debug mode.
	Debug bool

	// UseGzip, when set to true, enables the http client to use AcceptEncoding gzip.
	UseGzip bool

	// ClientMaxIdleConnections represents the maximum number of idle connections that the client can have.
	ClientMaxIdleConnections *int

	// ClientTLSHandshakeTimeout represents the timeout for the TLS handshake.
	ClientTLSHandshakeTimeout *time.Duration

	// Headers represents the headers to be sent with the request.
	Headers map[string]string

	// Logger represents the logger to be used by the client.
	Logger *log.Logger
}

// NewDefaultOptions creates a default configuration for the GOSRM client.
func NewDefaultOptions() *Options {
	return &Options{
		Version:                   VersionFirst,                  // Default protocol version
		RequestTimeout:            30,                            // Default timeout in seconds
		Debug:                     false,                         // Debug mode off by default
		UseGzip:                   true,                          // Gzip compression on by default
		ClientMaxIdleConnections:  intPtr(1024),                  // Default max idle connections
		ClientTLSHandshakeTimeout: durationPtr(10 * time.Second), // Default TLS handshake timeout
		BaseURL:                   url.URL{Scheme: "https", Host: "router.project-osrm.org"},
		Headers:                   make(map[string]string),
		Logger:                    log.New(log.Writer(), "GOSRM: ", log.Flags()),
	}
}

// NewCustomOptions allows customizing certain key aspects of the Options.
func NewCustomOptions(version string, requestTimeout int, debug, useGzip bool) *Options {
	opts := NewDefaultOptions()
	opts.Version = version
	opts.RequestTimeout = requestTimeout
	opts.Debug = debug
	opts.UseGzip = useGzip

	return opts
}

// intPtr is a helper function to create a pointer to an int.
func intPtr(i int) *int {
	return &i
}

// durationPtr is a helper function to create a pointer to a time.Duration.
func durationPtr(d time.Duration) *time.Duration {
	return &d
}
