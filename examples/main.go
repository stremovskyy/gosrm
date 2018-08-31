package main

import "github.com/paulmach/go.geo"
import (
	"encoding/json"
	"fmt"
	"github.com/Karmadon/gosrm"
	"net/url"
)

func main() {

	options := &gosrm.Options{
		Url:            url.URL{Host: "https://router.project-osrm.org/"},
		Service:        gosrm.ServiceRoute,
		Version:        gosrm.VersionFirst,
		Profile:        gosrm.ProfileDrivig,
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

	json, _ := json.Marshal(response)

	fmt.Printf("%s", json)
}
