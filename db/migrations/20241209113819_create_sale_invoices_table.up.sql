BEGIN;

CREATE TABLE sale_invoices (
    id uuid PRIMARY KEY,
    user_id INTEGER NOT NULL,
    number VARCHAR(50),
    subtotal DECIMAL(12, 2),
    service_fee DECIMAL(12, 2),
    payment_fee DECIMAL(12, 2),
    discount DECIMAL(12, 2),
    vat DECIMAL(12, 2),
    total DECIMAL(12, 2),
    status VARCHAR(50),
    transaction_at DATE,
    due_at DATE,
    completed_at DATE,
    canceled_at DATE,
    expired_at DATE,
    rejected_at DATE,
    refunded_at DATE,
    metadata JSON,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);

END;