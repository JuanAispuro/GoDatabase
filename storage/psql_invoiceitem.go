package storage

import (
	"database/sql"
	"fmt"

	"github.com/JuanAispuro/GoDatabase/pkg/invoiceitem"
)

const (
	psqlMigrateItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT  NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
		(invoice_header_id) REFERENCES invoice_headers(id) ON UPDATE
		RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY
		(product_id) REFERENCES products(id) ON UPDATE
		RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, creaded_at`
)

// PsqlInvoiceItem usado para trabajar con postgres y el paquete invoiceItem
type PsqlInvoiceItem struct {
	db *sql.DB
}

// Func Psqlproduct que retorna el nuevo puntero de PsqlInvoiceItem
func NewPsqlInvocieItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implementa la interfaz de invoiceItem.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de invoiceItem ejecutada correctamente")
	return nil
}

// CreateTx implementar un interfaz de invoiceItem.Storage
func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.productID).Scan(&item.ID, &item.CreadedAt)
	}
	if err != nil {
		return err
	}
	return nil
}
