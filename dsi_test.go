package dsi

import (
	"context"
	"testing"
)

func Test_Traceme(t *testing.T) {
	patchconfig = []*patch{
		{
			What:   "test",
			Match:  "1",
			Return: map[string]interface{}{"res": true},
		},
		{
			What:  "test2",
			Match: "1",
			Patch: map[string]interface{}{"res": true},
		},
		{
			What:   "test3",
			Match:  "1",
			Patch:  map[string]interface{}{"res": true},
			Repeat: 3,
		},
	}

	var res, called bool

	res, called = false, false
	p := 1
	Traceme(context.Background(), "test", A{"a": &p}, func() { called = true }, A{"res": &res}, A{})
	if !res {
		t.Errorf("res was not set to true")
	}
	if called {
		t.Errorf("return should not call")
	}

	res, called = false, false
	Traceme(context.Background(), "test2", A{}, func() { called = true }, A{"res": &res}, A{})
	if !res {
		t.Errorf("res was not set to true")
	}
	if !called {
		t.Errorf("patch should call")
	}

	t.Run("repeats", func(t *testing.T) {
		Traceme(context.Background(), "test3", A{}, func() { called = true }, A{"res": &res}, A{})
		Traceme(context.Background(), "test3", A{}, func() { called = true }, A{"res": &res}, A{})

		res, called = false, false
		Traceme(context.Background(), "test3", A{}, func() { called = true }, A{"res": &res}, A{})
		if !res {
			t.Errorf("res was not set to true")
		}
		if !called {
			t.Errorf("patch should call")
		}

		res, called = false, false
		Traceme(context.Background(), "test3", A{}, func() { called = true }, A{"res": &res}, A{})
		if res {
			t.Errorf("res should not be patched again")
		}
		if !called {
			t.Errorf("after repetition should call")
		}
	})
}
