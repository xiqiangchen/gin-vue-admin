package main

import (
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	var rtb = new(BidRequest)
	proto.SetExtension(rtb, E_B, true)
	var A bool
	var err error
	A, err = proto.GetExtension(rtb, E_B)
	log.Printf("A:%t,err:%v", A, err)
}
