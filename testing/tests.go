package testing

import (
	"github.com/crackcomm/go-actions/action"
	"github.com/crackcomm/go-actions/local"
	"io"
	"io/ioutil"
)

// Tests - List of tests.
type Tests []*Test

// Run - Runs all tests and returns results.
func (tests Tests) Run() (results *Results) {
	results = NewResults()
	for _, test := range tests {
		result := test.Run()
		results.Add(result)
	}
	results.End()
	return
}
