-- +Goose Up

CREATE TABLE organizations {
    metamask_address VARCHAR(42) PRIMARY KEY REFFERENCES accounts(metamask_address),
    acad_emal TEXT NOT NULL UNIQUE,
    org_name TEXT NOT NULL,
    registration_number TEXT,
    business_type TEXT,
    verification_status TEXT NOT NULL DEFAULT 'pending', -- "pending", "verified", "rejected"
};

-- +Goose Down
DROP TABLE organizations;