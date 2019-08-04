package main

import (
	"net/http"
	"time"
)

func (obj *Host) StartWebInterface() {
	h := http.NewServeMux()
	h.HandleFunc("/jquery.js", WebJquery)
	h.HandleFunc("/style.css", WebCSS)
	h.HandleFunc("/", obj.WebSettings)
	h.HandleFunc("/win", obj.WebWin)
	h.HandleFunc("/nonwin", obj.WebNonWin)

	for {
		http.ListenAndServe("192.168.10.1:80", h)
		time.Sleep(time.Second)
	}

}
