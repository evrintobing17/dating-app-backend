CREATE TABLE IF NOT EXISTS swipes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    profile_id INT NOT NULL,
    action VARCHAR(10) CHECK (action IN ('like', 'pass')),
    swiped_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
