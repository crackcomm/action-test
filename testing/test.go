package testing

import (
	"io"
	"io/ioutil"

	"github.com/crackcomm/go-actions/action"
	"github.com/crackcomm/go-actions/local"
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
