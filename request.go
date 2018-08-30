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
	"net/url"
)

const (
	/* Finds the fastest route between coordinates in the supplied order */
	ServiceRoute = "route"

	// Snaps a coordinate to the street network and returns the nearest n matches
	ServiceNearest = "nearest"

	// Computes the duration of the fastest route between all pairs of supplied coordinates
	ServiceTable = "table"

	// Map matching matches/snaps given GPS points to the road network in the most plausible way
	ServiceMatch = "match"

	// The trip plugin solves the Traveling Salesman Problem using a greedy heuristic (farthest-insertion algorithm) for 10 or more waypoints and uses brute force for less than 10 waypoints.
	ServiceTrip = "trip"

	// This service generates Mapbox Vector Tiles that can be viewed with a vector-tile capable slippy-map viewer.
	ServiceTile = "tile"

	// Standart profile
	ProfileDrivig = "driving"

	// Car profile
	ProfileCar = "car"

	// Foot profile
	ProfileFoot = "foot"

	// First and only (for now) version of OSRM api
	VersionFirst = "v1"
)

// Client object
type Options struct {
	// Url of osrm server with / on the end
	Url url.URL `json:"url"`

	// One of the following values:  route ,  nearest ,  table ,  match ,  trip ,  tile
	Service string `json:"service"`

	//	Version of the protocol implemented by the service
	Version string `json:"version"`

	// Mode of transportation, is determined statically by the Lua profile that is used to prepare the data using  osrm-extract
	Profile string `json:"profile"`

	// Timeout for request in seconds
	RequestTimeout int `json:"request_timeout"`
}

type UrlResponder interface {
	Url() (*url.URL, error)
}

type RouteRequest struct {
	//coordinates
	Coordinates geo.PointSet `json:"coordinates"`

	// Returned route steps for each route leg
	Steps *bool `json:"steps"`

	// Search for alternative routes. Passing a number alternatives=n searches for up to  n alternative routes
	Alternatives *string `json:"alternatives"`

	// Returns additional metadata for each coordinate along the route geometry
	Annotations *string `json:"annotations"`

	// Returned route geometry format (influences overview and per step)
	Geometries *string `json:"geometries"`

	// Add overview geometry either full, simplified according to highest zoom level it could be display on, or not at all.
	Overview *string `json:"overview"`

	// Forces the route to keep going straight at waypoints constraining uturns there even if it would be faster. Default value depends on the profile.
	ContinueStraight *string `json:"continue_straight"`
}
