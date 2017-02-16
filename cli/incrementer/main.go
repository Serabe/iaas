package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Serabe/iaas/iaas"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	str := os.Args[len(os.Args)-1]
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("%s is not a number, at least for me", str))
	}

	conn, err := grpc.Dial("localhost:5678", grpc.WithInsecure())
	if err != nil {
		panic("Oh no, the service is not available!")
	}

	client := iaas.NewIaasServiceClient(conn)
	result, err := client.Inc(context.Background(), &iaas.IncRequest{Arg1: num})
	if err != nil {
		panic("Oh no, I could not increment your number!")
	}

	fmt.Printf("%d", result.Result1)
}
