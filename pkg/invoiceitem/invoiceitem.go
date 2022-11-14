package invoiceitem

import (
	"database/sql"
	"time"
)

// Modelo de invoiceitem
type Model struct {
	ID              uint
	invoiceheaderID uint
	productID       uint
	CreadedAt       time.Time
	updateAt        time.Time
}

// Slice de models
type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, uint, Models) error
}

// Nueva estructura Logica de servicio de invoiceitem
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
func (s *Service) CreateTx(tx *sql.Tx, invoiceheaderID uint, m Models) error {
	return s.storage.CreateTx(tx, invoiceheaderID, m)
}
