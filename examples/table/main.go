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
		Service:        consts.ServiceTable,
		Version:        consts.VersionFirst,
		Profile:        consts.ProfileDriving,
		RequestTimeout: 5,
		Debug:          true,
	}

	client := gosrm.NewClient(options)

	sources := []int{1, 2}
	destinations := []int{2, 1}

	annotation := consts.TableAnnotationDurationDistance
	scaleFactor := 1.2

	tableRequest := &models.TableRequest{
		Coordinates:  geo.PointSet{{36.232051849365234, 49.98765584451778}, {36.22089385986328, 50.03718650830641}},
		Sources:      &sources,
		Destinations: &destinations,
		Annotations:  &annotation,
		ScaleFactor:  &scaleFactor,
	}

	response, err := client.Table(tableRequest)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v", response)
}
