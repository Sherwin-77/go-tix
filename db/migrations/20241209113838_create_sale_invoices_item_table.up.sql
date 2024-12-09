BEGIN;

CREATE TABLE sale_invoice_items (
    id uuid SERIAL PRIMARY KEY AUTOINCREMENT,
    sale_invoice_id INTEGER NOT NULL,
    invoiceable_id INTEGER,
    invoiceable_type VARCHAR(50),
    qty INTEGER,
    total DECIMAL(12, 2),
    metadata JSON,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);

END;