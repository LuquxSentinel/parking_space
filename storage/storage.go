package storage

import (
	"context"

	"github.com/luquxSentinel/parkingspace/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {
	// insert user into db
	CreateUser(ctx context.Context, user *types.User) error

	// count records matching email
	CountUserEmail(ctx context.Context, email string) (int64, error)

	// count records matching phone number
	CountUserPhoneNumber(ctx context.Context, phoneNumber string) (int64, error)

	// get user from db | query by email
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
}

type mongoStorage struct {
	userCollection *mongo.Collection
}

func NewMongoStorage(userCol *mongo.Collection) *mongoStorage {
	// new mongo collection

	return &mongoStorage{
		userCollection: userCol,
	}
}

func (s *mongoStorage) CreateUser(ctx context.Context, user *types.User) error {
	// insert user into db

	// create unique doc id && uid
	user.ID = primitive.NewObjectID()
	user.UID = user.ID.Hex()

	// insert user
	_, err := s.userCollection.InsertOne(ctx, user)

	return err
}

func (s *mongoStorage) CountUserEmail(ctx context.Context, email string) (int64, error) {
	// count records matching email {}

	filter := primitive.D{primitive.E{
		Key:   "email",
		Value: email,
	}}

	// count docs that match email
	return s.userCollection.CountDocuments(ctx, filter)
}

func (s *mongoStorage) CountUserPhoneNumber(ctx context.Context, phoneNumber string) (int64, error) {
	// count records matching phone number {}
	filter := primitive.D{primitive.E{
		Key:   "phone_number",
		Value: phoneNumber,
	}}

	// count docs that match phone number
	return s.userCollection.CountDocuments(ctx, filter)

}

func (s *mongoStorage) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	// get user from db | query by email {}

	user := new(types.User)

	filter := primitive.D{primitive.E{
		Key:   "email",
		Value: email,
	}}

	err := s.userCollection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
