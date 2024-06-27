package dataservices

import (
	"fiberframework/src/domain/interfaces"
	mongo_dataservices "fiberframework/src/frameworks/dataservices/mongo"
	"os"
)

type DataService struct {
	Type     string
	URI      string
	Database string
}

func (dataservice *DataService) NewDataService() interfaces.IDataService {

	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DATABASE")

	config := mongo_dataservices.MongoDataServiceConfig{
		Uri:      uri,
		Database: database,
	}

	return mongo_dataservices.NewMongoDataService(config)
}
