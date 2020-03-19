# GOSRM - GO Client for OSRM
[![Build Status](https://travis-ci.org/Karmadon/gosrm.svg?branch=master)](https://travis-ci.org/Karmadon/gosrm)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/karmadon/gosrm)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/karmadon/gosrm)](https://goreportcard.com/report/github.com/karmadon/gosrm)

Advanced OSRM client for golang.

## Features

- Route Client

## Installation

```bash
go get github.com/karmadon/gosrm
```

## Quick Start

####Route Service
```go
package main

import (
	"fmt"
	"net/url"
	"github.com/karmadon/gosrm"
	"github.com/paulmach/go.geo"
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
```
####Table Service
```go
package main

import (
	"fmt"
	"net/url"
	"github.com/karmadon/gosrm"
	"github.com/paulmach/go.geo"
)

func main() {
	options := &gosrm.Options{
		Url:            url.URL{Host: "https://router.project-osrm.org/"},
		Service:        consts.ServiceTable,
		Version:        consts.VersionFirst,
		Profile:        consts.ProfileDriving,
		RequestTimeout: 5,
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
```

### Full Example

See `examples` folder
