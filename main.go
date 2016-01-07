package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	testutils "github.com/crackcomm/action-test/utils"
	"github.com/crackcomm/go-actions/core"

	// Sources
	_ "github.com/crackcomm/go-actions/source/file"
	_ "github.com/crackcomm/go-actions/source/http"

	// Core functions
	_ "github.com/crackcomm/go-core"
)

// l - Logger
var l = log.New(os.Stdout, "[action-test] ", 0)

// dirname - json or yaml file containing tests
var dirname = "tests.json"

// sources - actions sources
var sources string

func main() {
	// Parse flags and print Usage if `-tests` flag is empty.
	parseFlags()

	// Load tests from file
	tests, err := testutils.ReadTests(dirname)
	if err != nil {
		l.Fatal(err)
	}

	// Print error when no tests
	if len(tests) == 0 {
		fmt.Printf("--- FAIL no tests")
		return
	}

	// Add actions sources
	for _, source := range strings.Split(sources, ",") {
		core.Source(source)
	}

	// Run tests
	results := tests.Run()

	// Print results
	results.Print()
}

var actionTestDesctiption = `Application action-test runs tests from JSON or YAML files against actions from different sources.`

func parseFlags() {
	// Print help if --help flag or --test empty
	flag.Parse()
	if dirname == "" {
		flag.Usage()
		os.Exit(0)
	}
}

// init - Adds usage to flag.Usage :)
func init() {
	flag.StringVar(&dirname, "tests", "", "files or directory containing YAML or JSON tests (can be glob pattern)")
	flag.StringVar(&sources, "sources", "", "actions sources (comma separated)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of action-test:\n  %s\n", actionTestDesctiption)
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
}
