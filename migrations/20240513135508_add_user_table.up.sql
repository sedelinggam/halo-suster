CREATE TYPE user_role AS ENUM ('it','nurse');
CREATE TABLE users (
  id VARCHAR (26) PRIMARY KEY,
  nip VARCHAR (15) UNIQUE NOT NULL,
  name VARCHAR (50) NOT NULL,
  password VARCHAR (64) NULL,
  identity_card_scan_img VARCHAR NULL,
  role user_role,
  created_at timestamptz NOT NULL,
  deleted_at timestamptz NULL
);

CREATE INDEX user_id ON users (id);
CREATE INDEX user_nip ON users (nip);
CREATE INDEX role_it ON users (role) WHERE role = 'it';
CREATE INDEX role_nurse ON users (role) WHERE role = 'nurse';
CREATE INDEX user_deleted_at_null ON users (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX user_deleted_at_not_null ON users (deleted_at) WHERE deleted_at IS NOT NULL;