CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name_task VARCHAR(255) NOT NULL,
    description TEXT,
    status_task VARCHAR(50),
    priority VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
