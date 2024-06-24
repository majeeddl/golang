package usecases

import "fiberframework/src/domain/interfaces"

type OrderUseCase struct {
	datasource interfaces.IDataService
}

func (u *OrderUseCase) CreateOrder() {

}
