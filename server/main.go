package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jradhima/grpc-demo/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(ctx context.Context, cr *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	amount := cr.Amount.GetAmount()
	currency := cr.Amount.GetCurrency()
	from := cr.GetFrom()
	to := cr.GetTo()
	msg := fmt.Sprintf("Request accepted: transfered %d%s from %s to %s", amount, currency, from, to)
	pdf := []byte(msg + "\tPdf")
	docx := []byte(msg + "\tDocx")
	return &invoicer.CreateResponse{
		Pdf:  pdf,
		Docx: docx,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	serverRegisrtar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegisrtar, service)

	err = serverRegisrtar.Serve(lis)
	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}
