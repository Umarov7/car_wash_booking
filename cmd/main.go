package main

import (
	"booking-service/config"
	pbb "booking-service/genproto/bookings"
	pbpa "booking-service/genproto/payments"
	pbp "booking-service/genproto/providers"
	pbr "booking-service/genproto/reviews"
	pbs "booking-service/genproto/services"
	"booking-service/service"
	mongodb "booking-service/storage/mongoDB"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	db, err := mongodb.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to mongodb: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", cfg.BOOKING_SERVICE_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	pbp.RegisterProvidersServer(server, service.NewProviderService(db))
	pbs.RegisterServicesServer(server, service.NewServiceService(db))
	pbb.RegisterBookingsServer(server, service.NewBookingService(db))
	pbpa.RegisterPaymentsServer(server, service.NewPaymentService(db))
	pbr.RegisterReviewsServer(server, service.NewReviewService(db))

	log.Printf("Service is listening on port %s...\n", cfg.BOOKING_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving booking service: %v", err)
	}
}
