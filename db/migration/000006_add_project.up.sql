CREATE TABLE project (
  ID serial PRIMARY KEY,
  project_name TEXT NOT NULL,
  comment TEXT,
  creator int4 NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  status text,
  address text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE TRIGGER project_updated_at BEFORE
INSERT
  OR
UPDATE ON project FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();