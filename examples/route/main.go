package main

import (
	"fmt"
	"net/url"

	"github.com/paulmach/go.geo"

	"github.com/karmadon/gosrm"
	"github.com/karmadon/gosrm/consts"
	"github.com/karmadon/gosrm/models"
)

func main() {

	options := &gosrm.Options{
		Url:            url.URL{Host: "https://router.project-osrm.org/"},
		Service:        consts.ServiceRoute,
		Version:        consts.VersionFirst,
		Profile:        consts.ProfileDriving,
		RequestTimeout: 5,
	}

	client := gosrm.NewClient(options)

	routeRequest := &models.RouteRequest{
		Coordinates: geo.PointSet{{36.232051849365234, 49.98765584451778}, {36.22089385986328, 50.03718650830641}},
	}

	response, err := client.Route(routeRequest)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v", response)
}
