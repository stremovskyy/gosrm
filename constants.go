package gosrm

// TableAnnotation Return the requested table or tables in response
type TableAnnotation string

const (
	TableAnnotationDuration         TableAnnotation = "duration" // default
	TableAnnotationDistance         TableAnnotation = "distance"
	TableAnnotationDurationDistance TableAnnotation = "duration,distance"
)

func (a *TableAnnotation) String() string {
	return string(*a)
}

// Return the requested table or tables in response
type FallbackCoordinate string

const (
	FallbackCoordinateInput   FallbackCoordinate = "input" // default
	FallbackCoordinateSnapped FallbackCoordinate = "snapped"
)

func (a *FallbackCoordinate) String() string {
	return string(*a)
}

const (
	// Version of lib
	Version = "0.2.0"

	// ServiceRoute Finds the fastest route between coordinates in the supplied order
	ServiceRoute = "route"

	// ServiceNearest Snaps a coordinate to the street network and returns the nearest n matches
	ServiceNearest = "nearest"

	// ServiceTable Computes the duration of the fastest route between all pairs of supplied coordinates
	ServiceTable = "table"

	// ServiceMatch Map matching matches/snaps given GPS points to the road network in the most plausible way
	ServiceMatch = "match"

	// ServiceTrip trip plugin solves the Traveling Salesman Problem using a greedy heuristic (farthest-insertion algorithm) for 10 or more waypoints and uses brute force for less than 10 waypoints.
	ServiceTrip = "trip"

	// ServiceTile This service generates Mapbox Vector Tiles that can be viewed with a vector-tile capable slippy-map viewer.
	ServiceTile = "tile"

	// ProfileDriving Standard profile
	ProfileDriving = "driving"

	// ProfileCar Car profile
	ProfileCar = "car"

	// ProfileFoot Foot profile
	ProfileFoot = "foot"

	// VersionFirst First and only (for now) version of OSRM api
	VersionFirst = "v1"
)
