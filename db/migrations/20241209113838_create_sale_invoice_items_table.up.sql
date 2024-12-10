BEGIN;

CREATE TABLE sale_invoice_items (
    id UUID PRIMARY KEY,
    sale_invoice_id UUID NOT NULL,
    invoiceable_id UUID NOT NULL,
    invoiceable_type VARCHAR(255) NOT NULL,
    qty INTEGER,
    total DECIMAL(16, 2),
    metadata JSONB,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT sale_invoice_items_sale_invoice_id_fkey FOREIGN KEY (sale_invoice_id) REFERENCES sale_invoices (id),
    CONSTRAINT sale_invoice_items_invoiceable_id_fkey FOREIGN KEY (invoiceable_id) REFERENCES tickets (id)
);

END;