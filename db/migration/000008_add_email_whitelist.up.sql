-- 000008_add_email_whitelist.up
CREATE TABLE email_whitelist (
  id SERIAL PRIMARY KEY,
  email VARCHAR(320) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
insert into email_whitelist (email)
values
  ('vincewang0101@gmail.com'),
  ('nuodi@hotmail.com'),
  ('ajvalino09@gmail.com'),
  ('superadmin@superadmin.com'),
  ('admin@admin.com'),
  ('liuguoxinn87@gmail.com');