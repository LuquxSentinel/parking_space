package storage

import (
	"context"

	"github.com/luquxSentinel/parkingspace/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage interface {

	// Generate primitive and hex id
	GenerateID() (primitive.ObjectID, string)

	// insert user into db
	CreateUser(ctx context.Context, user *types.User) error

	// count records matching uid
	CountUserID(ctx context.Context, uid string) (int64, error)

	// count records matching email
	CountUserEmail(ctx context.Context, email string) (int64, error)

	// count records matching phone number
	CountUserPhoneNumber(ctx context.Context, phoneNumber string) (int64, error)

	// get user from db | query by email
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)

	// get park with park ID
	GetParkByID(ctx context.Context, parkID string) (*types.Park, error)
}

type mongoStorage struct {
	userCollection    *mongo.Collection
	bookingCollection *mongo.Collection
	parkCollection    *mongo.Collection
}

func NewMongoStorage(userCol *mongo.Collection, bookingCollection *mongo.Collection, parkCollection *mongo.Collection) *mongoStorage {
	// new mongo collection

	return &mongoStorage{
		userCollection:    userCol,
		bookingCollection: bookingCollection,
		parkCollection:    parkCollection,
	}
}

func (s *mongoStorage) GenerateID() (primitive.ObjectID, string)

func (s *mongoStorage) CreateUser(ctx context.Context, user *types.User) error {
	// insert user into db

	// create unique doc id && uid
	user.ID = primitive.NewObjectID()
	user.UID = user.ID.Hex()

	// insert user
	_, err := s.userCollection.InsertOne(ctx, user)

	return err
}

func (s *mongoStorage) CountUserID(ctx context.Context, uid string) (int64, error) {
	// count records matching uid

	filter := primitive.D{primitive.E{
		Key:   "uid",
		Value: uid,
	}}

	return s.userCollection.CountDocuments(ctx, filter)
}

func (s *mongoStorage) CountUserEmail(ctx context.Context, email string) (int64, error) {
	// count records matching email

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

	// new user instance to decode to
	user := new(types.User)

	filter := primitive.D{primitive.E{
		Key:   "email",
		Value: email,
	}}

	// get user from user collection by email and decode into user instance
	err := s.userCollection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		// return nil, err if err != nil
		return nil, err
	}

	// return user and nil if no error
	return user, nil
}

func (s *mongoStorage) GetParkByID(ctx context.Context, parkID string) (*types.Park, error) {
	// get park from park collection by parkID

	// filter object by park_id with value parkID
	filter := primitive.D{
		primitive.E{
			Key:   "park_id",
			Value: parkID,
		},
	}

	// new park to decode to
	park := new(types.Park)

	// get park from park collection and decode in park instance
	err := s.parkCollection.FindOne(ctx, filter).Decode(park)
	if err != nil {
		// return nil, err if err != nil
		return nil, err
	}

	// return park and nil
	return park, nil
}
