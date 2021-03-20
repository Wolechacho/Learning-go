package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	//contextWithValue()
	// helloWorldHandler := http.HandlerFunc(HelloWorld)
	// http.Handle("/welcome", injectMsgId(helloWorldHandler))
	// http.ListenAndServe(":8083", nil)

	//context with cancel
	// rootCtx := context.Background()
	// canCtx, cancelFunc := context.WithCancel(rootCtx)
	// go task(canCtx)
	// time.Sleep(time.Second * 3)
	//explicitly cancellation you do not need a defer
	// cancelFunc()
	// time.Sleep(time.Second * 1)

	//cancel with timeout
	// ctx := context.Background()
	// cancelCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	// defer cancel()
	// go task(cancelCtx)
	// time.Sleep(time.Second * 4)

	//cancel with deadline
	// ctx := context.Background()
	// cancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
	// defer cancel()
	// go task(cancelCtx)
	// time.Sleep(time.Second * 6)

	callApi()
}

//understanding context_with_value
func contextWithValue() {
	rootCtx := context.Background()
	var id interface{} = "id"
	childCtx := context.WithValue(rootCtx, id, 2)
	childOfChildCtx, _ := context.WithCancel(childCtx)

	fmt.Println(rootCtx)
	fmt.Println(childCtx)
	fmt.Println(childOfChildCtx.Value("id"))

	//trying to get the value of root context print out <nil>
	fmt.Println(rootCtx.Value("id"))

}

//understanding context_with_value -- second part
func injectMsgId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := uuid.New().String()
		var msgid interface{} = "msgId"
		ctx := context.WithValue(r.Context(), msgid, u)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

//HelloWorld hellow world handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
}

//understanding context
func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}

	}
}

func callApi() {

	// Create a new request.
	req, err := http.NewRequest("GET", "https://traf.nibss-plc.com.ng:7443/traf/ajax?command=website&action=detail&order=loadNIP&clientCode=NIBSS&txnSubCat=ALL", nil)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Create a context with a timeout of 50 milliseconds.
	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	// Bind the new context into the request.
	req = req.WithContext(ctx)

	// Make the web call and return any error. Do will handle the
	// context level timeout.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Close the response body on the return.
	defer resp.Body.Close()

	// Write the response to stdout.
	io.Copy(os.Stdout, resp.Body)
}
