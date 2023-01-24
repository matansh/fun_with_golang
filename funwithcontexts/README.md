# Fun with contexts
yay!

## Some context about contexts
So, what are contexts? \
Contexts wear a couple of different hats, and can be used to accomplish quite a lot of things
- A context is a key-value storage that propagates throughout the applications method execution stack
- A context is a lifecycle control mechanism
- A context is a high level synchronization mechanism

If you are thinking "contexts are a golang thing" then you are only partially correct \
Most languages that proved asynchronous functionalities have similar constructs, \
sometimes as a tread object, perhaps a process context and such.

What is unique about the approach golang has opted for is that it is an explicit, non global solution \
Which some would say aligns perfectly with golang's communities general aversion for implicitly and global vars.

Meaning that we (yes you) need to pass the context around and check its state. \
Simply canceling a context does nothing unless you explicitly handle the cancellation, \
**There is no magic mechanism that will stop your execution on context cancellation**

## Usage
### Creating a context
Creating a new context from scratch is as simple as 
```go
ctx := context.Background()
```
From the context we just created we can create a cancelable child context
```go
childCtx, cancel := context.WithCancel(ctx)
```
A child context can also be set to auto cancel within a specific time-frame or at a specific time
```go
// both will cancel in 10 seconds
context.WithTimeout(ctx, 10 * time.Second)
context.WithDeadline(ctx, time.Now().Add(10 * time.Second))
```

### Using the key value store
Note: context values should be used sparingly and only for contextual info like a requests tracking id
#### Writing
```go
parentCtx := context.Background()
childCtx = context.WithValue(ctx, "k", "v")
// parentCtx does not have key "k"
```
#### Reading
```go
value := ctx.Value("k")
```

### Checking a contexts state
#### Sync
A context will report on its current state when we call `.Err()` \
If it is not canceled the method will return `nil` \
We can differentiate between explicitly canceled contexts and automatically expired contexts using the error that returns
```go
err := ctx.Err()
if errors.Is(err, context.Canceled) {
    // explicitly canceled
}
if errors.Is(err, context.DeadlineExceeded) {
    // timed out
}
// still valid
```
#### Async
For the purposes of this exercise you don't need to know anything about async routines or channels - we will get to that in a later session. \
For now just accept this info "as is"

We can block our execution until the context is done
```go
<-ctx.Done()
// think of it like `await ctx.Done()`
```

