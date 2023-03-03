package main

import (
	"fmt"
	geo "github.com/paulmach/go.geo"
	"net/url"

	"github.com/stremovskyy/gosrm"
)

func main() {

	options := &gosrm.Options{
		Url:            url.URL{Host: "https://router.project-osrm.org/"},
		Service:        gosrm.ServiceRoute,
		Version:        gosrm.VersionFirst,
		Profile:        gosrm.ProfileDriving,
		RequestTimeout: 5,
	}

	client := gosrm.NewClient(options)

	routeRequest := &gosrm.RouteRequest{
		Coordinates: geo.PointSet{{36.232051849365234, 49.98765584451778}, {36.22089385986328, 50.03718650830641}},
	}

	response, err := client.Route(routeRequest)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v", response)
}
