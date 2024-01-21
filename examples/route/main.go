package main

import (
	"fmt"

	geo "github.com/paulmach/go.geo"

	"github.com/stremovskyy/gosrm"
)

func main() {
	client := gosrm.NewDefaultClient()

	// Create a route request
	routeRequest := &gosrm.RouteRequest{
		GenerateHints: gosrm.BoolRef(false),
		Coordinates: geo.PointSet{
			{36.232051849365234, 49.98765584451778},
			{36.22089385986328, 50.03718650830641},
		},
	}

	// Perform the route request
	response, err := client.
		Debug().
		ForProfile(gosrm.ProfileCar).
		Route(routeRequest)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(response)
}
