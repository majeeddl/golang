package repositories

import (
	"context"
	"fiberframework/src/domain/entities"
	"fiberframework/src/domain/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepositoryConfig struct {
	Ctx      context.Context
	Database *mongo.Database
}

type MongoOrderRepository struct {
	MongoRepository[entities.Order]
}

func (r *MongoOrderRepository) Test() (string, error) {
	return "Test", nil
}

func NewMongoOrderRepository(config MongoRepositoryConfig) interfaces.IOrderRepository {
	// var orderRepository interfaces.IRepository[entities.Order] = &MongoRepository[entities.Order]{
	// 	Ctx:        config.Ctx,
	// 	Database:   config.Database,
	// 	Collection: "orders",
	// }

	var orderRepository interfaces.IOrderRepository = &MongoOrderRepository{
		MongoRepository[entities.Order]{
			Ctx:        config.Ctx,
			Database:   config.Database,
			Collection: "orders",
		},
	}

	return orderRepository
}
