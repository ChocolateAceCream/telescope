-- 000002_add_users_table.down

-- Drop the users table
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "uuid-ossp";

DROP TRIGGER IF EXISTS "users_updated_at" ON "public"."users";