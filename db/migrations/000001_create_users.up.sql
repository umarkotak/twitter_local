-- Plugin for uuid generation, ref: https://dba.stackexchange.com/questions/122623/default-value-for-uuid-column-in-postgres
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id UUID NOT NULL DEFAULT uuid_generate_v1() PRIMARY KEY,
  name TEXT NOT NULL,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  photo_url TEXT NOT NULL,
  about TEXT,

  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX users_username_unique_index ON users(UPPER(username));
CREATE UNIQUE INDEX users_email_unique_index ON users(email);
