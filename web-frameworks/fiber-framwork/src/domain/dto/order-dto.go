package dto

type CreateOrderDTO struct {
	ShipmentNumber string `json:"shipment_number" validate:"required"`
	CountryCode    string `json:"country_code" validate:"required"`
	OrderNumber    string `json:"order_number" validate:"required"`
	Age            int    `json:"age" validate:"required,oldAge"`
}
