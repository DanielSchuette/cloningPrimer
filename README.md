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


**NOTE**: You can install the CLI without installing **GO** by going to the `./bin` directory of this repository and downloading the respective binary for your operating system. However, this is not recommended because you will have to download the example files (in `./app/assets`) manually as well if you want to run the example code below and you cannot use the local web app or the API.


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

Install the *Cloning Primer* CLI (called `goprimer`) as described [above](#api). To see the available command line arguments, run:

```bash
$ goprimer --help

## Output:

#Usage of goprimer:
#  -3prime_start int
#    	3' position of the first complementary nucleotide in the provided sequence that the reverse primer should bind to
#    	see './doc' for more information on how to customize primer calculations (default 1)
#  -5prime_start int
#    	5' position of the first complementary nucleotide in the provided sequence that the forward primer should bind to
#    	see './doc' for more information on how to customize primer calculations (default 1)
#  -enzyme_file string
#    	valid file path to a *.re file with correctly formatted restriction enzyme information
#    	default is the file at 'github.com/DanielSchuette/app/assets/enzymes.re' (default "../app/assets/enzymes.re")
#  -enzyme_name_forward string
#    	name of the enzyme you want to use for the 5' end (must be in the '--enzyme_file') (default "BamHI")
#  -enzyme_name_reverse string
#    	name of the enzyme you want to use for the 3' end (must be in the '--enzyme_file') (default "EcoRI")
#  -length_forward int
#    	length of the complementary part of the forward primer (default 18)
#  -length_reverse int
#    	length of the complementary part of the reverse primer (default 18)
#  -overhang_forward int
#    	number of random nucleotides added to the forward primer (an integer between 2 - 10) (default 4)
#  -overhang_reverse int
#    	number of random nucleotides added to the reverse primer (an integer between 2 - 10) (default 4)
#  -seq_file string
#    	valid file path to a *.seq file with correctly formatted DNA sequence information
#    	default is the file at 'github.com/DanielSchuette/app/assets/tp53.seq' (default "../app/assets/tp53.seq")
#  -start_codon
#    	set this flag to 'false' if the input sequence does not have a start codon (an ATG will be added automatically) (default true)
#  -stop_codon
#    	set this flag to 'false' if the input sequence does not have a stop cdon (then, a TAA will be added automatically) (default true)
#  -verbose
#    	enable verbose output (defaults to false)
```

If you installed **GO** and `goprimer` correctly (i.e. the `bin/` directory of your workspace is in your `$PATH`), you should be able to run `$ goprimer ...` from every directory. If you want to run the following example, though, navigate to the `./cmd` directory of this repository first (otherwise `goprimer` cannot find the example files):

```bash
$ cd $GOPATH/src/github.com/DanielSchuette/cloningPrimer/cmd
$ goprimer
```

The output should be the forward and reverse primers for cloning the tp53 gene (sequence in `./app/assets/tp53.seq`). Additional command line flags (`--<argument>`) allow for further customization of these primers. If you want to use `goprimer` to design primers based on your own sequence and/or `.re` enzyme file, specify the `--seq_file` and `--enzyme_file` arguments. Be aware that your `.re` and `.seq` files have to follow the formats specified in the example files, otherwise the `goprimer` utility will not be able to parse your data and will throw an error.



### <a name="web_app"></a> Web Application

Documentation is coming soon.




## <a name="license"></a> License

This software is licensed under the MIT license, see *[LICENSE](./LICENSE.txt)* for more information.


