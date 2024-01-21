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
	"fmt"
	"net/url"
)

// RespCode Errors which could be returned from OSRM Server
var RespCode = map[string]string{
	"NoRoute":        "No route found",
	"NoTable":        "No route found in table method.",
	"NoMatch":        "No matching found.",
	"NoTrips":        "No trips found because input coordinates are not connected.",
	"NotImplemented": "This request is not supported",
	"InvalidUrl":     "URL string is invalid",
	"Ok":             "All OK",
	"InvalidService": "Service name is invalid.",
	"InvalidVersion": "OSRM Version is not found.",
	"InvalidOptions": "options are invalid.",
	"InvalidQuery":   "The query string is syntactically malformed.",
	"InvalidValue":   "The successfully parsed query parameters are invalid.",
	"NoSegment":      "One of the supplied input coordinates could not snap to street segment.",
	"TooBig":         "The request size violates one of the service specific request size restrictions",
}

const CodeOK = "Ok"

// Error represents an error response from the OSRM service.
type Error struct {
	URL         *url.URL
	Code        string
	Message     string
	RawResponse *[]byte
}

// Error returns the string representation of the error.
func (e *Error) Error() string {
	return fmt.Sprintf("[GOSRM][ERROR]: %s - %s", e.Code, e.Message)
}

func NewGOSRMError(url *url.URL, err error, rawResponse *[]byte) *Error {
	return &Error{
		URL:         url,
		Code:        "GOSRMError",
		Message:     err.Error(),
		RawResponse: rawResponse,
	}
}
