package storage

import (
	"database/sql"
	"fmt"

	"github.com/JuanAispuro/GoDatabase/pkg/product"
)

// interface de metodos scan, cualquier metodo sera un scanner.
type Scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateProduct  = ` INSERT INTO products(name, observations, price, created_at) VALUES($1,$2,$3,$4) RETURNING id `
	psqlGetAllProduct  = ` SELECT id, name, observations, price, created_at, updated_at FROM products `
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = ` UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5 `
	psqlDeleteProduct  = `DElETE FROM products WHERE id = $1`
)

// psqlProduct usado para trabajar con postgres y el paquete prodcut
type psqlProduct struct {
	db *sql.DB
}

// Func psqlProduct que retorna el nuevo puntero de psqlProduct
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

// Migrate implementa la interfaz de product.Storage
func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}

func (p *psqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreadedAt,
	).Scan(&m.ID) //Recueramos el id

	if err != nil {
		return err
	}
	fmt.Println("Se creo el producto correctamente")
	return nil
}

// Funcion para GetAll para consultar.
func (p *psqlProduct) GetAll() (product.Models, error) {

	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Escanear los valores.
	ms := make(product.Models, 0) //slice
	for rows.Next() {

		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m) //Agregamos el modelo al slice de modelos.

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return ms, nil
}

// Metodo para GetByID
func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()
	return scanRowProduct(stmt.QueryRow(id))
}

// Método Update
func (p *psqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdateAt),
		m.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	//if del error
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No existe el Product con el id: %d", m.ID)
	}

	fmt.Println("Se actualizo el producto correctamente")
	return nil
}

// Funcion para eliminar un producto
func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Se elimino el producto correctamente")
	return nil
}

// Funcion helper
func scanRowProduct(s Scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	UpdateAtNull := sql.NullTime{}
	//En donde queremos mapear el modelo.
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreadedAt,
		&UpdateAtNull,
	)
	if err != nil {
		return nil, err // nil slice y err para que se maneje.
	}
	m.Observations = observationNull.String
	m.UpdateAt = UpdateAtNull.Time
	return m, nil
}
