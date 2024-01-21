# GOSRM - GO Client for OSRM

[![Build Status](https://travis-ci.org/stremovskyy/gosrm.svg?branch=master)](https://travis-ci.org/stremovskyy/gosrm)
![Go](https://github.com/stremovskyy/gosrm/workflows/Go/badge.svg?branch=master)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/stremovskyy/gosrm)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/stremovskyy/gosrm)](https://goreportcard.com/report/github.com/stremovskyy/gosrm)

`GOSRM` is an advanced Go client library for interacting with the Open Source Routing Machine (OSRM) API. It simplifies the process of routing, table, nearest, match, and trip requests to the OSRM
server.

## Features

- Supports all OSRM services including Route, Nearest, Table, Match, Trip, and Tile.
- Customizable request options for advanced use cases.
- Convenient handling of OSRM responses.
- Extensible design for easy integration with existing Go applications.

## Installation

Install `GOSRM` by running:

```bash
go get github.com/stremovskyy/gosrm
```

## Quick Start

### Route Service

Get the fastest route between coordinates with the Route service:

```go
package main

import (
"fmt"
	"net/url"



"github.com/paulmach/go.geo"

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
```

### Table Service

Compute duration tables for pairs of coordinates:

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/paulmach/go.geo"

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
	response, err := client.Table(tableRequest)
	if err != nil {
		panic(err.Error())
	}

	// Print the response
	fmt.Println(response)
}

```

### Full Example

For more examples, see the `examples` folder in this repository.
