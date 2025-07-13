-- +goose Up

CREATE TABLE users {
    metamask_address VARCHAR(42) PRIMARY KEY REFERENCES accounts(metamask_address), -- Ethereum address from Metamask wallet
    email            TEXT NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    first_name       TEXT NOT NULL,
    last_name        TEXT NOT NULL,
    password         TEXT NOT NULL
};

-- +goose Down

DROP TABLE users;



