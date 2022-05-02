package domain

type OrderProduct struct {
	OrderID   int64
	ProductID int64
	Quantity  int64
	Name      string
	Image     string
	Price     int64
}
