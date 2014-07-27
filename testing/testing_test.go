package testing

import (
	"github.com/crackcomm/go-actions/action"
	. "github.com/smartystreets/goconvey/convey"
	. "testing"
)

func TestTestIsExpected(t *T) {
	Convey("Test.IsExpected()", t, func() {
		test := &Test{
			Expect: action.Map{"ok": true},
		}
		Convey("It should recognize unexpected value", func() {
			expected, ok := test.IsExpected("ok", false)
			So(ok, ShouldEqual, false)
			So(expected, ShouldEqual, true)
		})
		Convey("It should recognize expected value", func() {
			expected, ok := test.IsExpected("ok", true)
			So(ok, ShouldEqual, true)
			So(expected, ShouldEqual, true)
		})
	})
}

func TestTestIsExpectedResult(t *T) {
	Convey("Test.IsExpectedResult()", t, func() {
		test := &Test{
			Expect: action.Map{"ok": true},
		}
		Convey("It should recognize unexpected result", func() {
			expected, ok := test.IsExpectedResult(action.Map{"ok": false})
			So(ok, ShouldEqual, false)
			So(len(expected), ShouldEqual, 1)
			So(expected["ok"], ShouldEqual, true)
		})
		Convey("It should recognize expected value", func() {
			expected, ok := test.IsExpectedResult(action.Map{"ok": true})
			So(ok, ShouldEqual, true)
			So(len(expected), ShouldEqual, 0)
		})
	})
}
