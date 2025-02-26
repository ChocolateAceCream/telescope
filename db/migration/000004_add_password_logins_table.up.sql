-- 000004_add_a_users_table.up
-- Create a reusable trigger function for updating the updated_at column
DROP TABLE IF EXISTS "public"."password_logins";
CREATE TABLE "public"."password_logins" (
  "id" SERIAL PRIMARY KEY,
  "email" VARCHAR(320) NOT NULL UNIQUE REFERENCES users(email) ON DELETE CASCADE,
  "password" VARCHAR(64) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP (6) DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER "password_logins_updated_at_autocomplete" BEFORE
INSERT ON "public"."password_logins" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();