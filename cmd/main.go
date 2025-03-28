package main

import (
	"log"
	"net"

	"coupon-issuance-system.com/coupon-issuance-system/internal/service"
	pb "coupon-issuance-system.com/coupon-issuance-system/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterCouponServiceServer(server, service.GetService())

	log.Println("Starting server on :50051...")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
