package mongo_dataservices

import (
	"context"
	"fiberframework/src/domain/entities"
	"fiberframework/src/domain/interfaces"
	"fiberframework/src/frameworks/dataservices/mongo/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// import (
// 	"fiberframework/src/domain/entities"
// 	"fiberframework/src/domain/interfaces"
// )

type MongoDataService struct {
	Database *mongo.Database
	Context  context.Context
}

func (m *MongoDataService) OrderRepository() interfaces.Repository[entities.Order] {
	return repositories.MongoOrderRepository(m.Context, m.Database)
}

func NewMongoDataService(uri string, database string) interfaces.IDataService {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	mongodatabase := client.Database(database)
	return &MongoDataService{
		Database: mongodatabase,
		Context:  context.TODO(),
	}
}
