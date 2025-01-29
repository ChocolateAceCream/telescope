-- 000001_init.down

-- Drop the trigger associated with the test table
DROP TRIGGER IF EXISTS test_created_at ON "public"."test";
-- Drop the encryption extension
drop extension if exists pgcrypto;

-- Drop the test table
DROP TABLE IF EXISTS "public"."test";
-- Drop the reusable trigger function for updating the updated_at column
DROP FUNCTION IF EXISTS updated_at_autocomplete;