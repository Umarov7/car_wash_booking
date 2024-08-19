package main

import (
	"booking-service/config"
	pbb "booking-service/genproto/bookings"
	pbn "booking-service/genproto/notifications"
	pbpa "booking-service/genproto/payments"
	pbp "booking-service/genproto/providers"
	pbr "booking-service/genproto/reviews"
	pbs "booking-service/genproto/services"
	"booking-service/kafka"
	"booking-service/kafka/consumer"
	"booking-service/pkg/db"
	"booking-service/service"
	mongodb "booking-service/storage/mongoDB"
	"booking-service/storage/redis"
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	time.Sleep(15 * time.Second)

	cfg := config.Load()

	mongo, err := mongodb.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to mongoDB: %v", err)
	}
	defer mongo.Close()

	redis, err := redis.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to redis: %v", err)
	}
	defer redis.Close()

	if err := db.SeedData(context.Background(), cfg); err != nil {
		log.Fatalf("error while initializing mongoDB: %v", err)
	}

	lis, err := net.Listen("tcp", cfg.BOOKING_SERVICE_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	p := service.NewProviderService(mongo)
	s := service.NewServiceService(mongo, redis)
	b := service.NewBookingService(mongo, redis)
	pay := service.NewPaymentService(mongo)
	r := service.NewReviewService(mongo)
	n := service.NewNotificationService(mongo)
	server := grpc.NewServer()

	pbp.RegisterProvidersServer(server, p)
	pbs.RegisterServicesServer(server, s)
	pbb.RegisterBookingsServer(server, b)
	pbpa.RegisterPaymentsServer(server, pay)
	pbr.RegisterReviewsServer(server, r)
	pbn.RegisterNotificationsServer(server, n)

	consumers := map[string]func([]byte){
		cfg.KAFKA_TOPIC_BOOKING_CREATED:      kafka.ConsumeCreateBooking(cfg, b),
		cfg.KAFKA_TOPIC_BOOKING_UPDATED:      kafka.ConsumeUpdateBooking(cfg, b),
		cfg.KAFKA_TOPIC_BOOKING_CANCELLED:    kafka.ConsumeCancelBooking(cfg, b),
		cfg.KAFKA_TOPIC_PAYMENT_CREATED:      kafka.ConsumeCreatePayment(cfg, pay),
		cfg.KAFKA_TOPIC_REVIEW_CREATED:       kafka.ConsumeCreateReview(cfg, r),
		cfg.KAFKA_TOPIC_NOTIFICATION_CREATED: kafka.ConsumeCreateNotification(cfg, n),
	}

	for topic, handler := range consumers {
		consumer := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, topic)
		go func(t string, h func([]byte)) {
			log.Printf("Starting consumer for topic: %s", t)
			consumer.Consume(h)
		}(topic, handler)
	}

	log.Printf("Service is listening on port %s...\n", cfg.BOOKING_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving booking service: %v", err)
	}
}
