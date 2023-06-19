package main

import (
	"context"
	"log"
	"time"

	"github.com/jradhima/grpc-demo/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// printRequestResult gets the feature for the given point.
func printRequestResult(client invoicer.InvoicerClient, req *invoicer.CreateRequest) {
	log.Println("Sending request")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("client.Create failed: %v", err)
	}
	log.Printf("Response pdf: %s", string(resp.Pdf))
	log.Printf("Response docx: %s", string(resp.Docx))
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:8089", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := invoicer.NewInvoicerClient(conn)
	req := invoicer.CreateRequest{
		Amount: &invoicer.Amount{Amount: 100, Currency: "E"},
		From:   "yannis",
		To:     "ardalun",
	}

	printRequestResult(client, &req)
}
