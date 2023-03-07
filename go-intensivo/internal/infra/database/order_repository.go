package database

import (
	"database/sql"
	"go-intensivo/internal/entity"
)

type OrderReposity struct {
	Db *sql.DB
}

func NewOrderReposityr(db *sql.DB) *OrderReposity {
	return &OrderReposity{
		Db: db,
	}
}

func (r *OrderReposity) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES(?,?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderReposity) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
