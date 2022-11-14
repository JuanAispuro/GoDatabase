package storage

import (
	"database/sql"
	"fmt"

	"github.com/JuanAispuro/GoDatabase/pkg/product"
)

const (
	MySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	MySQLlCreateProduct = ` INSERT INTO products(name, observations, 
	price, created_at) VALUES(?,?,?,?)`
	MySQLGetAllProduct  = ` SELECT id, name, observations, price, created_at, updated_at FROM products `
	MySQLGetProductByID = psqlGetAllProduct + " WHERE id = ?"
	MySQLUpdateProduct  = ` UPDATE products SET name = ?, observations = ?, price = ?, updated_at = ? WHERE id = ? `
	MySQLDeleteProduct  = `DElETE FROM products WHERE id = ?`
)

// mySQLProduct usado para trabajar con postgres y el paquete prodcut
type mySQLProduct struct {
	db *sql.DB
}

// Func mySQLProduct que retorna el nuevo puntero de mySQLProduct
func newMySQLProduct(db *sql.DB) *mySQLProduct {
	return &mySQLProduct{db}
}

// Migrate implementa la interfaz de product.Storage
func (p *mySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(MySQLMigrateProduct)
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

func (p *mySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(MySQLlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//utilizamos Exec no query row
	results, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreadedAt,
	)
	if err != nil {
		return err
	}
	id, err := results.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id) //Recueramos el id
	fmt.Printf("Se creo el producto correctamente con ID: %d", m.ID)
	return nil
}

// Funcion para GetAll para consultar.
func (p *mySQLProduct) GetAll() (product.Models, error) {

	stmt, err := p.db.Prepare(MySQLGetAllProduct)
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
func (p *mySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(MySQLGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()
	return scanRowProduct(stmt.QueryRow(id))
}

// Método Update
func (p *mySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(MySQLUpdateProduct)
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
func (p *mySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(MySQLDeleteProduct)
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
