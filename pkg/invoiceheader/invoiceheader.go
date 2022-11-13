package invoiceheader

import "time"

//Modelo de invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreadedAt time.Time
	updateAt  time.Time
}
