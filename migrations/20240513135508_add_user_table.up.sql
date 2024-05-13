CREATE TYPE user_role AS ENUM ('it','nurse');
CREATE TABLE users (
  id uuid PRIMARY KEY,
  nip VARCHAR (13) UNIQUE NOT NULL,
  name VARCHAR (50) NOT NULL,
  password VARCHAR (33),
  created_at timestamptz NOT NULL,
  identity_card_scan_img VARCHAR,
  role user_role
);

CREATE INDEX user_id ON users (id);
CREATE INDEX user_nip ON users (nip);