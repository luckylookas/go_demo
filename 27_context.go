package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func wait(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "failed to finish in time."
	case <-time.After(time.Duration(rand.Intn(10)) * time.Second): //long running operation
		return "made it"
	}
}
/*
context is used to stop depending operations if the "partent" requests stopping.
eg.: a http request requires data from a DB and complex calculations,
if the requests is reset by peer, there is no point in bothering to read the results form the db
or in continuing the calculations when the database has already returned results
 */
func main() {
	//contexts always depend on other contexts, canceling a context also cancels any child contexts.

	//with timeouts
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	fmt.Println(wait(ctx))

	//with manually called cancelFunc
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	fmt.Println(wait(ctx))
}
