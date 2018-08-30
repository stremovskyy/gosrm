package gosrm

type RouteResponse struct {
	Code      string     `json:"code"`
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
	Distance   int     `json:"distance"`
	Duration   float64 `json:"duration"`
	WeightName string  `json:"weight_name"`
	Weight     float64 `json:"weight"`
}

type Leg struct {
	Annotation Annotation `json:"annotation"`
	Steps      []Step     `json:"steps"`
	Distance   int        `json:"distance"`
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
