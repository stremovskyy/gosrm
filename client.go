package gosrm

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"context"
	"strings"
	"net/url"
)

type (
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}

	client struct {
		httpClient HTTPClient
		serverURL  string
	}
)

func newClient(serverURL string, c HTTPClient) client {
	return client{c, serverURL}
}

// request contains parameters for OSRM query
type request struct {
	profile string
}

func (c client) doRequest(ctx context.Context, in *request, out interface{}) error {
	url, err := in.URL(c.serverURL)
	if err != nil {
		return err
	}

	resp, err := c.get(ctx, url)
	if err != nil {
		return err
	}
	defer closeSilently(resp.Body)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("unexpected http status code %d with body %q", resp.StatusCode, bytes)
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("failed to unmarshal body %q: %v", bytes, err)
	}

	return nil
}

func (c client) get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(req.WithContext(ctx))
}

func closeSilently(c io.Closer) {
	_ = c.Close()
}

// URL generates a url for OSRM request
func (r *request) URL(serverURL string) (string, error) {
	if r.profile == "" {
		return "", ErrEmptyProfileName
	}
	// http://{server}/{service}/{version}/{profile}/{coordinates}[.{format}]?option=value&option=value
	url := strings.Join([]string{
		serverURL, // server
		r.service, // service
		version,   // version
		r.profile, // profile
		"polyline(" + url.PathEscape(r.coords.Polyline()) + ")", // coordinates
	}, "/")
	if len(r.options) > 0 {
		url += "?" + r.options.encode() // options
	}
	return url, nil
}
