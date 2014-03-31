package main

import "github.com/go-martini/martini"

func main() {
	Open()
	defer Close()

	m := martini.Classic()
	m.Get("/", func() string {
		url := GetUrl(3)
		return "Hello rafał! " + url
	})
	m.Run()
}
