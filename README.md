# GOSRM - GO Client for OSRM
[![Build Status](https://travis-ci.org/Karmadon/gosrm.svg?branch=master)](https://travis-ci.org/Karmadon/gosrm)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/karmadon/gosrm)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/karmadon/gosrm)](https://goreportcard.com/report/github.com/Karmadon/gosrm)

Advanced OSRM client for golang.

## Features

- Route Client

## Installation

```bash
go get github.com/karmadon/gosrm
```

## Quick Start

```go
package gosrm

import (
    "github.com/karmadon/gosrm"
)

func main() {
	ra := &Request{
		Url:         "https://router.project-osrm.org/",
		Service:     ServiceRoute,
		Version:     1,
		Profile:     "driving",
		Coordinates: geo.PointSet{{36.232051849365234, 49.98765584451778}, {36.22089385986328, 50.03718650830641},},
	}

	resp, err := ra.GetOSRMResponse()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%v", resp)
}
```
