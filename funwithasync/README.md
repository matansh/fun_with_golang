# Fun with asynchronous synchronization
Without order there is only chaos

## First some concurrency syntax
In order to spin up a concurrent routine we simply add a `go` statement before the method call we want to run async
```go
// will block us for 1 second
time.Sleep(time.Second)

// wont block us at all
go time.Sleep(time.Second)
```

## why sync at all?
Lets look at an example 