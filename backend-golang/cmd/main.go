package main

import (
	"demo-application-opa-golang/data"
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

func main() {
	store := data.Store{}
	store.Initialise()


	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
