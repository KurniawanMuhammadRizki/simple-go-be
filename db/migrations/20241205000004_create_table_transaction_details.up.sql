CREATE TABLE transaction_details (
    id BIGSERIAL PRIMARY KEY,
    transaction_id BIGINT NOT NULL,
    voucher_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    sub_total_cost BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (voucher_id) REFERENCES vouchers(id)
);
