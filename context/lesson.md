The Context struct type has 4 fields

Done() - a method that returns a receive-only channel
Error() - method that returns an error message
Value() - method that returns an a value asserted as an interface
Deadline() - method that returns a time and a bool

Context can be chained together an pass a valued information that propagate from the parent to the child

Problem Statement:
Let’s say that you started a function and you need to pass some common parameters to the downstream functions. You cannot pass these common parameters each as an argument to all the downstream functions.
You started a goroutine which in turn start more goroutines and so on. Suppose the task that you were doing is no longer needed. Then how to inform all child goroutines to gracefully exit so that resources can be freed up
A task should be finished within a specified timeout of say 2 seconds. If not it should gracefully exit or return.
A task should be finished within a deadline eg it should end before 5 pm . If not finished then it should gracefully exit and return

Using Context Package in GO (Golang) – Complete Guide
Posted on December 23, 2019by admin
Introduction:
Definition:
Context is a package provided by GO. Let’s first understand some problems that existed already, and which context package tries to solve.

Problem Statement:
Let’s say that you started a function and you need to pass some common parameters to the downstream functions. You cannot pass these common parameters each as an argument to all the downstream functions.
You started a goroutine which in turn start more goroutines and so on. Suppose the task that you were doing is no longer needed. Then how to inform all child goroutines to gracefully exit so that resources can be freed up
A task should be finished within a specified timeout of say 2 seconds. If not it should gracefully exit or return.
A task should be finished within a deadline eg it should end before 5 pm . If not finished then it should gracefully exit and return
If you notice all the above problems are quite applicable to HTTP requests and but none the less these problems are also applicable to many different areas too.

For a web HTTP request, it needs to be canceled when the client has disconnected, or the request has to be finished within a specified timeout and also requests scope values such as request_id needs to be available to all downstream functions.

When to Use (Some Use Cases):
To pass data to the downstream. Eg. a HTTP request creates a request_id, request_user which needs to be passed around to all downstream functions for distributed tracing.
When you want to halt the operation in the midway – A HTTP request should be stopped because the client disconnected
When you want to halt the operation within a specified time from start i.e with timeout – Eg- a HTTP request should be completed in 2 sec or else should be aborted.
When you want to halt an operation before a certain time – Eg. A cron is running that needs to be aborted in 5 mins if not completed.

type Context interface {
//It retures a channel when a context is cancelled, timesout (either when deadline is reached or timeout time has finished)
Done() <-chan struct{}

    //Err will tell why this context was cancelled. A context is cancelled in three scenarios.
    // 1. With explicit cancellation signal
    // 2. Timeout is reached
    // 3. Deadline is reached
    Err() error

    //Used for handling deallines and timeouts
    Deadline() (deadline time.Time, ok bool)

    //Used for passing request scope values
    Value(key interface{}) interface{}

}


Resources
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html#:~:text=The%20Context%20package%20defines%20an,you%20will%20write%20in%20Go.
https://golangbyexample.com/using-context-in-golang-complete-guide/

