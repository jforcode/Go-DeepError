package deepError

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	fun1 := func() error {
		return New("fun1", "doing nothing", errors.New("some dummy nested error"))
	}

	err := fun1()
	actual := err.Error()
	expected := "IN fun1 WHILE doing nothing \nNested Error: some dummy nested error"

	if actual != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestNewFull(t *testing.T) {
	fun1 := func() error {
		return NewFull("fun1", "doing nothing", errors.New("some dummy nested error"), "SOME_CODE", "dummy message with %s", []interface{}{"some param"})
	}

	err := fun1()
	actual := err.Error()
	expected := `IN fun1 WHILE doing nothing GOT SOME_CODE (dummy message with some param)
Nested Error: some dummy nested error`

	if actual != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestNoMessageButParams(t *testing.T) {
	fun1 := func() error {
		return NewFull("fun1", "doing nothing", errors.New("some dummy nested error"), "NO_INTERFACE_FOUND_ERROR", "", []interface{}{"some param"})
	}

	err := fun1()
	actual := err.Error()
	expected := `IN fun1 WHILE doing nothing GOT NO_INTERFACE_FOUND_ERROR [some param]
Nested Error: some dummy nested error`

	if actual != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

func TestRandomError(t *testing.T) {
	fun1 := func() error {
		return NewFull("", "", errors.New("some dummy nested error"), "SOME_CODE", "dummy message with %s", []interface{}{"some param"})
	}

	err := fun1()
	actual := err.Error()
	expected := `GOT SOME_CODE (dummy message with some param)
Nested Error: some dummy nested error`

	if actual != expected {
		t.Fatalf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}
