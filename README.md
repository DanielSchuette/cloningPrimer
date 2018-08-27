# CloningPrimer

![GitHub (pre-)release](https://img.shields.io/badge/release-v0.0.3-green.svg) [![GoDoc](https://godoc.org/github.com/DanielSchuette/cloningPrimer?status.svg)](https://godoc.org/github.com/DanielSchuette/cloningPrimer) ![Packagist](https://img.shields.io/packagist/l/doctrine/orm.svg) [![Build Status](https://travis-ci.org/DanielSchuette/cloningPrimer.svg?branch=master)](https://travis-ci.org/DanielSchuette/cloningPrimer) [![codecov](https://codecov.io/gh/DanielSchuette/cloningPrimer/branch/master/graph/badge.svg)](https://codecov.io/gh/DanielSchuette/cloningPrimer)

## <a name="about"></a> About

Cloning Primer is a software tool that facilitates the design of primer pairs for gene cloning. Given a certain nucleotide sequence and a set of parameters (see [Documentation](#documentation)), it returns forward and reverse primers as well as useful statistics like GC content, the probability of the primer pair to form dimers, and much more.

The software is accessible via a [web application](http://www.cloningprimer.com), but it is recommended to download the command line application ([CLI](./bin)) to enable offline use. Also, an API (written in Go) is available through this GitHub repository. Cloning Primer is under active development, so source code, web interface, API, and CLI might change in the future!

A working version of Cloning Primer is currently available as a pre-release version v0.0.3. Please review the [documentation section](#documentation) of this README file for more information about available functionality.





## <a name="overview"></a> Overview

A pre-release version of the web application is available [here](http://www.cloningprimer.com).

CLI binaries can be downloaded from [./bin](./bin) in this repository.


### Table of Contents

1. [About](#about)

2. [Overview](#overview)

3. [Documentation](#documentation)

    * [Installation and API](#api)

    * [CLI](#cli)

    * [Web App](#web_app)

4. [License](#license)




## <a name="documentation"></a> Documentation

### <a name="api"></a> Installation and Application Programming Interface (API)

#### Installation

To use *Cloning Primer*, install **The Go Programming Language** from [here](https://golang.org/). Follow [these](https://golang.org/doc/code.html) instructions to set up a working **Go** programming environment (you need to set the `$GOPATH` environment variable to be able to `go get` third party packages). The, run the following command at a command prompt:

```bash
$ go get github.com/DanielSchuette/cloningPrimer
```

To install the `goprimer` command line tool (CLI), run:

```bash
$ cd $GOPATH/src/github.com/DanielSchuette/cloningPrimer
$ make
```

You can run a local version of the *Cloning Primer* [web app](http://cloningprimer.com), too:

```bash
$ cd $GOPATH/src/github.com/DanielSchuette/cloningPrimer/app
$ go run server.go # runs the web app at `localhost:8080'
```

#### Usage Example

More documentation regarding the Go API of *Cloning Primer* is available [here](https://godoc.org/github.com/DanielSchuette/cloningPrimer).

A minimal, working example:

```go
package main

import (
	"fmt"
	"log"

	cloningprimer "github.com/DanielSchuette/cloningPrimer"
)

func main() {
	// define an input string (must be a valid nucleotide sequence)
	input := "ATGCAAAAACGGGCGATTTATCCGGGTACTTTCGATCCCATTACCAATGGTCATATCGATATCGTGACGCGCGCCACGCAGATGTTCGATCACGTTATTCTGGCGATTGCCGCCAGCCCCAGTAAAAAACCGATGTTTACCCTGGAAGAGCGTGTGGCACTGGCACAGCAGGCAACCGCGCATCTGGGGAACGTGGAAGTGGTCGGGTTTAGTGATTTAATGGCGAACTTCGCCCGTAATCAACACGCTACGGTGCTGATTCGTGGCCTGCGTGCGGTGGCAGATTTTGAATATGAAATGCAGCTGGCGCATATGAATCGCCACTTAATGCCGGAACTGGAAAGTGTGTTTCTGATGCCGTCGAAAGAGTGGTCGTTTATCTCTTCATCGTTGGTGAAAGAGGTGGCGCGCCATCAGGGCGATGTCACCCATTTCCTGCCGGAGAATGTCCATCAGGCGCTGATGGCGAAGTTAGCGTAG"

	// find forward primer with EcoRI restriction site and 5 random nucleotides as an overhang
	forward, err := cloningprimer.FindForward(input, "GAATTC", 1, 18, 4, false)
	if err != nil {
		log.Fatalf("error finding forward primer: %s\n", err)
	}
	if forward == "" {
		log.Fatalf("no forward primer found\n")
	}

	// find reverse primer with BamHI restriction site and 3 random nucleotides as an overhang
	reverse, err := cloningprimer.FindReverse(input, "GGATCC", 1, 20, 3, true)
	if err != nil {
		log.Fatalf("error finding reverse primer: %s\n", err)
	}
	if reverse == "" {
		log.Fatalf("no reverse primer found\n")
	}

	// print results
	fmt.Printf("input: %s\nforward primer: %s\nreverse primer: %s\n", input, forward, reverse)
}
```


### <a name="cli"></a> Command Line Interface (CLI)

Documentation is coming soon.



### <a name="web_app"></a> Web Application

Documentation is coming soon.



## <a name="license"></a> License

This software is licensed under the MIT license, see *[LICENSE](./LICENSE.txt)* for more information.

















