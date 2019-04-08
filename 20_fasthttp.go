package main

import (
	"fmt"
	"github.com/valyala/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func Reply(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprintf(ctx, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := fasthttprouter.New()
	router.GET("/hello/:name", Reply)
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}