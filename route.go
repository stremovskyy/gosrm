package gosrm

import (
	"net/url"

	geo "github.com/paulmach/go.geo"
)

func (c *osrmClient) Route(r *RouteRequest) (*OSRMResponse, error) {
	Url, err := routeUrl(r, c.baseUrlForService(ServiceRoute))
	if err != nil {
		return nil, NewGOSRMError(nil, err, nil)
	}

	return c.doRequest(Url)
}

// URL generates a url for OSRM request
func routeUrl(r *RouteRequest, baseURL *url.URL) (*url.URL, error) {
	newURL := *baseURL

	path := geo.Path{PointSet: r.Coordinates}
	encodedPath := path.Encode()

	encodedPolyline := "polyline(" + url.PathEscape(encodedPath) + ")"
	newURL.Path = newURL.Path + "/" + encodedPolyline

	// Initialize and populate query parameters
	parameters := url.Values{}
	addOptionalBoolParam(parameters, "generate_hints", r.GenerateHints)
	addOptionalBoolParam(parameters, "steps", r.Steps)
	addOptionalStringParam(parameters, "alternatives", r.Alternatives)
	addOptionalStringParam(parameters, "annotations", r.Annotations)
	addOptionalStringParam(parameters, "geometries", r.Geometries)
	addOptionalStringParam(parameters, "continue_straight", r.ContinueStraight)
	addOptionalStringParam(parameters, "overview", r.Overview)

	// Assign the encoded query parameters to the URL
	newURL.RawQuery = parameters.Encode()

	return &newURL, nil
}
