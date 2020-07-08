package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response header for data
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ResponseWrite app
func ResponseWrite(r *Response, w http.ResponseWriter) {
	byteOfResponse, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Opps, Something Wrong"))
	}
	w.Write([]byte(byteOfResponse))
}
