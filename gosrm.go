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
	"github.com/paulmach/go.geo"
)

type OSRM interface {
	Route(r *RouteRequest) (*OSRMResponse, error)
	Table(r *TableRequest) (*OSRMResponse, error)
	Match(r *MatchRequest) (*OSRMResponse, error)
	Nearest(r *NearestRequest) (*OSRMResponse, error)
}

// Finds the fastest route between coordinates in the supplied order
type RouteRequest struct {
	// coordinates
	Coordinates geo.PointSet `json:"coordinates"`

	// Returned route steps for each route leg
	Steps *bool `json:"steps"`

	// Search for alternative routes. Passing a number alternatives=n searches for up to  n alternative routes
	// (n not working because it bool)
	Alternatives *bool `json:"alternatives"`

	// Returns additional metadata for each coordinate along the route geometry
	Annotations *string `json:"annotations"`

	// Returned route geometry format (influences overview and per step)
	Geometries *string `json:"geometries"`

	// Add overview geometry either full, simplified according to highest zoom level it could be display on, or not at all.
	Overview *string `json:"overview"`

	// Forces the route to keep going straight at waypoints constraining uturns there even if it would be faster. Default value depends on the profile.
	ContinueStraight *string `json:"continue_straight"`
}

// Snaps a coordinate to the street network and returns the nearest n matches.
type NearestRequest struct {
	// coordinates
	Coordinates geo.Point `json:"coordinates"`

	// Number of nearest segments that should be returned
	Number *int `json:"number"`
}

// Computes the duration of the fastest route between all pairs of supplied coordinates
type TableRequest struct {
	// coordinates
	Coordinates geo.PointSet `json:"coordinates"`

	// Use location with given index as source.
	Sources *[]int `json:"sources"`

	// Use location with given index as destination.
	Destinations *[]int `json:"destinations"`
}

type MatchRequest struct {
	// coordinates
	Coordinates geo.PointSet `json:"coordinates"`
}
