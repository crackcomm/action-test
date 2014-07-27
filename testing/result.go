package testing

import (
	"fmt"
	"github.com/crackcomm/go-actions/action"
	"time"
)

// Result - Test result.
type Result struct {
	Test      *Test         // test (parent)
	Passed    bool          // test passed?
	Start     time.Time     // test start
	Duration  time.Duration // test execution duration
	Error     error         // result error
	Context   action.Map    // result context
	Variables Variables
}

// NewResult - Creates a new test result.
func NewResult(test *Test) *Result {
	return &Result{
		Test:      test,
		Start:     time.Now(),
		Variables: Variables{},
	}
}

// End - Calculates result collection time and tests variables.
func (result *Result) End(ctx action.Map, err error) {
	result.Error = err
	result.Context = ctx
	result.Duration = time.Since(result.Start)

	for attr, expected := range result.Test.Expected {
		result.Variables[attr] = NewVariable(attr, nil, expected)
	}

	for attr, value := range result.Context {
		if _, has := result.Variables[attr]; has {
			result.Variables[attr].Value = value
		} else {
			result.Variables[attr] = NewVariable(attr, value, nil)
		}
	}

	result.Passed = result.Variables.Expected()
}

// Print - Print result information.
func (result *Result) Print() (passed bool) {
	fmt.Printf("=== RUN %s\n", result.Test.Name)
	fmt.Print("\n")
	fmt.Printf("  %s\n", result.Test.Description)
	fmt.Print("\n")

	for _, variable := range result.Variables {
		fmt.Printf("    %s\n", variable.String())
	}

	// Print result footer
	if result.Passed && result.Error == nil {
		result.printfoot("PASS")
	} else {
		result.printfoot("FAIL")
	}
	return
}

// printfoot - Prints foot information (state should be PASS/FAIL).
func (result *Result) printfoot(state string) {
	fmt.Print("\n")
	fmt.Printf("--- %s %s (%v)\n", state, result.Test.Name, result.Duration)
}
