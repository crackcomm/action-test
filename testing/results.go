package testing

import (
	"fmt"
	"github.com/crackcomm/go-actions/action"
	"time"
)

// Result - Test result.
type Result struct {
	Test     *Test         // result test (parent)
	Context  action.Map    // result context
	Error    error         // result error
	Duration time.Duration // action execution time
}

// Results - Tests results.
type Results []*Result

// Print - Print result information.
func (result *Result) Print() {
	fmt.Printf("=== RUN %s\n", result.Test.Name)
	fmt.Print("\n")
	fmt.Printf("  %s\n", result.Test.Description)
	fmt.Print("\n")

	// Check expected
	expected, passed := result.Test.IsExpectedResult(result.Context)

	// Range over result context
	for name := range result.Context {
		value := result.Context.Get(name)
		var strvalue string
		if m, ok := value.Map(); ok {
			body, _ := m.JSON()
			strvalue = string(body)
		} else {
			strvalue = fmt.Sprintf("%v", value.Interface())
		}

		// Cut value at 100 characters
		if len(strvalue) >= 100 {
			strvalue = strvalue[:100] + "..."
		}

		// If is in expected it means we got unexpected value
		if exp, didexpected := expected[name]; !didexpected {
			fmt.Printf("    √ %s => %s\n", name, strvalue)
		} else {
			fmt.Printf("    × %s => %s (expected %v)\n", name, strvalue, exp)
		}
	}

	// Print error if any
	if err := result.Error; err != nil || !passed {
		if err != nil {
			fmt.Printf("    ERROR: %v\n", result.Error)
		}
		result.printfoot("FAIL")
		return
	}

	result.printfoot("PASS")
	// fmt.Print("\n")
}

// printfoot - Prints foot information (state should be PASS/FAIL).
func (result *Result) printfoot(state string) {
	fmt.Print("\n")
	fmt.Printf("--- %s: %s (%v)\n", state, result.Test.Name, result.Duration)
}

// Print - Print results information.
func (results Results) Print() {
	for _, result := range results {
		result.Print()
	}
}
