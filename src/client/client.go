package client

import (
	"context"
	"google.golang.org/grpc"

	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewQueryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.ExecuteQuery(ctx, &pb.QueryRequest{Query: "select iccid from aggregates limit 1;"})
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	log.Printf("Query Result: %v", r.Result)
}
