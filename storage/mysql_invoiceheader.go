package storage

import (
	"database/sql"
	"fmt"

	"github.com/JuanAispuro/GoDatabase/pkg/invoiceheader"
)

const (
	MySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	MySQLCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)

// MySQLInvoiceHeader usado para trabajar con postgres y el paquete invoiceheader
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// Func Psqlproduct que retorna el nuevo puntero de MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implementa la interfaz de invoice_headers.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(MySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de invoiceheader ejecutada correctamente")
	return nil
}

// Crear la transaccion con la interface invoiceheader.Storage
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(MySQLCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)
	return nil
}
