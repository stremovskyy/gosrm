# GOSRM - GO Client for OSRM
[![Build Status](https://travis-ci.org/Karmadon/gosrm.svg?branch=master)](https://travis-ci.org/Karmadon/gosrm)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/Karmadon/gosrm)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/Karmadon/gosrm)](https://goreportcard.com/report/github.com/Karmadon/gosrm)

Advanced OSRM client for golang.

## Features

- Route Client

## Installation

```bash
go get github.com/karmadon/gosrm
```

## Quick Start

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
    		Url: url.URL{Host:"https://router.project-osrm.org/"},
    		Service:        gosrm.ServiceRoute,
    		Version:        gosrm.VersionFirst,
    		Profile:        gosrm.ProfileDrivig,
    		RequestTimeout: 5,
    	}
    
    	client := gosrm.NewClient(options)
    
    	routeRequest := &gosrm.RouteRequest{
    		Coordinates: geo.PointSet{{36.232051849365234, 49.98765584451778}, {36.22089385986328, 50.03718650830641},},
    	}
    
    	response,err := client.Route(routeRequest)
    	if err != nil {
    		panic(err.Error())
    	}
    
    	fmt.Printf("%v", response)
}
```

### Full Example

See `examples/main.go`
