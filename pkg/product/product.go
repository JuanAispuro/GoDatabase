package product

import "time"

//Modelo de producto
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreadedAt    time.Time
	UpdateAt     time.Time
}

//Modelo slice de modelos
type Models []*Model

type Storage interface {
	Migrate() error
	/*
		Create(*Model) error //Puntero del modelo y si hay error.
		Update(*Model) error
		GetAll() (Models, error)
		GetByID(uint) (*Model, error)
		Delete(uint) error
	*/

}

//Nueva structura Logida del producto
type Service struct {
	storage Storage
}

//Regresa un puntero y lo inicializa con el valor del constructor.
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate usada para migrar product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
