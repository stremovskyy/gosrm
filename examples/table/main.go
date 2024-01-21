package main

import (
	"fmt"

	geo "github.com/paulmach/go.geo"

	"github.com/stremovskyy/gosrm"
)

func main() {
	// Initialize the client with these options
	client := gosrm.NewDefaultClient()

	// Define sources and destinations
	sources := []int{0, 1}
	destinations := []int{1}

	// Define annotations and scale factor
	annotation := gosrm.TableAnnotationDurationDistance
	scaleFactor := 1.2

	// Create a table request
	tableRequest := &gosrm.TableRequest{
		Coordinates:  geo.PointSet{{36.232051849365234, 49.98765584451778}, {36.22089385986328, 50.03718650830641}},
		Sources:      &sources,
		Destinations: &destinations,
		Annotations:  &annotation,
		ScaleFactor:  &scaleFactor,
	}

	// Perform the table request
	response, err := client.
		Table(tableRequest)

	if err != nil {
		panic(err.Error())
	}

	// Print the response
	fmt.Println(response)
}
