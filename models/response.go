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

package models

// Main Response object
type OSRMResponse struct {
	// Response code
	Code string `json:"code,omitempty"`
	// Response message for errors
	Message string `json:"message,omitempty"`
	// Routes array object
	Routes []Route `json:"routes,omitempty"`
	// Waypoints array object
	Waypoints []Waypoint `json:"waypoints,omitempty"`

	Durations    [][]float64   `json:"durations,omitempty"`    // Only for Table Request
	Distances    [][]float64   `json:"distances,omitempty"`    // Only for Table Request
	Destinations []Destination `json:"destinations,omitempty"` // Only for Table Request
	Sources      []Source      `json:"sources,omitempty"`      // Only for Table Request
	Trips        []Trip        `json:"trips,omitempty"`        // Only for Match Request
}

// For Match Request
type Trip struct {
	Geometry   string  `json:"geometry,omitempty"`
	Legs       []Leg   `json:"legs,omitempty"`
	Distance   float64 `json:"distance,omitempty"`
	Duration   float64 `json:"duration,omitempty"`
	WeightName string  `json:"weight_name,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
}

// For Table Request
type Source struct {
	Hint     string    `json:"hint,omitempty"`
	Name     string    `json:"name,omitempty"`
	Location []float64 `json:"location,omitempty"`
}

// For Table Request
type Destination struct {
	Hint     string    `json:"hint,omitempty"`
	Name     string    `json:"name,omitempty"`
	Location []float64 `json:"location,omitempty"`
}

// OSRM Waypoint object
type Waypoint struct {
	Hint          string    `json:"hint,omitempty"`
	Name          string    `json:"name,omitempty"`
	Location      []float64 `json:"location,omitempty"`
	Nodes         []int     `json:"nodes,omitempty"`
	Distance      float64   `json:"distance,omitempty"`       // Nearest Request
	WaypointIndex int       `json:"waypoint_index,omitempty"` // Match Request
	TripsIndex    int       `json:"trips_index,omitempty"`    // Match Request
}

type Route struct {
	Geometry   string  `json:"geometry,omitempty"`
	Legs       []Leg   `json:"legs,omitempty"`
	Distance   float64 `json:"distance,omitempty"`
	Duration   float64 `json:"duration,omitempty"`
	WeightName string  `json:"weight_name,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
}

// OSRM Leg
type Leg struct {
	Annotation Annotation `json:"annotation,omitempty"`
	Steps      []Step     `json:"steps,omitempty"`
	Distance   float64    `json:"distance,omitempty"`
	Duration   float64    `json:"duration,omitempty"`
	Summary    string     `json:"summary,omitempty"`
	Weight     float64    `json:"weight,omitempty"`
}

type Annotation struct {
	Metadata    Metadata      `json:"metadata,omitempty"`
	Nodes       []interface{} `json:"nodes,omitempty"`
	Datasources []int         `json:"datasources,omitempty"`
	Speed       []float64     `json:"speed,omitempty"`
	Weight      []float64     `json:"weight,omitempty"`
	Duration    []float64     `json:"duration,omitempty"`
	Distance    []float64     `json:"distance,omitempty"`
}

type Metadata struct {
	DatasourceNames []string `json:"datasource_names,omitempty"`
}

type Step struct {
	Intersections []Intersection `json:"intersections,omitempty"`
	DrivingSide   string         `json:"driving_side,omitempty"`
	Geometry      string         `json:"geometry,omitempty"`
	Mode          string         `json:"mode,omitempty"`
	Duration      float64        `json:"duration,omitempty"`
	Maneuver      Maneuver       `json:"maneuver,omitempty"`
	Weight        float64        `json:"weight,omitempty"`
	Distance      float64        `json:"distance,omitempty"`
	Name          string         `json:"name,omitempty"`
	Ref           string         `json:"ref,omitempty,omitempty"`
}

type Intersection struct {
	Out      int       `json:"out,omitempty"`
	Entry    []bool    `json:"entry,omitempty"`
	Bearings []int     `json:"bearings,omitempty"`
	Location []float64 `json:"location,omitempty"`
}

type Maneuver struct {
	BearingAfter  int       `json:"bearing_after,omitempty"`
	Type          string    `json:"type,omitempty"`
	Modifier      string    `json:"modifier,omitempty"`
	BearingBefore int       `json:"bearing_before,omitempty"`
	Location      []float64 `json:"location,omitempty"`
}
