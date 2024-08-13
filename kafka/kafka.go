package kafka

import (
	"booking-service/config"
	pbb "booking-service/genproto/bookings"
	pbn "booking-service/genproto/notifications"
	pbp "booking-service/genproto/payments"
	pbr "booking-service/genproto/reviews"
	"booking-service/service"
	"context"
	"encoding/json"
	"log"
)

func ConsumeCreateBooking(cfg *config.Config, s *service.BookingService) func(message []byte) {
	return func(message []byte) {
		log.Printf("Received message from topic %s: %s", cfg.KAFKA_TOPIC_BOOKING_CREATED, string(message))

		var bk pbb.NewBooking
		err := json.Unmarshal(message, &bk)
		if err != nil {
			log.Printf("error while unmarshalling booking: %v", err)
		}

		log.Printf("Received booking: %v", &bk)

		_, err = s.CreateBooking(context.Background(), &bk)
		if err != nil {
			log.Printf("error while creating booking: %v", err)
		}
	}
}

func ConsumeUpdateBooking(cfg *config.Config, s *service.BookingService) func(message []byte) {
	return func(message []byte) {
		log.Printf("Received message from topic %s: %s", cfg.KAFKA_TOPIC_BOOKING_UPDATED, string(message))

		var bk pbb.NewData
		err := json.Unmarshal(message, &bk)
		if err != nil {
			log.Printf("error while unmarshalling booking: %v", err)
		}

		log.Printf("Received booking: %v", &bk)

		_, err = s.UpdateBooking(context.Background(), &bk)
		if err != nil {
			log.Printf("error while updating booking: %v", err)
		}
	}
}

func ConsumeCancelBooking(cfg *config.Config, s *service.BookingService) func(message []byte) {
	return func(message []byte) {
		log.Printf("Received message from topic %s: %s", cfg.KAFKA_TOPIC_BOOKING_CANCELLED, string(message))

		var bk pbb.ID
		err := json.Unmarshal(message, &bk)
		if err != nil {
			log.Printf("error while unmarshalling booking: %v", err)
		}

		log.Printf("Received booking: %v", &bk)

		_, err = s.CancelBooking(context.Background(), &bk)
		if err != nil {
			log.Printf("error while canceling booking: %v", err)
		}
	}
}

func ConsumeCreatePayment(cfg *config.Config, s *service.PaymentService) func(message []byte) {
	return func(message []byte) {
		log.Printf("Received message from topic %s: %s", cfg.KAFKA_TOPIC_PAYMENT_CREATED, string(message))

		var pay pbp.NewPayment
		err := json.Unmarshal(message, &pay)
		if err != nil {
			log.Printf("error while unmarshalling payment: %v", err)
		}

		log.Printf("Received payment: %v", &pay)

		_, err = s.CreatePayment(context.Background(), &pay)
		if err != nil {
			log.Printf("error while creating payment: %v", err)
		}
	}
}

func ConsumeCreateReview(cfg *config.Config, s *service.ReviewService) func(message []byte) {
	return func(message []byte) {
		log.Printf("Received message from topic %s: %s", cfg.KAFKA_TOPIC_REVIEW_CREATED, string(message))

		var rev pbr.NewReview
		err := json.Unmarshal(message, &rev)
		if err != nil {
			log.Printf("error while unmarshalling review: %v", err)
		}

		log.Printf("Received review: %v", &rev)

		_, err = s.CreateReview(context.Background(), &rev)
		if err != nil {
			log.Printf("error while creating review: %v", err)
		}
	}
}

func ConsumeCreateNotification(cfg *config.Config, s *service.NotificationService) func(message []byte) {
	return func(message []byte) {
		log.Printf("Received message from topic %s: %s", cfg.KAFKA_TOPIC_NOTIFICATION_CREATED, string(message))

		var not pbn.NewNotification
		err := json.Unmarshal(message, &not)
		if err != nil {
			log.Printf("error while unmarshalling notification: %v", err)
		}

		log.Printf("Received notification: %v", &not)

		_, err = s.CreateNotification(context.Background(), &not)
		if err != nil {
			log.Printf("error while creating notification: %v", err)
		}
	}
}
