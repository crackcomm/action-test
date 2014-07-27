package testing

import (
	"fmt"
	"github.com/crackcomm/go-actions/action"
	"reflect"
	"strings"
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
		return fmt.Sprintf("√ %s => %s", v.Name, toString(v.Value))
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

	// Trim space from value
	res = strings.TrimSpace(res)

	// Cut value at 100 characters
	if len(res) >= MaxVariableLength {
		res = res[:MaxVariableLength] + "..."
	}
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
	default:
		ok = reflect.DeepEqual(a, b)
	}

	return
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
