package gosrm

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// doRequest performs the HTTP request and handles the response.
func (c *osrmClient) doRequest(url *url.URL) (*OSRMResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		c.logger.Printf("Request Error: %s\n", err.Error())

		return nil, err
	}

	req.Header.Set("User-Agent", "GOSRM/"+Version)
	req.Header.Set("Accept", "application/json")
	if c.options.UseGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	if c.options.Debug {
		c.logger.Printf("Request %s\n", url.String())
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Printf("Response Error: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := c.handleResponseBody(resp)
	if err != nil {
		c.logger.Printf("Response Error: %s\n", err.Error())
		return nil, err
	}

	return c.parseResponse(body)
}

// handleResponseBody handles the response body based on Content-Encoding.
func (c *osrmClient) handleResponseBody(resp *http.Response) ([]byte, error) {
	var reader io.Reader = resp.Body

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	return io.ReadAll(reader)
}

// parseResponse parses the raw response body into OSRMResponse.
func (c *osrmClient) parseResponse(body []byte) (*OSRMResponse, error) {
	if c.options.Debug {
		c.logger.Printf("Response: %s\n", string(body))
	}

	var response OSRMResponse
	if err := json.Unmarshal(body, &response); err != nil {
		c.logger.Printf("Error: %s\n", err.Error())
		return nil, err
	}

	if response.Code != CodeOK {
		errMsg := response.Message
		if msg, ok := RespCode[response.Code]; ok {
			errMsg = msg
		}
		c.logger.Printf("Error: %s\n", errMsg)

		return &response, errors.New(errMsg)
	}

	return &response, nil
}
