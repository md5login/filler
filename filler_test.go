package filler

import (
	"errors"
	"testing"
)

var demoFiller = Filler{
	Tag: "demoFiller1",
	Fn: func(obj interface{}) (interface{}, error) {
		return "hello", nil
	},
}

var errDemoFiller = Filler{
	Tag: "demoFiller1",
	Fn: func(obj interface{}) (interface{}, error) {
		return nil, errors.New("some error")
	},
}

type demoStruct struct {
	Name           string  `fill:"demoFiller1:Val"`
	Val            string  `fill:"demoFiller2"`
	Ignore1        string  `fill:"-"`
	Ignore2        string  `fill:""`
	DefaultString  string  `defaults:"DefaultValue"`
	DefaultInt     int     `defaults:"5"`
	DefaultFloat64 float64 `defaults:"42.6"`
	DefaultBool    bool    `defaults:"true"`
}

// RegFiller - register new filler into []fillers
func TestRegFiller(t *testing.T) {
	RegFiller(demoFiller)
	v1, err1 := fillers[0].Fn("hello")
	v2, err2 := demoFiller.Fn("hello")
	if fillers[0].Tag != demoFiller.Tag || v1 != v2 || err1 != err2 {
		t.FailNow()
	}
}

// Fill - fill the object with all the current fillers
func TestFill(t *testing.T) {
	RegFiller(demoFiller)
	RegFiller(errDemoFiller)
	m := demoStruct{
		Name: "nameVal",
		Val:  "valVal",
	}
	// check non ptr - should panic
	func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		Fill(m)
		t.FailNow()
	}()
	// check if got filled
	Fill(&m)
	Defaults(&m)
	// should be filled
	if m.Name != "hello" || m.Val != "valVal" || m.DefaultString != "DefaultValue" || m.DefaultInt != 5 || m.DefaultFloat64 != 42.6 || m.DefaultBool != true {
		t.FailNow()
	}
}
