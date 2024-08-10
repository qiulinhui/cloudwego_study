package main

import (
	"cloudwego_study/demo/demo_proto/kitex_gen/pbapi"
	"cloudwego_study/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("localhost:8500")

	if err != nil {
		log.Fatal(err)
	}

	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r))

	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v, %v", res, err)
}
