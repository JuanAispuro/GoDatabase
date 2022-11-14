package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrIDNotFound = errors.New("El producto no contiene un ID")
)

// Modelo de producto
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreadedAt    time.Time
	UpdateAt     time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreadedAt.Format("2022-11-13"), m.UpdateAt.Format("2022-11-13"))
}

// Modelo slice de modelos
type Models []*Model

// Func models
func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"id", "name", "observations", "price", "created_at", "update_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

type Storage interface {
	Migrate() error
	Create(*Model) error //Puntero del modelo y si hay error.
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Update(*Model) error
	Delete(uint) error
}

// Nueva structura Logida del producto
type Service struct {
	storage Storage
}

// Regresa un puntero y lo inicializa con el valor del constructor.
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate usada para migrar producto
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create es usado para crear el producto.
func (s *Service) Create(m *Model) error {
	m.CreadedAt = time.Now() //Mandamos el tiempo de hoy.
	return s.storage.Create(m)
}

// Getall es usado para mostrar los productos
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetByID es usado para obtener un solo producto.
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

// Update actualiza los datos de la tabla productos.
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdateAt = time.Now()
	return s.storage.Update(m)
}

// DELETE elimina un dato del producto.
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
