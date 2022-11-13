package invoiceitem

import "time"

//Modelo de invoiceitem
type Model struct {
	ID              uint
	invoiceheaderID uint
	productID       uint
	CreadedAt       time.Time
	updateAt        time.Time
}
