package domain

type OrderProduct struct {
	OrderID   int64
	ProductID int64
	Name      string
	Price     int64
	Quantity  int64
}
