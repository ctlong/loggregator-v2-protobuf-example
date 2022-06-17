package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"code.cloudfoundry.org/go-loggregator/v8/rpc/loggregator_v2"
	"github.com/golang/protobuf/jsonpb"
)

var (
	port      = flag.Int("port", 8080, "port to listen on")
	marshaler = jsonpb.Marshaler{}
)

func init() {
	log.SetPrefix("old server: ")
}

func envelopeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("no body provided")
		fmt.Fprintln(w, "error: proper request body must be provided")
		return
	}

	var e loggregator_v2.Envelope
	err := jsonpb.Unmarshal(r.Body, &e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("could not unmarshal body")
		fmt.Fprintln(w, "error: could not parse body")
		return
	}
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	log.Printf("unmarshaled event: %#v\n", e)
	err = marshaler.Marshal(w, &e)
	if err != nil {
		log.Printf("error: %s", err)
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/envelope", envelopeHandler)

	portStr := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
