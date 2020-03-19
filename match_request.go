package gosrm

import geo "github.com/paulmach/go.geo"

type MatchRequest struct {
	// coordinates
	Coordinates geo.PointSet `json:"coordinates"`
}
