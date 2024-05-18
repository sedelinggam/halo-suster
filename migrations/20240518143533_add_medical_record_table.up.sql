CREATE TABLE medical_records (
  id VARCHAR (26) PRIMARY KEY,
  created_at timestamptz NOT NULL,
  symptoms VARCHAR(2000) NOT NULL,
  medications VARCHAR(2000) NOT NULL,
  identity_number VARCHAR(16) NOT NULL,
  user_id VARCHAR(26) NOT NULL,
  CONSTRAINT fk_identity_number FOREIGN KEY (identity_number) REFERENCES patients (identity_number),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);