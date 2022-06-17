package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"test/pkg/rpc/loggregator_v2"

	"google.golang.org/protobuf/encoding/protojson"
)

var (
	port = flag.Int("port", 8080, "port to send messages to")
)

func init() {
	log.SetPrefix("old client: ")
}

func makeEnvelope() *loggregator_v2.Envelope {
	return &loggregator_v2.Envelope{
		SourceId:   "my-source-id",
		InstanceId: "my-instance-id",
		Message:    &loggregator_v2.Envelope_Log{},
	}
}

func main() {
	flag.Parse()

	url := fmt.Sprintf("http://0.0.0.0:%d/envelope", *port)

	e := makeEnvelope()
	b, err := protojson.Marshal(e)
	if err != nil {
		log.Panicf("error marshaling: %s\n", err)
	}

	br := bytes.NewBuffer(b)
	resp, err := http.Post(url, "application/json", br)
	if err != nil {
		log.Panicf("error making request: %s\n", err)
	}

	log.Printf("status code: %d\n", resp.StatusCode)

	if resp.Body == nil {
		log.Panicln("error: no body provided")
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("could not parse body %s\n", err)
		return
	}

	var respE loggregator_v2.Envelope
	err = protojson.Unmarshal(b, &respE)
	if err != nil {
		log.Printf("error unmarshaling body: %s\n", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("unmarshaled response envelope: %#v\n", respE)
}
