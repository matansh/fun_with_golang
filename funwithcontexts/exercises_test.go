package funwithcontexts_test

/*
Below you will find a couple of fun exercises to use what you have just learned

The exercises range from the easiest on the top to the most complex at the bottom
Feel free to pick an choose, you don't have to run through them all (unless you are having fun ;P )

Solutions can be found in `solutions_test.go`
You can run your solutions directly here as unit tests via `go test -v funwithcontexts/solutions_test.go`
Or in Golang's hosted playground: https://play.golang.com/
*/

import (
	"context"
	"testing"
)

func Test_Create_A_Context_And_Check_Its_State(t *testing.T) {
	// below is a variable deceleration that var `ctx` is a context
	var ctx context.Context
	// assign a value to var `ctx`
	ctx = nil

	// print out its state
	t.Log(ctx)
}

func Test_Cancel_A_Context(t *testing.T) {
	/*
		todo:
		- start a context
		- create a cancelable child context from it
		- cancel the child context
		- print its state
	*/
}

func Test_Synchronously_Wait_For_Context_Timeout(t *testing.T) {
	/*
		todo:
		- start a context
		- create a child context that will automatically cancel in 2 seconds
		- poll the context state until it is canceled
	*/
}

func Test_Asynchronously_Wait_For_Context_Timeout(t *testing.T) {
	/*
		todo:
		- start a context
		- create a child context that will automatically cancel in 2 seconds
		- spin up an async routine that will print the context's state on cancellation
	*/
}
