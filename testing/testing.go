package testing

import (
	"github.com/crackcomm/go-actions/action"
	"github.com/crackcomm/go-actions/local"
	"io/ioutil"
	"io"
)

// DefaultRunner - Test action default runner.
var DefaultRunner = local.DefaultRunner

// Test - Structure describing expected behaviour of action.
type Test struct {
	Name        string     `json:"name" yaml:"name"`               // action name
	Context     action.Map `json:"ctx" yaml:"ctx"`                 // action context
	Description string     `json:"description" yaml:"description"` // expected behaviour description
	Expected    action.Map `json:"expect" yaml:"expect"`           // expected action result values
}

// Tests - List of tests.
type Tests []*Test

// Run - Runs action and returns result.
func (t *Test) Run() *Result {
	a := &action.Action{
		Name: t.Name,
		Ctx:  t.Context,
	}

	// Create new result
	result := NewResult(t)

	// Run action using default runner
	ctx, err := DefaultRunner.Run(a)

	for key, value := range ctx {
		if rc, ok := value.(io.ReadCloser); ok {
			arr, _ := ioutil.ReadAll(rc)
			rc.Close()
			ctx[key] = arr
		}
	}

	ctx.Close()

	// Result collection end
	result.End(ctx, err)

	return result
}

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
