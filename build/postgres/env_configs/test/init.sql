CREATE DATABASE test_university;
CREATE ROLE test_user WITH LOGIN PASSWORD 'test_password';
GRANT ALL PRIVILEGES ON DATABASE test_university TO test_user;