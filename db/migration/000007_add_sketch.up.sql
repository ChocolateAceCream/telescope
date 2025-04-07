CREATE TABLE sketch (
  ID serial PRIMARY KEY,
  project_name TEXT NOT NULL,
  comment TEXT,
  project_id int4 NOT NULL REFERENCES project(id) ON DELETE CASCADE,
  uploader_id int4 NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  thumbnail_url TEXT,
  full_image_url TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
CREATE TRIGGER sketch_updated_at BEFORE
INSERT
  OR
UPDATE ON sketch FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();