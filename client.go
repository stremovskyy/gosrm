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
	"encoding/json"
	"errors"
	"github.com/paulmach/go.geo"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Osrm Client Object
type OsrmClient struct {
	httpClient *http.Client

	Options *Options
}

// NewClient creates a client with options
func NewClient(options *Options) *OsrmClient {
	timeout := time.Duration(time.Duration(options.RequestTimeout) * time.Second)

	t := &http.Transport{
		MaxIdleConnsPerHost: 1024,
		TLSHandshakeTimeout: 0 * time.Second,
	}

	c := http.Client{
		Transport: t,
		Timeout:   timeout,
	}

	return &OsrmClient{&c, options}
}

func (c OsrmClient) Route(r *RouteRequest) (*OSRMResponse, error) {
	Url, err := r.Url(c)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    Url,
		Header: http.Header{
			"User-Agent": {"GOSRM/1.0.0"},
			"Accept":     {"application/json"},
		},
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	routeResponse := &OSRMResponse{}
	if err := json.Unmarshal(raw, &routeResponse); err != nil {
		return nil, err
	}

	if routeResponse.Code != CodeOK {

		i, ok := RespCode[routeResponse.Code]
		if !ok {
			i = routeResponse.Message
		}

		return routeResponse, errors.New(i)
	}

	return routeResponse, nil
}

func (o *Options) BaseUrl() (*url.URL, error) {
	u, err := url.Parse(o.Url.Host)
	if err != nil {
		return nil, errors.New("gosrm: bad server Url string")
	}

	u.Path += strings.Join([]string{
		o.Service,
		o.Version,
		o.Profile,
	}, "/")

	return u, nil
}

// URL generates a url for OSRM request
func (r RouteRequest) Url(c OsrmClient) (*url.URL, error) {
	Url, err := c.Options.BaseUrl()
	if err != nil {
		return nil, err
	}

	path := geo.Path{r.Coordinates}

	Url.Path += "/" + "polyline(" + url.PathEscape(path.Encode()) + ")"

	parameters := url.Values{}

	if r.Steps != nil {
		parameters.Add("steps", strconv.FormatBool(*r.Steps))
	}
	if r.Alternatives != nil {
		parameters.Add("alternatives", *r.Alternatives)
	}
	if r.Annotations != nil {
		parameters.Add("annotations", *r.Annotations)
	}
	if r.Geometries != nil {
		parameters.Add("geometries", *r.Geometries)
	}
	if r.ContinueStraight != nil {
		parameters.Add("continue_straight", *r.ContinueStraight)
	}
	if r.Overview != nil {
		parameters.Add("overview", *r.Overview)
	}

	Url.RawQuery = parameters.Encode()

	return Url, nil
}
