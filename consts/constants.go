package consts

import "time"

const (
	ClientMaxIdleConnections  = 1024
	ClientTLSHandshakeTimeout = 1 * time.Second
)

// Return the requested table or tables in response
type TableAnnotation string

const (
	// default
	TableAnnotationDuration         TableAnnotation = "duration"
	TableAnnotationDistance         TableAnnotation = "distance"
	TableAnnotationDurationDistance TableAnnotation = "duration,distance"
)

func (a *TableAnnotation) String() string {
	return string(*a)
}

// Return the requested table or tables in response
type FallbackCoordinate string

const (
	// default
	FallbackCoordinateInput   FallbackCoordinate = "input"
	FallbackCoordinateSnapped FallbackCoordinate = "snapped"
)

func (a *FallbackCoordinate) String() string {
	return string(*a)
}

const (
	// Lib Version
	Version = "0.1.0"

	// Finds the fastest route between coordinates in the supplied order
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

	// Standard profile
	ProfileDriving = "driving"

	// Car profile
	ProfileCar = "car"

	// Foot profile
	ProfileFoot = "foot"

	// First and only (for now) version of OSRM api
	VersionFirst = "v1"
)
