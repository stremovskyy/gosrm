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
	"net/http"
	"time"
)

const (
	clientMaxIdleConnectionsFallBack  = 1024
	clientTLSHandshakeTimeoutFallBack = 1 * time.Second
)

// Client is an Osrm Client
type Client struct {
	httpClient *http.Client
	options    *Options
}

// NewClient creates a client with options
func NewClient(options *Options) *Client {
	if options == nil {
		panic("no client options provided")
	}

	maxIdleClients := clientMaxIdleConnectionsFallBack
	tlsHandshakeTimeout := clientTLSHandshakeTimeoutFallBack

	if options.ClientMaxIdleConnections != nil {
		maxIdleClients = *options.ClientMaxIdleConnections
	}

	if options.ClientTLSHandshakeTimeout != nil {
		tlsHandshakeTimeout = *options.ClientTLSHandshakeTimeout
	}

	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdleClients,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
	}

	c := http.Client{
		Transport: transport,
		Timeout:   time.Duration(options.RequestTimeout) * time.Second,
	}

	return &Client{&c, options}
}
