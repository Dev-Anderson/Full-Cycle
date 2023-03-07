package entity

type OrderReposityInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
}
