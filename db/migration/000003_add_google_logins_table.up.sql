CREATE TABLE google_logins (
  id SERIAL PRIMARY KEY,
  username text NOT NULL,
  email VARCHAR(320) NOT NULL UNIQUE REFERENCES users(email) ON DELETE CASCADE,
  access_token VARCHAR(255) NOT NULL,
  issued_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  expired_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY (email) REFERENCES users(email) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TRIGGER "google_logins_updated_at" BEFORE
INSERT
  OR
UPDATE ON "public"."google_logins" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete();