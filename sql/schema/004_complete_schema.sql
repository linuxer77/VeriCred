-- Complete VeriCred Database Schema
-- Generated from Go models in internal/models/models.go
-- Date: 2025-08-07

-- ============================================================================
-- DROP TABLES (in reverse dependency order)
-- ============================================================================

DROP TABLE IF EXISTS verification_requests CASCADE;
DROP TABLE IF EXISTS credentials CASCADE;
DROP TABLE IF EXISTS degrees CASCADE;
DROP TABLE IF EXISTS credential_templates CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS organizations CASCADE;
DROP TABLE IF EXISTS accounts CASCADE;

-- ============================================================================
-- CREATE TABLES
-- ============================================================================

-- Accounts table (base table for all account types)
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    metamask_address VARCHAR(42) UNIQUE NOT NULL,
    account_type VARCHAR(255) NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    credentials INTEGER DEFAULT 0,
    last_login_at TIMESTAMP NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Organizations table (universities and other institutions)
CREATE TABLE organizations (
    metamask_address VARCHAR(42) PRIMARY KEY,
    acad_email VARCHAR(255) UNIQUE NOT NULL,
    org_name VARCHAR(255) NOT NULL,
    org_type VARCHAR(100),
    org_url VARCHAR(255),
    logo_ipfs_hash VARCHAR(100),
    org_desc TEXT,
    country VARCHAR(100),
    state VARCHAR(100),
    city VARCHAR(100),
    address TEXT,
    postal_code VARCHAR(20),
    phone_number VARCHAR(20),
    established_year INTEGER,
    accreditation VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    total_credentials INTEGER DEFAULT 0,
    active_credentials INTEGER DEFAULT 0,
    total_students INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraint
    CONSTRAINT fk_org_account FOREIGN KEY (metamask_address) 
        REFERENCES accounts(metamask_address) ON DELETE CASCADE
);

-- Users table (students)
CREATE TABLE users (
    metamask_address VARCHAR(42) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    student_id VARCHAR(100),
    current_university_id VARCHAR(100),
    profile_picture VARCHAR(255),
    
    -- Foreign key constraint
    CONSTRAINT fk_user_account FOREIGN KEY (metamask_address) 
        REFERENCES accounts(metamask_address) ON DELETE CASCADE
);

-- Credential Templates table (for degree programs)
CREATE TABLE credential_templates (
    id SERIAL PRIMARY KEY,
    degree_type VARCHAR(100) NOT NULL,
    department VARCHAR(255),
    course_duration VARCHAR(100),
    description TEXT,
    university_wallet VARCHAR(42) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraint
    CONSTRAINT fk_template_university FOREIGN KEY (university_wallet) 
        REFERENCES organizations(metamask_address) ON DELETE CASCADE
);

-- Degrees table (specific degree programs offered by universities)
CREATE TABLE degrees (
    id SERIAL PRIMARY KEY,
    degree_name VARCHAR(255) NOT NULL,
    degree_type VARCHAR(100) NOT NULL, -- Bachelor's, Master's, PhD, etc.
    description TEXT,
    university_wallet VARCHAR(42) NOT NULL,
    
    -- Foreign key constraint
    CONSTRAINT fk_degree_university FOREIGN KEY (university_wallet) 
        REFERENCES organizations(metamask_address) ON DELETE CASCADE
);

-- Credentials table (issued certificates/diplomas)
CREATE TABLE credentials (
    id VARCHAR(100) PRIMARY KEY,
    degree_id INTEGER NOT NULL,
    student_wallet VARCHAR(42) NOT NULL,
    university_wallet VARCHAR(42) NOT NULL,
    degree_name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(100) NOT NULL,
    major VARCHAR(255),
    issued_date TIMESTAMP NOT NULL,
    graduation_date VARCHAR(50),
    status VARCHAR(50) DEFAULT 'Active',
    gpa VARCHAR(10),
    template_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    CONSTRAINT fk_credential_student FOREIGN KEY (student_wallet) 
        REFERENCES users(metamask_address) ON DELETE CASCADE,
    CONSTRAINT fk_credential_university FOREIGN KEY (university_wallet) 
        REFERENCES organizations(metamask_address) ON DELETE CASCADE,
    CONSTRAINT fk_credential_degree FOREIGN KEY (degree_id) 
        REFERENCES degrees(id) ON DELETE CASCADE,
    CONSTRAINT fk_credential_template FOREIGN KEY (template_id) 
        REFERENCES credential_templates(id) ON DELETE SET NULL
);

-- Verification Requests table (for employers to verify student credentials)
CREATE TABLE verification_requests (
    id SERIAL PRIMARY KEY,
    employer_wallet VARCHAR(42) NOT NULL,
    student_wallet VARCHAR(42) NOT NULL,
    credential_id VARCHAR(100) NOT NULL,
    request_message TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    responded_at TIMESTAMP NULL,
    
    -- Foreign key constraints
    CONSTRAINT fk_verification_student FOREIGN KEY (student_wallet) 
        REFERENCES users(metamask_address) ON DELETE CASCADE,
    CONSTRAINT fk_verification_credential FOREIGN KEY (credential_id) 
        REFERENCES credentials(id) ON DELETE CASCADE
);

-- ============================================================================
-- CREATE INDEXES
-- ============================================================================

