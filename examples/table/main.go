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
		Service:        gosrm.ServiceTable,
		Version:        gosrm.VersionFirst,
		Profile:        gosrm.ProfileDriving,
		RequestTimeout: 5,
	}

	client := gosrm.NewClient(options)

	sources := []int{0, 1}
	destinations := []int{1}

	annotation := gosrm.TableAnnotationDurationDistance
	scaleFactor := 1.2

	tableRequest := &gosrm.TableRequest{
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
