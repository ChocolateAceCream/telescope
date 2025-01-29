CREATE TABLE locale (
  ID serial PRIMARY KEY,
  LANGUAGE TEXT NOT NULL,
  raw_message TEXT NOT NULL,
  translated_message TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE TRIGGER "locale_updated_at" BEFORE
INSERT
  OR
UPDATE ON "public"."locale" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete ();


insert into "public"."locale" (LANGUAGE, raw_message, translated_message) values
('en', 'error.missing.params', 'Missing required parameters'),
('en', 'error.failed.operation', 'Operation failed'),
('en', 'error.invalid.credentials', 'Invalid credentials');