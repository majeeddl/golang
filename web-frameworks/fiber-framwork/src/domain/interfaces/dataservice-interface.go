package interfaces

import "fiberframework/src/domain/entities"

type IDataService interface {
	OrderRepository() Repository[entities.Order]
}
