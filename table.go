package gosrm

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"

	geo "github.com/paulmach/go.geo"

	"github.com/karmadon/gosrm/models"
)

func (c *OsrmClient) Table(t *models.TableRequest) (*models.OSRMResponse, error) {
	baseURL, err := c.Options.BaseUrl()
	Url, err := tableUrl(t, baseURL, c.Options.GenerateHints)
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
func tableUrl(r *models.TableRequest, baseURL *url.URL, hints bool) (*url.URL, error) {
	path := geo.Path{PointSet: r.Coordinates}
	baseURL.Path += "/" + "polyline(" + url.PathEscape(path.Encode()) + ")"

	parameters := url.Values{}

	parameters.Add("generate_hints", strconv.FormatBool(hints))

	if r.Annotations != nil {
		parameters.Add("annotations", r.Annotations.String())
	}

	if r.Sources != nil {
		sources := ""

		if len(*r.Sources) != len(r.Coordinates) {
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

		if len(*r.Destinations) != len(r.Coordinates) {
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
