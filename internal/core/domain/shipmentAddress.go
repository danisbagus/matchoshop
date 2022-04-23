package domain

type ShipmentAddress struct {
	ShipmentAddressID int64  `db:"shipment_address_id"`
	OrderID           int64  `db:"order_id"`
	Address           string `db:"address"`
	City              string `db:"city"`
	PostalCode        string `db:"postal_code"`
	Country           string `db:"country"`
}
