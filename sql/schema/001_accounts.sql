-- +goose Up

CREATE TABLE accounts {
    metamask_address PRIMARY KEY VARCHAR(42) NOT NULL, -- Ethereum address from Metamask wallet
    account_type     VARCHAR(20) NOT NULL,    -- "individual" or "organization"
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    password         TEXT NOT NULL,
    is_active        BOOLEAN NOT NULL DEFAULT TRUE
};

-- +goose Down
    DROP TABLE accounts;