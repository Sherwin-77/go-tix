BEGIN;

CREATE TABLE snap_payments (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    sale_invoice_id UUID NOT NULL,
    external_id VARCHAR(255) NOT NULL,
    amount DECIMAL(16, 2) NOT NULL,
    snap_token VARCHAR(255),
    invoice_url VARCHAR(512),
    status VARCHAR(255) NOT NULL,
    method VARCHAR(255) NOT NULL,
    transaction_at TIMESTAMP(6) WITH TIME ZONE,
    expired_at TIMESTAMP(6) WITH TIME ZONE,
    completed_at TIMESTAMP(6) WITH TIME ZONE,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (sale_invoice_id) REFERENCES sale_invoices(id),
    CONSTRAINT snap_payments_external_id_key UNIQUE (external_id)
);

END;