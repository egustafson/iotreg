package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/", rootHandler)
	log.Print("rootHandler registerd for /")
}

type InfoMsg struct {
	Msg string    `json:"msg"`
	Now time.Time `json:"now"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	msg := InfoMsg{
		Msg: "helo",
		Now: time.Now(),
	}
	b, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}
	w.Write(b)
}
