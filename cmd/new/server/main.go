package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"test/pkg/rpc/loggregator_v2"

	"google.golang.org/protobuf/encoding/protojson"
)

var (
	port = flag.Int("port", 8080, "port to listen on")
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
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("could not parse body")
		fmt.Fprintln(w, "error: could not parse body")
		return
	}

	var e loggregator_v2.Envelope
	err = protojson.Unmarshal(b, &e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("could not unmarshal body")
		fmt.Fprintln(w, "error: could not parse body")
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("unmarshaled event: %#v\n", e)

	b, err = protojson.Marshal(&e)
	if err != nil {
		log.Printf("error: %s", err)
	}

	_, err = w.Write(b)
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
