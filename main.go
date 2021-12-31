package main

import (
	"context"
	"fmt"

	"github.com/Pinablink/proto/conciergef"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("xxxxxxxxxx", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	client := conciergef.NewConciergefClient(connection)

	request := &conciergef.Hello{Status: true}

	res, mErr := client.CallTest(context.Background(), request)

	if mErr != nil {
		panic(mErr)
	}

	fmt.Println(res.Status)

}
