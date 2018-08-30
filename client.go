package gosrm

import (
	"encoding/json"
	"errors"
	"github.com/paulmach/go.geo"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// URL generates a url for OSRM request
func (r *Request) URL(serverURL string) (*url.URL, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		panic("gosrm: bad server Url string")
	}

	path := geo.Path{r.Coordinates}

	u.Path += strings.Join([]string{
		r.Service,
		"v" + strconv.Itoa(r.Version),
		r.Profile,
		"polyline(" + url.PathEscape(path.Encode()) + ")",
	}, "/")

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

	u.RawQuery = parameters.Encode()

	return u, nil
}

func (r *Request) GetOSRMResponse() (*OSRMResponse, error) {
	raw, err := r.fire()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(raw.Body)
	if err != nil {
		return nil, err
	}

	resp := &OSRMResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if resp.Code != CodeOK {

		i, ok := RespCode[resp.Code]
		if !ok {
			i = resp.Message
		}

		return nil, errors.New(i)
	}

	return resp, nil
}

func (r *Request) fire() (*http.Response, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	u, _ := r.URL(r.Url)

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: nil,
	}

	//req.Header.Set("User-Agent","GOSRM/1.0.0")
	//req.Header.Set("Accept","application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
