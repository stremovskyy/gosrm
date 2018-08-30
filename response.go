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

type OSRMResponse struct {
	Code      string     `json:"code"`
	Message   string     `json:"message"`
	Routes    []Route    `json:"routes"`
	Waypoints []Waypoint `json:"waypoints"`
}

type Waypoint struct {
	Hint     string    `json:"hint"`
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
}

type Route struct {
	Geometry   string  `json:"geometry"`
	Legs       []Leg   `json:"legs"`
	Distance   float64 `json:"distance"`
	Duration   float64 `json:"duration"`
	WeightName string  `json:"weight_name"`
	Weight     float64 `json:"weight"`
}

type Leg struct {
	Annotation Annotation `json:"annotation"`
	Steps      []Step     `json:"steps"`
	Distance   float64    `json:"distance"`
	Duration   float64    `json:"duration"`
	Summary    string     `json:"summary"`
	Weight     float64    `json:"weight"`
}

type Annotation struct {
	Metadata    Metadata      `json:"metadata"`
	Nodes       []interface{} `json:"nodes"`
	Datasources []int         `json:"datasources"`
	Speed       []float64     `json:"speed"`
	Weight      []float64     `json:"weight"`
	Duration    []float64     `json:"duration"`
	Distance    []float64     `json:"distance"`
}

type Metadata struct {
	DatasourceNames []string `json:"datasource_names"`
}

type Step struct {
	Intersections []Intersection `json:"intersections"`
	DrivingSide   string         `json:"driving_side"`
	Geometry      string         `json:"geometry"`
	Mode          string         `json:"mode"`
	Duration      float64        `json:"duration"`
	Maneuver      Maneuver       `json:"maneuver"`
	Weight        float64        `json:"weight"`
	Distance      float64        `json:"distance"`
	Name          string         `json:"name"`
	Ref           string         `json:"ref,omitempty"`
}

type Intersection struct {
	Out      int       `json:"out"`
	Entry    []bool    `json:"entry"`
	Bearings []int     `json:"bearings"`
	Location []float64 `json:"location"`
}

type Maneuver struct {
	BearingAfter  int       `json:"bearing_after"`
	Type          string    `json:"type"`
	Modifier      string    `json:"modifier"`
	BearingBefore int       `json:"bearing_before"`
	Location      []float64 `json:"location"`
}
