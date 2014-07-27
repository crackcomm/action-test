package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/crackcomm/action-test/testing"
	"github.com/crackcomm/go-actions/core"
	"github.com/crackcomm/go-actions/source/file"
	"github.com/crackcomm/go-actions/source/http"
	_ "github.com/crackcomm/go-core"
	"log"
	"net/url"
	"os"
	"strings"
)

// l - Logger
var l = log.New(os.Stdout, "[action-test] ", 0)

// flags
var (
	testfile = "tests.json" // json file containing tests
	sources  []string       // actions sources (comma separated)
)

func main() {
	var srcs string
	// var printhelp bool
	flag.StringVar(&testfile, "tests", "", "File containing json ")
	flag.StringVar(&srcs, "sources", "", "Actions sources (comma separated directories & urls)")
	// flag.BoolVar(&printhelp, "help", false, "Print help")
	flag.Parse()

	// Split comma separated sources into a list
	sources = strings.Split(srcs, ",")

	// Print help if --help flag or --test empty
	if testfile == "" {
		flag.Usage()
		return
	}

	// Open file containing tests
	f, err := os.Open(testfile)
	if err != nil {
		l.Fatal(err)
	}

	// Load tests from file
	tests := testing.Tests{}
	err = json.NewDecoder(f).Decode(&tests)
	if err != nil {
		l.Fatal(err)
	}

	// Close file (ignore error, continue testing)
	f.Close()

	// Add actions sources
	for _, source := range sources {
		// If source is a valid url - create http source
		if isURL(source) {
			core.AddSource(&http.Source{Path: source})
		} else {
			// Add file source to default core registry
			core.AddSource(&file.Source{source})
		}
	}

	// Run tests
	results := tests.Run()

	// Print results
	results.Print()
}

// isURL - Returns true if value url scheme is a `http` or `https`.
func isURL(value string) (yes bool) {
	if uri, err := url.Parse(value); err == nil {
		if uri.Scheme == "http" || uri.Scheme == "https" {
			yes = true
		}
	}
	return
}

var actionTestDesctiption = `Application action-test runs tests from JSON files against actions from different sources.`

var exampleTestJSON = `
  [
    {
      "action": "filmweb.find",
      "description": "Should find movie by title",
      "arguments": {
        "title": "Pulp Fiction"
      },
      "expect": {
        "writers": "Quentin Tarantino",
        "directors": "Quentin Tarantino",
        "title": "Pulp Fiction",
        "year": "1994"
      }
    }
  ]
`

// init - Adds usage to flag.Usage :)
func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of action-test:\n  %s\n", actionTestDesctiption)
		// fmt.Fprint(os.Stderr, actionTestDesctiption)
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Example JSON tests:\n")
		fmt.Fprint(os.Stderr, exampleTestJSON)
	}
}
