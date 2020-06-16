CREATE DATABASE university;
CREATE ROLE stage_admin WITH LOGIN PASSWORD 'stage_password';
GRANT ALL PRIVILEGES ON DATABASE university TO stage_admin;