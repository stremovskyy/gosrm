package gosrm

import (
	geo "github.com/paulmach/go.geo"
)

// RouteRequest Finds the fastest route between coordinates in the supplied order
type RouteRequest struct {
	// coordinates
	Coordinates geo.PointSet `json:"coordinates"`

	// Returned route steps for each route leg
	Steps *bool `json:"steps"`

	// Search for alternative routes. Passing a number alternatives=n searches for up to  n alternative routes
	// (n not working because it bool)
	Alternatives *string `json:"alternatives"`

	// Returns additional metadata for each coordinate along the route geometry
	Annotations *string `json:"annotations"`

	// Returned route geometry format (influences overview and per step)
	Geometries *string `json:"geometries"`

	// Add overview geometry either full, simplified according to highest zoom level it could be display on, or not at all.
	Overview *string `json:"overview"`

	// Forces the route to keep going straight at waypoints constraining uturns there even if it would be faster. Default value depends on the profile.
	ContinueStraight *string `json:"continue_straight"`

	// Adds a Hint to the response which can be used in subsequent requests
	GenerateHints *bool `json:"generate_hints"`
}
