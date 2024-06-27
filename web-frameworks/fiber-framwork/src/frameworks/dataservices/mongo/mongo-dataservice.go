package mongo_dataservices

import (
	"context"
	"fiberframework/src/domain/interfaces"
	"fiberframework/src/frameworks/dataservices/mongo/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// import (
// 	"fiberframework/src/domain/entities"
// 	"fiberframework/src/domain/interfaces"
// )

type MongoDataServiceConfig struct {
	Uri      string
	Database string
}

type MongoDataService struct {
	Uri      string
	Database string
	database *mongo.Database
	context  context.Context
}

func (m *MongoDataService) OrderRepository() interfaces.IOrderRepository {
	return repositories.NewMongoOrderRepository(
		repositories.MongoRepositoryConfig{
			Ctx:      m.context,
			Database: m.database,
		},
	)
}

func NewMongoDataService(config MongoDataServiceConfig) interfaces.IDataService {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Uri))

	if err != nil {
		panic(err)
	}

	mongodatabase := client.Database(config.Database)

	return &MongoDataService{
		database: mongodatabase,
		context:  context.TODO(),
	}
}
