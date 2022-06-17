package main_test

import (
	"test/pkg/rpc/loggregator_v2"
	"testing"

	"google.golang.org/protobuf/proto"
)

var (
	newEnvelope = &loggregator_v2.Envelope{
		SourceId:   "my-source-id",
		InstanceId: "my-instance-id",
		Message:    &loggregator_v2.Envelope_Log{},
	}
)

// Benchmark new proto marshalling
func BenchmarkNewProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(newEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// TODO: fix this
// Benchmark new proto gogo marshalling
// func BenchmarkNewProtoGogoMarshal(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_, err := gogoproto.Marshal(newEnvelope)
// 		if err != nil {
// 			b.Fatal("Marshaling error:", err)
// 		}
// 	}
// }

// Benchmark new proto unmarshalling
func BenchmarkNewProtoUnmarshal(b *testing.B) {
	buf, err := proto.Marshal(newEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e loggregator_v2.Envelope
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// TODO: fix this
// Benchmark new proto gogo unmarshalling
// func BenchmarkNewProtoGogoUnmarshal(b *testing.B) {
// 	buf, err := gogoproto.Marshal(newEnvelope)
// 	if err != nil {
// 		b.Fatal("Marshaling error:", err)
// 	}

// 	var e events.Envelope
// 	for i := 0; i < b.N; i++ {
// 		err := gogoproto.Unmarshal(buf, &e)
// 		if err != nil {
// 			b.Fatal("Unmarshaling error:", err)
// 		}
// 	}
// }
