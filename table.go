package gosrm

import (
	"net/url"
	"path"
	"strings"
)

func (c *osrmClient) Table(t *TableRequest) (*OSRMResponse, error) {
	Url, err := tableUrl(t, c.baseUrlForService(ServiceTable))
	if err != nil {
		return nil, NewGOSRMError(nil, err, nil)
	}

	return c.doRequest(Url)
}

// tableUrl generates a URL for an OSRM Table request based on the TableRequest and the base URL.
func tableUrl(t *TableRequest, baseURL *url.URL) (*url.URL, error) {
	newURL := *baseURL

	// Construct the coordinates part of the path
	coordinates := make([]string, len(t.Coordinates))
	for i, coord := range t.Coordinates {
		coordinates[i] = formatCoordinate(coord)
	}
	newURL.Path = path.Join(newURL.Path, strings.Join(coordinates, ";"))

	// Initialize and populate query parameters
	parameters := url.Values{}
	addOptionalBoolParam(parameters, "generate_hints", t.GenerateHints)
	addOptionalIndexListParam(parameters, "sources", t.Sources)
	addOptionalIndexListParam(parameters, "destinations", t.Destinations)
	addOptionalFloatParam(parameters, "scale_factor", t.ScaleFactor)

	if t.Annotations != nil {
		addOptionalStringParam(parameters, "annotations", StringRef(t.Annotations.String()))
	}

	if t.FallbackCoordinate != nil {
		addOptionalStringParam(parameters, "fallback_coordinate", StringRef(t.FallbackCoordinate.String()))
	}

	// Assign the encoded query parameters to the URL
	newURL.RawQuery = parameters.Encode()

	return &newURL, nil
}
