-- 000004_add_a_users_table.up
-- Create a reusable trigger function for updating the updated_at column
DROP TABLE IF EXISTS "public"."a_users";
CREATE TABLE "public"."a_users" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(50) NOT NULL UNIQUE,
  "email" VARCHAR(320) NOT NULL UNIQUE,
  "password" VARCHAR(64) NOT NULL,
  "role" VARCHAR(50) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP (6) DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER "a_users_updated_at_autocomplete" BEFORE
INSERT ON "public"."a_users" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();