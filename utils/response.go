package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response header for data
type Response struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	TotalField int         `json:"totalField"`
	Data       interface{} `json:"result"`
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
