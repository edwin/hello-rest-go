package main

import (
	"fmt"
	"log"
	"net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

			tracer.Start(tracer.WithServiceName("my-service"))
			http.HandleFunc("/", handle)
			http.HandleFunc("/_ah/health", healthCheckHandler)
			log.Print("Listening on port 8099")
			log.Fatal(http.ListenAndServe(":8099", nil))

			defer tracer.Stop()
}

func handle(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, "Hello world!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")

}
