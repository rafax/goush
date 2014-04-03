package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

func main() {
	InitDb()

	m := martini.Classic()
	m.Get("/benchmark/:uid", func(params martini.Params, req *http.Request) string {
		return VisitUrl(params["uid"], VisitData(req))
	})
	m.Run()
}

func VisitData(req *http.Request) []byte {
	json, _ := json.Marshal(map[string]interface{}{
		"time":    time.Now(),
		"values":  req.Form,
		"method":  req.Method,
		"url":     req.URL,
		"headers": req.Header,
		"ip":      req.RemoteAddr,
	})
	return json
}
