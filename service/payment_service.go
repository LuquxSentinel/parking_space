package service

import (
	"net/http"

	"github.com/rpip/paystack-go"
)

type PaymentService interface {
	FinishBooking()
}

type paymentService struct {
	client *paystack.Client
}

func NewPaymentService() *paymentService {
	paystackapiKey := "sk_test_b748a89ad84f35c2f1a8b81681f956274de048bb"
	return &paymentService{
		client: paystack.NewClient(paystackapiKey, http.DefaultClient),
	}
}
