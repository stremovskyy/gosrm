package gosrm

import (
	geo "github.com/paulmach/go.geo"
)

// Computes the duration of the fastest route between all pairs of supplied coordinates
type TableRequest struct {
	// coordinates
	Coordinates geo.PointSet `json:"coordinates"`

	// Use location with given index as source.
	Sources *[]int `json:"sources"`

	// Use location with given index as destination.
	Destinations *[]int `json:"destinations"`

	// Returns additional metadata for each coordinate along the route geometry
	Annotations *TableAnnotation `json:"annotations"`

	//If no route found between a source/destination pair, calculate the as-the-crow-flies distance, then use this speed to estimate duration.
	FallbackSpeed *float64 `json:"fallback_speed"`

	//When using a  fallback_speed , use the user-supplied coordinate ( input ), or the snapped location ( snapped ) for calculating distances.
	FallbackCoordinate *FallbackCoordinate `json:"fallback_coordinate"`

	//Use in conjunction with  annotations=durations . Scales the table duration values by this number.
	ScaleFactor *float64 `json:"scale_factor"`
}
