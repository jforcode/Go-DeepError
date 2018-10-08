package deepErr

import (
	"fmt"
)

func ExampleSimpleFunction() {
	fun1 := func() error {
		return DeepErr{
			Function: "fun1",
			Action:   "doing nothing",
			Code:     "test",
			Message:  "no message",
			Params:   []interface{}{"param1", "param2"},
		}
	}

	err := fun1()
	err1 := err.(DeepErr)
	fmt.Println(err1)

	// Output:
	// IN fun1 WHILE doing nothing GOT test (no message): [param1 param2]
}

func ExampleNestedFunction() {
	fun1 := func() error {
		return DeepErr{
			Action: "returning error",
		}
	}

	fun2 := func() error {
		err := fun1()
		return DeepErr{
			Function: "fun2",
			Cause:    err,
		}
	}

	err := fun2()
	fmt.Println(err)

	// Output:
	// IN fun2 WHILE <NO_ACTION> GOT <NO_CODE> (<NO_MESSAGE>): <NO_PARAMS>
	// Nested Error: IN <NO_FUNCTION> WHILE returning error GOT <NO_CODE> (<NO_MESSAGE>): <NO_PARAMS>
}
