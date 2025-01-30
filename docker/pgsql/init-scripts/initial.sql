-- Create the database
CREATE DATABASE telescope_dev;
-- Connect to the database to run the following commands within its context
\connect telescope_dev;
-- Create the user with the specified password
-- ALTER SCHEMA public OWNER TO admin;
CREATE USER admin WITH PASSWORD '123qwe';
ALTER USER admin WITH SUPERUSER;
GRANT ALL PRIVILEGES ON DATABASE telescope_dev TO admin;
-- Grant usage and create privileges on the public schema to the user

GRANT ALL PRIVILEGES ON SCHEMA public TO admin;

GRANT USAGE,
  CREATE ON SCHEMA public TO admin;
GRANT CONNECT ON DATABASE telescope_dev TO admin;
