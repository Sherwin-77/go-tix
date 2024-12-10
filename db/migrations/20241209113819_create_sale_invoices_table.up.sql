BEGIN;

CREATE TABLE sale_invoices (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
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
    metadata JSONB,
    created_at TIMESTAMP(6) WITH TIME ZONE,
    updated_at TIMESTAMP(6) WITH TIME ZONE,

    CONSTRAINT sale_invoices_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

END;