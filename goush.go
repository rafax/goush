package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

func main() {
	InitDb()

	g := gin.Default()
	runtime.GOMAXPROCS(runtime.NumCPU())
	g.GET("/benchmark/:uid", func(c *gin.Context) {
		c.String(200, VisitUrl(c.Params.ByName("uid"), VisitData(c.Request)))
	})
	g.Run(":3000")
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
