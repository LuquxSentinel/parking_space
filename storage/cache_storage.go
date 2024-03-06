package storage

import (
	"context"

	"github.com/luquxSentinel/parkingspace/types"
	"github.com/redis/go-redis/v9"
)

type CacheStorage interface {
	CreateBooking(ctx context.Context, booking *types.Booking) error
	etBooking(ctx context.Context, bookingID string) (*types.Booking, error)
}

type redisCacheStorage struct {
	client *redis.Client
}

func NewRedisCacheStorage() *redisCacheStorage {
	return &redisCacheStorage{
		client: redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
			},
		),
	}
}

func (c *redisCacheStorage) CreateBooking(ctx context.Context, booking *types.Booking) error {
	// TODO: set expiry date for the cache
	err := c.client.Set(ctx, booking.BookingID, booking, 0).Err()
	return err
}

func (c *redisCacheStorage) GetBooking(ctx context.Context, bookingID string) (*types.Booking, error) {
	// get booking from cache
	booking := new(types.Booking)

	// get booking and scan into booking of [types.Booking] type
	err := c.client.Get(ctx, bookingID).Scan(booking)

	// return booking[nil if err != nil] and err[nil if no error]
	return booking, err
}
