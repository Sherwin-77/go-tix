BEGIN;

CREATE TABLE sale_invoices (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    number VARCHAR(50),
    subtotal DECIMAL(16, 2) NOT NULL,
    service_fee DECIMAL(16, 2) NOT NULL,
    payment_fee DECIMAL(16, 2) NOT NULL,
    discount DECIMAL(16, 2) NOT NULL,
    vat DECIMAL(16, 2) NOT NULL,
    total DECIMAL(16, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    transaction_at TIMESTAMP(6) WITH TIME ZONE,
    due_at TIMESTAMP(6) WITH TIME ZONE,
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