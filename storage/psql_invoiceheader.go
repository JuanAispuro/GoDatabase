package storage

import (
	"database/sql"
	"fmt"

	"github.com/JuanAispuro/GoDatabase/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`
)

// PsqlINvoiceHeader usado para trabajar con postgres y el paquete invoiceheader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// Func Psqlproduct que retorna el nuevo puntero de PsqlInvoiceHeader
func NewPsqlInvocieHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implementa la interfaz de invoice_headers.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
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
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreadedAt)
}
