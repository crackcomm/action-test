package testing

import (
	"flag"
	"fmt"
	"time"
)

// Results - Tests results.
type Results struct {
	Duration time.Duration
	Passed   bool
	Start    time.Time
	List     []*Result
}

// verbose - Verbose `-v` flag, when true prints tests information.
var verbose bool

// NewResults - Creates a new results list.
func NewResults() *Results {
	return &Results{
		Passed: true,
		Start:  time.Now(),
		List:   []*Result{},
	}
}

// Add - Adds result to the list.
func (results *Results) Add(result *Result) {
	if !result.Passed || result.Error != nil {
		results.Passed = false
	}
	results.List = append(results.List, result)
}

// End - Calculate duration of results collection.
func (results *Results) End() {
	results.Duration = time.Since(results.Start)
}

// Print - Print results information.
func (results *Results) Print() {
	if verbose {
		for _, result := range results.List {
			result.Print()
		}
	}

	var state string
	if results.Passed {
		state = "OK"
	} else {
		state = "FAIL"
	}

	if !verbose {
		fmt.Println(state)
	} else {
		fmt.Printf("--- %s   %v", state, results.Duration)
	}
}

func init() {
	flag.BoolVar(&verbose, "p", false, "Verbose output: log all tests")
}
