package gosrm

import (
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) Table(t *TableRequest) (*OSRMResponse, error) {
	baseURL, err := c.options.BaseUrl()
	Url, err := tableUrl(t, baseURL)

	if err != nil {
		return nil, NewGOSRMError(nil, err, nil)
	}

	return c.http(Url)
}

// URL generates a url for OSRM request
func tableUrl(r *TableRequest, baseURL *url.URL) (*url.URL, error) {
	locations := []string{}

	for _, coordinate := range r.Coordinates {
		locations = append(locations, strconv.FormatFloat(coordinate.Lng(), 'f', 9, 64)+","+strconv.FormatFloat(coordinate.Lat(), 'f', 9, 64))
	}

	baseURL.Path += "/" + strings.Join(locations, ";")

	parameters := url.Values{}

	if r.GenerateHints != nil {
		parameters.Add("generate_hints", strconv.FormatBool(*r.GenerateHints))
	}

	if r.Annotations != nil {
		parameters.Add("annotations", r.Annotations.String())
	}

	if r.Sources != nil {
		sources := ""

		if r.Sources != nil && len(*r.Sources) > 0 {
			var t []string
			for _, i := range *r.Sources {
				t = append(t, strconv.Itoa(i))
			}
			sources = strings.Join(t, ";")
		}

		if sources != "" {
			parameters.Add("sources", sources)
		}
	}
	if r.Destinations != nil {
		destinations := ""

		if r.Destinations != nil && len(*r.Destinations) > 0 {
			var t []string
			for _, i := range *r.Destinations {
				t = append(t, strconv.Itoa(i))
			}
			destinations = strings.Join(t, ";")
		}

		if destinations != "" {
			parameters.Add("destinations", destinations)
		}

	}

	if r.ScaleFactor != nil {
		parameters.Add("scale_factor", strconv.FormatFloat(*r.ScaleFactor, 'f', 4, 32))
	}

	if r.FallbackCoordinate != nil {
		parameters.Add("fallback_coordinate", r.FallbackCoordinate.String())
	}

	baseURL.RawQuery = parameters.Encode()

	return baseURL, nil
}
