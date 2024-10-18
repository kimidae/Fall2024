DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(50),
    age VARCHAR(50)
);
INSERT INTO users (user_name, age) VALUES
('Alice', 24),
('Bob', 48),
('Cathy', 36),
('Jake', 25),
('David', 54);

SELECT * FROM users;
