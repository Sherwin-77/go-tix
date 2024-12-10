BEGIN;

CREATE TABLE sale_invoices (
    id uuid PRIMARY KEY,
    user_id INTEGER NOT NULL,
    number VARCHAR(50),
    subtotal DECIMAL(16, 2),
    service_fee DECIMAL(16, 2),
    payment_fee DECIMAL(16, 2),
    discount DECIMAL(16, 2),
    vat DECIMAL(16, 2),
    total DECIMAL(16, 2),
    status VARCHAR(50),
    transaction_at DATE,
    due_at DATE,
    completed_at TIMESTAMP(6) WITH TIME ZONE,
    canceled_at TIMESTAMP(6) WITH TIME ZONE,
    expired_at TIMESTAMP(6) WITH TIME ZONE,
    rejected_at TIMESTAMP(6) WITH TIME ZONE,
    refunded_at TIMESTAMP(6) WITH TIME ZONE,
    metadata JSON,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE
);

END;