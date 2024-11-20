package main

import (
	"io/ioutil"
	"testing"

	pb "github.com/brotherlogic/system/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

func TestLoad(t *testing.T) {
	// Validate the existing schema
	file, err := ioutil.ReadFile("base.textproto")
	if err != nil {
		t.Fatalf("Bad read: %v", err)
	}

	system := &pb.Bootstrap{}
	err = prototext.Unmarshal(file, system)
	if err != nil {
		t.Fatalf("Cannot parse file: %v", err)
	}
}
