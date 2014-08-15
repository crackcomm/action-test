package main

import "os"
import "fmt"
import "log"
import "flag"
import "github.com/crackcomm/go-actions/core"
import "github.com/crackcomm/go-actions/source/utils"
import testutils "github.com/crackcomm/action-test/utils"

// Core functions
import _ "github.com/crackcomm/go-core"

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
	core.AddSources(utils.GetSources(sources)...)

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
	flag.StringVar(&dirname, "tests", "", "Files or directory containing YAML or JSON tests (can be glob pattern)")
	flag.StringVar(&sources, "sources", "", "Actions sources (comma separated directories & urls)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of action-test:\n  %s\n", actionTestDesctiption)
		// fmt.Fprint(os.Stderr, actionTestDesctiption)
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
}
