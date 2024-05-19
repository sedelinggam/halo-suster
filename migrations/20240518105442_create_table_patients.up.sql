CREATE TYPE patient_gender AS ENUM ('male', 'female');

CREATE TABLE patients (
    id VARCHAR (26) PRIMARY KEY,
    identity_number VARCHAR (16) UNIQUE NOT NULL,
    phone_number VARCHAR (15) NOT NULL,
    name VARCHAR (30) NOT NULL,
    birth_date DATE NOT NULL,
    gender patient_gender NOT NULL,
    identity_card_scan_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ NULL
);

CREATE INDEX patients_identity_number ON patients (identity_number);
