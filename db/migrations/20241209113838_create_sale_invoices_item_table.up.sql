BEGIN;

CREATE TABLE sale_invoice_items (
    id uuid PRIMARY KEY,
    sale_invoice_id INTEGER NOT NULL,
    invoiceable_id INTEGER,
    invoiceable_type VARCHAR(50),
    qty INTEGER,
    total DECIMAL(12, 2),
    metadata JSON,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT sale_invoice_items_sale_invoice_id_fkey FOREIGN KEY (sale_invoice_id) REFERENCES sale_invoices (id),
    CONSTRAINT sale_invoice_items_invoiceable_id_fkey FOREIGN KEY (invoiceable_id) REFERENCES tickets (id)
);

END;