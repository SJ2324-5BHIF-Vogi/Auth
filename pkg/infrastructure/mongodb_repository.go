package infrastructure

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
)

type mongodbUserRepository struct {
	collection *mongo.Collection
}

func NewMongoDBUserRepository(collection *mongo.Collection) repository.UserRepository {
	return &mongodbUserRepository{
		collection: collection,
	}
}

func (r *mongodbUserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *mongodbUserRepository) Read(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	var user domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // TODO: return not found error
		}
		return nil, err
	}
	return &user, nil
}

func (r *mongodbUserRepository) ReadByName(ctx context.Context, username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	var user domain.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // TODO: return not found error
		}
		return nil, err
	}
	return &user, nil
}

func (r *mongodbUserRepository) Update(ctx context.Context, id uuid.UUID, user *domain.User) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *mongodbUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
