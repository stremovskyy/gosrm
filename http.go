package gosrm

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/karmadon/gosrm/consts"
)

func (c *OsrmClient) http(Url *url.URL) ([]byte, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    Url,
		Header: http.Header{
			"User-Agent": {"GOSRM/" + consts.Version},
			"Accept":     {"application/json"},
		},
	}

	if c.Options.Debug {
		fmt.Printf("[GOSRM][URL]: %s\n", Url.String())
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return nil, err
	}

	if c.Options.Debug {
		fmt.Printf("[GOSRM][RESPONCE]: %s\n", raw)
	}
	return raw, nil
}
