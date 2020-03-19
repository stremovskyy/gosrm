package gosrm

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	geo "github.com/paulmach/go.geo"

	"github.com/karmadon/gosrm/models"
)

func (c *OsrmClient) Route(r *models.RouteRequest) (*models.OSRMResponse, error) {
	baseURL, err := c.Options.BaseUrl()
	Url, err := routeUrl(r, baseURL, c.Options.GenerateHints)
	if err != nil {
		return nil, err
	}

	raw, err := c.http(Url)
	if err != nil {
		return nil, err
	}

	response := &models.OSRMResponse{}
	if err := json.Unmarshal(raw, &response); err != nil {
		return nil, err
	}

	if response.Code != CodeOK {

		i, ok := RespCode[response.Code]
		if !ok {
			i = response.Message
		}

		return response, errors.New(i)
	}

	return response, nil
}

// URL generates a url for OSRM request
func routeUrl(r *models.RouteRequest, baseURL *url.URL, hints bool) (*url.URL, error) {
	path := geo.Path{PointSet: r.Coordinates}

	baseURL.Path += "/" + "polyline(" + url.PathEscape(path.Encode()) + ")"

	parameters := url.Values{}

	parameters.Add("generate_hints", strconv.FormatBool(hints))

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

	baseURL.RawQuery = parameters.Encode()

	return baseURL, nil
}
