Exposed - minimal high performant reflectionless RPC Server

## Features

the following features are currently available:

- requests with timeout/cancellation 

- codecs - specify your own codec to marshal/unmarshal messages

- [**`expose`**](https://github.com/thesyncim/expose) - codegen utility to generate exposed services from interface definitions

### Usage

example server:

```go
package main

import (
        "log"
        "net"

        "github.com/thesyncim/exposed"
        "github.com/thesyncim/exposed/encoding/codec/json"
)

func main() {
        ln, err := net.Listen("tcp", "127.0.0.1:8888")
        if err != nil {
                panic(err)
        }

        server := exposed.NewServer(
                exposed.ServerCodec(json.CodecName),
                //exposed.ServerCompression(exposed.CompressSnappy),
        )

        server.HandleFunc("echo",
                func(ctx *exposed.Context, req exposed.Message, resp exposed.Message) (err error) {
                        resp.(*string) = req.(*string)
                        return nil
                },
                &exposed.OperationTypes{
                        ReplyType: func() exposed.Message {
                                return new(string)
                        },
                        ArgsType: func() exposed.Message {
                                return new(string)
                        },
                })

        log.Fatalln(server.Serve(ln))
}
```

example client:
```go
package main

import (
        "fmt"

        "github.com/thesyncim/exposed"
        "github.com/thesyncim/exposed/encoding/codec/json"
)

func main() {
        client := exposed.NewClient("127.0.0.1:8888",
                exposed.ClientCodec(json.CodecName),
                //exposed.ServerCompression(exposed.CompressSnappy),
        )

        var req = "ping"
        var resp string

        err := client.Call("echo", &req, &resp)
        if err != nil {
                panic(err)
        }

        fmt.Println(resp)
}
```
### Generate service from interface definition

lets looks at the example

```go
package echo

type Echoer interface {
	Echo(msg []byte) (ret []byte)
}

type Echo struct {
}

func (Echo) Echo(msg []byte) []byte {
	return msg
}

```

```sh
go get github.com/thesyncim/expose
```

download and install *expose*.
 A codegen tool to generate an exposed service from your interface definition

 ```sh 
 expose gen -i  Echoer -p github.com/thesyncim/exposed/examples/echo -s echoservice -o echoservice
 ```
 this will generate all the [boilerplate code](https://github.com/thesyncim/exposed/tree/master/examples/echo/echoservice) to expose your interface as an service 
 