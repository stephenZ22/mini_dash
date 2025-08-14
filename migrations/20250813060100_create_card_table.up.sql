CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description  TEXT,
    card_type SMAllINT,
    parent_id INT REFERENCES cards(id) ON DELETE SET NULL,
    creater_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
)