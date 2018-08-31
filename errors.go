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

// Errors which could be returned from OSRM Server
var RespCode = map[string]string{
	"NoRoute":        "No route found",
	"NoTable":        "No route found.",
	"NoMatch":        "No matchings found.",
	"NoTrips":        "No trips found because input coordinates are not connected.",
	"NotImplemented": "This request is not supported",
	"InvalidUrl":     "URL string is invalid",
	"Ok":             "All OK",
	"InvalidService": "Service name is invalid.",
	"InvalidVersion": "OSRM Version is not found.",
	"InvalidOptions": "Options are invalid.",
	"InvalidQuery":   "The query string is synctactically malformed.",
	"InvalidValue":   "The successfully parsed query parameters are invalid.",
	"NoSegment":      "One of the supplied input coordinates could not snap to street segment.",
	"TooBig":         "The request size violates one of the service specific request size restrictions",
}

const CodeOK = "Ok"
