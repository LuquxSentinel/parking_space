package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID        primitive.ObjectID `bson:"_id"`
	BookingID string             `bson:"booking_id"`
	ParkingID string             `bson:"parking_id"`
	From      time.Time          `bson:"from"`
	To        time.Time          `bson:"to"`
	CreatedAt time.Time          `bson:"created_at"`
	UID       string             `bson:"uid"`
	Balance   float64            `bson:"balance"`
	Payment   Payment            `bson:"payment"`
	Status    string             `bson:"status"` // CANCELLED | IN | OUT | PENDING | RESERVED
	QRCode    string             `bson:"qr_code"`
}

type CreateBookingData struct {
	From   string `json:"from"`
	To     string `json:"to"`
	ParkID string `json:"park_id"`
}

type Payment struct {
	ID            primitive.ObjectID `bson:"_id"`
	PaymentID     string             `bson:"payment_id"`
	PaymentStatus string             `bson:"payment_status"` // PENDING | FAILED | SUCCESSFUL
	PaymentType   string             `bson:"payment_type"`   // ONLINE | CREDIT CARD | CASH
}
