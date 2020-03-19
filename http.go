package gosrm

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *OsrmClient) http(Url *url.URL) (*OSRMResponse, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    Url,
		Header: http.Header{
			"User-Agent": {"GOSRM/" + Version},
			"Accept":     {"application/json"},
		},
	}

	if c.Options.Debug {
		fmt.Printf("[GOSRM][URL]: %s\n", Url.String())
	}

	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NewGOSRMError(Url, err, nil)
	}
	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, NewGOSRMError(Url, err, nil)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	raw, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewGOSRMError(Url, err, &raw)
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return nil, NewGOSRMError(Url, err, &raw)
	}

	if c.Options.Debug {
		fmt.Printf("[GOSRM][RESPONCE]: %s\n", raw)
	}

	response := &OSRMResponse{}
	if err := json.Unmarshal(raw, &response); err != nil {
		return nil, NewGOSRMError(Url, err, &raw)
	}

	if response.Code != CodeOK {

		i, ok := RespCode[response.Code]
		if !ok {
			i = response.Message
		}

		e := errors.New(i)

		return response, NewGOSRMError(Url, e, &raw)
	}

	return response, nil
}
