package main

import (
	"github.com/jiangjiali/exposed"
	"github.com/jiangjiali/exposed/examples/echo"
	"github.com/jiangjiali/exposed/examples/echo/echoservice"
	"github.com/jiangjiali/exposed/examples/echo/ecodec"
	"log"
	"net"
)

func main() {

	log.SetFlags(log.Lshortfile)
	ln, err := net.Listen("tcp", "127.0.0.1:5555")
	if err != nil {
		panic(err)
	}
	s := exposed.NewServer(exposed.ServerCodec(ecodec.CodecName), exposed.ServerCompression(exposed.CompressNone))
	if err != nil {
		panic(err)
	}

	simpleService := echoservice.NewServer(echo.Echo{})
	s.RegisterService(simpleService)

	log.Print(s.Serve(ln))

}
