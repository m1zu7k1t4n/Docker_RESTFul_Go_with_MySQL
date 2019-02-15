USE test;
-- Insert initlal data
START TRANSACTION;
INSERT INTO tasks(title, body) VALUES('First task', 'test data1');
INSERT INTO tasks(title, body) VALUES('Second task', 'test data2');
INSERT INTO tasks(title, body) VALUES('Third task', 'test data3');
INSERT INTO tasks(title, body) VALUES('Dummy Data4', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data5', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data6', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data7', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data8', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data9', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data10', 'long long long long long');
INSERT INTO tasks(title, body) VALUES('Dummy Data11', 'long long long long long');
COMMIT;
