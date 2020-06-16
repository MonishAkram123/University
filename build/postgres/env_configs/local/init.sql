CREATE DATABASE university;
CREATE ROLE local_admin WITH LOGIN PASSWORD 'local_password';
GRANT ALL PRIVILEGES ON DATABASE university TO local_admin;