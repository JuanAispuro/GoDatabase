package invoiceheader

import (
	"database/sql"
	"time"
)

// Modelo de invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreadedAt time.Time
	updateAt  time.Time
}

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
}

// Nueva structura Logica de servicio del invoiceheader
type Service struct {
	storage Storage
}

// Regresa un puntero y lo inicializa con el valor del constructor.
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate usada para migrar product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