-- Accounts indexes
CREATE INDEX idx_accounts_metamask_address ON accounts(metamask_address);
CREATE INDEX idx_accounts_account_type ON accounts(account_type);
CREATE INDEX idx_accounts_verified ON accounts(verified);

-- Organizations indexes
CREATE INDEX idx_organizations_org_name ON organizations(org_name);
CREATE INDEX idx_organizations_org_type ON organizations(org_type);
CREATE INDEX idx_organizations_is_verified ON organizations(is_verified);
CREATE INDEX idx_organizations_country ON organizations(country);

-- Users indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_student_id ON users(student_id);
CREATE INDEX idx_users_first_name ON users(first_name);
CREATE INDEX idx_users_last_name ON users(last_name);

-- Degrees indexes
CREATE INDEX idx_degrees_university_wallet ON degrees(university_wallet);
CREATE INDEX idx_degrees_degree_type ON degrees(degree_type);
CREATE INDEX idx_degrees_degree_name ON degrees(degree_name);

-- Credentials indexes
CREATE INDEX idx_credentials_student_wallet ON credentials(student_wallet);
CREATE INDEX idx_credentials_university_wallet ON credentials(university_wallet);
CREATE INDEX idx_credentials_degree_id ON credentials(degree_id);
CREATE INDEX idx_credentials_status ON credentials(status);
CREATE INDEX idx_credentials_type ON credentials(type);
CREATE INDEX idx_credentials_template_id ON credentials(template_id);
CREATE INDEX idx_credentials_issued_date ON credentials(issued_date);

-- Verification Requests indexes
CREATE INDEX idx_verification_employer_wallet ON verification_requests(employer_wallet);
CREATE INDEX idx_verification_student_wallet ON verification_requests(student_wallet);
CREATE INDEX idx_verification_credential_id ON verification_requests(credential_id);
CREATE INDEX idx_verification_status ON verification_requests(status);
CREATE INDEX idx_verification_requested_at ON verification_requests(requested_at);

-- Credential Templates indexes
CREATE INDEX idx_credential_templates_university_wallet ON credential_templates(university_wallet);
CREATE INDEX idx_credential_templates_degree_type ON credential_templates(degree_type);

-- ============================================================================
-- TRIGGERS FOR UPDATED_AT TIMESTAMPS
-- ============================================================================

-- Function to update the updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Triggers for tables with updated_at columns
CREATE TRIGGER update_accounts_updated_at 
    BEFORE UPDATE ON accounts 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_organizations_updated_at 
    BEFORE UPDATE ON organizations 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_credentials_updated_at 
    BEFORE UPDATE ON credentials 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_credential_templates_updated_at 
    BEFORE UPDATE ON credential_templates 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================================
-- SAMPLE DATA INSERTS (Optional - for testing)
-- ============================================================================

-- Sample account data
INSERT INTO accounts (metamask_address, account_type, verified) VALUES
('0x1234567890123456789012345678901234567890', 'organization', true),
('0x2345678901234567890123456789012345678901', 'student', false),
('0x3456789012345678901234567890123456789012', 'student', false);

-- Sample organization data
INSERT INTO organizations (
    metamask_address, acad_email, org_name, org_type, country, 
    state, city, is_verified
) VALUES (
    '0x1234567890123456789012345678901234567890',
    'admin@university.edu',
    'Sample University',
    'University',
    'United States',
    'California',
    'San Francisco',
    true
);

-- Sample user data
INSERT INTO users (
    metamask_address, email, first_name, last_name, 
    password, student_id
) VALUES 
(
    '0x2345678901234567890123456789012345678901',
    'student1@email.com',
    'John',
    'Doe',
    'hashed_password_here',
    'STU001'
),
(
    '0x3456789012345678901234567890123456789012',
    'student2@email.com',
    'Jane',
    'Smith',
    'hashed_password_here',
    'STU002'
);

-- Sample degree data
INSERT INTO degrees (degree_name, degree_type, university_wallet, description) VALUES
(
    'Bachelor of Computer Science',
    'Bachelor''s',
    '0x1234567890123456789012345678901234567890',
    'Undergraduate degree in Computer Science'
);

-- Sample credential template data
INSERT INTO credential_templates (degree_type, department, course_duration, university_wallet) VALUES
(
    'Bachelor''s',
    'Computer Science',
    '4 years',
    '0x1234567890123456789012345678901234567890'
);

-- ============================================================================
-- COMMENTS AND DOCUMENTATION
-- ============================================================================

COMMENT ON TABLE accounts IS 'Base table for all user accounts (students, organizations, employers)';
COMMENT ON TABLE organizations IS 'Universities and educational institutions';
COMMENT ON TABLE users IS 'Student users of the platform';
COMMENT ON TABLE degrees IS 'Degree programs offered by universities';
COMMENT ON TABLE credentials IS 'Issued credentials/certificates to students';
COMMENT ON TABLE verification_requests IS 'Requests from employers to verify student credentials';
COMMENT ON TABLE credential_templates IS 'Templates for different types of credentials';

COMMENT ON COLUMN accounts.metamask_address IS 'Ethereum wallet address (primary identifier)';
COMMENT ON COLUMN accounts.account_type IS 'Type of account: student, organization, employer';
COMMENT ON COLUMN credentials.status IS 'Status of credential: Active, Revoked, Suspended, etc.';
COMMENT ON COLUMN verification_requests.status IS 'Status: pending, approved, denied';
