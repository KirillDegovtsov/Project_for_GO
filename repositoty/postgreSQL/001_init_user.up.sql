CREATE TABLE users (
    id TEXT PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

-- Для поиска по логину (ускоряет поиск и проверку уникальности)
CREATE INDEX idx_users_login ON users (login);