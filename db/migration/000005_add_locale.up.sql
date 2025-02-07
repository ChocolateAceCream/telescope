CREATE TABLE locale (
  ID serial PRIMARY KEY,
  LANGUAGE TEXT NOT NULL,
  raw_message TEXT NOT NULL,
  translated_message TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT unique_language_raw_message UNIQUE (language, raw_message)
);
CREATE TRIGGER "locale_updated_at" BEFORE
INSERT
  OR
UPDATE ON "public"."locale" FOR EACH ROW EXECUTE FUNCTION updated_at_autocomplete ();


insert into "public"."locale" (LANGUAGE, raw_message, translated_message) values
('en', 'error.missing.params', 'Missing required parameters'),
('cn', 'error.missing.params', '缺少参数'),
('en', 'error.failed.operation', 'Operation failed'),
('cn', 'error.failed.operation', '操作失败'),
('en', 'error.invalid.credentials', 'Invalid credentials'),
('cn', 'error.invalid.credentials', '身份验证失败'),
('en', 'error.session.expired', 'You have been signed out'),
('cn', 'error.session.expired', '身份信息已过期，请重新登录');
