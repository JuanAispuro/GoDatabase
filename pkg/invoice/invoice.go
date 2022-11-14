package invoice

import (
	"github.com/JuanAispuro/GoDatabase/pkg/invoiceheader"
	"github.com/JuanAispuro/GoDatabase/pkg/invoiceitem"
)

// Modelo de la factura con el encabezado y el slice de los items.
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage interfaz de la base de datos.
type Storage interface {
	Create(*Model) error
}
type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

// Crear una nueva factura.
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
