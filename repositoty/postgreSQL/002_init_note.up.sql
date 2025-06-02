CREATE TABLE notes (
    title TEXT NOT NULL,
    text TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_time TIMESTAMP WITH TIME ZONE NOT NULL,
    last_change TIMESTAMP WITH TIME ZONE NOT NULL,
    
    -- Уникальный композитный индекс (пользователь + заголовок)
    CONSTRAINT unique_user_note UNIQUE (user_id, title)
);

-- Индекс для поиска по пользователю
CREATE INDEX idx_notes_user ON notes (user_id);