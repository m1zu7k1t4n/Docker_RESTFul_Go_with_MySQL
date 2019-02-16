CREATE DATABASE IF NOT EXISTS test;
CREATE TABLE IF NOT EXISTS test.tasks(
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100),
  body VARCHAR(1000),
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp on update current_timestamp
);
