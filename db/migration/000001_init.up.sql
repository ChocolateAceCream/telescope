-- 000001_init.up
-- Create a reusable trigger function for updating the updated_at column
CREATE OR REPLACE FUNCTION updated_at_autocomplete() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- add encryption extension
CREATE EXTENSION IF NOT EXISTS pgcrypto;
DROP TABLE IF EXISTS "public"."test";
CREATE TABLE "public"."test" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "created_at" timestamp(6) NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP (6) DEFAULT CURRENT_TIMESTAMP,
    "started_at" timestamp(6) NOT NULL,
    "balance" NUMERIC(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0)
);
CREATE TRIGGER "test_created_at" BEFORE
INSERT ON "public"."test" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();