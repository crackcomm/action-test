package testing

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/crackcomm/go-actions/action"
)

// Variables - Variables by name.
type Variables map[string]*Variable

// Variable - Context variable structure.
type Variable struct {
	Name     string
	Value    interface{}
	Expected interface{}
	Passed   bool
}

// MaxVariableLength - Max length of Variable value.
var MaxVariableLength = 100

// NewVariable - Creates a new variable.
func NewVariable(name string, value, expected interface{}) *Variable {
	return &Variable{Name: name, Value: value, Expected: expected}
}

// IsExpected - Checks if value was expected.
func (v *Variable) IsExpected() bool {
	// Return true if no expectations at all
	if v.Expected == nil {
		return true
	}

	// If there was expectations and no value - false.
	if v.Value == nil {
		return false
	}

	return compareValues(v.Value, v.Expected)
}

// String - Returns string for console printing.
func (v *Variable) String() string {
	if v.IsExpected() {
		symbol := "√"
		if v.Expected == nil {
			symbol = "?"
		}
		return fmt.Sprintf("%s %s => %s", symbol, v.Name, toString(v.Value))
	}

	// If not expected because value was nil
	if v.Value == nil {
		return fmt.Sprintf("× %s is empty (expected %v)", v.Name, toString(v.Expected))
	}

	return fmt.Sprintf("× %s => %s (expected %v)", v.Name, toString(v.Value), toString(v.Expected))
}

// Expected - Checks if all variables were expected.
func (variables Variables) Expected() (ok bool) {
	ok = true
	for _, variable := range variables {
		if expected := variable.IsExpected(); !expected {
			ok = false
		}
	}
	return
}

// toString - Formats interface to string and cuts to `MaxVariableLength`.
func toString(v interface{}) (res string) {
	value := action.Format{v}

	if m, ok := value.Map(); ok {
		body, _ := m.JSON()
		res = string(body)
	} else if b, ok := value.Bytes(); ok {
		res = string(b)
	} else {
		res = fmt.Sprintf("%v", value.Interface())
	}

	// Cut value at 100 characters
	if len(res) >= MaxVariableLength {
		res = res[:MaxVariableLength] + "..."
	}

	// Break by new line and trim every line
	var lines []string
	for _, line := range strings.Split(res, "\n") {
		lines = append(lines, strings.TrimSpace(line))
	}

	// Join to one line
	res = strings.Join(lines, "")

	return
}

// compareValues - Compares two values - returns true if they are equal.
func compareValues(a, b interface{}) (ok bool) {
	// Check if values are match
	switch a.(type) {
	case string, []byte:
		// Dumb strings comparator
		ok = dumbStringsCompare(a, b)
	case float32, float64,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		// Dumb numbers comparator
		ok = dumbNumbersCompare(a, b)
	case action.Map, map[string]interface{}, map[interface{}]interface{}:
		ok = mapCompare(a, b)
	default:
		ok = reflect.DeepEqual(a, b)
	}

	return
}

// Comparing maps, tough - it's bad implementation
func mapCompare(a, b interface{}) bool {
	am, _ := action.Format{a}.Map()
	bm, _ := action.Format{b}.Map()
	if len(am) != len(bm) {
		return false
	}
	for key, value := range am {
		if !compareValues(value, bm[key]) {
			return false
		}
	}
	return true
}

// dumbStringsCompare - Dumb strings comparator.
func dumbStringsCompare(a, b interface{}) bool {
	stra := fmt.Sprintf("%s", a)
	strb := fmt.Sprintf("%s", b)
	return stra == strb
}

// dumbNumbersCompare - Dumb numbers comparator.
func dumbNumbersCompare(a, b interface{}) bool {
	stra := fmt.Sprintf("%v", a)
	strb := fmt.Sprintf("%v", b)
	return stra == strb
}
