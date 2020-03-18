package models

import geo "github.com/paulmach/go.geo"

// Snaps a coordinate to the street network and returns the nearest n matches.
type NearestRequest struct {
	// coordinates
	Coordinates geo.Point `json:"coordinates"`

	// Number of nearest segments that should be returned
	Number *int `json:"number"`
}
