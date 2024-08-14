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
	"booking-service/service"
	mongodb "booking-service/storage/mongoDB"
	"booking-service/storage/redis"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
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

	// if err := db.SeedData(context.Background(), cfg); err != nil {
	// 	log.Fatalf("error while initializing mongoDB: %v", err)
	// }

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

	consumer1 := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC_BOOKING_CREATED)
	consumer2 := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC_BOOKING_UPDATED)
	consumer3 := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC_BOOKING_CANCELLED)
	consumer4 := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC_PAYMENT_CREATED)
	consumer5 := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC_REVIEW_CREATED)
	consumer6 := consumer.NewKafkaConsumer([]string{cfg.KAFKA_HOST, cfg.KAFKA_PORT}, cfg.KAFKA_TOPIC_NOTIFICATION_CREATED)

	go consumer1.Consume(kafka.ConsumeCreateBooking(cfg, b))
	go consumer2.Consume(kafka.ConsumeUpdateBooking(cfg, b))
	go consumer3.Consume(kafka.ConsumeCancelBooking(cfg, b))
	go consumer4.Consume(kafka.ConsumeCreatePayment(cfg, pay))
	go consumer5.Consume(kafka.ConsumeCreateReview(cfg, r))
	go consumer6.Consume(kafka.ConsumeCreateNotification(cfg, n))

	log.Printf("Service is listening on port %s...\n", cfg.BOOKING_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving booking service: %v", err)
	}
}
