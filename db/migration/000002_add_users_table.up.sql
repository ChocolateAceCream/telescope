-- 000002_add_users_table.up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
  -- id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(320) NOT NULL UNIQUE,
  info JSONB DEFAULT '{}',
  -- given_name VARCHAR(255),
  -- avatar TEXT,
  -- family_name VARCHAR(255),
  -- locale VARCHAR(10),
  -- role VARCHAR(10) DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER "users_updated_at" BEFORE
INSERT OR UPDATE ON "public"."users" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();
