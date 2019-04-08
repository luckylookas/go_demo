package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/julienschmidt/sse"
	"log"
	"net/http"
)

var streamer sse.Streamer

type Payload struct {
	Name string `json: name`
	Age uint `json: age`
}

func Stream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	streamer.ServeHTTP(w,nil)
}

func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payload := new(Payload)
	json.NewDecoder(r.Body).Decode(payload)
	streamer.SendJSON("0", "person", *payload)
}

func main() {
	streamer = sse.Streamer{}
	router := httprouter.New()
	router.GET("/socket", Stream)
	router.POST("/socket", Post)

	log.Fatal(http.ListenAndServe(":8080", router))
}