package testing

import (
	"github.com/crackcomm/go-actions/action"
	"github.com/crackcomm/go-actions/local"
	"reflect"
	"time"
)

// DefaultRunner - Test action default runner.
var DefaultRunner = local.DefaultRunner

// Test - Structure describing expected behaviour of action.
type Test struct {
	Name        string     `json:"name"`        // action name
	Context     action.Map `json:"ctx"`         // action context
	Description string     `json:"description"` // expected behaviour description
	Expect      action.Map `json:"expect"`      // expected action result values
}

// Tests - List of tests.
type Tests []*Test

// IsExpected - Checks if key and value was expected in action result.
func (t *Test) IsExpected(key string, value interface{}) (expected interface{}, ok bool) {
	// If expect is empty - treat all as expected
	if t.Expect == nil {
		ok = true
		return
	}

	// Get expected value
	var has bool
	expected, has = t.Expect[key]

	// If not in expected - ignore
	if !has {
		ok = true
		return
	}

	// Check if values are match
	ok = reflect.DeepEqual(value, expected)

	return
}

// IsExpectedResult - Checks if result is expected output of action.
func (t *Test) IsExpectedResult(result action.Map) (expected action.Map, ok bool) {
	expected = make(action.Map)
	ok = true

	// Iterate over all values in result
	// for key, value := range result {
	// 	if e, good := t.IsExpected(key, value); !good {
	// 		expected[key] = e
	// 		ok = false
	// 	}
	// }

	// Range over expectations and check match
	for attr, value := range t.Expect {
		if !reflect.DeepEqual(value, result.Get(attr).Interface()) {
			expected[attr] = value
			ok = false
		}
	}

	return
}

// Run - Runs action and returns result.
func (t *Test) Run() *Result {
	a := &action.Action{
		Name: t.Name,
		Ctx:  t.Context,
	}

	// Register start time
	start := time.Now()

	// Run action using default runner
	ctx, err := DefaultRunner.Run(a)

	// Calculate execution duration
	duration := time.Since(start)

	return &Result{
		Test:     t,
		Error:    err,
		Context:  ctx,
		Duration: duration,
	}
}

// Run - Runs all tests and returns results.
func (tests Tests) Run() (results Results) {
	results = Results{}
	for _, test := range tests {
		result := test.Run()
		results = append(results, result)
	}
	return
}
