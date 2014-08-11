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
	"github.com/golang/glog"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// l - Logger
var l = log.New(os.Stdout, "[action-test] ", 0)

// flags
var (
	testfile = "tests.json" // json or yaml file containing tests
	sources  []string       // actions sources (comma separated)
)

func main() {
	// Parse flags and print Usage if `-tests` flag is empty.
	parseFlags()

	// Load tests from file
	tests, err := readTests()
	if err != nil {
		l.Fatal(err)
	}

	// Add actions sources
	for _, source := range sources {
		// pass if empty
		if source == "" {
			continue
		}
		
		// If source is a valid url - create http source
		if isURL(source) {
			glog.V(3).Infof("New HTTP source: %#v", source)
			core.AddSource(&http.Source{Path: source})
		} else {
			// Add file source to default core registry
			glog.V(3).Infof("New File source: %#v", source)
			core.AddSource(&file.Source{source})
		}
	}

	// Print error when no tests
	if len(tests) == 0 {
		fmt.Printf("--- FAIL no tests")
		return
	}

	// Run tests
	results := tests.Run()

	// Print results
	results.Print()
}

func readTests() (tests testing.Tests, err error) {
	ext := filepath.Ext(testfile)
	if ext == "" {
		glog.V(3).Infof("Tests flag is a directory: %v", testfile)
		var files []string
		var dirname string
		// If `-tests` flag contains `*` it's already a pattern
		if strings.Contains(testfile, "*") {
			dirname = testfile
		} else {
			dirname = filepath.Join(testfile, "*")
		}
		glog.V(3).Infof("Looking for tests in %s", dirname)
		files, err = filepath.Glob(dirname)
		if err != nil {
			return
		}
		glog.V(3).Infof("Reading files %v", files)
		tests, err = readFiles(files)
		return
	}

	tests, err = readFiles([]string{testfile})
	return
}

func readFiles(files []string) (tests testing.Tests, err error) {
	tests = testing.Tests{}
	for _, fname := range files {
		more := testing.Tests{}
		ext := filepath.Ext(fname)
		switch ext {
		case ".json":
			// Read json file
			var body []byte
			glog.V(3).Infof("Reading json test %s", fname)
			body, err = ioutil.ReadFile(fname)
			if err != nil {
				return
			}

			// Unmarshal json file
			err = json.Unmarshal(body, &more)
			if err != nil {
				return
			}
		case ".yaml":
			// Read yaml file
			var body []byte
			glog.V(3).Infof("Reading yaml test %s", fname)
			body, err = ioutil.ReadFile(fname)
			if err != nil {
				return
			}

			// Unmarshal yaml file
			err = yaml.Unmarshal(body, &more)
			if err != nil {
				return
			}
		default:
			glog.Warningf("Ignoring file %s (ext=%#v)", fname, ext)
		}
		tests = append(tests, more...)
	}
	return
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

var actionTestDesctiption = `Application action-test runs tests from JSON or YAML files against actions from different sources.`

var exampleTestYAML = `
  - 
    name: "filmweb.find"
    description: "Should find movie by title"
    arguments: 
      title: "Pulp Fiction"
    expect: 
      writers: "Quentin Tarantino"
      directors: "Quentin Tarantino"
      title: "Pulp Fiction"
      year: "1994"
`

func parseFlags() {
	var srcs string
	flag.StringVar(&testfile, "tests", "", "Files or directory containing YAML or JSON tests (can be glob pattern)")
	flag.StringVar(&srcs, "sources", "", "Actions sources (comma separated directories & urls)")
	flag.Parse()

	// Split comma separated sources into a list
	sources = strings.Split(srcs, ",")

	// Print help if --help flag or --test empty
	if testfile == "" {
		flag.Usage()
		os.Exit(0)
	}
}

// init - Adds usage to flag.Usage :)
func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of action-test:\n  %s\n", actionTestDesctiption)
		// fmt.Fprint(os.Stderr, actionTestDesctiption)
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Example YAML tests:\n")
		fmt.Fprint(os.Stderr, exampleTestYAML)
	}
}
