package service

import (
	"context"
	"errors"
	"time"

	"github.com/luquxSentinel/parkingspace/storage"
	"github.com/luquxSentinel/parkingspace/types"
)

type BookingService interface {
	CreateBooking(ctx context.Context, uid string, data *types.CreateBookingData) error
	CancelBooking(ctx context.Context, bookingID string) error
	GetBooking(ctx context.Context, bookingID string) (*types.Booking, error)
	CommitBooking(ctx context.Context, bookingID string) error
}

type bookingService struct {
	storage      storage.Storage
	cacheStorage storage.CacheStorage
}

func NewBookingService(storage storage.Storage) *bookingService {
	return &bookingService{
		storage: storage,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, uid string, data *types.CreateBookingData) error {
	// create a booking cache and save in redis cache awaiting  payment and confirmation

	// check if uid exists
	count, err := s.storage.CountUserID(ctx, uid)
	if err != nil {
		return err
	}

	// TODO: count should be ZERO or ONE 0 or 1
	if count < 1 {
		return errors.New("user not found")
	}

	// get park from storage
	park, err := s.storage.GetParkByID(ctx, data.ParkID)
	if err != nil {
		return err
	}

	if park.FreeSpace == 0 {
		return errors.New("no space available for this park")
	}

	balance := park.PricePerHour * time.Hour.Hours()

	newBooking := new(types.Booking)

	newBooking.Balance = balance
	newBooking.ParkingID = park.ParkID

	//from time string to time.Time
	from, _ := time.Parse(time.RFC3339, data.From)

	//to time string to time.Time
	to, _ := time.Parse(time.RFC3339, data.To)

	// assign Booking [From | To] to from and to time.Time
	newBooking.From = from
	newBooking.To = to

	// generate booking id
	newBooking.ID, newBooking.BookingID = s.storage.GenerateID()

	// new payment
	payment := types.Payment{}

	// generate payment id
	payment.ID, payment.PaymentID = s.storage.GenerateID()

	payment.PaymentType = ""
	payment.PaymentStatus = "PENDING"

	// set booking payment to payment
	newBooking.Payment = payment
	newBooking.CreatedAt = time.Now().UTC()
	newBooking.Status = "PENDING"
	newBooking.UID = uid
	newBooking.QRCode = ""

	// insert booking into cache storage to await confirmation and payment
	err = s.cacheStorage.CreateBooking(ctx, newBooking)
	if err != nil {
		return err
	}

	return nil
}
