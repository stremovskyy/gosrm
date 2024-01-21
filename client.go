/*
 * Copyright 2018 Anthony Stremovskyy <stremovskyy@me.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software
 * and associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
 * OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR
 * IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package gosrm

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultMaxIdleConnections  = 1024
	defaultTLSHandshakeTimeout = 1 * time.Second
)

// Client represents an OSRM HTTP client.
type osrmClient struct {
	httpClient *http.Client
	baseURL    string
	headers    http.Header
	logger     *log.Logger
	options    *Options

	profile *string
}

func (c *osrmClient) SetOptions(options *Options) {
	c.options = options
}

func (c *osrmClient) ForProfile(profile string) OsrmClient {
	newClient := *c

	newClient.profile = &profile

	return &newClient
}

func (c *osrmClient) Debug() OsrmClient {
	newClient := *c

	newClient.logger = log.New(log.Writer(), "GOSRM DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	newClient.options.Debug = true

	return &newClient
}

// NewClient creates a new OSRM Client with the provided options.
// If options is nil, default values are used.
func NewClient(options *Options) OsrmClient {
	if options == nil {
		options = NewDefaultOptions()
		options.Logger.Println("Using default options")
	}

	maxIdleConnections := defaultMaxIdleConnections
	if options.ClientMaxIdleConnections != nil {
		maxIdleConnections = *options.ClientMaxIdleConnections
	}

	tlsHandshakeTimeout := defaultTLSHandshakeTimeout
	if options.ClientTLSHandshakeTimeout != nil {
		tlsHandshakeTimeout = *options.ClientTLSHandshakeTimeout
	}

	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdleConnections,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
	}

	client := http.Client{
		Transport: transport,
		Timeout:   time.Duration(options.RequestTimeout) * time.Second,
	}

	headers := make(http.Header)
	for key, value := range options.Headers {
		headers.Set(key, value)
	}

	if options.Debug {
		log.SetOutput(log.Writer())
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	}

	return &osrmClient{
		httpClient: &client,
		baseURL:    options.BaseURL.String(),
		headers:    headers,
		logger:     options.Logger,
		options:    options,
	}
}

func NewDefaultClient() OsrmClient {
	return NewClient(NewDefaultOptions())
}

func (c *osrmClient) baseUrlForService(service string) *url.URL {
	u := c.options.BaseURL

	u.Path += service + "/" + VersionFirst

	if c.profile != nil {
		u.Path += "/" + *c.profile
	} else {
		u.Path += "/" + ProfileCar
	}

	return &u
}
