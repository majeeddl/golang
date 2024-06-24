package dataservices

import (
	"fiberframework/src/domain/interfaces"
	mongo_dataservices "fiberframework/src/frameworks/dataservices/mongo"
)

type DataService struct {
	URI      string
	Database string
}

func (dataservice *DataService) NewDataService() interfaces.IDataService {
	return mongo_dataservices.NewMongoDataService(
		dataservice.URI,
		dataservice.Database,
	)
}
