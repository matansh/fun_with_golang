package funwithcontexts

import (
	"context"
	"testing"
	"time"
)

/*
Below are solutions for the exercises in `exercises_test.go`
*/

func Test_Create_A_Context_And_Check_Its_State(t *testing.T) {
	// below is a variable deceleration that var `ctx` is a context
	var ctx context.Context
	// assign a value to var `ctx`
	ctx = context.Background()

	// print out its state
	if err := ctx.Err(); err != nil {
		t.Log("context is canceled for reason: ", err)
		return
	}
	t.Log("context is valid")
}

func Test_Cancel_A_Context(t *testing.T) {
	/*
		todo:
		- start a context
		- create a cancelable child context from it
		- cancel the child context
		- print its state
	*/
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	t.Log("context is canceled for reason: ", ctx.Err())
}

func Test_Synchronously_Wait_For_Context_Timeout(t *testing.T) {
	/*
		todo:
		- start a context
		- create a child context that will automatically cancel in 2 seconds
		- poll the context state until it is canceled
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	for {
		if err := ctx.Err(); err != nil {
			break
		}
		time.Sleep(time.Second / 2) // 0.5s
	}

	t.Log("context is canceled for reason: ", ctx.Err())
}

func Test_Asynchronously_Wait_For_Context_Timeout(t *testing.T) {
	/*
		todo:
		- start a context
		- create a child context that will automatically cancel in 2 seconds
		- spin up an async routine that will print the context's state on cancellation
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	t.Cleanup(cancel)

	<-ctx.Done()

	t.Log("context is canceled for reason: ", ctx.Err())
}
