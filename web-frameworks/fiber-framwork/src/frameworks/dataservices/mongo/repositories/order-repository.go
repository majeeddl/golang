package repositories

import (
	"context"
	"fiberframework/src/domain/entities"
	"fiberframework/src/domain/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

func MongoOrderRepository(ctx context.Context, database *mongo.Database) interfaces.Repository[entities.Order] {
	var orderRepository interfaces.Repository[entities.Order] = &MongoRepository[entities.Order]{
		Ctx:        ctx,
		Database:   database,
		Collection: "orders",
	}

	return orderRepository
}
