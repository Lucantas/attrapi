package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//response is a basic struct cotaining a name and statuscode to an error
type response struct {
	Message string
}

func makeResponse(n string, w http.ResponseWriter, statusCode int) {
	err := response{Message: n}
	log.Println(err)
	e := json.NewEncoder(w)
	e.Encode(err)
}

func errorResponse(n string, e *json.Encoder) {
	err := response{Message: n}
	log.Println(err)
	e.Encode(err)
}

func successResponse(n string, e *json.Encoder) {
	err := response{Message: n}
	log.Println(err)
	e.Encode(err)
}
