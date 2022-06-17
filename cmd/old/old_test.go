package main_test

import (
	"io"
	"testing"

	"code.cloudfoundry.org/go-loggregator/v8/rpc/loggregator_v2"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

var (
	marshaler = jsonpb.Marshaler{}

	oldEnvelope = &loggregator_v2.Envelope{
		SourceId:   "my-source-id",
		InstanceId: "my-instance-id",
		Message:    &loggregator_v2.Envelope_Log{},
	}
)

// Benchmark old proto marshalling
func BenchmarkOldProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := marshaler.Marshal(io.Discard, oldEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark old proto gogo marshalling
func BenchmarkOldProtoGogoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := gogoproto.Marshal(oldEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark old proto unmarshalling
func BenchmarkOldProtoUnmarshal(b *testing.B) {
	buf := &buffer{}
	err := marshaler.Marshal(buf, oldEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e loggregator_v2.Envelope
	for i := 0; i < b.N; i++ {
		err := jsonpb.Unmarshal(buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark old proto gogo unmarshalling
func BenchmarkOldProtoGogoUnmarshal(b *testing.B) {
	buf, err := gogoproto.Marshal(oldEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e loggregator_v2.Envelope
	for i := 0; i < b.N; i++ {
		err := gogoproto.Unmarshal(buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

type buffer struct {
	buf []byte
}

func (b *buffer) Write(p []byte) (n int, err error) {
	b.buf = make([]byte, len(p))
	return copy(b.buf, p), nil
}

func (b *buffer) Read(p []byte) (n int, err error) {
	n = copy(p, b.buf)
	return n, nil
}
